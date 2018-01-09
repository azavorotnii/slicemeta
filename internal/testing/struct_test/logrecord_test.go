package struct_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/azavorotnii/slicemeta/internal/testing/struct_test/logrecord"
	"github.com/azavorotnii/slicemeta/internal/testing/struct_test/logrecordutil"
)

var PST *time.Location

func init() {
	PST, _ = time.LoadLocation("America/Los_Angeles")
}

func setup() []logrecord.LogRecord {
	return []logrecord.LogRecord{
		logrecord.LogRecord{
			Timestamp: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			User:      "root",
			Level:     logrecord.DEBUG,
			Message:   "System log level DEBUG test.",
		},
		logrecord.LogRecord{
			Timestamp: time.Date(2000, 1, 1, 0, 0, 20, 0, time.UTC),
			User:      "root",
			Level:     logrecord.INFO,
			Message:   "System started sucessfully",
		},
		logrecord.LogRecord{
			Timestamp: time.Date(2017, 1, 9, 8, 15, 20, 0, PST),
			User:      "root",
			Level:     logrecord.INFO,
			Message:   "Time updated with NTP server.",
		},
		logrecord.LogRecord{
			Timestamp: time.Date(2017, 1, 9, 8, 15, 30, 0, PST),
			User:      "root",
			Level:     logrecord.WARNING,
			Message:   "Could not find  ",
		},
	}
}

func TestContains(t *testing.T) {
	input := setup()

	// have timestamp in UTC, not in PST but still same moment of time
	log1 := logrecord.LogRecord{
		Timestamp: time.Date(2017, 1, 9, 16, 15, 20, 0, time.UTC),
		User:      "root",
		Level:     logrecord.INFO,
		Message:   "Time updated with NTP server.",
	}
	assert.True(t, logrecordutil.Contains(input, log1))

	// have same time and date but in different timezone
	log2 := logrecord.LogRecord{
		Timestamp: time.Date(2000, 1, 1, 0, 0, 0, 0, PST),
		User:      "root",
		Level:     logrecord.DEBUG,
		Message:   "System log level DEBUG test.",
	}
	assert.False(t, logrecordutil.Contains(input, log2))

	assert.True(t, logrecordutil.ContainsAny(input, log2, log1))

	isFatal := func(lr logrecord.LogRecord) bool {
		return lr.Level == logrecord.FATAL
	}

	isWarning := func(lr logrecord.LogRecord) bool {
		return lr.Level == logrecord.WARNING
	}
	isDebug := func(lr logrecord.LogRecord) bool {
		return lr.Level == logrecord.WARNING
	}

	assert.True(t, logrecordutil.ContainsFunc(input, isWarning))
	assert.False(t, logrecordutil.ContainsFunc(input, isFatal))
}
