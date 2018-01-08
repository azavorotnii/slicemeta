package flagutil

import (
	"strings"
	"flag"
)

type StringList []string

var _ flag.Value = &StringList{} // ensure interface

func (sl StringList) String() string {
	return strings.Join(sl, ",")
}

func (sl *StringList) Set(value string) error {
	*sl = append(*sl, value)
	return nil
}