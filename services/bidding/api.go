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
	adObject := &AdObject{
		AdID:     "AD-ID",
		BidPrice: rand.IntN(1500) + 200,
	}

	w.Header().Set("content-type", "application/json")

	if err := json.NewEncoder(w).Encode(adObject); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
