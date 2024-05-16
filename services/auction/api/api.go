package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type APIResponce struct {
	AdPlacementID string `json:"ad_placement_id"`
	AdLink        string `json:"ad_link"`
	Price         int    `json:"price"`
}

func AuctionClientAPI(w http.ResponseWriter, r *http.Request) {
	responce := &APIResponce{
		AdPlacementID: r.URL.Query().Get("ad_placement_id"),
		Price:         1234,
		AdLink:        "http://some.link.to.the.ad",
	}

	w.Header().Set("content-type", "application/json")

	if err := json.NewEncoder(w).Encode(responce); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type AuctionResponce struct {
	Status string `json:"status"`
}

func AuctionHandler(w http.ResponseWriter, r *http.Request) {
	adPlacementID := r.URL.Query().Get("ad_placement_id")
	w.Header().Set("Content-Type", "application/json")

	auctionRes := &AuctionResponce{
		Status: adPlacementID,
	}
	if err := json.NewEncoder(w).Encode(auctionRes); err != nil {
		http.Error(w, "UPSI", http.StatusInternalServerError)
	}
}

type Acum struct {
	list []any
	sync.RWMutex
}

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}