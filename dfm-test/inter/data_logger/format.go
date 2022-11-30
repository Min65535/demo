package logger

type MatchFormatter struct {
}

func (mf *MatchFormatter) Format(msg []byte) ([]byte, error) {
	return msg, nil
}
