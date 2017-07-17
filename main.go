package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"encoding/json"
)

func main() {
	var fileaddr string
	var err error
	taskId := flag.Int("task", 0, "an int")
	fileJSON := flag.String("file", "", "a string")
	link := flag.String("link", "http://localhost", "a string")
	port := flag.Int("port", 8080, "an int")
	url := *link+":"+strconv.Itoa(*port)
	flag.Parse()
	if *fileJSON == "" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Please, write an address to json file with initial data.")
		fileaddr, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Something went wrong \n", err.Error())
			os.Exit(0)
		}
	} else {
		fileaddr = *fileJSON
	}
	res, err := ioutil.ReadFile(fileaddr)
	fmt.Println(fileaddr)
	if err != nil {
		fmt.Println("Problem with json file opening with initial data. \n", err.Error())
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
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))
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
			fmt.Println("Blin")
		}
		fmt.Println("Task: ", returnTask.Task, "\nResult: ", returnTask.Resp, "\nReason: \n", returnTask.Reason)
	}
	for _, results := range returnData {
		fmt.Println("\nTask: ", results.Task, "\nResult: ", results.Resp, "\nReason: ", results.Reason)
	}
}
