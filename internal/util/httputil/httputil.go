package httputil

import (
	"encoding/json"
	"github.com/thehaung/golang-oracle-example/domain"
	"github.com/thehaung/golang-oracle-example/internal/util/typeutil"
	"log"
	"net/http"
)

type HttpUtil struct {
	Request *http.Request
	Writer  http.ResponseWriter
}

func NewHttpUtil(w http.ResponseWriter, r *http.Request) *HttpUtil {
	return &HttpUtil{
		Request: r,
		Writer:  w,
	}
}

func (h *HttpUtil) WriteJson(status int, data interface{}) {
	h.Writer.Header().Set(typeutil.ContentType, typeutil.ContentTypeJson)
	h.Writer.WriteHeader(status)

	res, err := json.Marshal(data)

	if err != nil {
		log.Println()
	}

	if _, err = h.Writer.Write(res); err != nil {
		log.Println(err.Error())
	}
}

func (h *HttpUtil) WriteError(status int, message string, err error) {
	h.Writer.Header().Set(typeutil.ContentType, typeutil.ContentTypeJson)
	h.Writer.WriteHeader(status)
	data := &domain.ErrorResponse{
		ErrorCode:    status,
		ErrorMessage: message,
		RawMessage:   err.Error(),
	}

	res, err := json.Marshal(data)
	if err != nil {
		log.Println(err.Error())
	}

	if _, err = h.Writer.Write(res); err != nil {
		log.Println(err.Error())
	}
}
