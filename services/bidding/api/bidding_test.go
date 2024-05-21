// Package bidding handles the AdBidding Server logic
package bidding

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

func TestBiddingService(t *testing.T) {
	t.Run("Happy path", func(t *testing.T) {
		w := httptest.NewRecorder()
		query := url.Values{"ad_placement_id": {"1234-1234-1234-1234"}}
		req := httptest.NewRequest(http.MethodGet, "/bid?"+query.Encode(), nil)

		BidService(w, req)

		adObject := new(AdObject)
		resp := w.Result()

		if resp.Header.Get("Content-Type") != "application/json" {
			t.Error("Wrong Content-Type")
		}

		if err := json.NewDecoder(w.Body).Decode(&adObject); err != nil {
			t.Fatalf("Could not decode to JSON\n\t %s", w.Body.String())
		}

		if w.Result().StatusCode != http.StatusOK {
			t.Error("StatusOK is expected => got:\n\t", w.Result().StatusCode)
		}

		if adObject.AdID == "" {
			t.Error("expected not empty field")
		}
		if reflect.TypeOf(adObject.AdID) != reflect.TypeOf("string") {
			t.Error("expected field type `string`, got:\n\t", reflect.TypeOf(adObject.AdID))
		}

		if adObject.BidPrice < 200 && adObject.BidPrice > 1500 {
			t.Error("BidPrice out of range")
		}
	})

	t.Run("On BidService Rejection Should receive StatusNoContent", func(t *testing.T) {
		w := httptest.NewRecorder()
		query := url.Values{"ad_placement_id": {"1234-1234"}}
		req := httptest.NewRequest(http.MethodGet, "/bid?"+query.Encode(), nil)

		BidService(w, req)

		got := w.Result().StatusCode

		if got != http.StatusNoContent {
			t.Error("Expected StatusNoContent, got:\n\t", got)
		}
	})
}
