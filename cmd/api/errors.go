package main

import (
	"log"
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {

	log.Printf("internal server error: %s\npath: %s\n error %s\n", r.Method, r.URL.Path, err.Error())

	writeJSONError(w, http.StatusInternalServerError, "the server encountered a problem")
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {

	log.Printf("bad request error: %s\npath: %s\n error %s\n", r.Method, r.URL.Path, err.Error())

	writeJSONError(w, http.StatusBadRequest, err.Error())
}

func (app *application) conflictResponse(w http.ResponseWriter, r *http.Request, err error) {

	log.Printf("confict error: %s\npath: %s\n error %s\n", r.Method, r.URL.Path, err.Error())

	writeJSONError(w, http.StatusConflict, err.Error())
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {

	log.Printf("not found error: %s\npath: %s\n error %s\n", r.Method, r.URL.Path, err.Error())

	writeJSONError(w, http.StatusNotFound, "resource not found")
}
