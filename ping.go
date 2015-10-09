package marketdata

import (
	"fmt"
	"net/http"
)

func Ping(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "pong")
}