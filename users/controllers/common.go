package controllers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"user/users/models"
)

//renderJSON is use to render JSON in response
func renderJSON(w http.ResponseWriter, status int, res interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	if status == http.StatusNoContent {
		return
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("ERROR: renderJson - %q\n", err)
	}
}

//parseJSON function is use to parse JSON to model
func parseJSON(w http.ResponseWriter, body io.ReadCloser, model interface{}) bool {
	defer body.Close()

	b, _ := ioutil.ReadAll(body)
	err := json.Unmarshal(b, model)
	if err != nil {
		e := &models.ErrorData{}
		e.Message = "Error in parsing json"
		e.Err = err
		renderERROR(w, e)
		return false
	}

	return true
}

func renderERROR(w http.ResponseWriter, err *models.ErrorData) {
	err.Set()
	renderJSON(w, err.Code, err)
}

//getIDFromParams function is use to get ID from request
func getIDFromParams(w http.ResponseWriter, r *http.Request, key string) (int64, bool) {
	params, _ := r.Context().Value("params").(httprouter.Params)
	idStr := params.ByName(key)
	id, err := strconv.ParseInt(idStr, 10, 64)
	isErr := true

	if err != nil {
		isErr = false
		e := &models.ErrorData{}
		e.Message = "Invalid ID"
		e.Code = http.StatusBadRequest
		e.Err = err
		renderERROR(w, e)
	}

	return id, isErr
}
