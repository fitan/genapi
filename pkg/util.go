package pkg

import (
	"fmt"
	"golang.org/x/mod/modfile"
	"golang.org/x/tools/imports"
	"io/ioutil"
	"log"
	"os"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

type GoMod struct {
	Module Module
}

type Module struct {
	Path    string
	Version string
}

func GetModuleName() string {
	goModBytes, err := ioutil.ReadFile("go.mod")
	if err != nil {
		log.Fatalln(err.Error())
	}

	modName := modfile.ModulePath(goModBytes)
	fmt.Fprintf(os.Stdout, "modName=%+v\n", modName)

	return modName
}

type (
	file struct {
		path    string
		content []byte
	}
	assets struct {
		dirs  []string
		files []file
	}
)

func (a assets) write() error {
	for _, dir := range a.dirs {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return fmt.Errorf("create dir %q: %w", dir, err)
		}
	}
	for _, file := range a.files {
		if err := ioutil.WriteFile(file.path, file.content, 0644); err != nil {
			return fmt.Errorf("write file %q: %w", file.path, err)
		}
	}
	return nil
}

// formatGo runs "goimports" on all assets.
func (a assets) formatGo() error {
	for _, file := range a.files {
		path := file.path
		src, err := imports.Process(path, file.content, nil)
		if err != nil {
			return fmt.Errorf("formatGo file %s: %v", path, err)
		}
		if err := ioutil.WriteFile(path, src, 0644); err != nil {
			return fmt.Errorf("write file %s: %v", path, err)
		}
	}
	return nil
}
