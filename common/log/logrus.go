package log

import (
	"github.com/sirupsen/logrus"
)

type RedirectHook struct{}

func (RedirectHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (RedirectHook) Fire(entry *logrus.Entry) error {
	return nil
}
