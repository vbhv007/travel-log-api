package api

import (
	"encoding/json"
	"fmt"
	"github.com/vbhv007/travel-log-api/database"
	"github.com/vbhv007/travel-log-api/dto"
	"net/http"
)

type LogsRequest struct {
	Condition string
}

type LogsResponse struct {
	Message string
	Error   error
	Logs    []*dto.LogEntity
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
		response.Error = err
		response.Message = "db query failed"
	} else {
		response.Error = nil
		response.Message = "Found some logs"
	}
	response.Logs = logs
	resp, err := json.Marshal(response)
	if err != nil {
		fmt.Errorf("unable to marshal struct, response=%v", response)
	}
	w.WriteHeader(200)
	_, err = w.Write(resp)
	if err != nil {
		fmt.Errorf("failed to generate response, err=%v", err.Error())
	}
}
