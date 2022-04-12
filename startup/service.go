package startup

import (
	"net/http"

	"github.com/possawang/go-persist-lib-common/connection"
	"github.com/possawang/go-service-lib-common/routerutils"
	"gorm.io/gorm"
)

type PersistMainComponent struct {
	Endpoints map[string]routerutils.Endpoint
	Models    []gorm.Model
}

func StartingPersistService(component PersistMainComponent) {
	_, err := connection.ConnectionAndMigration(component.Models)
	if err != nil {
		panic(err)
	}
	routerutils.StartingService(component.Endpoints, func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { h.ServeHTTP(w, r) })
	})
}
