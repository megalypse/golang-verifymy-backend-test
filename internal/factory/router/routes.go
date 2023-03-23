package factory

import (
	"log"
	"net/http"

	controllerFactory "github.com/megalypse/golang-verifymy-backend-test/internal/factory/controller"
	"github.com/megalypse/golang-verifymy-backend-test/internal/infra/middlewares"
)

func BootControllers() {
	router := GetRouter()

	controllers := controllerFactory.GetControllers()

	for _, controller := range controllers {
		for _, routeDefinition := range controller.GetHandlers() {
			handlingFunc := func() http.Handler {
				if routeDefinition.Unprotected {
					return routeDefinition.HandlingFunc
				} else {
					return middlewares.VerifyJwt(routeDefinition.HandlingFunc)
				}
			}().(http.HandlerFunc)
			route := routeDefinition.Route

			switch routeDefinition.Method {
			case http.MethodGet:
				router.Get(route, handlingFunc)
			case http.MethodPost:
				router.Post(route, handlingFunc)
			case http.MethodPut:
				router.Put(route, handlingFunc)
			case http.MethodPatch:
				router.Patch(route, handlingFunc)
			case http.MethodDelete:
				router.Delete(route, handlingFunc)
			default:
				log.Fatalf("Http method not supported: %q", routeDefinition.Method)
			}
		}
	}

}
