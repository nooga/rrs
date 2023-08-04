package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func Start(host, route, rickroll, fake string, previews []string) error {
	if !strings.HasPrefix(route, "/") {
		route = "/" + route
	}

	router := mux.NewRouter()

	router.HandleFunc("/r/{bust}"+route, func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		bust := params["bust"]

		useragent := r.Header.Get("User-Agent")

		shouldFake := false
		for _, p := range previews {
			if strings.Contains(useragent, p) {
				shouldFake = true
				break
			}
		}

		fmt.Print("Request ", bust, " from: ", useragent)
		if shouldFake {
			fmt.Print(", showing fake\n")
			http.Redirect(w, r, fake, http.StatusFound)
		} else {
			fmt.Print(", showing rickroll\n")
			http.Redirect(w, r, rickroll, http.StatusFound)
		}
	})

	return http.ListenAndServe(host, router)
}
