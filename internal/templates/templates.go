package templates

import "sort"

var methodsTemplates = map[string]string{
	containsMethodName:      containsMethodTemplate,
	containsAnyMethodName:   containsAnyMethodTemplate,
	containsFuncMethodName:  containsFuncMethodTemplate,
	countMethodName:         countMethodTemplate,
	countAnyMethodName:      countAnyMethodTemplate,
	countFuncMethodName:     countFuncMethodTemplate,
	equalMethodName:         equalMethodTemplate,
	filterMethodName:        filterMethodTemplate,
	indexMethodName:         indexMethodTemplate,
	indexAnyMethodName:      indexAnyMethodTemplate,
	indexFuncMethodName:     indexFuncMethodTemplate,
	lastIndexMethodName:     lastIndexMethodTemplate,
	lastIndexAnyMethodName:  lastIndexAnyMethodTemplate,
	lastIndexFuncMethodName: lastIndexFuncMethodTemplate,
	mapMethodName:           mapMethodTemplate,
	reduceMethodName:        reduceMethodTemplate,
}

var methodNames []string

func init() {
	for key := range methodsTemplates {
		methodNames = append(methodNames, key)
	}
	sort.Strings(methodNames)
}

const (
	packageHeaderTemplate = `
		package {{.PackageName}}

		// {{.Comment}}

		{{if .Imports}}
		import (
			{{- range .Imports}}
			"{{.}}"
			{{- end}}
		)
		{{end}}
	`

	containsMethodName     = "Contains"
	containsMethodTemplate = `
		func ` + containsMethodName + `(in []{{.TypeName}}, value {{.TypeName}}) bool {
			for _, v := range in {
				if {{equal "v" "value"}} {
					return true
				}
			}
			return false
		}
	`

	containsAnyMethodName     = "ContainsAny"
	containsAnyMethodTemplate = `
		func ` + containsAnyMethodName + `(in []{{.TypeName}}, values ...{{.TypeName}}) bool {
			for _, v := range in {
				for _, value := range values {
					if {{equal "v" "value"}} {
						return true
					}
				}
			}
			return false
		}
	`

	containsFuncMethodName     = "ContainsFunc"
	containsFuncMethodTemplate = `
		func ` + containsFuncMethodName + `(in []{{.TypeName}}, f func({{.TypeName}}) bool) bool {
			for _, v := range in {
				if f(v) {
					return true
				}
			}
			return false
		}
	`

	countMethodName     = "Count"
	countMethodTemplate = `
		func ` + countMethodName + `(in []{{.TypeName}}, value {{.TypeName}}) int {
			result := 0
			for _, v := range in {
				if {{equal "v" "value"}} {
					result++
				}
			}
			return result
		}
	`

	countAnyMethodName     = "CountAny"
	countAnyMethodTemplate = `
		func ` + countAnyMethodName + `(in []{{.TypeName}}, values ...{{.TypeName}}) int {
			result := 0
			for _, v := range in {
				for _, value := range values {
					if {{equal "v" "value"}} {
						result++
						break
					}
				}
			}
			return result
		}
	`

	countFuncMethodName     = "CountFunc"
	countFuncMethodTemplate = `
		func ` + countFuncMethodName + `(in []{{.TypeName}}, f func({{.TypeName}}) bool) int {
			result := 0
			for _, v := range in {
				if f(v) {
					result++
				}
			}
			return result
		}
	`

	equalMethodName     = "Equal"
	equalMethodTemplate = `
		func ` + equalMethodName + `(a, b []{{.TypeName}}) bool {
			if len(a) != len(b) {
				return false
			}
			for i := 0; i < len(a); i++ {
				if !{{equal "a[i]" "b[i]"}} {
					return false
				}
			}
			return true
		}
	`

	filterMethodName     = "Filter"
	filterMethodTemplate = `
		func ` + filterMethodName + `(in []{{.TypeName}}, f func({{.TypeName}}) bool) []{{.TypeName}} {
			var result []{{.TypeName}}
			for _, v := range in {
				if f(v) {
					result = append(result, v)
				}
			}
			return result
		}
	`

	mapMethodName     = "Map"
	mapMethodTemplate = `
		func ` + mapMethodName + `(in []{{.TypeName}}, f func({{.TypeName}}) {{.TypeName}}) []{{.TypeName}} {
			out := make([]{{.TypeName}}, len(in))
			for i, v := range in {
				out[i] = f(v)
			}
			return out
		}
	`

	reduceMethodName     = "Reduce"
	reduceMethodTemplate = `
		func ` + reduceMethodName + `(in []{{.TypeName}}, f func({{.TypeName}}, {{.TypeName}}) {{.TypeName}}) {{.TypeName}} {
			var accumulator {{.TypeName}}
			if len(in) == 0 {
				return accumulator
			}
			accumulator = in[0]
			for i := 1; i < len(in); i++ {
				accumulator = f(accumulator, in[i])
			}
			return accumulator
		}
	`

	indexMethodName     = "Index"
	indexMethodTemplate = `
		func ` + indexMethodName + `(in []{{.TypeName}}, value {{.TypeName}}) int {
			for i, v := range in {
				if {{equal "v" "value"}} {
					return i
				}
			}
			return -1
		}
	`

	lastIndexMethodName     = "LastIndex"
	lastIndexMethodTemplate = `
		func ` + lastIndexMethodName + `(in []{{.TypeName}}, value {{.TypeName}}) int {
			for i := len(in)-1; i >= 0; i-- {
				if {{equal "in[i]" "value"}} {
					return i
				}
			}
			return -1
		}
	`

	indexAnyMethodName     = "IndexAny"
	indexAnyMethodTemplate = `
		func ` + indexAnyMethodName + `(in []{{.TypeName}}, values ...{{.TypeName}}) int {
			for i, v := range in {
				for _, value := range values {
					if {{equal "v" "value"}} {
						return i
					}
				}
			}
			return -1
		}
	`

	lastIndexAnyMethodName     = "LastIndexAny"
	lastIndexAnyMethodTemplate = `
		func ` + lastIndexAnyMethodName + `(in []{{.TypeName}}, values ...{{.TypeName}}) int {
			for i := len(in)-1; i >= 0; i-- {
				for _, value := range values {
					if {{equal "in[i]" "value"}} {
						return i
					}
				}
			}
			return -1
		}
	`

	indexFuncMethodName     = "IndexFunc"
	indexFuncMethodTemplate = `
		func ` + indexFuncMethodName + `(in []{{.TypeName}}, f func ({{.TypeName}}) bool) int {
			for i, v := range in {
				if f(v) {
					return i
				}
			}
			return -1
		}
	`

	lastIndexFuncMethodName     = "LastIndexFunc"
	lastIndexFuncMethodTemplate = `
		func ` + lastIndexFuncMethodName + `(in []{{.TypeName}}, f func ({{.TypeName}}) bool) int {
			for i := len(in)-1; i >= 0; i-- {
				if f(in[i]) {
					return i
				}
			}
			return -1
		}
	`
)
