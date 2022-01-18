package infrastructure

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"key-value-store/domain"
)

func WriteData(data domain.Memory) {
	store := GetData()

	store[data.Key] = data.Value
	fmt.Println(store)
	writeFile, _ := os.Create("tmp/TIMESTAMP-data.json")
	json.NewEncoder(writeFile).Encode(store)
}

func ReadData(key string) string {
	readFile, err := os.Open("tmp/TIMESTAMP-data.json")
	if err != nil {
		log.Fatal(err)
		return "err"
	} else {
		var store map[string]string
		json.NewDecoder(readFile).Decode(&store)

		fmt.Println(store[key])
		return store[key]
	}

}
func DeleteAll() {
	blank := GetData()

	for k := range blank {
		delete(blank, k)
	}
	writeFile, _ := os.Create("tmp/TIMESTAMP-data.json")
	json.NewEncoder(writeFile).Encode(blank)

}
func DeleteKey(key string) {
	blank := GetData()

	delete(blank, key)

	writeFile, _ := os.Create("tmp/TIMESTAMP-data.json")
	json.NewEncoder(writeFile).Encode(blank)

}

func CreateFile() {
	firstData := make(map[string]string)
	writeFile, _ := os.Create("tmp/TIMESTAMP-data.json")
	json.NewEncoder(writeFile).Encode(firstData)
}

func GetData() map[string]string {
	readFile, _ := os.Open("tmp/TIMESTAMP-data.json")
	var data map[string]string
	json.NewDecoder(readFile).Decode(&data)

	return data
}

func GetAllData() string {
	datas := GetData()
	jsonString, err := json.Marshal(datas)
	if err != nil {
		return "err"
	}

	return string(jsonString)
}

func CopyFile() {
	GetAllData()
}
