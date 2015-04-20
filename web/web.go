package web

import (
	"fmt"
	aresource "github.com/AprilFool/AprilFool/resource"
	aapi "github.com/AprilFool/AprilFool/web/api"
	"net/http"
)

func Start(port int) error {
	api := aapi.NewAPI("/api/")
	api.AddResource(new(aresource.TagResource))
	return http.ListenAndServe(fmt.Sprintf(":%d", port), api.Mux())
}
