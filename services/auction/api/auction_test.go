package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestAuctionDSPHappyPath(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api?ad_placement_id=THE-COOLEST-ID-ETHER", nil)
	w := httptest.NewRecorder()
	AuctionClientAPI(w, req)

	t.Run("StatusOK & content-Type = application/json", func(t *testing.T) {
		if w.Result().StatusCode != http.StatusOK {
			t.Error("got: ", w.Result().StatusCode)
		}

		if w.Result().Header.Get("content-type") != "application/json" {
			t.Error("Content Type is not as expected")
		}
	})

	t.Run("Body should contain", func(t *testing.T) {
		field1 := `"ad_placement_id":"THE-COOLEST-ID-ETHER"`
		field2 := `"price":`
		field3 := `"ad_link":`
		body := w.Body.String()

		if !strings.Contains(body, field1) && !strings.Contains(body, field2) && !strings.Contains(body, field3) {
			t.Error("Body does not contain expected content", body)
		}
	})

	t.Run("Unmarshaled body should contain fields", func(t *testing.T) {
		r := &APIResponse{}

		if err := json.NewDecoder(w.Body).Decode(r); err != nil {
			t.Fatalf("Failed to decode body %v", err)
		}

		if r.AdLink == "" {
			t.Error("AdLink should not be empty")
		}

		if r.Price < 10 || r.Price > 3000 {
			t.Error("Price out of range")
		}

		if r.AdPlacementID != "THE-COOLEST-ID-ETHER" {
			t.Error("AdPlacementID not as expected:\n\t", "`r.AdPlacementID` != THE-COOLEST-ID-ETHER")
		}
	})
}

func TestAuctionServiceClientSingleRequest(t *testing.T) {
	w := httptest.NewRecorder()
	query := url.Values{"ad_placement_id": {"THE-COOLEST-ID-ETHER"}}
	req := httptest.NewRequest(http.MethodGet, "/bid?"+query.Encode(), nil)
	AuctionHandler(w, req)

	decod := new(AuctionResponse)

	t.Run("The response should be JSON", func(t *testing.T) {
		err := json.NewDecoder(w.Body).Decode(&decod)
		res := w.Result()
		contentType := res.Header.Get("content-type")
		if err != nil {
			t.Errorf("no error was expected... but: %v", err)
		}

		if contentType != "application/json" {
			t.Errorf("Expecting `application/json`... got: %s", contentType)
		}
	})
	t.Run("Body should include `THE-COOLEST-ID-ETHER`", func(t *testing.T) {
		decodeStatus := decod.Status
		if decodeStatus != "THE-COOLEST-ID-ETHER" {
			t.Error("got: ", decodeStatus, " wanted: THE-COOLEST-ID-ETHER")
		}
	})
}

func MakeBiddRequest(bidURI string, output chan<- any) {
	data := map[string]any{}
	client := http.DefaultClient
	req, err := http.NewRequest(http.MethodGet, bidURI, nil)
	if err != nil {
		close(output)
		return
	}
	go func() {
		resp, err := client.Do(req)
		if err != nil {
			return
		}

		_ = json.NewDecoder(resp.Body).Decode(&data)
		output <- data
	}()
}
