package main

import (
	"encoding/json"
	"os"
)

type PackageJSON struct {
	Name         string            `json:"name"`
	Author       Author            `json:"author"`
	Contributors []Author          `json:"contributors"`
	Version      string            `json:"version"`
	Description  string            `json:"description"`
	Keywords     []string          `json:"keywords"`
	Homepage     string            `json:"homepage"`
	License      string            `json:"license"`
	Files        []string          `json:"files"`
	Main         string            `json:"main"`
	Scripts      map[string]string `json:"scripts"`
	Os           []string          `json:"os"`
	Cpu          []string          `json:"cpu"`
	Private      bool              `json:"private"`
	Dependencies map[string]string `json:"dependencies"`
}

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func PackageLoad() (*PackageJSON, error) {
	file, _ := os.ReadFile("package.json")

	var packagejson *PackageJSON

	err := json.Unmarshal([]byte(file), &packagejson)
	return packagejson, err
}
