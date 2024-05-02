package controller

import (
	"net/http"
	"time"

	"github.com/Mau005/MyPet/models"
)

type ControllerException struct {
	ResponseWriter http.ResponseWriter
}

func (ce *ControllerException) NewException(error, message string, status int, timeStamp *time.Time) (excep models.Exception) {
	ce.ResponseWriter.WriteHeader(status)
	excep.Error = error
	excep.Message = message
	excep.Status = status
	if timeStamp == nil {
		excep.TimeStamp = time.Now()
		return
	}
	excep.TimeStamp = *timeStamp
	return
}
