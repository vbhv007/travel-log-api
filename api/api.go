package api

import (
	"encoding/json"
	"fmt"
	"github.com/vbhv007/travel-log-api/database"
	"github.com/vbhv007/travel-log-api/dto"
	"io/ioutil"
	"net/http"
	"time"
)

type AddLogRequest struct {
	ID          uint
	Title       string
	Description string
	Rating      int
	ImageUrl    string
	Latitude    int
	Longitude   int
	UpdatedAt   time.Time
}

type LogsResponse struct {
	BaseResponse
	Logs    []*dto.LogEntity
}

type BaseResponse struct {
	Message string
	Error   error
}

type ErrorResponse struct {
	Message   string
	ErrorCode int
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	_, err := w.Write([]byte(`{"Message": "Not Found"}`))
	if err != nil {
		fmt.Errorf("failed to generate response, err=%v", err.Error())
	}
}

func Logs(w http.ResponseWriter, r *http.Request) {
	condition := struct{}{}
	response := LogsResponse{}
	logs, err := database.LogEntityDao.Find(condition)
	if err != nil {
		wrapError(500, "db query failed", w)
	}
	response.Error = nil
	response.Message = "Found some logs"
	response.Logs = logs
	wrapResponse(response, w)
}

func AddLog(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		wrapError(400, "Not able to read body", w)
		return
	}
	log := dto.LogEntity{}
	err = json.Unmarshal(body, &log)
	if err != nil {
		wrapError(500, "Not able to unmarshal body", w)
		return
	}
	err = database.LogEntityDao.Save(&log)
	if err != nil {
		wrapError(500, "unable to save into db", w)
		return
	}
	response := BaseResponse{}
	response.Message = "Log added"
	response.Error = nil
	wrapResponse(response, w)
}

func wrapError(statusCode int, message string, w http.ResponseWriter) {
	response := ErrorResponse{}
	w.WriteHeader(statusCode)
	response.Message = message
	response.ErrorCode = statusCode
	resp, err := json.Marshal(response)
	if err != nil {
		fmt.Errorf("unable to marshal response=%v", response)
		return
	}
	_, err = w.Write(resp)
	if err != nil {
		fmt.Errorf("failed to generate response, err=%v", err.Error())
		return
	}
}

func wrapResponse(respStruct interface{}, w http.ResponseWriter) {
	resp, err := json.Marshal(respStruct)
	if err != nil {
		wrapError(500, "unable to marshal response", w)
		return
	}
	w.WriteHeader(200)
	_, err = w.Write(resp)
	if err != nil {
		wrapError(500, "failed to generate response", w)
		return
	}
}
