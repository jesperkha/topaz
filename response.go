package topaz

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
)

var (
	errJsonMarshalFail = errors.New("topaz: failed to marshal json data")
)

type response struct {
	response http.ResponseWriter
	status   int
}

func (r *response) Write(c []byte) {

}

func (r *response) JSON(content any) error {
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

func (r *response) File(file *os.File) error {

	return nil
}

func (r *response) Response() http.ResponseWriter {
	return r.response
}
