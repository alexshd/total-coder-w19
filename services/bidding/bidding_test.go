// Package bidding handles the AdBidding Server logic
package bidding

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBiddingService(t *testing.T) {
	Convey("Scenario - Bidding Sever receives HTTP request", t, func() {
		Convey("Given request handler", func() {
			Convey("Then service receives request", func() {
				Convey("When request contains AdPlacmentID", func() {
					w := httptest.NewRecorder()
					query := url.Values{"ad_placement_id": {"1234-1234-1234-1234"}}
					req := httptest.NewRequest(http.MethodGet, "/bid?"+query.Encode(), nil)

					BidService(w, req)
					Convey("Then service should respond with JSON, AdOpbject", func() {
						adObject := new(AdObject)
						resp := w.Result()
						So(resp.Header.Get("Content-Type"), ShouldEqual, "application/json")

						err := json.NewDecoder(w.Body).Decode(&adObject)
						So(err, ShouldBeNil)

						Convey("Then AdObject should contain AdID", func() {
							So(adObject.AdID, ShouldNotBeEmpty)
							So(adObject.AdID, ShouldHaveSameTypeAs, "string")
						})
						Convey("And  AdObject should contain BidPrice ( random for now )", func() {
							So(adObject.BidPrice, ShouldNotBeEmpty)
							Convey("Then BidPrice should be in cents ( avoid floats for currency !!! )", func() {
								So(adObject.BidPrice, ShouldHaveSameTypeAs, 111)
								So(adObject.BidPrice, ShouldBeBetween, 200, 1500)
							})
						})

						Convey("And StatusOK (200)", func() {
							So(w.Result().StatusCode, ShouldEqual, http.StatusOK)
						})
					})
					Reset(func() { w.Result().Body.Close() })
				})
				Convey("When service not interested in the spot", func() {
					w := httptest.NewRecorder()

					query := url.Values{"ad_placement_id": {"1234-1234"}}
					req := httptest.NewRequest(http.MethodGet, "/bid?"+query.Encode(), nil)

					BidService(w, req)

					Convey("Then return StatusNoContent (204)", func() {
						So(w.Result().StatusCode, ShouldEqual, http.StatusNoContent)
					})
					Reset(func() { w.Result().Body.Close() })
				})
			})
		})
	})
}

// ShouldPass a way to integrate `testify.assertion` with goconvey
func ShouldPass(actual any, expected ...any) string {
	if actual == true {
		return ""
	}
	return "suite test failed"
}

// Then rapper around So() for readability
func Then(assertion any) {
	So(assertion, ShouldPass)
}
