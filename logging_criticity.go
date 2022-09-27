package main

import (
	"log"
	"os"
)

var (
	ErrorLogger   *log.Logger
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
)

func initLoggers() {
	f, err := os.OpenFile(
		"logs.txt",
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0666,
	)
	if err != nil {
		log.Fatalln(err)
	}

	ErrorLogger = log.New(f, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(f, "[WARNING] ", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLogger = log.New(f, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	initLoggers()
	InfoLogger.Println("Message d'info")
	WarningLogger.Println("Attention ceci est un warning!")
	ErrorLogger.Println("Oh mon Dieu ! T'as tout pété !!!!")
}

// Bibliothèque de logging sympa:
// https://github.com/siruspen/logrus
