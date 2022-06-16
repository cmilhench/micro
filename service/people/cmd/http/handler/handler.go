package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"people"
	"people/model"
)

type peopleHandler struct {
	service *people.Service
	routes  []struct {
		method  string
		path    string
		handler http.Handler
	}
}

func PeopleHandler(service *people.Service) http.Handler {
	h := peopleHandler{service: service}
	h.routes = []struct {
		method  string
		path    string
		handler http.Handler
	}{
		{http.MethodGet, "/", h.GetPeople()},
		{http.MethodPost, "/", h.AddPerson()},
	}
	return h
}

func (h peopleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	httpHandler := http.NotFoundHandler()
	for _, route := range h.routes {
		if r.Method == route.method && strings.EqualFold(r.URL.Path, route.path) {
			httpHandler = route.handler
			break
		}
	}
	httpHandler.ServeHTTP(w, r)
}

func (h *peopleHandler) GetPeople() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		out, err := h.service.GetPeople()
		if err != nil {
			writeText(w, http.StatusBadRequest, err.Error())
			log.Printf("failed to read people: %v\n", err)
			return
		}
		writeJSON(w, http.StatusOK, out)
	}
}

func (h *peopleHandler) AddPerson() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		buf, err := ioutil.ReadAll(req.Body)
		if err != nil {
			writeText(w, http.StatusBadRequest, err.Error())
			log.Printf("failed to read body: %v\n", err)
			return
		}
		payload := &model.Person{}
		err = json.Unmarshal(buf, payload)
		if err != nil {
			writeText(w, http.StatusBadRequest, err.Error())
			log.Printf("failed to unmarshal json: %v\n", err)
			return
		}
		resp, err := h.service.AddPerson(payload)
		if err != nil {
			writeText(w, http.StatusBadRequest, err.Error())
			log.Printf("failed to add Person: %v\n", err)
			return
		}
		writeJSON(w, http.StatusAccepted, resp)
	}
}

func writeText(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	fmt.Fprintln(w, msg)
}

func writeJSON(w http.ResponseWriter, code int, val interface{}) {
	var buf []byte
	var err error

	buf, err = json.Marshal(val)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	w.Write(buf)
}
