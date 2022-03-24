package topaz

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

var (
	errJsonUnmarshalFail = errors.New("topaz: failed to unmarshal json data")
	errParamNotFound     = errors.New("topaz: could not find parameter '%s' in request")
)

type request struct {
	request *http.Request
	params  map[string]string
}

func (r *request) JSON(dest any) error {
	data, err := io.ReadAll(r.request.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, dest); err != nil {
		return errJsonUnmarshalFail
	}
	return nil
}

func (r *request) URL() string {
	return r.request.URL.Path
}

func (r *request) Query(key string) string {
	return r.request.URL.Query().Get(key)
}

func (r *request) Param(param string) string {
	if val, ok := r.params[param]; ok {
		return val
	}
	return ""
}

func (r *request) Request() *http.Request {
	return r.request
}
