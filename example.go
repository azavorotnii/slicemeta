package slicemeta

//go:generate slicemeta -type MyStruct -import "github.com/azavorotnii/slicemeta" -outputDir .
type MyStruct struct {
	A, B int
}
