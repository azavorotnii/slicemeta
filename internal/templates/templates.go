package templates

import "sort"

var methodsMap = map[string]struct {
	templateText string
	imports      []string
}{
	containsMethodName:      {containsMethodTemplate, nil},
	containsAnyMethodName:   {containsAnyMethodTemplate, nil},
	containsFuncMethodName:  {containsFuncMethodTemplate, nil},
	countMethodName:         {countMethodTemplate, nil},
	countAnyMethodName:      {countAnyMethodTemplate, nil},
	countFuncMethodName:     {countFuncMethodTemplate, nil},
	equalMethodName:         {equalMethodTemplate, nil},
	filterMethodName:        {filterMethodTemplate, nil},
	indexMethodName:         {indexMethodTemplate, nil},
	indexAnyMethodName:      {indexAnyMethodTemplate, nil},
	indexFuncMethodName:     {indexFuncMethodTemplate, nil},
	lastIndexMethodName:     {lastIndexMethodTemplate, nil},
	lastIndexAnyMethodName:  {lastIndexAnyMethodTemplate, nil},
	lastIndexFuncMethodName: {lastIndexFuncMethodTemplate, nil},
	mapMethodName:           {mapMethodTemplate, nil},
	reduceMethodName:        {reduceMethodTemplate, nil},
	shuffleMethodName:       {shuffleMethodTemplate, []string{"math/rand"}},
}

var methodsMapKeys []string

func init() {
	for key := range methodsMap {
		methodsMapKeys = append(methodsMapKeys, key)
	}
	sort.Strings(methodsMapKeys)
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

	sortableWrapper = `
		type {{`+ titleMixin +` .TypeName}}Slice []{{.TypeName}}
		func (s {{`+ titleMixin +` .TypeName}}Slice) Len() int {
			return len(s)
		}
		func (s {{`+ titleMixin +` .TypeName}}Slice) Swap(i, j int) {
			s[i], s[j] = s[j], s[i]
		}
		func (s {{`+ titleMixin +` .TypeName}}Slice) Less(i, j int) bool {
			return {{` + lessMixin + ` "s[i]" "s[j]"}}
		}
	`

	containsMethodName     = "Contains"
	containsMethodTemplate = `
		func ` + containsMethodName + `(in []{{.TypeName}}, value {{.TypeName}}) bool {
			for _, v := range in {
				if {{` + equalMixin + ` "v" "value"}} {
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
					if {{` + equalMixin + ` "v" "value"}} {
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
				if {{` + equalMixin + ` "v" "value"}} {
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
					if {{` + equalMixin + ` "v" "value"}} {
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
				if !{{` + equalMixin + ` "a[i]" "b[i]"}} {
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
				if {{` + equalMixin + ` "v" "value"}} {
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
				if {{` + equalMixin + ` "in[i]" "value"}} {
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
					if {{` + equalMixin + ` "v" "value"}} {
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
					if {{` + equalMixin + ` "in[i]" "value"}} {
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

	shuffleMethodName     = "Shuffle"
	shuffleMethodTemplate = `
		func ` + shuffleMethodName + `(in []{{.TypeName}}) {
			for i := 0; i < len(in); i++ {
				j := i + rand.Int()%(len(in)-i)
				in[i], in[j] = in[j], in[i]
			}
		}
	`
)
