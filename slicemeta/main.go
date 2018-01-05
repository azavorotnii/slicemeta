package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"strings"
	"text/template"
	"time"
)

func main() {
	var (
		typeName       string
		importTypeName string
		useDeepEqual   bool
		packageName    string
		outputDir      string
	)

	flag.StringVar(&typeName, "type", "", "")
	flag.StringVar(&importTypeName, "import", "", "")
	flag.BoolVar(&useDeepEqual, "deepequal", false, "")
	flag.StringVar(&packageName, "package", "", "")
	flag.StringVar(&outputDir, "outputDir", "", "")
	flag.Parse()

	if typeName == "" {
		flag.Usage()
		fmt.Println("'type' is mandatory argument.")
		os.Exit(1)
	}
	if packageName == "" {
		packageName = strings.ToLower(typeName) + "util"
	}
	if importTypeName != "" {
		typeName = path.Base(importTypeName) + "." + typeName
	}

	containsPackageTemplate, err := template.New("contains").Parse(containsTemplate)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	formatter := exec.Command("gofmt", "-e")
	if outputDir == "" {
		formatter.Stdout = os.Stdout
	} else {
		dir := path.Join(outputDir, packageName)
		if err := os.MkdirAll(dir, 0777); err != nil {
			fmt.Println(err)
		}
		f, err := os.Create(path.Join(dir, packageName+".go"))
		if err != nil {
			fmt.Println(err)
			os.Exit(3)
		}
		defer f.Close()
		formatter.Stdout = f
	}
	formatter.Stderr = os.Stderr

	out, in := io.Pipe()
	formatter.Stdin = out

	if err := formatter.Start(); err != nil {
		fmt.Println(err)
		os.Exit(4)
	}

	config := struct {
		TypeName       string
		ImportTypeName string
		UseDeepEqual   bool
		PackageName    string
		Now            string
	}{
		TypeName:       typeName,
		ImportTypeName: importTypeName,
		UseDeepEqual:   useDeepEqual,
		PackageName:    packageName,
		Now:            time.Now().Format(time.RFC3339),
	}
	err = containsPackageTemplate.Execute(in, config)
	if err != nil {
		fmt.Println(err)
		os.Exit(5)
	}
	if err := in.Close(); err != nil {
		fmt.Println(err)
	}
	if err := formatter.Wait(); err != nil {
		fmt.Println(err)
	}
}
