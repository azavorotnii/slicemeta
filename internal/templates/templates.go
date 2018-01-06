package templates

var PkgTemplates = map[string]string{
	"contains.go": containsTemplate,
	"equal.go":    equalTemplate,
	"filter.go":   filterTemplate,
	"index.go":    indexTemplate,
}

const packageHeaderTemplate = `
package {{.PackageName}}

// {{.Comment}}

{{if .UseDeepEqual}}
import "reflect"
{{end}}
{{if .ImportTypeName}}
import "{{.ImportTypeName}}"
{{end}}

`

const containsTemplate = packageHeaderTemplate + `
func Contains(in []{{.TypeName}}, value {{.TypeName}}) bool {
	for _, v := range in {
	{{- if .UseDeepEqual -}}
		if reflect.DeepEqual(v, value) {
	{{- else if .UseEqual -}}
		if v.Equal(value) {
	{{- else -}}
		if v == value {
	{{- end -}}
			return true
		}
	}
	return false
}
`
const equalTemplate = packageHeaderTemplate + `
func Equal(a, b []{{.TypeName}}) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if {{if .UseDeepEqual}} !reflect.DeepEqual(a[i], b[i]) {{else}} a[i] != b[i] {{end}} {
			return false
		}
	}
	return true
}
`
const filterTemplate = packageHeaderTemplate + `
func Filter(in []{{.TypeName}}, filter func({{.TypeName}}) bool) []{{.TypeName}} {
	var result []{{.TypeName}}
	for _, v := range in {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}
`

const indexTemplate = packageHeaderTemplate + `
func Index(in []{{.TypeName}}, value {{.TypeName}}) int {
	for i, v := range in {
		if {{if .UseDeepEqual}} reflect.DeepEqual(v, value) {{else}} v == value {{end}} {
			return i
		}
	}
	return -1
}

func IndexAny(in []{{.TypeName}}, values []{{.TypeName}}) int {
	for i, v := range in {
		for _, value := range values {
			if {{if .UseDeepEqual}} reflect.DeepEqual(v, value) {{else}} v == value {{end}} {
				return i
			}
		}
	}
	return -1
}

func IndexFunc(in []{{.TypeName}}, f func ({{.TypeName}}) bool) int {
	for i, v := range in {
		if f(v) {
			return i
		}
	}
	return -1
}
`
