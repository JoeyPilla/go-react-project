/*Package api ...
Provides the api functionality through graphql
can be accessed at the root route of /api
*/
package api

import (
	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
)

// AddAPIHandler ... adds the api routes to the mux
func AddAPIHandler(mux *http.ServeMux) {
	mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: r.URL.Query().Get("query"),
		})
		json.NewEncoder(w).Encode(result)
	})
}
