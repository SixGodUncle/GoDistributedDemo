package log

import (
	stlog "log"
	"os"
)

var log *stlog.Logger

type fileLog string

// 写入日志信息
func (fl fileLog) write(data []byte) (int, error) {
	f, err := os.OpenFile(string(fl), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)

	if err != nil {
		return 0, err
	}

	defer f.Close()
	return f.Write(data)
}

func Run(de)
