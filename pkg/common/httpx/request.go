package httpx

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//readRequest Return error if request is invalid
func readRequest(w http.ResponseWriter, r *http.Request, req interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		code, errs := buildInvalidRequestErrorsResponse(fmt.Sprintf("invalid request: %s", err.Error()))
		WriteErrorResponse(w, code, errs...)
		return err
	}
	return nil
}
