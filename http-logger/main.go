package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", hanlder)
	http.ListenAndServe(":8080", nil)
}

func hanlder(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		errorMessage := fmt.Sprintf("error reading request body %s", err.Error())
		logRequest(errorMessage)
		http.Error(w, errorMessage, http.StatusInternalServerError)
		return
	}

	logRequest(string(body))
}

func logRequest(str string) {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("Failed to determine executable path: %v", err)
	}
	exeDir := filepath.Dir(exePath)

	logFilePath := filepath.Join(exeDir, "app.log")
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	log.Println(str)
}
