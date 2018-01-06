package main

const containsTemplate = `
package {{.PackageName}}

// generated by slicemeta-0.0.1 ({{.Now}})

{{if .UseDeepEqual}}
import "reflect"
{{end}}
{{if .ImportTypeName}}
import "{{.ImportTypeName}}"
{{end}}

func Equal(a, b []{{.TypeName}}) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if {{if .UseDeelEqual}} !reflect.DeepEqual(a[i], b[i]) {{else}} a[i] != b[i] {{end}} {
			return false
		}
	}
	return true
}

func Contains(in []{{.TypeName}}, value {{.TypeName}}) bool {
	for _, v := range in {
		if {{if .UseDeepEqual}} reflect.DeepEqual(v, value) {{else}} v == value {{end}} {
			return true
		}
	}
	return false
}

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
