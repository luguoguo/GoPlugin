package GoPlugin

import "net/http"

type AddHeader struct {
}

func New() http.Handler {
	return &AddHeader{}
}

func (this *AddHeader) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("name", "myjobs")
}
