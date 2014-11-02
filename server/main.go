package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ajstarks/svgo"
	"github.com/benwebber/bitboard"
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

func invadiconSVGHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	seed := vars["seed"] // will be empty if not defined
	w.Header().Set("Content-Type", "image/svg+xml")
	i, _ := invadicon.New(seed)
	i.Width, i.Height = getSize(r), getSize(r)
	s := svg.New(w)
	s.Start(int(i.Width), int(i.Height))
	style := fmt.Sprintf("fill: #000000;")
	dx, dy := int(i.Width/10), int(i.Height/10)
	for y := 1; y < 9; y++ {
		for x := 1; x < 9; x++ {
			p := (y-1)*8 + (x - 1)
			if bitboard.GetBit(&i.Bitmap, p) == 1 {
				s.Rect(x*dx, y*dx, dx, dy, style)
			}
		}
	}
	s.End()
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
	r.HandleFunc("/{seed}.svg", invadiconSVGHandler)
	r.HandleFunc("/{seed}.png", invadiconPNGHandler)
	r.HandleFunc("/{seed}", invadiconPNGHandler)
	http.Handle("/", r)
	http.ListenAndServe(":3001", nil)
}
