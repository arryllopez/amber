// Handles environment profile HTTP requests.
package environment

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return Handler{
		service: service,
	}
}

func (h Handler) Preview(w http.ResponseWriter, r *http.Request) {
	var request PreviewRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON request body")
		return
	}

	response, err := h.service.Preview(request)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, response)
}

func (h Handler) PreviewByPath(w http.ResponseWriter, r *http.Request) {
	latitude, err := strconv.ParseFloat(chi.URLParam(r, "lat"), 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "lat must be a valid number")
		return
	}

	longitude, err := strconv.ParseFloat(chi.URLParam(r, "lng"), 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "lng must be a valid number")
		return
	}

	response, err := h.service.Preview(PreviewRequest{
		Latitude:  latitude,
		Longitude: longitude,
		EraID:     chi.URLParam(r, "era"),
	})
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, response)
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{
		"error": message,
	})
}
