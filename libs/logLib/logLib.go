package loglib

import (
	"fmt"
	"log"
	"os"
)

func LogError(text string, err error) {
	formatedString := fmt.Sprintf("%s : %v \n", text, err)
	log.Println(formatedString)
	os.WriteFile("logs.txt", []byte(formatedString), 0666)
}
