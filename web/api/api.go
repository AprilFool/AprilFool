package api

import (
	"encoding/json"
	aresource "github.com/AprilFool/AprilFool/resource"
	gmux "github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
	HEAD   = "HEAD"
	PATCH  = "PATCH"
)

type API struct {
	path string
	mux  *gmux.Router
}

func NewAPI(path string, router *gmux.Router) *API {
	mux := router.PathPrefix(path).Subrouter()
	return &API{path: path, mux: mux}
}

func (api *API) requestHandler(resource aresource.Resource) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		if req.ParseForm() != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		var handler func(map[string][]string) (interface{}, error)
		log.Printf("Method: %v\n", req.Method)
		switch req.Method {
		case GET:
			if resource, ok := resource.(aresource.GetSupported); ok {
				handler = resource.Get
			}
		case POST:
			if resource, ok := resource.(aresource.PostSupported); ok {
				handler = resource.Post
			}
		case PUT:
			if resource, ok := resource.(aresource.PutSupported); ok {
				handler = resource.Put
			}
		case DELETE:
			if resource, ok := resource.(aresource.DeleteSupported); ok {
				handler = resource.Delete
			}
		case HEAD:
			if resource, ok := resource.(aresource.HeadSupported); ok {
				handler = resource.Head
			}
		case PATCH:
			if resource, ok := resource.(aresource.PatchSupported); ok {
				handler = resource.Patch
			}
		}

		if handler == nil {
			rw.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		vars := make(map[string][]string)
		for k, v := range gmux.Vars(req) {
			vars[k] = []string{v}
		}
		for k, v := range req.Form {
			vars[k] = v
		}
		data, _ := handler(vars)

		content, err := json.MarshalIndent(data, "", "  ")

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.Header().Add("Content-type", "application/json")

		rw.WriteHeader(200)
		rw.Write(content)
	}
}

func (api *API) Mux() *gmux.Router {
	return api.mux
}

func (api *API) AddResource(resource aresource.Resource) {
	var path string

	path = "/" + resource.Name()
	api.Mux().HandleFunc(path, api.requestHandler(resource))
	api.Mux().HandleFunc(path+"/{id}", api.requestHandler(resource))
}
