package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
	"text/template"
)

type Project struct {
	Name     string
	Registry string
}

func main() {

	projectName := flag.String("name", "", "Project name")
	registryName := flag.String("registry", "fakeregistry", "Registry name")
	compose := flag.Bool("compose", false, "Generate docker-compose.yml")
	output := flag.String("output", ".", "Output directory")
	flag.Parse()

	if *projectName == "" {
		log.Fatal("project name is required")
	}

	var project Project
	project.Name = *projectName
	project.Registry = *registryName

	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	templatesDir := os.DirFS(home + "/jjtemplate")
	err = walkProject(templatesDir, "templates/core", project, output)

	if err != nil {
		log.Fatal(err)
	}

	if *compose {
		templatesDir = os.DirFS(home + "/jjtemplate")
		err = walkProject(templatesDir, "templates/compose", project, output)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func walkProject(templatesDir fs.FS, templateDir string, project Project, output *string) error {
	err := fs.WalkDir(templatesDir, templateDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("failed walking directory %v", err)
		}
		if path == templateDir {
			return nil
		}
		target := strings.TrimPrefix(path, templateDir+"/")
		target = strings.TrimSuffix(target, ".template")

		template1 := template.Must(template.New("dir").Parse(target))

		var buf bytes.Buffer
		err = template1.Execute(&buf, project)
		if err != nil {
			return fmt.Errorf("failed executing template: %v", err)
		}

		outputPath := *output + "/" + buf.String()

		if d.IsDir() {
			log.Println("Creating directory", outputPath)

			err := os.MkdirAll(outputPath, 0777)
			if err != nil {
				return fmt.Errorf("failed creating directory: %v", err)
			}
		} else {
			log.Println("Creating file", outputPath)
			data, err := fs.ReadFile(templatesDir, path)
			if err != nil {
				return fmt.Errorf("failed reading data from file: %v", err)
			}

			template2 := template.Must(template.New("file").Parse(string(data)))
			var result bytes.Buffer
			err = template2.Execute(&result, project)
			if err != nil {
				return fmt.Errorf("failed executing template: %v", err)
			}
			err = os.WriteFile(outputPath, result.Bytes(), 0644)
			if err != nil {
				return fmt.Errorf("failed writing data to file: %v", err)
			}
		}

		return nil
	})
	return err
}
