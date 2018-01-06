package main

// avoid recursion //go:generate slicemeta -type string -outputDir ./internal

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"github.com/azavorotnii/slicemeta/internal/templates"
	"github.com/azavorotnii/slicemeta/internal/stringutil"
)

const version = "0.0.1a"

func main() {
	var (
		typeName       string
		importTypeName string
		equalityOp     string
		packageName    string
		outputDir      string
	)

	flag.StringVar(&typeName, "type", "", "")
	flag.StringVar(&importTypeName, "import", "", "")
	flag.StringVar(&equalityOp, "equalityOp", "operator", "can be 'operator' (default), 'equal' or 'deepequal'.")
	flag.StringVar(&packageName, "package", "", "")
	flag.StringVar(&outputDir, "outputDir", "", "")
	flag.Parse()

	if typeName == "" {
		flag.Usage()
		log.Fatal("type is mandatory argument")
	}
	if packageName == "" {
		packageName = strings.ToLower(typeName) + "util"
	}
	if importTypeName != "" {
		typeName = path.Base(importTypeName) + "." + typeName
	}

	if !stringutil.Contains([]string{"equal", "deepequal", "operator"}, equalityOp) {
		log.Fatalf("unknown equality option: %v", equalityOp)
	}

	if outputDir == "" {
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		outputDir = cwd
	}
	outputDir = path.Join(outputDir, packageName)
	if err := os.MkdirAll(outputDir, 0777); err != nil {
		log.Fatal(err)
	}

	config := templates.Config{
		Comment:        fmt.Sprintf("Generated by %v-%v (%v)", path.Base(os.Args[0]), version, time.Now().Format(time.RFC3339)),
		TypeName:       typeName,
		ImportTypeName: importTypeName,
		UseDeepEqual:   equalityOp == "deepequal",
		UseEqual:       equalityOp == "equal",
		PackageName:    packageName,
	}

	for filename, templateText := range templates.PkgTemplates {
		goCode, err := templates.FormatPackageCode(templateText, config)
		if err != nil {
			log.Printf("%v: %+v\n", filename, err)
			continue
		}
		if err := ioutil.WriteFile(path.Join(outputDir, filename), goCode, 0755); err != nil {
			log.Printf("%v: %v\n", filename, err)
		}
	}
}
