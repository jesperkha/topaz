package test

import (
	"testing"

	"github.com/jesperkha/topaz"
)

var server topaz.Server

func TestAll(t *testing.T) {
	server = topaz.NewServer()
	server.Get("/", func(req topaz.Request, res topaz.Response) {
		res.JSON("{'hello': 'world'}")
	})

	go server.Listen(":3000")
}
