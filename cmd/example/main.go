package main

import (
	"flag"
	"io"
	"lab2"
	"os"
	"strings"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	// TODO: Add other flags support for input and output configuration.
)

func main() {

	fStringPtr := flag.String("f", "", "enter the filename with expression")
	oStringPtr := flag.String("o", "", "enter the name of output file")
	flag.Parse()

	fString := *fStringPtr
	oString := *oStringPtr
	eString := flag.Lookup("e").Value.String()

	// fmt.Println(eString)
	// fmt.Println(fString)

	if (eString != "" && fString != "") || (eString == "" && fString == "") {
		os.Stderr.WriteString("Wrong arguments")
		return
	}

	// var handler lab2.ComputeHandler
	var reader io.Reader
	var writer io.Writer
	var fileDesc *os.File
	if eString != "" {
		reader = strings.NewReader(eString)
	} else {
		file, err := os.Open(fString)
		if err != nil {
			os.Stderr.WriteString("File does not exist")
			return
		}
		reader = file
	}
	if oString == "" {
		writer = os.Stdout
	} else {
		file, err := os.OpenFile(oString, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			os.Stderr.WriteString("Cannot create file")
			return
		}
		writer = file
		fileDesc = file
	}

	handler := &lab2.ComputeHandler{reader, writer}
	err := handler.Compute()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	if fileDesc != nil {
		defer fileDesc.Close()
	}
}
