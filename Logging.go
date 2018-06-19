
package main

import (
	"os"
	"fmt"
	"log"
	"io/ioutil"
	"path"
	"runtime"
	"strings"
)

type BaseConfig struct {
	LogLevel, LogDir, WriteLog string
}

var (
	Debug *log.Logger
	Info *log.Logger
	Warn *log.Logger
	Error *log.Logger
)

func gettingConfig(finalConfig map[string]string, configs BaseConfig) map[string]string {

	// For Loglevel

	if configs.LogLevel != "" {
		if configs.LogLevel == "Info" || configs.LogLevel == "Debug" || configs.LogLevel == "Warn" || configs.LogLevel == "Error" {
			finalConfig["LogLevel"] = configs.LogLevel
			} else {
				panic(fmt.Sprintf("%v does not exists", configs.LogLevel))
		}
	}

	// For LogFile

	_, executionFilePath, _, ok := runtime.Caller(0)
	if ok != true {
		panic("YOOO")
	}
	dirName, executableFileName := path.Split(executionFilePath)
	arrayStrings := strings.Split(executableFileName, ".")
	logFilename := strings.Join(arrayStrings[:len(arrayStrings) -1], "")

	if configs.LogDir != "" {
		stat, err := os.Stat(configs.LogDir)
		if err != nil {
			panic("does not exist")
		}
		if !stat.IsDir() {
			panic("not a dir")
		}
		dirName = configs.LogDir
	}
	finalConfig["logFile"] = path.Join(dirName, logFilename + ".log")

	// For writing logs

	if configs.WriteLog == "True" {
		finalConfig["WriteLog"] = "True"
	} else if configs.WriteLog != "False" {
		panic("Invalid argument")
	}


	return finalConfig
}

func Init(configs ...BaseConfig) {
	if len(configs) > 1 { panic("Paramaters adsas") }

	order := [4]string{"Debug", "Info", "Warn", "Error"}
	logOrder := [4]*log.Logger{}

	finalConfig := make(map[string]string)
	finalConfig["LogLevel"] = "Info"
	finalConfig["WriteLog"] = "False"
	if len(configs) != 0 {	finalConfig = gettingConfig(finalConfig, configs[0])}

	logIoWriter := ioutil.Discard
	writeLogs := os.Stdout
	if finalConfig["WriteLog"] == "True" {
		writeLogs, _ = os.OpenFile(finalConfig["logFile"], os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
	}

	for  count, value := range order {
		if value == finalConfig["LogLevel"] {
			logIoWriter = writeLogs
		}
		logOrder[count] = log.New(logIoWriter, value + " ", log.Ldate|log.Ltime|log.Lshortfile)
	}

	Debug, Info, Warn, Error = logOrder[0], logOrder[1], logOrder[2], logOrder[3]
}

func main()  {
	Init(BaseConfig{WriteLog:"True", LogDir:"/home/ace"})
	//Init(BaseConfig{LogLevel:"Debug"})
	Debug.Println("YOOO")
	Info.Println("HELLO")
	Warn.Println("HELLO")
	Error.Println("HELLO")
}
