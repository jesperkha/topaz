package test

import (
	"net/http"
	"testing"

	"github.com/jesperkha/topaz"
)

var server topaz.Server

func TestAll(t *testing.T) {
	server = topaz.NewServer()
	server.Get("/users/:id", func(req topaz.Request, res topaz.Response) {
		id, err := req.Param("id")
		if err != nil {
			t.Error(err)
		}

		if id != "123" {
			t.Errorf("expected id of 123, got %s", id)
		}
	})

	go server.Listen(":3000")
	t.Run("GET request", TestGet)
}

func TestGet(t *testing.T) {
	if _, err := http.Get("http://localhost:3000/hello/123"); err != nil {
		t.Error(err)
	}
}
