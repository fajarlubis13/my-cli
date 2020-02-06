package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Page struct {
	Title string
}

func getWD() string {
	dir, _ := os.Getwd()
	return dir
}

func main() {
	targetPath := getWD() + "/result"
	sourcePath := "mold/src/golang"
	p := Page{Title: "Heading Test"}

	err := filepath.Walk(sourcePath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if strings.Contains(path, ".go") {
				_, fileName := filepath.Split(path)

				funcMap := template.FuncMap{
					"Truncate": func(values ...interface{}) string {
						s := values[0].(string)
						l := 5
						if len(values) > 1 {
							l = values[1].(int)
						}
						return fmt.Sprintf("%s ...", s[:l])
					},
				}

				// templates, err := template.ParseFiles(path)
				templates, err := template.New(fileName).Funcs(funcMap).ParseFiles(path)
				if err != nil {
					return err
				}

				__targetpath := strings.Replace(path, sourcePath, targetPath, -1)
				log.Println(__targetpath)

				if _, err := os.Stat(__targetpath); os.IsNotExist(err) {
					os.MkdirAll(strings.Replace(__targetpath, fileName, "", -1), os.ModePerm)
				}

				f, err := os.Create(__targetpath)
				if err != nil {
					return err
				}

				if err := templates.Execute(f, p); err != nil {
					return err
				}
				f.Close()
			}

			return nil
		})

	if err != nil {
		log.Println(err)
	}
}
