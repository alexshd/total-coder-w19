package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	. "github.com/alexshd/total-coder-w19/internal"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestAuctionDSPHappyPath(t *testing.T) {
	Convey("When the client sends request", t, func() {
		Convey("Then the request includes AdPlacementID", func() {
			req := httptest.NewRequest(http.MethodGet, "/api?ad_placement_id=THE-COOLEST-ID-ETHER", nil)
			w := httptest.NewRecorder()
			Convey("Given a Handler", func() {
				AuctionClientAPI(w, req)
				Convey("Then StatusOK (200)", func() {
					So(w.Result().StatusCode, ShouldEqual, http.StatusOK)
				})

				Convey("response body should contain JSON", func() {
					So(w.Result().Header.Get("content-type"), ShouldEqual, "application/json")
				})

				Convey("the body should include fields", func() {
					field1 := `"ad_placement_id":"THE-COOLEST-ID-ETHER"`
					field2 := `"price":`
					field3 := `"ad_link":`
					So(w.Body.String(), ShouldContainSubstring, field1)
					So(w.Body.String(), ShouldContainSubstring, field2)
					So(w.Body.String(), ShouldContainSubstring, field3)
				})

				Convey("And it should be in a struct", func() {
					r := &APIResponse{}
					Then(assert.NoError(t, json.NewDecoder(w.Body).Decode(r)))
					Convey("And the struct should have", func() {
						So(r.AdLink, ShouldNotBeEmpty)
						So(r.Price, ShouldBeBetween, 10, 3000)
						So(r.AdPlacementID, ShouldEqual, "THE-COOLEST-ID-ETHER")
					})
				})
			})
		})
	})
}

func TestAuctionServiceClientExposedAPI(t *testing.T) {
	t.Helper()
	Convey("Given publicly exposed API", t, func() {
		w := httptest.NewRecorder()
		Convey("Featureing client facing API for single request", func() {
			query := url.Values{"ad_placement_id": {"THE-COOLEST-ID-ETHER"}}
			req := httptest.NewRequest(http.MethodGet, "/bid?"+query.Encode(), nil)
			AuctionHandler(w, req)

			Convey("When handling the client request", func() {
				res := w.Result()
				decod := new(AuctionResponse)

				So(json.NewDecoder(w.Body).Decode(&decod), ShouldBeNil)
				Convey("Then the response is JSON", func() {
					So(res.Header.Get("content-type"), ShouldEqual, "application/json")
					So(decod.Status, ShouldEqual, "THE-COOLEST-ID-ETHER")
				})
			})
		})
		Reset(func() { w.Result().Body.Close() })

		Convey("On external client request", func() {
			Convey("When contains AdPlacementID", func() {
				Convey("Then \"FanOut\" ( optimize ) client request", func() {
					Convey("When On Success", func() {
						So(1, ShouldEqual, 1)
					})
				})
			})
		})
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
