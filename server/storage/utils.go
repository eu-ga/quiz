package storage

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"sort"
)

func InsertSort(data []UserStatistics, el UserStatistics) ([]UserStatistics, int) {
	index := sort.Search(len(data), func(i int) bool { return data[i].SuccessRate > el.SuccessRate })
	data = append(data, UserStatistics{})
	copy(data[index+1:], data[index:])
	data[index] = el
	return data, index
}

func readConf(fileName string, i interface{}) error {
	file, err := os.OpenFile(fileName, os.O_RDWR, 0)
	defer file.Close()
	if err != nil {
		return errors.New("readConf error: error while reading conf.yaml file, " + err.Error())
	} else {
		decoder := json.NewDecoder(file)
		err := decoder.Decode(i)
		if err != nil {
			log.Println("Error: cannot decode conf.yaml file to json, ", err)
			//			PrintToLog("Error: cannot decode conf.yaml file to json, ", err.Error())
			return err
		}
	}
	return nil
}
