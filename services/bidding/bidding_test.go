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
				w := httptest.NewRecorder()

				Convey("When request contains AdPlacmentID", func() {
					query := url.Values{"ad_placement_id": {"1234-1234-1234"}}
					req := httptest.NewRequest(http.MethodGet, "/bid"+query.Encode(), nil)

					BidService(w, req)
					Convey("Then service should respond with JSON, AdOpbject", func() {
						adObject := new(AdObject)
						resp := w.Result()
						So(resp.Header.Get("Content-Type"), ShouldEqual, "application/json")

						err := json.NewDecoder(w.Body).Decode(&adObject)
						So(err, ShouldBeNil)
						// So(adObject, ShouldEqualJSON, "")
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
							So(resp.StatusCode, ShouldEqual, http.StatusOK)
						})
					})

					Convey("When service not interested in the spot", func() {
						Convey("Then return StatusNoContent (204)", nil)
					})
				})

				Convey("But the request doe's not contains AdPlacementID", func() {
					Convey("Then Status Forbidden", nil)
				})
			})
		})
	})
}
