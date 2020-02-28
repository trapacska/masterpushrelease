// Package p contains an HTTP Cloud Function.
package p

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

// Deploy ...
func Deploy(w http.ResponseWriter, r *http.Request) {
	b, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(string(b))
	w.WriteHeader(http.StatusOK)
}
