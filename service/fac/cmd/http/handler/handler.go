package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"fac"
)

type FactorialHandler struct {
}

func (h *FactorialHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	part := r.URL.Path[1:]
	num, err := strconv.Atoi(part)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintln(w, fac.Factorial(num))
}
