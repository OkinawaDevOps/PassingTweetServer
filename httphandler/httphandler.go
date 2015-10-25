package httpdhandler

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.Method)
}

func StartServer(route string, port string) {
	http.HandleFunc(route, IndexHandler)
	http.ListenAndServe(port, nil)
}
