package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func main() {
	var fileaddr string
	var err error
	taskId := flag.Int("task", 0, "an int")
	fileJSON := flag.String("file", "", "a string")
	link := flag.String("link", "http://localhost", "a string")
	port := flag.Int("port", 8080, "an int")
	url := *link + ":" + strconv.Itoa(*port)
	flag.Parse()
	if *fileJSON == "" {
		fmt.Print("Please, write an address to json file with initial data: ")
		fmt.Scanln(&fileaddr)
		if err != nil {
			fmt.Println("Something went wrong \n", err.Error())
			os.Exit(0)
		}
	} else {
		fileaddr = *fileJSON
	}
	res, err := ioutil.ReadFile(fileaddr)
	if err != nil {
		fmt.Println("Problem with json file opening with initial data. \n", err.Error())
		os.Exit(0)
	}
	var resource string
	dataJson := res
	if *taskId == 0 {
		resource = url + "/tasks"
	} else {
		var params map[string]json.RawMessage
		json.Unmarshal(res, &params)
		resource = url + "/task/" + strconv.Itoa(*taskId)
		dataJson = params[strconv.Itoa(*taskId)]
	}
	req, err := http.NewRequest("POST", resource, bytes.NewBuffer(dataJson))
	if err != nil {
		fmt.Println("Some problem with your server: ", resource, "\n", err.Error())
		os.Exit(0)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Some problem with your server: ", resource, "\n", err.Error())
		os.Exit(0)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	// Technical info
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	fmt.Println("response Body:", string(body))

	type WriteBack struct {
		Task   int    `json:"task"`
		Resp   string `json:"resp"`
		Reason string `json:"reason"`
	}
	var returnData []WriteBack

	err = json.Unmarshal(body, &returnData)
	if err != nil {
		var returnTask WriteBack
		err = json.Unmarshal(body, &returnTask)
		if err != nil {
			fmt.Println("I can't unmarshal json: ", err.Error())
		}
		if returnTask.Reason != "" {
			fmt.Println("Task:", returnTask.Task, " failed\nReason: \n", returnTask.Reason)
		} else {
			fmt.Println("Task:", returnTask.Task, "\nResult:")
			fmt.Println(returnTask.Resp)
		}
	}
	for _, results := range returnData {
		if results.Reason != "" {
			fmt.Println("\nTask:", results.Task, "failed\nReason: ", results.Reason)
		} else {
			fmt.Println("\nTask:", results.Task, "\nResult:")
			fmt.Println(results.Resp)
		}
	}
}
