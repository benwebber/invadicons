package main

import (
	"net/http"
	"strconv"

	"github.com/benwebber/invadicon"
	"github.com/gorilla/mux"
)

// Retrieve the invadicon size from the query string.
// Default to the default size defined in the invadicon package.
func getSize(r *http.Request) uint {
	size, err := strconv.Atoi(r.URL.Query().Get("size"))
	if err != nil {
		return invadicon.DefaultSize
	}
	return uint(size)
}

// Render invadicon as a PNG and wrap in an HTTP response.
func invadiconPNGHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	seed := vars["seed"] // will be empty if not defined
	w.Header().Set("Content-Type", "image/png")
	i, _ := invadicon.New(seed)
	i.Width, i.Height = getSize(r), getSize(r)
	i.Write(w)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", invadiconPNGHandler)
	r.HandleFunc("/{seed}.png", invadiconPNGHandler)
	r.HandleFunc("/{seed}", invadiconPNGHandler)
	http.Handle("/", r)
	http.ListenAndServe(":3001", nil)
}
