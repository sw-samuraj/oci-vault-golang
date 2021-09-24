package logging

import log "github.com/sirupsen/logrus"

func FuncLog(f string) *log.Entry {
	return log.WithFields(log.Fields{
		"func": f,
	})
}
