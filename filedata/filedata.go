package filedata

import (

	"io/ioutil"
	"fmt"
)

func FileData(fileAddress string) ([]byte, error) {
	res, err := ioutil.ReadFile(fileAddress)
	if err != nil {
		return []byte{}, fmt.Errorf("Problem with json file opening with initial data. \n%v", err)
	}
	return res, nil
}

func ChooseTask([]byte, int) ([]byte, error) {
	
}