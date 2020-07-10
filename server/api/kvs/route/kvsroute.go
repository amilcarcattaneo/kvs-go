package route

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	"kvs-go/server/api/kvs/controller"
)

func Init(router *mux.Router, db *gorm.DB) {
	newcontroller := controller.NewController(db)

	router.HandleFunc("/key/{key}", newcontroller.GetValue).Methods(http.MethodGet)
	router.HandleFunc("/key/{key}/value/{value}", newcontroller.SetValue).Methods(http.MethodPost)
}
