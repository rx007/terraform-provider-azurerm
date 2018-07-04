package response

import (
	"net/http"
)

func WasConflict(resp *http.Response) bool {
	return responseWasStatusCode(resp, http.StatusConflict)
}

func WasNotFound(resp *http.Response) bool {
	return responseWasStatusCode(resp, http.StatusNotFound)
}

func WasBadRequest(resp *http.Response) bool {
	return responseWasStatusCode(resp, http.StatusBadRequest)
}

func WasOK(resp *http.Response) bool {
	return responseWasStatusCode(resp, http.StatusOK)
}

func WasNoContent(resp *http.Response) bool {
	return responseWasStatusCode(resp, http.StatusNoContent)
}

func responseWasStatusCode(resp *http.Response, statusCode int) bool {
	if r := resp; r != nil {
		if r.StatusCode == statusCode {
			return true
		}
	}

	return false
}
