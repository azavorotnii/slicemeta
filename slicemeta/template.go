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

func Contains(in []{{.TypeName}}, value {{.TypeName}}) bool {
	for _, v := range in {
		if {{if .UseDeepEqual}} reflect.DeepEqual(v, value) {{else}} v == value {{end}} {
			return true
		}
	}
	return false
}
`