package libs

import (
	"database/sql"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/accuknox/knoxAutoPolicy/src/types"

	_ "github.com/go-sql-driver/mysql"
)

// ConnectMySQL function
func ConnectMySQL(cfg types.ConfigDB) (db *sql.DB) {
	db, err := sql.Open(cfg.DBDriver, cfg.DBUser+":"+cfg.DBPass+"@tcp("+cfg.DBHost+":"+cfg.DBPort+")/"+cfg.DBName)
	for err != nil {
		log.Error().Msg("connection error :" + err.Error())
		time.Sleep(time.Second * 1)
		db, err = sql.Open(cfg.DBDriver, cfg.DBUser+":"+cfg.DBPass+"@tcp("+cfg.DBHost+":"+cfg.DBPort+")/"+cfg.DBName)
	}
	db.SetMaxIdleConns(0)
	return db
}

// ================= //
// == Network Log == //
// ================= //

var networkLogQueryBase string = "select id,time,cluster_name,traffic_direction,verdict,policy_match_type,drop_reason,event_type,source,destination,ip,l4,l7 from "

func convertDateTimeToUnix(dateTime string) (int64, error) {
	thetime, err := time.Parse(time.RFC3339, dateTime)
	if err != nil {
		return 0, err
	}
	return thetime.Unix(), nil
}

func convertJSONRawToString(raw json.RawMessage) string {
	j, _ := json.Marshal(&raw)
	return string(j)
}

// ScanNetworkLogs scans the db records
func ScanNetworkLogs(results *sql.Rows) ([]map[string]interface{}, error) {
	networkLogs := []map[string]interface{}{}
	var err error

	for results.Next() {
		var id, time, policyMatchTypeInt, dropReasonInt uint32
		var policyMatchType, dropReason sql.NullInt32
		var verdict, direction, clusterName string
		var srcByte, destByte, eventTypeByte []byte
		var ipByte, l4Byte, l7Byte []byte

		err = results.Scan(
			&id,
			&time,
			&clusterName,
			&direction,
			&verdict,
			&policyMatchType,
			&dropReason,
			&eventTypeByte,
			&srcByte,
			&destByte,
			&ipByte,
			&l4Byte,
			&l7Byte,
		)

		if policyMatchType.Valid {
			policyMatchTypeInt = uint32(policyMatchType.Int32)
		}

		if dropReason.Valid {
			dropReasonInt = uint32(dropReason.Int32)
		}

		if err != nil {
			log.Error().Msg("Error while scanning network logs :" + err.Error())
			return nil, err
		}

		log := map[string]interface{}{
			"id":                id,
			"time":              time,
			"cluster_name":      clusterName,
			"traffic_direction": direction,
			"verdict":           verdict,
			"policy_match_type": policyMatchTypeInt,
			"drop_reason":       dropReasonInt,
			"event_type":        eventTypeByte,
			"source":            srcByte,
			"destination":       destByte,
			"ip":                ipByte,
			"l4":                l4Byte,
			"l7":                l7Byte,
		}

		networkLogs = append(networkLogs, log)
	}

	return networkLogs, nil
}

// GetNetworkLogByTimeFromMySQL function
func GetNetworkLogByTimeFromMySQL(cfg types.ConfigDB, startTime, endTime int64) ([]map[string]interface{}, error) {
	db := ConnectMySQL(cfg)
	defer db.Close()

	QueryBase := networkLogQueryBase + cfg.TableNetworkFlow

	rows, err := db.Query(QueryBase+" WHERE time >= ? and time <= ?", int(startTime), int(endTime))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ScanNetworkLogs(rows)
}

