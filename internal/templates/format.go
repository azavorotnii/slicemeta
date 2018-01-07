package templates

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"text/template"

	"github.com/pkg/errors"
	"regexp"
)

type Config struct {
	PackageName string
	Comment     string
	Imports     []string

	TypeName string
	Funcs    template.FuncMap
	MethodsRegex string
}

func FormatPackageCode(config Config) ([]byte, error) {
	packageTemplate := template.New("")
	if config.Funcs != nil {
		packageTemplate = packageTemplate.Funcs(config.Funcs)
	}

	templateText := packageHeaderTemplate
	for methodName, methodTemplate := range PackageMethodsTemplates {
		if config.MethodsRegex != "" {
			matches, err := regexp.MatchString(config.MethodsRegex, methodName)
			if err != nil {
				return nil, errors.Wrap(err, config.MethodsRegex)
			}
			if !matches {
				continue
			}
		}
		templateText += methodTemplate
	}

	packageTemplate, err := packageTemplate.Parse(templateText)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	formatter := exec.Command("gofmt", "-e")

	out, in := io.Pipe()
	formatter.Stdin = out

	buffer := bytes.NewBuffer(nil)
	formatter.Stdout = buffer
	formatter.Stderr = os.Stderr

	if err := formatter.Start(); err != nil {
		return nil, errors.Wrap(err, "")
	}

	err = packageTemplate.Execute(in, config)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	if err := in.Close(); err != nil {
		return nil, errors.Wrap(err, "")
	}
	if err := formatter.Wait(); err != nil {
		return nil, errors.Wrap(err, "")
	}
	return buffer.Bytes(), nil
}
