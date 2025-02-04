package xalan

import (
	"fmt"
	"os"
	"sync"

	go_xslt "github.com/wamuir/go-xslt"
)

func WorkerXML(file_path string, wg *sync.WaitGroup) {
	defer wg.Done()
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

	_, err = processor.Transform([]byte(file_data))
	if err != nil {
		panic(err)
	}

}

func Run() {
	files_dir := "workload/"
	files := []string{
		"acks.xml",
		"binding.xml",
		"changes.xml",
		"concepts.xml",
		"controls.xml",
		"datatypes.xml",
		"expr.xml",
		"index.xml",
		"intro.xml",
		"model.xml",
		"prod-notes.xml",
		"references.xml",
		"rpm.xml",
		"schema.xml",
		"structure.xml",
		"template.xml",
		"terms.xml",
		"acks.xml",
		"binding.xml",
		"changes.xml",
		"concepts.xml",
		"controls.xml",
		"datatypes.xml",
		"expr.xml",
		"index.xml",
		"intro.xml",
		"model.xml",
		"prod-notes.xml",
		"references.xml",
		"rpm.xml",
		"schema.xml",
		"structure.xml",
		"template.xml",
		"terms.xml",
		"acks.xml",
		"binding.xml",
		"changes.xml",
		"concepts.xml",
		"controls.xml",
		"datatypes.xml",
		"expr.xml",
		"index.xml",
		"intro.xml",
		"model.xml",
		"prod-notes.xml",
		"references.xml",
		"rpm.xml",
		"schema.xml",
		"structure.xml",
		"template.xml",
		"terms.xml",
		"acks.xml",
		"binding.xml",
		"changes.xml",
		"concepts.xml",
		"controls.xml",
		"datatypes.xml",
		"expr.xml",
		"index.xml",
		"intro.xml",
		"model.xml",
		"prod-notes.xml",
		"references.xml",
		"rpm.xml",
		"schema.xml",
		"structure.xml",
		"template.xml",
		"terms.xml",
		"acks.xml",
		"binding.xml",
		"changes.xml",
		"concepts.xml",
		"controls.xml",
		"datatypes.xml",
		"expr.xml",
		"index.xml",
		"intro.xml",
		"model.xml",
		"prod-notes.xml",
		"references.xml",
		"rpm.xml",
		"schema.xml",
		"structure.xml",
		"template.xml",
		"terms.xml",
	}

	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1)
		WorkerXML(files_dir+file, &wg)
	}

	wg.Wait()
}
