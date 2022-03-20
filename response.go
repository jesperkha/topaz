package topaz

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	errJsonMarshalFail = errors.New("topaz: failed to marshal json data")
)

type response struct {
	response http.ResponseWriter
	status   int
}

func (r *response) JSON(content interface{}) error {
	payload, err := json.Marshal(content)
	if err != nil {
		return errJsonMarshalFail
	}

	r.response.Write(payload)
	r.status = http.StatusOK
	return nil
}

func (r *response) Status(status int) {
	r.response.WriteHeader(status)
	r.status = status
}

func (r *response) Response() http.ResponseWriter {
	return r.response
}
