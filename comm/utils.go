package comm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func LoadJson(jsonPath string) []map[string]string {
	file, err := os.Open(jsonPath)
	if err != nil {
		Logger.Error().Any("err", err)
		return nil
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		Logger.Error().Msg(err.Error())
		return nil
	}
	var result []map[string]string
	err = json.Unmarshal(data, &result)
	if err != nil {
		Logger.Error().Msg(err.Error())
		return nil
	}

	return result
}

func WriteJson(data []map[string]string, filepath string) {
	// 将data转换为json字符串
	//jsonBytes, err := json.Marshal(data)
	jsonBytes, err := json.MarshalIndent(data, "", "    ")

	if err != nil {
		fmt.Println("json.Marshal failed:", err)
		return
	}
	err = os.WriteFile(filepath, jsonBytes, 0666)

	if err != nil {
		fmt.Println("WriteFile failed:", err)
		return
	}
}