// GetNetworkLogByIDTimeFromMySQL function
func GetNetworkLogByIDTimeFromMySQL(cfg types.ConfigDB, id, endTime int64) ([]map[string]interface{}, error) {
	db := ConnectMySQL(cfg)
	defer db.Close()

	QueryBase := networkLogQueryBase + cfg.TableNetworkFlow

	rows, err := db.Query(QueryBase+" WHERE id > ? ORDER BY id ASC ", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ScanNetworkLogs(rows)
}

// InsertNetworkLogToMySQL function
func InsertNetworkLogToMySQL(cfg types.ConfigDB, nfe []types.NetworkFlowEvent) error {
	db := ConnectMySQL(cfg)
	defer db.Close()

	sqlStr := "INSERT INTO " + cfg.TableNetworkFlow + "(time,cluster_name,verdict,drop_reason,ethernet,ip,l4,l7,reply,source,destination,type,node_name,event_type,source_service,destination_service,traffic_direction,policy_match_type,trace_observation_point,summary) VALUES "
	vals := []interface{}{}

	for _, e := range nfe {
		sqlStr += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?),"
		unixTime, err := convertDateTimeToUnix(e.Time)
		if err != nil {
			log.Error().Msgf("Error converting date time to timestamp: %s", err.Error())
		}

		vals = append(vals,
			unixTime,
			e.ClusterName,
			e.Verdict,
			e.DropReason,
			convertJSONRawToString(e.Ethernet),
			convertJSONRawToString(e.IP),
			convertJSONRawToString(e.L4),
			convertJSONRawToString(e.L7),
			e.Reply,
			convertJSONRawToString(e.Source),
			convertJSONRawToString(e.Destination),
			e.Type,
			e.NodeName,
			convertJSONRawToString(e.EventType),
			convertJSONRawToString(e.SourceService),
			convertJSONRawToString(e.DestinationService),
			e.TrafficDirection,
			e.PolicyMatchType,
			e.TraceObservationPoint,
			e.Summary)
	}

	//trim the last ','
	sqlStr = strings.TrimSuffix(sqlStr, ",")

	//prepare the statement
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	//format all vals at once
	_, err = stmt.Exec(vals...)
	if err != nil {
		return err
	}

	return nil
}

// ================ //
// == System Log == //
// ================ //

var systemLogQueryBase string = "select id,time,cluster_name,traffic_direction,verdict,policy_match_type,drop_reason,event_type,source,destination,ip,l4,l7 from "

// ScanSystemLogs scans the db records
func ScanSystemLogs(results *sql.Rows) ([]map[string]interface{}, error) {
	systemLogs := []map[string]interface{}{}
	var err error

	for results.Next() {
		var id, time, hostPid, ppid, pid, uid uint32
		var clusterName, nodeName, namespace, podName, containerID, containerName, types, source, operation, resource, data, result string

		err = results.Scan(
			&id,
			&time,
			&clusterName,
			&nodeName,
			&namespace,
			&podName,
			&containerID,
			&containerName,
			&hostPid,
			&ppid,
			&pid,
			&uid,
			&types,
			&source,
			&operation,
			&resource,
			&data,
			&result,
		)

		if err != nil {
			log.Error().Msg("Error while scanning system logs :" + err.Error())
			return nil, err
		}

		log := map[string]interface{}{
			"id":           id,
			"time":         time,
			"cluster_name": clusterName,
		}

		systemLogs = append(systemLogs, log)
	}

	return systemLogs, nil
}

// GetSystemLogByTimeFromMySQL function
func GetSystemLogByTimeFromMySQL(cfg types.ConfigDB, startTime, endTime int64) ([]map[string]interface{}, error) {
	db := ConnectMySQL(cfg)
	defer db.Close()

	QueryBase := systemLogQueryBase + cfg.TableSystemLog

	rows, err := db.Query(QueryBase+" WHERE time >= ? and time <= ?", int(startTime), int(endTime))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ScanSystemLogs(rows)
}

