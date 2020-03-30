package machine

import (
	"os"
)

// Recorder contains methods for recording machine learning training data.
type Recorder struct {
	data *os.File
}

// NewRecorder will create a new Recorder instance with an open file for storing training data.
func NewRecorder(file string) *Recorder {
	data, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	return &Recorder{
		data: data,
	}
}

// Close the open file handle. This should always be called.
func (r *Recorder) Close() {
	r.data.Close()
}

// Write a record to a file.
func (r *Recorder) Write(record *Record) {
	if r == nil || record == nil {
		return
	}
	record.Write(r.data)
}
