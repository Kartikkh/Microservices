package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func ReadJsonFile(dirName, fileName string, target interface{}) error {
	fileRead, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", dirName, fileName))
	if err != nil {
		return err
	}
	err = json.Unmarshal(fileRead, target)
	if err != nil {
		return err
	}
	return nil
}
