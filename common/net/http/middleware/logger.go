package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"

	"solo/solo-go-common/ecode"
	logger "solo/solo-go-common/log"
	"github.com/sirupsen/logrus"

)

// Logger is logger  middleware
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		now := time.Now()
		req := c.Request
		ip := c.Request.RemoteAddr
		path := req.URL.Path
		params := req.Form

		c.Next()

		err := c.MustGet("err")
		var cerr ecode.Codes
		if err == nil {
			cerr = ecode.Cause(nil)
		} else {
			cerr = ecode.Cause(err.(ecode.Codes))
		}

		dt := time.Since(now)

		log := logger.Log

		_log := log.WithFields(logrus.Fields{
			"method": req.Method,
			"ip": ip,
			"path": path,
			"params": params.Encode(),
			"ret": cerr.Code(),
			"message": cerr.Message(),
			"stack": fmt.Sprintf("%+v", err),
			"ts": dt.Seconds(),
		})

		lf := _log.Info
		isSlow := dt >= (time.Millisecond * 500)
		if err != nil {
			//errmsg = err.Error()
			lf = _log.Error
			if cerr.Code() > 0 {
				lf = _log.Info
			}
		} else {
			if isSlow {
				lf = _log.Warn
			}
		}
		lf("This is log info.")
	}
}
