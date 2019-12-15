package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// simply post some demo values
	out, err := exec.Command("git", "rev-parse", "HEAD").Output()
	if err != nil {
		panic(err)
	}

	// create custom request
	data := map[string]interface{}{
		"sha": strings.TrimSpace(string(out)),
		"values": []map[string]interface{}{
			{
				"value": "99 %",
				"line":  "a",
			},
			{
				"value": "50 %",
				"line":  "b",
			},
			{
				"value": "82 %",
				"line":  "c",
			},
		},
	}

	var b bytes.Buffer
	if err := json.NewEncoder(&b).Encode(data); err != nil {
		panic(err)
	}

	uri := "https://seriesci.com/api/repos/seriesci/multiple/demo/multi"
	req, err := http.NewRequest(http.MethodPost, uri, &b)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Token %s", os.Getenv("SERIESCI_TOKEN")))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
}
