package transport

import "net/http"

type HTTPHandler interface {
	RegisterUser(w http.ResponseWriter, r *http.Request)
}
