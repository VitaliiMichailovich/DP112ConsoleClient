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
)

func main() {
	var fileaddr string
	var err error
	taskId := flag.Int("task", 0, "an int")
	fileJSON := flag.String("file", "", "a string")
	flag.Parse()
	if *fileJSON == "" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Please, write an address to json file with initial data.")
		fileaddr, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
	} else {
		fileaddr = *fileJSON
	}

	res, err := ioutil.ReadFile(fileaddr)
	if err != nil {
		return []byte{}, fmt.Errorf("Problem with json file opening with initial data. \n%v", err)
	}

	if err != nil {
		fmt.Println("Something went wrong.\n", err.Error())
		os.Exit(0)
	}

	url := "http://localhost:8080"
	var resource string
	if *taskId == 0 {
		resource = url + "/tasks"
	} else {

		resource = url + "/task/" + strconv.Itoa(*taskId)
	}
	fmt.Println(resource, string(res))
	req, err := http.NewRequest("POST", resource, bytes.NewBuffer(res))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

}
