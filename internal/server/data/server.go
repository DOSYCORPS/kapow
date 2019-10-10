package data

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/BBVA/kapow/internal/server/model"
	"github.com/gorilla/mux"
)

var WriteSafe func(string, HandlerFunction) error = Handlers.WriteSafe

func setStatus(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	hID := vars["handler_id"]

	value, _ := ioutil.ReadAll(req.Body)
	status, _ := strconv.Atoi(string(value))

	var operation HandlerFunction = func(m *model.Handler) error {
		m.Writer.WriteHeader(status)
		return nil
	}

	err := WriteSafe(hID, operation)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
}

func setHeader(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	hID := vars["handler_id"]
	key := vars["key"]

	value, _ := ioutil.ReadAll(req.Body)
	header := string(value)

	var operation HandlerFunction = func(m *model.Handler) error {
		m.Writer.Header().Set(key, header)
		return nil
	}

	err := WriteSafe(hID, operation)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
}
