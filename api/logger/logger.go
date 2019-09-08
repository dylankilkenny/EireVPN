package logger

import (
	"fmt"
	"strings"
	"time"

	. "github.com/logrusorgru/aurora"
)

var loggingEnabled bool

type Fields struct {
	Code  string
	Loc   string
	Err   string
	Extra map[string]interface{}
}

func Init(enabled bool) {
	loggingEnabled = enabled
}

func Log(fields Fields) {
	if loggingEnabled {
		var str strings.Builder
		str.WriteString(fmt.Sprintf("%s %v - ", Bold(BrightRed("ERROR")), time.Now().Format("2006-03-02 15:04:05")))
		for k, v := range fields.Extra {
			if k == "Err" {
				v = strings.Replace(v.(string), "\n", " ", -1)
			}
			field := fmt.Sprintf("%s: %v    ", Bold(BrightRed(k)), v)
			str.WriteString(field)
		}
		if fields.Code != "" {
			str.WriteString(fmt.Sprintf("%s: %v    ", Bold(BrightRed("Code")), fields.Code))
		}
		if fields.Loc != "" {
			str.WriteString(fmt.Sprintf("%s: %v    ", Bold(BrightRed("Loc")), fields.Loc))
		}
		if fields.Err != "" {
			str.WriteString(fmt.Sprintf("%s: %v    ", Bold(BrightRed("Err")), fields.Err))
		}
		fmt.Println(str.String())
	}
}
