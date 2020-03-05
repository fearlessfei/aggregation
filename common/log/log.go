package log

import (
	"github.com/sirupsen/logrus"

)

var Log *logrus.Logger

func init()  {
	Log = logrus.New()
	Log.SetFormatter(&logrus.TextFormatter{
		//ForceColors: true,
		//DisableColors: true,
		FullTimestamp: true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	Log.AddHook(&RedirectHook{})
}