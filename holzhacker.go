package holzhacker

import (
	"flag"
	"io"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"gopkg.in/natefinch/lumberjack.v2"
)


var holzhackerStdout = flag.Bool("log", false, "Duplicate all log messages to Stdout")

// Create creates a logger with the specified filename
func Create(fileName string) *log.Logger {
	if !flag.Parsed() {
		flag.Parse()
	}

	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf("[FATAL] holzhacker can't get currentUser: %s\n", err)
	}

	logFolder := currentUser.HomeDir + "/logs/" + filepath.Base(os.Args[0]) + "/"

	var writer io.Writer = &lumberjack.Logger{
		Filename:   logFolder + fileName + ".log",
		MaxSize:    1, // mb
		MaxBackups: 10,
		MaxAge:     8, // days
	}
	if *holzhackerStdout {
		writer = io.MultiWriter(os.Stdout, writer)
	}

	logger := log.New(writer, "", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("***********************")

	return logger
}
