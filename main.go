package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

// DataSource ...
type DataSource struct {
	ProjectName   string
	Entities      []*Entity
	TableStatus   bool
	Function      []string
	TableCreation bool
}

// Entity ...
type Entity struct {
	Name    string
	Type    string
	Binding bool
}

func main() {
	entities := []*Entity{
		&Entity{
			Name:    "ID Pengiriman",
			Type:    "int64",
			Binding: true,
		},
		&Entity{
			Name: "ID Surat",
			Type: "string",
		},
	}

	p := DataSource{
		ProjectName:   "HK Pengiriman",
		Entities:      entities,
		TableStatus:   true,
		Function:      []string{"create() => direct-db", "updated"},
		TableCreation: false,
	}

	workDir, _ := os.Getwd()
	targetPath := fmt.Sprintf("%s/%s", workDir, strcase.ToDelimited(p.ProjectName, '-'))
	sourcePath := "skeleton/src/golang"
	allowedExt := []string{".go", ".mod", ".env"}

	err := filepath.Walk(sourcePath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			for _, v := range allowedExt {
				if strings.Contains(path, v) {
					_, fileName := filepath.Split(path)

					funcMap := template.FuncMap{
						"truncate": func(values ...interface{}) string {
							s := values[0].(string)

							l := 0
							if len(values) > 1 {
								l = values[1].(int)
							}
							return fmt.Sprintf("%s", s[:l])
						},
						"toSnake": func(values ...interface{}) string {
							return fmt.Sprintf("%s", strcase.ToSnake(values[0].(string)))
						},
						"toScreamingSnake": func(values ...interface{}) string {
							return fmt.Sprintf("%s", strcase.ToScreamingSnake(values[0].(string)))
						},
						"toKebab": func(values ...interface{}) string {
							return fmt.Sprintf("%s", strcase.ToKebab(values[0].(string)))
						},
						"toScreamingKebab": func(values ...interface{}) string {
							return fmt.Sprintf("%s", strcase.ToScreamingKebab(values[0].(string)))
						},
						"toDelimeted": func(values ...interface{}) string {
							s := values[0].(string)

							l := uint8(45)
							if len(values) > 1 {
								l = uint8(values[1].(int))
							}
							return fmt.Sprintf("%s", strcase.ToDelimited(s, l))
						},
						"toScreamingDelimeted": func(values ...interface{}) string {
							return fmt.Sprintf("%s", strcase.ToScreamingDelimited(values[0].(string), '-', '-', false))
						},
						"toCamel": func(values ...interface{}) string {
							return fmt.Sprintf("%s", strcase.ToCamel(values[0].(string)))
						},
						"toLowerCamel": func(values ...interface{}) string {
							return fmt.Sprintf("%s", strcase.ToLowerCamel(values[0].(string)))
						},
					}

					// templates, err := template.ParseFiles(path)
					templates, err := template.New(fileName).Funcs(funcMap).ParseFiles(path)
					if err != nil {
						return err
					}

					_targetPath := strings.Replace(path, sourcePath, targetPath, -1)

					if _, err := os.Stat(_targetPath); os.IsNotExist(err) {
						os.MkdirAll(strings.Replace(_targetPath, fileName, "", -1), os.ModePerm)
					}

					f, err := os.Create(_targetPath)
					if err != nil {
						return err
					}

					if err := templates.Execute(f, p); err != nil {
						return err
					}
					f.Close()
				}
			}

			return nil
		})

	if err != nil {
		log.Println(err)
	}
}
