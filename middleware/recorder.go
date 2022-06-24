package middleware

import (
    "net/http"
    "strconv"
	"time"

	"github.com/rs/zerolog"

    "github.com/michaelact/firstApi/model/log"
)

type RecorderMiddleware struct {
	Handler http.Handler
	Logger  *zerolog.Logger
}

func NewRecorderMiddleware(handler http.Handler, logger *zerolog.Logger) *RecorderMiddleware {
	return &RecorderMiddleware{
		Handler: handler, 
		Logger:  logger, 
	}
}

func (self *RecorderMiddleware) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	start := time.Now()
    logField := &log.HTTP{
        ReceivedTime:      start,
        RequestMethod:     req.Method,
        RequestURL:        req.URL.String(),
        UserAgent:         req.UserAgent(),
        Referer:           req.Referer(),
        Proto:             req.Proto,
        RemoteIP:          req.RemoteAddr,
    }

    self.Handler.ServeHTTP(res, req)

    statusCode, err := strconv.Atoi(res.Header().Get("StatusCode"))
    if err != nil {
        self.Logger.Warn().Err(err)
    }

    logField.StatusCode = statusCode
    logField.Latency = time.Since(start)
    self.Logger.Info().
	    Time("received_time", logField.ReceivedTime).
        Str("method", logField.RequestMethod).
        Str("url", logField.RequestURL).
        Str("agent", logField.UserAgent).
        Str("referer", logField.Referer).
        Str("proto", logField.Proto).
        Str("remote_ip", logField.RemoteIP).
        Int("status", logField.StatusCode).
        Dur("latency", logField.Latency).
        Msg("")
}
