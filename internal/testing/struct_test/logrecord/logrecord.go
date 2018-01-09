package logrecord

import "time"

const (
	DEBUG   = 0
	INFO    = 10
	WARNING = 20
	ERROR   = 30
	FATAL   = 100
)

//go:generate slicemeta -type logrecord.LogRecord -import "github.com/azavorotnii/slicemeta/internal/testing/struct_test/logrecord" -equal method -outputDir ./..
type LogRecord struct {
	Timestamp time.Time
	Level     int
	User      string
	Message   string
}

func (r LogRecord) Equal(other LogRecord) bool {
	return r.Timestamp.Equal(other.Timestamp) &&
		r.Level == other.Level &&
		r.User == other.User &&
		r.Message == other.Message
}
