package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"github.com/VitaliiMichailovich/DP112ConsoleClient/filedata"
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
	err = filedata.FileData(fileaddr)
	fmt.Println(taskId,err)
}
