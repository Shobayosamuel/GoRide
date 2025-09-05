package http

import (
	// "context"
	"encoding/json"
	"log"
	"net/http"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/shared/types"
)

type HttpHandler struct {
	Service domain.TripService
}

func (s *HttpHandler) HandleTripPreview(w http.ResponseWriter, r *http.Request) {
	var reqBody previewTripRequest
	log.Printf("I am here")
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "failed to pass json data", http.StatusBadRequest)
		return
	}
	ctx := r.Context()
	t, err := s.Service.GetRoute(ctx, &reqBody.Pickup, &reqBody.Destination)
	log.Printf("Route: %v", t)
	if err != nil {
		log.Println(err)
	}
	writeJSON(w, http.StatusOK, t)
}

type previewTripRequest struct {
	UserID      string           `json:"userID"`
	Pickup      types.Coordinate `json:"pickup"`
	Destination types.Coordinate `json:"destination"`
}

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
