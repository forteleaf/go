package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

type TempFiles struct {
	Files map[string]string
}

// Loop through view files and load them
func LoadTempFiles() {
	t := new(TempFiles)

	// Load template files
	filepath.Walk("application/views", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			content, _ := ioutil.ReadFile(path)
			t.Files[path] = string(content)
		}
		return nil
	})
}

func ViewTemp(w http.ResponseWriter, path string) {
	t := new(TempFiles)

	temp, err := template.New().Parse(t.Files[path])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		temp.Execute(w, nil)
	}
}

func main() {

}
