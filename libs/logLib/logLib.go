package loglib

import (
	"fmt"
	"log"
	"os"
	"time"
)

func LogError(text string, err error) {
	formatedString := fmt.Sprintf("%s : %v \n", text, err)
	log.Println(formatedString)
	formatedString = fmt.Sprintf("%s:%s: %s", time.DateOnly, time.TimeOnly, formatedString)
	os.WriteFile("logs.txt", []byte(formatedString), 0666)
}
