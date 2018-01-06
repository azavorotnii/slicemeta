package templates

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"text/template"

	"github.com/pkg/errors"
)

type Config struct {
	Comment        string
	TypeName       string
	ImportTypeName string
	UseDeepEqual   bool
	UseEqual       bool
	PackageName    string
}

func FormatPackageCode(templateText string, config Config) ([]byte, error) {
	packageTemplate, err := template.New("").Parse(templateText)
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

