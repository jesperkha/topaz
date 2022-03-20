package topaz

import (
	"bufio"
	"encoding/json"
	"errors"
	"net/http"
)

var (
	errJsonUnmarshalFail = errors.New("topaz: failed to unmarshal json data")
	errParamNotFound     = errors.New("topaz: could not find parameter '%s' in request")
)

type request struct {
	request *http.Request
	query   map[string]string
	params  map[string]string
}

func (r *request) JSON(dest interface{}) error {
	reader := bufio.NewReader(r.request.Body)
	data := make([]byte, reader.Size())
	reader.Read(data)

	if err := json.Unmarshal(data, dest); err != nil {
		return errJsonUnmarshalFail
	}

	return nil
}

func (r *request) URL() string {
	return r.request.URL.Path
}

func (r *request) Query(param string) (value string, err error) {
	if val, ok := r.query[param]; ok {
		return val, err
	}

	return value, errorf(errParamNotFound, param)
}

func (r *request) Param(param string) (value string, err error) {
	if val, ok := r.query[param]; ok {
		return val, err
	}

	return value, errorf(errParamNotFound, param)
}

func (r *request) Request() *http.Request {
	return r.request
}
