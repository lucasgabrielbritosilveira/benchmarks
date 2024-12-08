package xalan

import (
	"fmt"
	"os"

	go_xslt "github.com/wamuir/go-xslt"
)

func WorkerXML(file_path string) {

	fmt.Println("Working in", file_path)

	xslPath := "template/xmlspec.xsl"

	xslData, err := os.ReadFile(xslPath)
	if err != nil {
		panic(err)
	}

	file_data, err := os.ReadFile(file_path)

	if err != nil {
		panic(err)
	}

	processor, err := go_xslt.NewStylesheet(xslData)
	if err != nil {
		panic(err)
	}

	result, err := processor.Transform([]byte(file_data))
	if err != nil {
		panic(err)
	}

	fmt.Println(string(result))

}

func Run() {
	files_dir := "workload/"
	files := []string{
		"template.xml",
		"index.xml",
		"prod-notes.xml",
		"schema.xml",
		"terms.xml",
		"references.xml",
	}

	for _, file := range files {
		WorkerXML(files_dir + file)
	}

}
