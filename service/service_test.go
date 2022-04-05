package service

import (
	"testing"
	"encoding/json"
)

func TestJoinNames(t *testing.T) {

	var dataset Data
	jsonData := `{"data":{"projects":{"nodes":[{"name":"hcs_utils","description":"","forksCount":1},{"name":"K","description":null,"forksCount":1},{"name":"Heroes of Wesnoth","description":null,"forksCount":5},{"name":"Leiningen","description":"","forksCount":1},{"name":"TearDownWalls","description":null,"forksCount":5}]}}}`
	json.Unmarshal([]byte(jsonData), &dataset)

	expectedData := NodeData {
		Name: "hcs_utils, K, Heroes of Wesnoth, Leiningen, TearDownWalls",
		ForksCount: 13,
	}
	actualData := JoinNames(dataset)
	
	if actualData != expectedData {
		t.Errorf("Actual data and expected data dont match")
	}
}