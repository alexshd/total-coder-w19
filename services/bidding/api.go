package bidding

import (
	"encoding/json"
	"log/slog"
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

	switch {
	case VerifyMock(adPlacementID) > 10: // All Good
		slog.Info("GOOD", "verify", VerifyMock(adPlacementID))
		adObject := &AdObject{
			AdID:     "AD-ID",
			BidPrice: rand.IntN(1300) + 200,
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(adObject); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case VerifyMock(adPlacementID) < 10: // Regect
		slog.Info("REJECT", "verify", VerifyMock(adPlacementID))

		w.WriteHeader(http.StatusNoContent)
	}
}

func VerifyMock(s string) int {
	return len(s)
}
