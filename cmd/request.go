package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func postRequest(url string, method string, params []string) {
	var paramsString []string

	for _, param := range params {
		var quoted string

		if param != "true" && param != "false" {
			quoted = fmt.Sprintf(`"%s"`, param)
		} else {
			quoted = param
		}

		paramsString = append(paramsString, quoted)
	}

	jsonStr := []byte(fmt.Sprintf(
		`{"id": 0, "jsonrpc":"2.0", "method": "%s", "params": [%s]}`,
		method, strings.Join(paramsString, ","),
	))

	if verboseFlag {
		fmt.Println(string(jsonStr))
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
