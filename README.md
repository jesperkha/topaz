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

  // On a GET request to for example /users/1234
  server.Get("/users/:id", func(req topaz.Request, res topaz.Response) {
    // Gets the id value from the URL
    userId := req.Param("id")
    newUser := user{id: userId}

    // Send back data as JSON
    if err := res.JSON(newUser); err != nil {
      res.Status(500)
    }

    // If a status is not set 200 is assumed
  })

  server.Listen(":3000")
}
```
