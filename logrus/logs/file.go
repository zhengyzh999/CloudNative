package logs

import "os"

const LOGPATH = "runtime/logs/logs.log"

type FileWriter struct {
	*os.File
}

func (fw *FileWriter) Flush() {
	err := fw.Sync()
	if err != nil {
		return
	}
}
func newFileWriter() LogWriter {
	file, err := os.OpenFile(LOGPATH, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		file = os.Stderr
	}
	return &FileWriter{
		file,
	}
}
func init() {
	RegisterInitLogWriterFunc("file", newFileWriter)
}
