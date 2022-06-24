package log

import (
	"time"
)

type HTTP struct {
    ReceivedTime      time.Time
    RequestMethod     string
    RequestURL        string
    UserAgent         string
    Referer           string
    Proto             string

    RemoteIP string

    StatusCode         int
    Latency            time.Duration
}
