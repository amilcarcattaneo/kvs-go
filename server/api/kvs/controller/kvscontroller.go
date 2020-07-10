package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"kvs-go/server/api/utils"
)

func (controller *Controller) GetValue(w http.ResponseWriter, r *http.Request) {
	key, ok := mux.Vars(r)["key"]
	if !ok || len(strings.TrimSpace(key)) == 0 {
		utils.HandleError(w, &utils.ApiError{
			Error: fmt.Errorf("missing 'key' url param"),
			Type:  http.StatusBadRequest,
		})
		return
	}
	keyvalue, err := controller.dao.GetValueByKey(key)
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&keyvalue)
	return
}

func (controller *Controller) SetValue(w http.ResponseWriter, r *http.Request) {
	key, ok := mux.Vars(r)["key"]
	if !ok || len(strings.TrimSpace(key)) == 0 {
		utils.HandleError(w, &utils.ApiError{
			Error: fmt.Errorf("missing 'key' url param"),
			Type:  http.StatusBadRequest,
		})
		return
	}
	value, ok := mux.Vars(r)["value"]
	if !ok || len(strings.TrimSpace(value)) == 0 {
		utils.HandleError(w, &utils.ApiError{
			Error: fmt.Errorf("missing 'value' url param"),
			Type:  http.StatusBadRequest,
		})
		return
	}

	if err := controller.dao.SetKeyValue(key, value); err != nil {
		utils.HandleError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&key)
	return
}
