package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Mau005/MyPet/models"
)

type ControllerException struct {
}

func (ce *ControllerException) NewException(w http.ResponseWriter, handlerError, message string, status int, timeStamp *time.Time) (err error) {
	var excep models.Exception
	w.WriteHeader(status)
	excep.Error = handlerError
	excep.Message = message
	excep.Status = status
	if timeStamp == nil {
		excep.TimeStamp = time.Now()
	} else {
		excep.TimeStamp = *timeStamp
	}
	return json.NewEncoder(w).Encode(excep)
}