// GetSystemLogByIDTimeFromMySQL function
func GetSystemLogByIDTimeFromMySQL(cfg types.ConfigDB, id, endTime int64) ([]map[string]interface{}, error) {
	db := ConnectMySQL(cfg)
	defer db.Close()

	QueryBase := systemLogQueryBase + cfg.TableSystemLog

	rows, err := db.Query(QueryBase+" WHERE id > ? ORDER BY id ASC ", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return ScanSystemLogs(rows)
}

// InsertSystemLogToMySQL function
func InsertSystemLogToMySQL(cfg types.ConfigDB, sle []types.SystemLogEvent) error {
	db := ConnectMySQL(cfg)
	defer db.Close()

	sqlStr := "INSERT INTO " + cfg.TableSystemLog + "(time,cluster_name,node_name,namespace_name,pod_name,container_id,container_name,hostpid,ppid,pid,uid,type,source,operation,resource,data,result) VALUES "
	vals := []interface{}{}

	for _, e := range sle {
		sqlStr += "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?),"

		vals = append(vals,
			e.Timestamp,
			e.ClusterName,
			e.HostName,
			e.NamespaceName,
			e.PodName,
			e.ContainerID,
			e.ContainerName,
			e.HostPID,
			e.PPID,
			e.PID,
			e.UID,
			e.Type,
			e.Source,
			e.Operation,
			e.Resource,
			e.Data,
			e.Result)
	}

	//trim the last ','
	sqlStr = strings.TrimSuffix(sqlStr, ",")

	//prepare the statement
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	//format all vals at once
	_, err = stmt.Exec(vals...)
	if err != nil {
		return err
	}

	return nil
}

// ==================== //
// == Network Policy == //
// ==================== //

// GetNetworkPoliciesFromMySQL function
func GetNetworkPoliciesFromMySQL(cfg types.ConfigDB, namespace, status string) ([]types.KnoxNetworkPolicy, error) {
	db := ConnectMySQL(cfg)
	defer db.Close()

	policies := []types.KnoxNetworkPolicy{}
	var results *sql.Rows
	var err error

	query := "SELECT apiVersion,kind,flow_ids,name,cluster_name,namespace,type,rule,status,outdated,spec,generatedTime FROM " + cfg.TableDiscoveredPolicies
	if namespace != "" && status != "" {
		query = query + " WHERE namespace = ? and status = ? "
		results, err = db.Query(query, namespace, status)
	} else if namespace != "" {
		query = query + " WHERE namespace = ? "
		results, err = db.Query(query, namespace)
	} else if status != "" {
		query = query + " WHERE status = ? "
		results, err = db.Query(query, status)
	} else {
		results, err = db.Query(query)
	}

	defer results.Close()

	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}

	for results.Next() {
		policy := types.KnoxNetworkPolicy{}

		var name, clusterName, namespace, policyType, rule, status string
		specByte := []byte{}
		spec := types.Spec{}

		flowIDsByte := []byte{}
		flowIDs := []int{}

		if err := results.Scan(
			&policy.APIVersion,
			&policy.Kind,
			&flowIDsByte,
			&name,
			&clusterName,
			&namespace,
			&policyType,
			&rule,
			&status,
			&policy.Outdated,
			&specByte,
			&policy.GeneratedTime,
		); err != nil {
			return nil, err
		}

		if err := json.Unmarshal(specByte, &spec); err != nil {
			return nil, err
		}

		if err := json.Unmarshal(flowIDsByte, &flowIDs); err != nil {
			return nil, err
		}

		policy.Metadata = map[string]string{
			"name":         name,
			"cluster_name": clusterName,
			"namespace":    namespace,
			"type":         policyType,
			"rule":         rule,
			"status":       status,
		}

		policy.FlowIDs = flowIDs
		policy.Spec = spec

		policies = append(policies, policy)
	}

	return policies, nil
}

// UpdateOutdatedPolicyFromMySQL ...
func UpdateOutdatedPolicyFromMySQL(cfg types.ConfigDB, outdatedPolicy string, latestPolicy string) error {
	db := ConnectMySQL(cfg)
	defer db.Close()

	var err error

	// set status -> outdated
	stmt1, err := db.Prepare("UPDATE " + cfg.TableDiscoveredPolicies + " SET status=? WHERE name=?")
	if err != nil {
		return err
	}
	defer stmt1.Close()

	_, err = stmt1.Exec("outdated", outdatedPolicy)
	if err != nil {
		return err
	}

	// set outdated -> latest' name
	stmt2, err := db.Prepare("UPDATE " + cfg.TableDiscoveredPolicies + " SET outdated=? WHERE name=?")
	if err != nil {
		return err
	}
	defer stmt2.Close()

	_, err = stmt2.Exec(latestPolicy, outdatedPolicy)
	if err != nil {
		return err
	}

	return nil
}

