package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/codeclysm/extract"
	"github.com/elsaland/elsa/util"
)

func ExtractTarGz(buffer *bytes.Buffer, name string) {
	var shift = func(path string) string {
		parts := strings.Split(path, string(filepath.Separator))
		parts = parts[0:]
		parts[0] = name
		return strings.Join(parts, string(filepath.Separator))
	}
	extract.Gz(context.TODO(), buffer, "node_modules", shift)
}

func fetchPackage(name, reference string) []byte {
	registry := "https://registry.yarnpkg.com"
	r, err := http.Get(fmt.Sprintf("%v/%v/-/%v-%v.tgz", registry, name, name, reference))
	util.Check(err)
	resp, err := io.ReadAll(r.Body)
	util.Check(err)
	return resp
}

func install() {
	json, _ := PackageLoad()
	for dep, version := range json.Dependencies {
		fmt.Println("Package:", dep, "=>", "Version:", version)
		resp := fetchPackage(dep, version)
		buffer := bytes.NewBuffer(resp)
		ExtractTarGz(buffer, dep)
	}
}
