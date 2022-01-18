package application

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"key-value-store/infrastructure"
	"key-value-store/interfaces"
)

func LaunchApp() {
	infrastructure.CreateFile()
	TransferData("mainData.json", "tmp/TIMESTAMP-data.json")
	go SaveFileDuration()
	fmt.Println("Listening 8080")
	interfaces.HandlerMain()
}
func SaveFileDuration() {
	for range time.Tick(time.Minute * 1) {

		TransferData("tmp/TIMESTAMP-data.json", "mainData.json")

	}
}
func TransferData(src string, dest string) {

	bytesRead, err := ioutil.ReadFile(src)

	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(dest, bytesRead, 0644)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Transfer Successful.")
	}
}
