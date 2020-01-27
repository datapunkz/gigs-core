package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// our main function
func main() {
	jsonFile, err := os.Open("pkg/models/users.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("JSON File Opened")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)
	fmt.Println(result["data"])
}
