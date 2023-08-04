package server

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func Start(host, rickroll string, previews []string) error {
	router := mux.NewRouter()

	router.HandleFunc("/r/{bust}/{url:.*}", func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		bust := params["bust"]
		url := params["url"]

		if !strings.HasPrefix(url, "http") {
			url = "https://" + url
		}

		useragent := r.Header.Get("User-Agent")

		shouldFake := false
		for _, p := range previews {
			if strings.Contains(useragent, p) {
				shouldFake = true
				break
			}
		}

		fmt.Print(time.Now().Round(0), ": ", "Request ", bust, " from: ", useragent)
		if shouldFake {
			fmt.Print(", showing fake ", url, "\n")
			http.Redirect(w, r, url, http.StatusFound)
		} else {
			fmt.Print(", showing rickroll\n")
			http.Redirect(w, r, rickroll, http.StatusFound)
		}
	})

	return http.ListenAndServe(host, router)
}
