package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/jesperkha/topaz"
)

var server topaz.Server

type person struct {
	name    string
	address string
}

func TestAll(t *testing.T) {
	server = topaz.NewServer()
	server.Get("/users/:id", func(req topaz.Request, res topaz.Response) {
		if req.Param("id") != "123" {
			t.Errorf("expected id of 123, got %s", req.Param("id"))
		}
		if err := res.JSON("{'hello': 'world'}"); err != nil {
			t.Error(err)
		}
	})

	server.Post("/people/:name/:address", func(req topaz.Request, res topaz.Response) {
		var data person
		err := req.JSON(&data)
		if err != nil {
			t.Error(err)
		}
		name, addr := req.Param("name"), req.Param("address")
		if name != "james" || addr != "road" {
			t.Errorf("POST expected 'james' and 'road', got '%s' and '%s'", name, addr)
		}
	})

	go server.Listen(":3000")
	t.Run("GET request", TestGet)
}

func TestGet(t *testing.T) {
	res, err := http.Get("http://localhost:3000/users/123")
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("GET got status %d", res.StatusCode)
	}
}

func TestPost(t *testing.T) {
	data, err := json.Marshal(person{name: "name", address: "address"})
	if err != nil {
		t.Error(err)
	}
	res, err := http.Post("http://localhost:3000/people/james/road", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("POST got status %d", res.StatusCode)
	}
}
