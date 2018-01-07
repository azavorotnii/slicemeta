package templates

var PackageMethodsTemplates = map[string]string{
	containsMethodName:  containsMethodTemplate,
	equalMethodName:     equalMethodTemplate,
	filterMethodName:    filterMethodTemplate,
	indexMethodName:     indexMethodTemplate,
	indexAnyMethodName:  indexAnyMethodTemplate,
	indexFuncMethodName: indexFuncMethodTemplate,
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
		func ` + filterMethodName + `(in []{{.TypeName}}, filter func({{.TypeName}}) bool) []{{.TypeName}} {
			var result []{{.TypeName}}
			for _, v := range in {
				if filter(v) {
					result = append(result, v)
				}
			}
			return result
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

	indexAnyMethodName     = "IndexAny"
	indexAnyMethodTemplate = `
		func ` + indexAnyMethodName + `(in []{{.TypeName}}, values []{{.TypeName}}) int {
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
)
