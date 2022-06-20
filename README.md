<br />
<div align="center">
  <img src=".github/topaz.png" alt="Logo" width="80">
  <br>
  <img src=".github/title.svg" alt="Logo" width="80">

  <p align="center">
    Simplified HTTP server package. Lets you  set up<br> a web server in minimal time.
  </p>
</div>

<br>
<br>

## Installation

```console
$ go get github.com/jesperkha/topaz
```

<br>

## Example

```go
...

type user struct {
  id string
}

func main() {
  server := topaz.NewServer()

  // Serve a static directory
  server.Static("/", "pages")

  // Example of a handler to create a new user json object with a given id
  server.Get("/create/:id", func(req topaz.Request, res topaz.Response) {
    // Gets the id value from the URL
    userId := req.Param("id")
    newUser := user{id: userId}

    // Send back data as JSON
    if err := res.JSON(newUser); err != nil {
      res.Status(500)
    }

    // If a status is not set 200 is assumed
  })

  // Will serve to localhost:3000 unless the PORT env variable is set
  server.Listen(server.EnvPort(":3000"))
}
```