// insertNetworkPolicy function
func insertNetworkPolicy(cfg types.ConfigDB, db *sql.DB, policy types.KnoxNetworkPolicy) error {
	stmt, err := db.Prepare("INSERT INTO " + cfg.TableDiscoveredPolicies + "(apiVersion,kind,flow_ids,name,cluster_name,namespace,type,rule,status,outdated,spec,generatedTime) values(?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	flowIDsPointer := &policy.FlowIDs
	flowids, err := json.Marshal(flowIDsPointer)
	if err != nil {
		return err
	}

	specPointer := &policy.Spec
	spec, err := json.Marshal(specPointer)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(policy.APIVersion,
		policy.Kind,
		flowids,
		policy.Metadata["name"],
		policy.Metadata["cluster_name"],
		policy.Metadata["namespace"],
		policy.Metadata["type"],
		policy.Metadata["rule"],
		policy.Metadata["status"],
		policy.Outdated,
		spec,
		policy.GeneratedTime)
	if err != nil {
		return err
	}

	return nil
}

// InsertNetworkPoliciesToMySQL function
func InsertNetworkPoliciesToMySQL(cfg types.ConfigDB, policies []types.KnoxNetworkPolicy) error {
	db := ConnectMySQL(cfg)
	defer db.Close()

	for _, policy := range policies {
		if err := insertNetworkPolicy(cfg, db, policy); err != nil {
			return err
		}
	}

	return nil
}

// =================== //
// == Configuration == //
// =================== //

// CountConfigByName ...
func CountConfigByName(db *sql.DB, tableName, configName string) int {
	var count int
	db.QueryRow("SELECT COUNT(*) FROM "+tableName+" WHERE config_name=?", configName).Scan(&count)
	return count
}

// AddConfiguration function
func AddConfiguration(cfg types.ConfigDB, newConfig types.Configuration) error {
	db := ConnectMySQL(cfg)
	defer db.Close()

	if CountConfigByName(db, cfg.TableConfiguration, newConfig.ConfigName) > 0 {
		return errors.New("Already exist config name: " + newConfig.ConfigName)
	}

	stmt, err := db.Prepare("INSERT INTO " +
		cfg.TableConfiguration +
		"(config_name," +
		"status," +
		"config_db," +
		"config_cilium_hubble," +
		"operation_mode," +
		"cronjob_time_interval," +
		"one_time_job_time_selection," +
		"network_log_from," +
		"discovered_policy_to," +
		"policy_dir," +
		"discovery_policy_types," +
		"discovery_rule_types," +
		"cidr_bits," +
		"ignoring_flows," +
		"l3_aggregation_level," +
		"l4_aggregation_level," +
		"l7_aggregation_level) " +
		"values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")

	if err != nil {
		return err
	}

	defer stmt.Close()

	configDBPtr := &newConfig.ConfigDB
	configDB, err := json.Marshal(configDBPtr)
	if err != nil {
		return err
	}

	configHubblePtr := &newConfig.ConfigCiliumHubble
	configCilium, err := json.Marshal(configHubblePtr)
	if err != nil {
		return err
	}

	ignoringFlowsPtr := &newConfig.IgnoringFlows
	ignoringFlows, err := json.Marshal(ignoringFlowsPtr)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(newConfig.ConfigName,
		newConfig.Status,
		configDB,
		configCilium,
		newConfig.OperationMode,
		newConfig.CronJobTimeInterval,
		newConfig.OneTimeJobTimeSelection,
		newConfig.NetworkLogFrom,
		newConfig.NetworkPolicyTo,
		newConfig.NetworkPolicyDir,
		newConfig.DiscoveryPolicyTypes,
		newConfig.DiscoveryRuleTypes,
		newConfig.CIDRBits,
		ignoringFlows,
		newConfig.L3AggregationLevel,
		newConfig.L4Compression,
		newConfig.L7AggregationLevel,
	)

	if err != nil {
		return err
	}

	return nil
}

// GetConfigurations function
func GetConfigurations(cfg types.ConfigDB, configName string) ([]types.Configuration, error) {
	db := ConnectMySQL(cfg)
	defer db.Close()

	configs := []types.Configuration{}
	var results *sql.Rows
	var err error

	query := "SELECT * FROM " + cfg.TableConfiguration
	if configName != "" {
		query = query + " WHERE config_name = ? "
		results, err = db.Query(query, configName)
	}

	defer results.Close()

	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}

	for results.Next() {
		cfg := types.Configuration{}

		id := 0
		configDBByte := []byte{}
		configDB := types.ConfigDB{}

		hubbleByte := []byte{}
		hubble := types.ConfigCiliumHubble{}

		ignoringFlowByte := []byte{}
		ignoringFlows := []types.IgnoringFlows{}

		if err := results.Scan(
			&id,
			&cfg.ConfigName,
			&cfg.Status,
			&configDBByte,
			&hubbleByte,
			&cfg.OperationMode,
			&cfg.CronJobTimeInterval,
			&cfg.OneTimeJobTimeSelection,
			&cfg.NetworkLogFrom,
			&cfg.NetworkPolicyTo,
			&cfg.NetworkPolicyDir,
			&cfg.DiscoveryPolicyTypes,
			&cfg.DiscoveryRuleTypes,
			&cfg.CIDRBits,
			&ignoringFlowByte,
			&cfg.L3AggregationLevel,
			&cfg.L4Compression,
			&cfg.L7AggregationLevel,
		); err != nil {
			return nil, err
		}

		if err := json.Unmarshal(configDBByte, &configDB); err != nil {
			return nil, err
		}
		cfg.ConfigDB = configDB

		if err := json.Unmarshal(hubbleByte, &hubble); err != nil {
			return nil, err
		}
		cfg.ConfigCiliumHubble = hubble

		if err := json.Unmarshal(ignoringFlowByte, &ignoringFlows); err != nil {
			return nil, err
		}
		cfg.IgnoringFlows = ignoringFlows

		configs = append(configs, cfg)
	}

	return configs, nil
}

