package templates

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/pkg/errors"
)

const (
	equalMixin = "equal"
	lessMixin  = "less"
	titleMixin = "title"
)

func UseEqualFormat(config *Config, format string) error {
	if config.Funcs == nil {
		config.Funcs = make(template.FuncMap)
	}
	// check format is correct somehow?

	config.Funcs[equalMixin] = func(l, r string) string {
		return fmt.Sprintf(format, l, r)
	}
	return nil
}

func UseDeepEqual(config *Config) error {
	if err := UseEqualFormat(config, "reflect.DeepEqual(%v, %v)"); err != nil {
		return errors.Wrap(err, "")
	}
	config.Imports = append(config.Imports, "reflect")
	return nil
}

func UseEqualMethod(config *Config) error {
	if err := UseEqualFormat(config, "%v.Equal(%v)"); err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}

func UseEqualOperator(config *Config) error {
	// needs parenteses so "!equal" will work as well
	if err := UseEqualFormat(config, "(%v == %v)"); err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}

func UseLessFormat(config *Config, format string) error {
	if config.Funcs == nil {
		config.Funcs = make(template.FuncMap)
	}
	// check format is correct somehow?

	config.Funcs[titleMixin] = func(typeName string) string {
		return strings.Title(typeName[strings.LastIndex(typeName, ".")+1:])
	}
	config.Funcs[lessMixin] = func(l, r string) string {
		return fmt.Sprintf(format, l, r)
	}
	config.Sortable = true
	return nil
}

func UseLessOperator(config *Config) error {
	// will need parenteses for "!less" usage
	if err := UseLessFormat(config, "%v < %v"); err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}
