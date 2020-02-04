package logger

import (
	"fmt"
	"os"
	"strings"
	"time"
	// . "github.com/logrusorgru/aurora"
)

var loggingEnabled bool

const LogFilePath = "./logs/default.log"

type Fields struct {
	Code  string
	Loc   string
	Err   string
	Extra map[string]interface{}
}

func Init(enabled bool) {
	loggingEnabled = enabled
}

func createLogFile() {
	// detect if file exists
	var _, err = os.Stat(LogFilePath)
	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(LogFilePath)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer file.Close()
	}
	fmt.Println("==> created log file", LogFilePath)
}

func writeFile(msg string) {
	// open file using READ & WRITE permission
	var file, err = os.OpenFile(LogFilePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}
	// defer file.Close()

	// write to file
	_, err = file.WriteString(msg + "\n")
	if err != nil {
		fmt.Println(err.Error())
	}
	// save changes
	err = file.Sync()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func Log(fields Fields) {
	if loggingEnabled {
		var msg strings.Builder
		msg.WriteString(fmt.Sprintf("%s %v | ", "ERROR", time.Now().Format("2006-01-02 15:04:05")))
		if fields.Err != "" {
			msg.WriteString(fmt.Sprintf("%s: %v | ", "ERROR", fields.Err))
		}
		if fields.Code != "" {
			msg.WriteString(fmt.Sprintf("%s: %v | ", "CODE", fields.Code))
		}
		for k, v := range fields.Extra {
			if k == "Err" {
				v = strings.Replace(v.(string), "\n", " ", -1)
			}
			field := fmt.Sprintf("%s: %v | ", k, v)
			msg.WriteString(field)
		}
		if fields.Loc != "" {
			msg.WriteString(fmt.Sprintf("%s: %v", "LOC", fields.Loc))
		}
		fmt.Println(msg.String())
		writeFile(msg.String())
	}
}
