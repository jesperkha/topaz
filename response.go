package topaz

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
)

var (
	errJsonMarshalFail = errors.New("failed to marshal json data")
)

type response struct {
	response      http.ResponseWriter
	request       *http.Request
	statusWritten bool
}

func (r *response) Write(c []byte) {
	r.response.Write(c)
	r.statusWritten = true
}

func (r *response) JSON(content any) error {
	payload, err := json.Marshal(content)
	if err != nil {
		return errJsonMarshalFail
	}

	r.response.Write(payload)
	r.statusWritten = true
	return nil
}

func (r *response) Status(status int) {
	r.response.WriteHeader(status)
	r.statusWritten = true
}

func (r *response) File(filename string) error {
	if _, err := os.Open(filename); err != nil {
		return err
	}
	http.ServeFile(r.response, r.request, filename)
	r.statusWritten = true
	return nil
}

func (r *response) Response() http.ResponseWriter {
	return r.response
}
