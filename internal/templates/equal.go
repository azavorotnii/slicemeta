package templates

import (
	"fmt"
	"text/template"

	"github.com/pkg/errors"
)

const equalFunc = "equal"

func UseEqualFormat(config *Config, format string) error {
	if config.Funcs == nil {
		config.Funcs = make(template.FuncMap)
	}
	// check format is correct somehow?

	config.Funcs[equalFunc] = func(l, r string) string {
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
	// needs parenteses so not_equal will work as well
	if err := UseEqualFormat(config, "(%v == %v)"); err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}
