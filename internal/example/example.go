package example

import "time"

//go:generate slicemeta -type "example.Example" -import "github.com/azavorotnii/slicemeta/internal/example" -equalityOp equal -outputDir .
type Example struct {
	Int    int
	String string
	Bool   bool
	Time   time.Time
}

func (e Example) Equal(other Example) bool {
	if !e.Time.Equal(other.Time) {
		return false
	}
	e.Time = other.Time
	return e == other
}