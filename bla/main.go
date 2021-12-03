package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func c(e *error) func(...interface{}) bool {
	return func(i ...interface{}) bool {
		if len(i) < 1 {
			return false
		}

		v, ok := i[len(i)-1].(error)

		if ok {
			*e = v
		}

		return ok
	}
}

type path string

func (p path) next() (path, string) {
	if len(p) < 1 {
		return "", ""
	}

	p = p[1:]

	i := strings.IndexRune(string(p), '/')

	if i < 0 {
		i = len(p)
	}

	return p[i:], string(p[0:i])
}

type MethodRouter map[string]http.Handler
type SegmentRouter map[string]http.Handler
type HostRouter map[string]http.Handler

func main() {

	fmt.Println(HostRouter{
		"loclahos": nil,
	})

	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := path(r.URL.Path)
		p, s := p.next()

		switch s {
		case "":
			w.Write([]byte("INDEX"))
		case "a":
			switch p, s = p.next(); s {
			case "":
				w.Write([]byte("A"))
			case "b":
				switch r.Method {
				case http.MethodGet:
					w.Write([]byte("B: " + string(p)))
				default:
					json.MarshalIndent()
					http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
				}
			default:
				http.NotFound(w, r)
			}
		case "b":
			w.Write([]byte("B: " + string(p)))
		default:
			http.NotFound(w, r)
		}
	}))
}
