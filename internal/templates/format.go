package templates

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
	"text/template"

	"github.com/pkg/errors"
)

type Config struct {
	PackageName string
	Comment     string
	Imports     []string

	TypeName     string
	Funcs        template.FuncMap
	MethodsRegex string
}

func FormatPackageCode(config Config) ([]byte, error) {
	packageTemplate := template.New("")
	if config.Funcs != nil {
		packageTemplate = packageTemplate.Funcs(config.Funcs)
	}

	templateText := packageHeaderTemplate
	for _, methodName := range methodNames {
		if config.MethodsRegex != "" {
			matches, err := regexp.MatchString(config.MethodsRegex, methodName)
			if err != nil {
				return nil, errors.Wrap(err, config.MethodsRegex)
			}
			if !matches {
				continue
			}
		}
		templateText += methodsTemplates[methodName]
	}

	packageTemplate, err := packageTemplate.Parse(templateText)
	if err != nil {
		return nil, errors.Wrap(err, templateText)
	}

	rendered := bytes.NewBuffer(nil)
	err = packageTemplate.Execute(rendered, config)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	debugRendered := addLineNumbers(rendered.String())

	gofmt := exec.Command("gofmt", "-e")
	gofmt.Stdin = rendered

	formatted, err := gofmt.CombinedOutput()
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error: %s\nTemplate: %v", formatted, debugRendered))
	}
	return formatted, nil
}

func addLineNumbers(s string) string {
	lines := strings.Split(s, "\n")
	var out string
	for i, line := range lines {
		out += fmt.Sprintf("%v %v\n", i, line)
	}
	return out
}
