package holzhacker

import (
	"flag"
	"io"
	"log"
	"os"
	"os/user"

	"lumberjack"
)

var console = flag.Bool("log", false, "Log auch auf Stdout ausgeben")

func Create(fileName string) *log.Logger {
	if !flag.Parsed() {
		flag.Parse()
	}
	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf("[FATAL] holzhacker can't get currentUser: %s\n", err)
	}
	logFolder := currentUser.HomeDir + "/logs/" + os.Args[0] + "/"
	var writer io.Writer = &lumberjack.Logger{
		Filename:   logFolder + fileName + ".log",
		MaxSize:    1, // mb
		MaxBackups: 10,
		MaxAge:     8, // days
	}

	if *console {
		writer = io.MultiWriter(os.Stdout, writer)
	}
	logger := log.New(writer, "", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("***********************")
	return logger
}
