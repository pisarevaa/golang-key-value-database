package api

import (
	"net/http"
)

//	@Summary	Get value by key
//	@Tags		Database
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	map[string]string
//	@Failure	409	{object}	ErrorResponse
//	@Router		/v1/key/{key} [get].
func (h *Handler) GetValue(w http.ResponseWriter, r *http.Request) {
	// err := h.Service.UpdateDatabase(r.Context())
	// if err != nil {
	// 	log.Error().Err(err).Msg("database sync failed")
	// 	writeError(w, "Failed to start sync", http.StatusInternalServerError)
	// 	return
	// }
	writeJSON(w, map[string]string{"status": "database sync completed successfully"})
}

//	@Summary	Set value by key
//	@Tags		Database
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	map[string]string
//	@Failure	409	{object}	ErrorResponse
//	@Router		/v1/key/{key} [post].
func (h *Handler) SetValue(w http.ResponseWriter, r *http.Request) {
	// err := h.Service.UpdateDatabase(r.Context())
	// if err != nil {
	// 	log.Error().Err(err).Msg("database sync failed")
	// 	writeError(w, "Failed to start sync", http.StatusInternalServerError)
	// 	return
	// }
	writeJSON(w, map[string]string{"status": "database sync completed successfully"})
}

//	@Summary	Delete value by key
//	@Tags		Database
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	map[string]string
//	@Failure	409	{object}	ErrorResponse
//	@Router		/v1/key/{key} [delete].
func (h *Handler) DeleteValue(w http.ResponseWriter, r *http.Request) {
	// err := h.Service.UpdateDatabase(r.Context())
	// if err != nil {
	// 	log.Error().Err(err).Msg("database sync failed")
	// 	writeError(w, "Failed to start sync", http.StatusInternalServerError)
	// 	return
	// }
	writeJSON(w, map[string]string{"status": "database sync completed successfully"})
}
