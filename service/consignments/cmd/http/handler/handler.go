package handler

import (
	"consignments/exceptions"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"consignments"
	"consignments/model"
)

type consignmentsHandler struct {
	service *consignments.Service
	routes  []struct {
		method  string
		path    string
		handler http.Handler
	}
}

func ConsignmentsHandler(service *consignments.Service) http.Handler {
	h := &consignmentsHandler{service: service}
	h.routes = []struct {
		method  string
		path    string
		handler http.Handler
	}{
		{http.MethodGet, "/", h.GetConsignments()},
		{http.MethodPost, "/", h.CreateConsignment()},
	}
	return h
}

func (h *consignmentsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	httpHandler := http.NotFoundHandler()
	for _, route := range h.routes {
		if r.Method == route.method && strings.EqualFold(r.URL.Path, route.path) {
			httpHandler = route.handler
			break
		}
	}
	httpHandler.ServeHTTP(w, r)
}

func (h *consignmentsHandler) GetConsignments() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		out, err := h.service.GetConsignments(r.Context())
		if err != nil {
			err := &exceptions.Error{Kind: exceptions.Internal, Inner: err}
			WriteError(w, err)
			log.Printf("failed to retrieve consignments, %v\n", err)
			return
		}
		writeJSON(w, http.StatusOK, out)
	}
}

func (h *consignmentsHandler) CreateConsignment() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			err := &exceptions.Error{Kind: exceptions.InvalidRequest, Inner: errors.New("method not allowed")}
			WriteError(w, err)
			log.Printf("failed process request, %v\n", err)
			return
		}
		buf, err := ioutil.ReadAll(r.Body)
		if err != nil {
			err := &exceptions.Error{Kind: exceptions.InvalidRequest, Inner: exceptions.MissingField("body")}
			WriteError(w, err)
			log.Printf("failed to read body, %v\n", err)
			return
		}
		payload := &model.Consignment{}
		err = json.Unmarshal(buf, payload)
		if err != nil {
			err := &exceptions.Error{Kind: exceptions.InvalidRequest, Inner: err}
			WriteError(w, err)
			log.Printf("failed to unmarshal json, %v\n", err)
			return
		}
		if payload.Weight == 0 {
			err := &exceptions.Error{Kind: exceptions.Validation, Param: "weight", Inner: exceptions.MissingField("weight")}
			log.Printf("failed validate body, %v\n", err)
			WriteError(w, err)
			return
		}
		resp, err := h.service.CreateConsignment(r.Context(), payload)
		if err != nil {
			// if it's something that the caller can handle such as a unique db
			// constraint return an appropriate error e.g. exceptions.Exist
			err := &exceptions.Error{Kind: exceptions.Internal, Inner: err}
			WriteError(w, err)
			log.Printf("failed to create consignment, %v\n", err)
			return
		}
		writeJSON(w, http.StatusAccepted, resp)
	}
}

func writeJSON(w http.ResponseWriter, code int, val interface{}) {
	buf, err := json.Marshal(val)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	w.Write(buf)
}
