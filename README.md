`publ` generates root path urls to the files at the designated folder in golang web app. It also generates the urls to the files in descendant folders.

## Installation

```bash
go get github.com/marusker/publi
```

## Sample usage(net/http)

```go
import (
  "github.com/marusker/publ"
  "net/http"
)

func init(){
  //give a path of public folder in func init()
  //all descendant folders will also be public.
  publ.SetFolder("./public")
}

func main() {
  //need to be activated once
  publ.Activate()
  http.ListenAndServe(":8080", nil)
}

```

## gorilla mux sample

It can be used for [gorilla mux](https://github.com/gorilla/mux) too.

```go
import (
  "github.com/gorilla/mux"
  "github.com/marusker/publ"
  "net/http"
)

func init(){
  publ.SetFolder("./public")
}

func main() {
  r := mux.NewRouter()
  publ.ActivateGorilla(r)

  http.ListenAndServe(":8080", r)
}
```

