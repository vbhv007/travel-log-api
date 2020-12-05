package api

import (
	"encoding/json"
	"github.com/vbhv007/travel-log-api/dto"
	"github.com/vbhv007/travel-log-api/storage"
	"io/ioutil"
	"log"
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

func NotFound(w http.ResponseWriter, r *http.Request) {
	logTag := "NotFound"
	response := dto.BaseResponse{}
	response.Message = PageNotFound
	wrapResponse(logTag, NotFoundStatusCode, response, w)
}

func EmptyResponse(w http.ResponseWriter, r *http.Request) {
	logTag := "EmptyResponse"
	response := ""

	wrapResponse(logTag, OkStatusCode, response, w)
}

func Logs(w http.ResponseWriter, r *http.Request) {
	logTag := "Logs"
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		wrapError(logTag, BadRequestStatusCode, BodyReadError, err, w)
		return
	}
	condition := storage.LogEntity{}
	err = json.Unmarshal(body, &condition)
	if err != nil {
		wrapError(logTag, InternalServerErrorStatusCode, UnmarshalError, err, w)
		return
	}
	response := dto.LogsResponse{}
	logs, err := storage.LogEntityDao.Find(condition)
	if err != nil {
		wrapError(logTag, InternalServerErrorStatusCode, DBQueryError, err, w)
	}
	response.Message = "Found some logs"
	response.Logs = logs
	wrapResponse(logTag, OkStatusCode, response, w)
}

func AddLog(w http.ResponseWriter, r *http.Request) {
	logTag := "AddLog"
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		wrapError(logTag, BadRequestStatusCode, BodyReadError, err, w)
		return
	}
	log := storage.LogEntity{}
	err = json.Unmarshal(body, &log)
	if err != nil {
		wrapError(logTag, InternalServerErrorStatusCode, UnmarshalError, err, w)
		return
	}
	err = storage.LogEntityDao.Save(&log)
	if err != nil {
		wrapError(logTag, InternalServerErrorStatusCode, DBQueryError, err, w)
		return
	}
	response := dto.BaseResponse{}
	response.Message = "Log added"
	wrapResponse(logTag, OkStatusCode, response, w)
}

func wrapError(logTag string, statusCode int, message string, err error, w http.ResponseWriter) {
	response := dto.ErrorResponse{}
	w.WriteHeader(statusCode)
	response.Message = message
	response.Error = err.Error()
	resp, err := json.Marshal(response)
	if err != nil {
		log.Printf("%v | %v=%v", logTag, MarshalError, response)
		return
	}
	_, err = w.Write(resp)
	if err != nil {
		log.Printf("%v | %v, err=%v", logTag, ResponseError, err.Error())
		return
	}
}

func wrapResponse(logTag string, statusCode int, respStruct interface{}, w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	resp, err := json.Marshal(respStruct)
	if err != nil {
		wrapError(logTag, InternalServerErrorStatusCode, MarshalError, err, w)
		return
	}
	w.WriteHeader(statusCode)
	_, err = w.Write(resp)
	if err != nil {
		wrapError(logTag, InternalServerErrorStatusCode, ResponseError, err, w)
		return
	}
}
