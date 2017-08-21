package util

// GetMapList - Get a list of maps out of a list of interfaces
func GetMapList(v []interface{}) []map[string]interface{} {
	serverList := make([]map[string]interface{}, 0)
	if len(v) > 0 {
		for _, server := range v {
			serverList = append(serverList, server.(map[string]interface{}))
		}
	}
	return serverList
}
