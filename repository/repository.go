package repository

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
	"fmt"
	"strings"
	"bd_test_task_three/service"
)

// this is the repo layer
// sends a request to the url https://gitlab.com/api/graphql
// takes an input of number of records that are needed from the user
// sends a call to service layer to get the results and return sum of all fork counts and names with a comma delimeter
func RepoLayer() service.Data{

	var dataset service.Data
	var i int
	var payloadReader *strings.Reader
	fmt.Println("Enter number of records you want to retrieve: ")
	_, err := fmt.Scanf("%d", &i)

	// url to send the request
	url := "https://gitlab.com/api/graphql"
	method := "POST"

	payload := fmt.Sprintf("{\"query\":\"query last_projects($n: Int = %v) {\\n  projects(last:$n) {\\n    nodes {\\n      name\\n      description\\n      forksCount\\n    }\\n  }\\n}\",\"variables\":{}}", i)
	payloadReader = strings.NewReader(payload)
	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payloadReader)

	if err != nil {
		fmt.Println(err)
		return dataset
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", "experimentation_subject_id=eyJfcmFpbHMiOnsibWVzc2FnZSI6IklqVTVNR1EyTlRjMExXUmxaV0l0TkRCbE5pMDRPV0UxTFRFM01USmtZMlptTUdabU5DST0iLCJleHAiOm51bGwsInB1ciI6ImNvb2tpZS5leHBlcmltZW50YXRpb25fc3ViamVjdF9pZCJ9fQ%3D%3D--5727874767db5bcef406cb77f37832df33bf0bb9")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return dataset
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return dataset
	}
	fmt.Println(string(body))

	// records unmarshaled and returned to main and sent to process in the service layer
	json.Unmarshal([]byte(body), &dataset)
	return dataset
}