package utils

import (
	"fmt"
	"log"
	"os"
	"portScanner/models"
	"strconv"
)

func WriteTxt(result models.ScanResult, fileName string) {
	var file *os.File

	//open file
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	//write to file
	file.WriteString(strconv.Itoa(result.Port) + "           " + result.State + "\n")
}

func CreateTxt(hostname string) string {
	//if file exists, create new .txt
	fileName := CreateFileName()
	fmt.Println("filename",fileName)
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		file, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}
		openFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			log.Fatal(err)
		}
		openFile.WriteString("     " + hostname + "     \n")
		openFile.WriteString("PORT <------> STATUS\n")
		defer openFile.Close()
		defer file.Close()
	}
	return fileName
}
