package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"consignments/exceptions"
)

func WriteError(w http.ResponseWriter, err error) {
	if err == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var e *exceptions.Error
	if errors.As(err, &e) {
		switch e.Kind {
		case exceptions.Unauthenticated:
			w.WriteHeader(http.StatusUnauthorized)
			return
		case exceptions.Unauthorized:
			w.WriteHeader(http.StatusForbidden)
			return
		default:
			writeError(w, e)
			return
		}
	}
	w.WriteHeader(http.StatusInternalServerError)
	return
}

func writeError(w http.ResponseWriter, error *exceptions.Error) {
	buf, _ := json.Marshal(exceptions.NewErrorResponse(error))
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(statusCode(error.Kind))
	fmt.Fprintf(w, string(buf))
}

func statusCode(kind exceptions.Kind) int {
	switch kind {
	case exceptions.Invalid, exceptions.Exist, exceptions.NotExist, exceptions.Private, exceptions.Validation, exceptions.InvalidRequest:
		return http.StatusBadRequest
	case exceptions.Other, exceptions.IO, exceptions.Internal, exceptions.Database, exceptions.Unanticipated:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
