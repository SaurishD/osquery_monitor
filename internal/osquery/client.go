// internal/osquery/client.go

package osquery

import (
	"encoding/json"
	"os/exec"
)

func runQuery(query string) ([]map[string]string, error) {
	cmd := exec.Command("osqueryi", "--json", query)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var result []map[string]string
	err = json.Unmarshal(output, &result)
	return result, err
}

func GetOSInfo() (map[string]string, error) {
	data, err := runQuery("SELECT * FROM os_version;")
	if err != nil || len(data) == 0 {
		return nil, err
	}
	return data[0], nil
}

func GetOsqueryVersion() (map[string]string, error) {
	data, err := runQuery("SELECT * FROM osquery_info;")
	if err != nil || len(data) == 0 {
		return nil, err
	}
	return data[0], nil
}

func GetInstalledApps() ([]map[string]string, error) {
	// Use best available table for your OS
	return runQuery("SELECT bundle_identifier as id, display_name as name, bundle_version as version, path from apps limit 1;")
}
