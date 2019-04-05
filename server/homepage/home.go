package homepage

import "net/http"

// adds the homepage routes to the mux
func AddHomepageHandler(mux *http.ServeMux) {
	fs := http.FileServer(http.Dir("../client/build"))
	mux.Handle("/", fs)
}
