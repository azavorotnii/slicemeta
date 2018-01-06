package example_test

import (
	"testing"

	"github.com/azavorotnii/slicemeta/internal/example"
	"github.com/azavorotnii/slicemeta/internal/example/exampleutil"
	"time"
)

func TestExampleContains(t *testing.T) {
	pst, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		t.Fatal(err)
	}

	in := []example.Example{
		{Int: 10},
		{String: "qwerty"},
		{Bool: true},
		{Time: time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC)},
	}

	expectedTrue := []example.Example{
		{String: "qwerty"},
		{Int: 10},
		{Bool: true},
		{Time: time.Date(2001, 1, 1, -8, 0, 0, 0, pst)},
	}
	for _, v := range expectedTrue {
		if !exampleutil.Contains(in, v) {
			t.Error(v)
		}
	}

	expectedFalse := []example.Example{
		{String: "QWERTY"},
		{Int: 20},
		{Bool: false},
		{Time: time.Date(2001, 1, 1, 0, 0, 0, 0, pst)},
	}
	for _, v := range expectedFalse {
		if exampleutil.Contains(in, v) {
			t.Error(v)
		}
	}
}
