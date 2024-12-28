package problem

import (
	"encoding/json"
	"fmt"
	"lets-go/libs/dockerController"
	localTypes "lets-go/types"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/docker/docker/api/types/container"
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

type RunCodeRequstData struct {
	programmingLanguage string `json:"programmingLanguage"`
	code                string `json:"code"`
}

func RunCode(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var dataObj RunCodeRequstData

	if err := json.NewDecoder(r.Body).Decode(&dataObj); err != nil {
		http.Error(w, "invalid Json format", http.StatusBadRequest)
		return
	}

	runningContainers, err := dockerController.GetContainersMap()

	if err != nil {
		log.Println("Error while trying to get the docker controller")
		http.Error(w, "server error", http.StatusInternalServerError)
	}

	programmingLanguage := localTypes.ProgrammingLanguage{
		Name: dataObj.programmingLanguage,
	}

	localContainer, err := runningContainers.Get(programmingLanguage)

	if err != nil {
		log.Println("Error while trying to get the docker container")
		http.Error(w, "server error", http.StatusInternalServerError)
	}

	formatedCode := fmt.Sprintf("\"%s\"", dataObj.code)

	execOptions := container.ExecOptions{
		Cmd: []string{"node", "-e", formatedCode},
	}

	execAttachOptions := container.ExecAttachOptions{
		Tty: true,
	}

	output, err := localContainer.ExecuteCommand(execOptions, execAttachOptions)

	if err != nil {
		log.Println("Error while trying to run code in the docker container")
		http.Error(w, "server error", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strings.Join(output, "\n")))
}
