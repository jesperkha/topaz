package topaz

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

var (
	errJsonUnmarshalFail = errors.New("failed to unmarshal json data")
)

type request struct {
	request    *http.Request
	response   http.ResponseWriter
	params     map[string]string
	redirected bool
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

func (r *request) Redirect(path string) {
	http.Redirect(r.response, r.request, path, http.StatusTemporaryRedirect)
	r.redirected = true
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