// UpdateConfiguration ...
func UpdateConfiguration(cfg types.ConfigDB, configName string, updateConfig types.Configuration) error {
	db := ConnectMySQL(cfg)
	defer db.Close()

	var err error

	stmt, err := db.Prepare("UPDATE " + cfg.TableConfiguration + " SET " +
		"config_name=?," +
		"config_db=?," +
		"config_cilium_hubble=?," +
		"operation_mode=?," +
		"cronjob_time_interval=?," +
		"one_time_job_time_selection=?," +
		"network_log_from=?," +
		"discovered_policy_to=?," +
		"policy_dir=?," +
		"discovery_policy_types=?," +
		"discovery_rule_types=?," +
		"cidr_bits=?," +
		"ignoring_flows=?," +
		"l3_aggregation_level=?," +
		"l4_aggregation_level=?," +
		"l7_aggregation_level=? " +
		"WHERE config_name=?")

	if err != nil {
		return err
	}
	defer stmt.Close()

	configDBPtr := &updateConfig.ConfigDB
	configDB, err := json.Marshal(configDBPtr)
	if err != nil {
		return err
	}

	configHubblePtr := &updateConfig.ConfigCiliumHubble
	configCilium, err := json.Marshal(configHubblePtr)
	if err != nil {
		return err
	}

	ignoringFlowsPtr := &updateConfig.IgnoringFlows
	ignoringFlows, err := json.Marshal(ignoringFlowsPtr)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		updateConfig.ConfigName,
		configDB,
		configCilium,
		updateConfig.OperationMode,
		updateConfig.CronJobTimeInterval,
		updateConfig.OneTimeJobTimeSelection,
		updateConfig.NetworkLogFrom,
		updateConfig.NetworkPolicyTo,
		updateConfig.NetworkPolicyDir,
		updateConfig.DiscoveryPolicyTypes,
		updateConfig.DiscoveryRuleTypes,
		updateConfig.CIDRBits,
		ignoringFlows,
		updateConfig.L3AggregationLevel,
		updateConfig.L4Compression,
		updateConfig.L7AggregationLevel,
		configName,
	)

	if err != nil {
		return err
	}

	return nil
}

