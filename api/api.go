package api

import (
	"encoding/json"
	"fmt"
	"github.com/vbhv007/travel-log-api/database"
	"github.com/vbhv007/travel-log-api/dto"
	"io/ioutil"
	"net/http"
)

const (
	InternalServerErrorStatusCode = 500
	NotFoundStatusCode            = 404
	BadRequestStatusCode          = 400
	OkStatusCode                  = 200
	PageNotFound                  = "Page Not Found"
	MarshalError                  = "unable to marshal response"
	UnmarshalError                = "unable to unmarshal body"
	BodyReadError                 = "unable to read body"
	DBQueryError                  = "db query failed"
	ResponseError                 = "failed to generate response"
)

type LogsResponse struct {
	BaseResponse
	Logs []*dto.LogEntity	`json:"logs"`
}

type BaseResponse struct {
	Message string `json:"msg"`
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	response := BaseResponse{}
	response.Message = PageNotFound
	wrapResponse(NotFoundStatusCode, response, w)
}

func Logs(w http.ResponseWriter, r *http.Request) {
	logtag := "Logs"
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		wrapError(logtag, BadRequestStatusCode, BodyReadError, w)
		return
	}
	condition := dto.LogEntity{}
	err = json.Unmarshal(body, &condition)
	if err != nil {
		wrapError(logtag, InternalServerErrorStatusCode, UnmarshalError, w)
		return
	}
	response := LogsResponse{}
	logs, err := database.LogEntityDao.Find(condition)
	if err != nil {
		wrapError(logtag, InternalServerErrorStatusCode, DBQueryError, w)
	}
	response.Message = "Found some logs"
	response.Logs = logs
	wrapResponse(OkStatusCode, response, w)
}

func AddLog(w http.ResponseWriter, r *http.Request) {
	logtag := "AddLog"
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		wrapError(logtag, BadRequestStatusCode, BodyReadError, w)
		return
	}
	log := dto.LogEntity{}
	err = json.Unmarshal(body, &log)
	if err != nil {
		wrapError(logtag, InternalServerErrorStatusCode, UnmarshalError, w)
		return
	}
	err = database.LogEntityDao.Save(&log)
	if err != nil {
		wrapError(logtag, InternalServerErrorStatusCode, DBQueryError, w)
		return
	}
	response := BaseResponse{}
	response.Message = "Log added"
	wrapResponse(OkStatusCode, response, w)
}

func wrapError(logtag string, statusCode int, message string, w http.ResponseWriter) {
	response := BaseResponse{}
	w.WriteHeader(statusCode)
	response.Message = message
	resp, err := json.Marshal(response)
	if err != nil {
		fmt.Errorf("%v | %v=%v", logtag, MarshalError, response)
		return
	}
	_, err = w.Write(resp)
	if err != nil {
		fmt.Errorf("%v | %v, err=%v", logtag, ResponseError, err.Error())
		return
	}
}

func wrapResponse(statusCode int, respStruct interface{}, w http.ResponseWriter) {
	logtag := "Response"
	resp, err := json.Marshal(respStruct)
	if err != nil {
		wrapError(logtag, InternalServerErrorStatusCode, MarshalError, w)
		return
	}
	w.WriteHeader(statusCode)
	_, err = w.Write(resp)
	if err != nil {
		wrapError(logtag, InternalServerErrorStatusCode, ResponseError, w)
		return
	}
}
