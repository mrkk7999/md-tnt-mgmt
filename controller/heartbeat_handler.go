package controller

import "net/http"

// HeartBeatHandler
func (c *Controller) HeartBeatHandler(w http.ResponseWriter, r *http.Request) {
	response := c.svc.HeartBeat()
	encodeJSONResponse(w, http.StatusOK, response, nil)
}
