package libmodel

import (
	"time"
)

type SystemReturn struct {
	Code          int         `json:"code"`
	Message       string      `json:"message"`
}

type Response struct {
	Code          int         `json:"code"`
	Message       string      `json:"message"`
	Data          interface{} `json:"data"`
	RequestedTime time.Time   `json:"requested_time"`
	ResponsedTime time.Time   `json:"responsed_time"`
}