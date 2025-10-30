package api

import (
	"net/http"
)

func (h *Handler) DeployHandler(w http.ResponseWriter, r *http.Request) {
	log := h.Log
	log.Info("deploy endpoint called")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"deploy triggered"}`))
}
