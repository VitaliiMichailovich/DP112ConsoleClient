package filedata

import (

	"io/ioutil"
	"fmt"
)

func FileData(fileAddress string) (error) {
	res, err := ioutil.ReadFile(fileAddress)
	if err != nil {
		return fmt.Errorf("Problem with json file opening with initial data. \n%v", err)
	}

	return nil
}