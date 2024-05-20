package bidding

import (
	"encoding/json"
	"math/rand/v2"
	"net/http"
)

type AdObject struct {
	AdID     string `json:"ad_id"`
	BidPrice int    `json:"bid_amount"`
}

func BidService(w http.ResponseWriter, r *http.Request) {
	adPlacementID := r.URL.Query().Get("ad_placement_id")
	w.Header().Set("Content-Type", "application/json")

	// For 2 cases `if` enough ... It should be more than 2 cases ...
	switch {
	case VerifyMock(adPlacementID) > 10: // All Good
		adObject := &AdObject{
			AdID:     "AD-ID",
			BidPrice: rand.IntN(1300) + 200,
		}

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(adObject)

	case VerifyMock(adPlacementID) < 10: // Reject
		w.WriteHeader(http.StatusNoContent)
	}
}

func VerifyMock(s string) int {
	return len(s)
}
