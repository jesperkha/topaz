package test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/jesperkha/topaz"
)

func TestAll(t *testing.T) {
	server := topaz.NewServer()

	server.Get("/users/:id", func(req topaz.Request, res topaz.Response) {
		if req.Param("id") != "123" {
			t.Errorf("GET expected id of 123, got %s", req.Param("id"))
		}
		if req.Query("name") != "john" {
			t.Errorf("GET expected query to be 'john', got %s", req.Query("name"))
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

	server.Get("/", func(req topaz.Request, res topaz.Response) {
		err := res.File("test.txt")
		if err != nil {
			t.Error(err)
		}
	})

	go server.Listen(":3000")
	t.Run("GET request", testGet)
	t.Run("POST request", testPost)
	t.Run("FILE request", testFile)
}

func testGet(t *testing.T) {
	res, err := http.Get("http://localhost:3000/users/123?name=john")
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("GET got status %d", res.StatusCode)
	}
}

type person struct {
	name    string
	address string
}

func testPost(t *testing.T) {
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

func testFile(t *testing.T) {
	res, err := http.Get("http://localhost:3000/")
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("FILE got status %d", res.StatusCode)
	}
	c, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	if string(c) != "hello" {
		t.Errorf("FILE expected 'hello', got %s", string(c))
	}
}
