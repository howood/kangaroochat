package requestid

import (
	"net/http"

	"github.com/howood/kangaroochat/infrastructure/uuid"
)

// KeyRequestID is XRequestId key
const KeyRequestID = "X-Request-Id"

func generateRequestID() string {
	return uuid.GetUUID(uuid.SatoriUUID)
}

// GetRequestID returns XRequestId
func GetRequestID(r *http.Request) string {
	if r.Header.Get(KeyRequestID) != "" {
		return r.Header.Get(KeyRequestID)
	}
	return generateRequestID()
}
