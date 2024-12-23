package problem

import (
	"fmt"
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
