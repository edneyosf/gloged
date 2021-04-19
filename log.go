package gloged

//
// by Edney O. S. Filho
//

import (
	"fmt"
	"os"
	"time"
)

var DebugMode = false
var WriteFile = true
var Path = "gloged"

const filePermission = 0755
const dirPermission = 0755
const resetColor = string("\033[0m")
const greenColor = string("\033[32m")
const redColor = string("\033[31m")
const yellowColor = string("\033[33m")
const blueColor = string("\033[34m")

func I(text string){
	time := getTime()
	line := time + " ★ " + text

	writeLog(time, line)
	fmt.Println(blueColor, line, resetColor)
}

func S(text string) {
	time := getTime()
	line := time + " • " + text

	writeLog(time, line)
	fmt.Println(greenColor, line, resetColor)
}

func D(method string, key string, value string) {
	if DebugMode {
		time := getTime()
		line := time + " ▸ (" + method + ") " + key + " -> |" + value + "|"

		writeLog(time, line)
		fmt.Println(resetColor, line, resetColor)
	}
}

func W(text string) {
	time := getTime()
	line := time + " ! " + text

	writeLog(time, line)
	fmt.Println(yellowColor, line, resetColor)
}

func E(text string) {
	time := getTime()
	line := time + " × " + text

	writeLog(time, line)
	fmt.Println(redColor, line, resetColor)
}

func getFileName() string {
	date := time.Now().Format("02-01-2006")
	return "log-" + date + ".ed"
}

func getTime() string {
	return time.Now().Format("2 Jan 2006 15:04:05")
}

func writeLog(time string, text string) {
	if WriteFile {
		line := time + " " + text + "\n"
		fileName := getFileName()

		if _, err := os.Stat(Path); os.IsNotExist(err) {
			os.Mkdir(Path, 0755)
		}

		f, err := os.OpenFile(Path+"/"+fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, filePermission)

		if err != nil {
			panic(err)
		}

		defer f.Close()

		if _, err = f.WriteString(line); err != nil {
			panic(err)
		}
	}
}
