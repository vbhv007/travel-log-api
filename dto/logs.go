package dto

import (
	"github.com/vbhv007/travel-log-api/storage"
)

type LogsResponse struct {
	BaseResponse
	Logs []*storage.LogEntity `json:"logs"`
}

type BaseResponse struct {
	Message string `json:"msg"`
}
