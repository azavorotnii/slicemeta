package exampleutil

// Generated by slicemeta-0.0.1a (2018-01-05T20:40:44-08:00)

import "github.com/azavorotnii/slicemeta/internal/example"

func Filter(in []example.Example, filter func(example.Example) bool) []example.Example {
	var result []example.Example
	for _, v := range in {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}