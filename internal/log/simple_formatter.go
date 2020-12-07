package log

import (
	"bytes"

	"github.com/sirupsen/logrus"
)

// SimpleFormatter is the most simple formatter you have ever seen!
type SimpleFormatter struct{}

// Format applies the default format for log entries.
func (f *SimpleFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var buffer *bytes.Buffer

	if entry.Buffer != nil {
		buffer = entry.Buffer
	} else {
		buffer = &bytes.Buffer{}
	}

	buffer.Write([]byte(entry.Message + "\n"))

	return buffer.Bytes(), nil
}