// DeleteConfiguration ...
func DeleteConfiguration(cfg types.ConfigDB, configName string) error {
	db := ConnectMySQL(cfg)
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM " + cfg.TableConfiguration + " WHERE config_name=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(configName)
	if err != nil {
		return err
	}

	return nil
}

// ApplyConfiguration ...
func ApplyConfiguration(cfg types.ConfigDB, oldConfigName, configName string) error {
	db := ConnectMySQL(cfg)
	defer db.Close()

	var err error
	stmt1, err := db.Prepare("UPDATE " + cfg.TableConfiguration + " SET status=? WHERE config_name=?")
	if err != nil {
		return err
	}
	defer stmt1.Close()

	_, err = stmt1.Exec(0, oldConfigName)
	if err != nil {
		return err
	}

	stmt2, err := db.Prepare("UPDATE " + cfg.TableConfiguration + " SET status=? WHERE config_name=?")
	if err != nil {
		return err
	}
	defer stmt2.Close()

	_, err = stmt2.Exec(1, configName)
	if err != nil {
		return err
	}

	return nil
}

// =========== //
// == Table == //
// =========== //

// ClearDBTablesMySQL function
func ClearDBTablesMySQL(cfg types.ConfigDB) error {
	db := ConnectMySQL(cfg)
	defer db.Close()

	query := "DELETE FROM " + cfg.TableDiscoveredPolicies
	if _, err := db.Query(query); err != nil {
		return err
	}

	query = "DELETE FROM " + cfg.TableNetworkFlow
	if _, err := db.Query(query); err != nil {
		return err
	}

	query = "DELETE FROM " + cfg.TableSystemLog
	if _, err := db.Query(query); err != nil {
		return err
	}

	return nil
}

// CreateTableNetworkFlowMySQL function
func CreateTableNetworkFlowMySQL(cfg types.ConfigDB) error {
	db := ConnectMySQL(cfg)
	defer db.Close()

	tableName := cfg.TableNetworkFlow

	query :=
		"CREATE TABLE IF NOT EXISTS `" + tableName + "` (" +
			"	`id` int NOT NULL AUTO_INCREMENT," +
			"	`time` int DEFAULT NULL," +
			"	`cluster_name` varchar(100) DEFAULT NULL," +
			"	`verdict` varchar(50) DEFAULT NULL," +
			"	`drop_reason` INT DEFAULT NULL," +
			"	`ethernet` JSON DEFAULT NULL," +
			"	`ip` JSON DEFAULT NULL," +
			"	`l4` JSON DEFAULT NULL," +
			"	`l7` JSON DEFAULT NULL," +
			"	`reply` BOOLEAN," +
			"	`source` JSON DEFAULT NULL," +
			"	`destination` JSON DEFAULT NULL," +
			"	`type` varchar(50) DEFAULT NULL," +
			"	`node_name` varchar(100) DEFAULT NULL," +
			"	`event_type` JSON DEFAULT NULL," +
			"	`source_service` JSON DEFAULT NULL," +
			"	`destination_service` JSON DEFAULT NULL," +
			"	`traffic_direction` varchar(50) DEFAULT NULL," +
			"	`policy_match_type` int DEFAULT NULL," +
			"	`trace_observation_point` varchar(100) DEFAULT NULL," +
			"	`summary` varchar(1000) DEFAULT NULL," +
			"	PRIMARY KEY (`id`)" +
			");"

	if _, err := db.Query(query); err != nil {
		return err
	}

	return nil
}

