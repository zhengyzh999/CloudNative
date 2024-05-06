package logs

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"time"
)

type fileRotateWriter struct {
	*rotatelogs.RotateLogs
}

func (frw *fileRotateWriter) Flush() {
	err := frw.Close()
	if err != nil {
		return
	}
}

func newFileRotateWriter() LogWriter {
	writer, err := getWriter()
	if err != nil {
		return newStdWriter()
	}
	return &fileRotateWriter{
		writer,
	}
}

func getWriter() (*rotatelogs.RotateLogs, error) {
	path := LOGPATH
	logf, err := rotatelogs.New(
		path+".%Y%m%d%H%M",
		//rotatelogs.WithLinkName("/path/to/access_log"),
		rotatelogs.WithMaxAge(time.Second*180),
		rotatelogs.WithRotationTime(time.Second*60),
	)
	return logf, err
}

func init() {
	RegisterInitLogWriterFunc("fileRotate", newFileRotateWriter)
}
