package logger

import log "github.com/sirupsen/logrus"

type MatchFormatter struct {
}

func (mf *MatchFormatter) Format(entry *log.Entry) ([]byte, error) {
	return []byte(entry.Message), nil
}