// CreateTableDiscoveredPoliciesMySQL function
func CreateTableDiscoveredPoliciesMySQL(cfg types.ConfigDB) error {
	db := ConnectMySQL(cfg)
	defer db.Close()

	tableName := cfg.TableDiscoveredPolicies

	query :=
		"CREATE TABLE IF NOT EXISTS `" + tableName + "` (" +
			"	`id` int NOT NULL AUTO_INCREMENT," +
			"	`apiVersion` varchar(20) DEFAULT NULL," +
			"	`kind` varchar(20) DEFAULT NULL," +
			"	`flow_ids` JSON DEFAULT NULL," +
			"	`name` varchar(50) DEFAULT NULL," +
			"	`cluster_name` varchar(50) DEFAULT NULL," +
			"	`namespace` varchar(50) DEFAULT NULL," +
			"	`type` varchar(10) DEFAULT NULL," +
			"	`rule` varchar(30) DEFAULT NULL," +
			"	`status` varchar(10) DEFAULT NULL," +
			"	`outdated` varchar(50) DEFAULT NULL," +
			"	`spec` JSON DEFAULT NULL," +
			"	`generatedTime` int DEFAULT NULL," +
			"	PRIMARY KEY (`id`)" +
			"  );"

	if _, err := db.Query(query); err != nil {
		return err
	}

	return nil
}

// CreateTableConfigurationMySQL function
func CreateTableConfigurationMySQL(cfg types.ConfigDB) error {
	db := ConnectMySQL(cfg)
	defer db.Close()

	tableName := cfg.TableConfiguration

	query :=
		"CREATE TABLE IF NOT EXISTS `" + tableName + "` (" +
			"	`id` int NOT NULL AUTO_INCREMENT," +
			"	`config_name` varchar(50) DEFAULT NULL," +
			"	`status` int DEFAULT '0'," +
			"	`config_db` JSON DEFAULT NULL," +
			"	`config_cilium_hubble` JSON DEFAULT NULL," +
			"	`operation_mode` int DEFAULT NULL," +
			"	`cronjob_time_interval` varchar(50) DEFAULT NULL," +
			"	`one_time_job_time_selection` varchar(50) DEFAULT NULL," +
			"	`network_log_from` varchar(50) DEFAULT NULL," +
			"	`discovered_policy_to` varchar(50) DEFAULT NULL," +
			"	`policy_dir` varchar(50) DEFAULT NULL," +
			"	`discovery_policy_types` int DEFAULT NULL," +
			"	`discovery_rule_types` int DEFAULT NULL," +
			"	`cidr_bits` int DEFAULT NULL," +
			"	`ignoring_flows` JSON DEFAULT NULL," +
			"	`l3_aggregation_level` int DEFAULT NULL," +
			"	`l4_aggregation_level` int DEFAULT NULL," +
			"	`l7_aggregation_level` int DEFAULT NULL," +
			"	PRIMARY KEY (`id`)" +
			"  );"

	if _, err := db.Query(query); err != nil {
		return err
	}

	return nil
}

// CreateTableSystemLogMySQL function
func CreateTableSystemLogMySQL(cfg types.ConfigDB) error {
	db := ConnectMySQL(cfg)
	defer db.Close()

	tableName := cfg.TableSystemLog

	query :=
		"CREATE TABLE IF NOT EXISTS `" + tableName + "` (" +
			"`id` int NOT NULL AUTO_INCREMENT," +
			"`time` int DEFAULT NULL," +
			"`cluster_name` varchar(100) DEFAULT NULL," +
			"`node_name` varchar(100) DEFAULT NULL," +
			"`namespace_name` varchar(100) DEFAULT NULL," +
			"`pod_name` varchar(100) DEFAULT NULL," +
			"`container_id` varchar(100) DEFAULT NULL," +
			"`container_name` varchar(100) DEFAULT NULL," +
			"`host_pid` int DEFAULT NULL," +
			"`ppid` int DEFAULT NULL," +
			"`pid` int DEFAULT NULL," +
			"`uid` int DEFAULT NULL," +
			"`type` varchar(20) DEFAULT NULL," +
			"`source` varchar(200) DEFAULT NULL," +
			"`operation` varchar(20) DEFAULT NULL," +
			"`resource` varchar(200) DEFAULT NULL," +
			"`data` varchar(100) DEFAULT NULL," +
			"`result` varchar(100) DEFAULT NULL," +
			"	PRIMARY KEY (`id`)" +
			");"

	if _, err := db.Query(query); err != nil {
		return err
	}

	return nil
}