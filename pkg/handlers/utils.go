package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strings"
	"time"
	"github.com/golang/glog"
)

const (
	randStringArray     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890~!@#$%^&*()|}{"
)

func WriteError(w http.ResponseWriter, e *error, code uint, message string) error {
	glog.Errorf("[%v] %v %v", code, message, *e)
	WriteResponse(w, &JsonResponse{
		Code:    code,
		Message: message,
	}, http.StatusBadRequest)
	return nil
}

func generateRandomString() string {
	var stringBuilder strings.Builder
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for stringBuilder.Len() < 16 {
		randomIndex := r.Int() % len(randStringArray)

		runeChar := randStringArray[randomIndex]
		stringBuilder.WriteByte(runeChar)
	}

	glog.V(3).Infof("generated string %v", stringBuilder.String())

	return stringBuilder.String()
}

func WriteResponse(w http.ResponseWriter, response interface{}, status int) {
	w.Header().Set(contentHeader, jsonUtf8)
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
