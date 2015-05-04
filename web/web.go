package web

import (
	"fmt"
	aresource "github.com/AprilFool/AprilFool/resource"
	aapi "github.com/AprilFool/AprilFool/web/api"
	gmux "github.com/gorilla/mux"
	"net/http"
)

func Start(port int) error {
	r := gmux.NewRouter()
	api := aapi.NewAPI("/api", r)
	api.AddResource(new(aresource.TagResource))
	return http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
