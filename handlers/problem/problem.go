package problem

import (
	"fmt"
	dockerclient "lets-go/libs/dockerClient"
	"log"
	"net/http"
	"os"
)

func GetProblemCode(w http.ResponseWriter, r *http.Request) {
	problemID := r.PathValue("problemID")
	programmingLanguage := r.PathValue("programmingLanguage")

	if problemID == "" || programmingLanguage == "" {
		http.Error(w, "Missing problemID or programmingLanguage", http.StatusBadRequest)
		return
	}

	filePath := fmt.Sprintf("./code-samples/%s/%s.txt", programmingLanguage, problemID)
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			http.Error(w, "Sample code not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Server error while reading the sample code", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write(fileContent)
}

func RunCode(w http.ResponseWriter, r *http.Request) {
	programmingLanguage := r.PathValue("programmingLanguage")

	if len(programmingLanguage) < 0 {
		log.Println("Error while trying to get the programming language")
		http.Error(w, "server error", http.StatusInternalServerError)
	}

	// runningContainer, err := docker
}
