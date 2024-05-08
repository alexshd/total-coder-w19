package auction

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type AuctionResponce struct {
	Status string `json:"status"`
}

func AuctionHandler(w http.ResponseWriter, r *http.Request) {
	adPlacementID := r.URL.Query().Get("ad_placement_id")
	slog.Info("recievd", "adPlacementID", adPlacementID)
	w.Header().Set("Content-Type", "application/json")

	auctionRes := &AuctionResponce{
		Status: "cool",
	}
	if err := json.NewEncoder(w).Encode(auctionRes); err != nil {
		http.Error(w, "UPSI", http.StatusInternalServerError)
	}
}
