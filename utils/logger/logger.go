package logUtils

import "log"

func Info(msg string) {
	log.Println(msg)
}

func Error(err error, msg ...string) {
	if len(msg) == 0 {
		msg[0] =  "Something went wrong"
	}

	log.Fatal(msg[0], err)
}