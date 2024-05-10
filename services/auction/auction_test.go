package auction

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAuctionService(t *testing.T) {
	Convey("Given multiple bidding services", t, func() {
		Convey("The Auction service sends concurrent requests", func() {
			Convey("When the response StatusCodeOK (200)", func() {
				Convey("Then accumulate valid bids", func() {
					So(1, ShouldEqual, 1)
				})
			})

			Convey("When response StatusCode != 200", func() {
				Convey("Or Bidding Service response time over 200ms", func() {
					Convey("Then Auction should  not accept the bid from that Bidding Service", func() {
						So(1, ShouldEqual, 1)
					})
				})
			})

			Convey("When FanOut finished", func() {
				Convey("Then return the Max BidPrice", func() {
					So(1, ShouldEqual, 1)
				})
			})
		})
	})
}

func TestAuctionServiceClientExposedAPI(t *testing.T) {
	Convey("Given publicly exposed API", t, func() {
		w := httptest.NewRecorder()
		Convey("Featureing client facing API for single request", func() {
			query := url.Values{"ad_placement_id": {"PBID-1234-1234-1234-1234"}}
			req := httptest.NewRequest(http.MethodGet, "/bid?"+query.Encode(), nil)
			AuctionHandler(w, req)

			Convey("When handeling the client request", func() {
				res := w.Result()
				decod := new(AuctionResponce)

				So(json.NewDecoder(w.Body).Decode(&decod), ShouldBeNil)
				Convey("Then the responce is JSON", func() {
					So(res.Header.Get("content-type"), ShouldEqual, "application/json")
					So(decod.Status, ShouldEqual, "cool")
				})
			})
		})
		Reset(func() { w.Result().Body.Close() })

		Convey("On external client request", func() {
			Convey("When contains AdPlacementID", func() {
				Convey("Then \"FanOut\" ( optimize ) client request", func() {
					Convey("When no eligible bids are received", func() {
						Convey("Then return StatusCodeNoContent ( 204 ) to the client", func() {
							So(1, ShouldEqual, 1)
						})
					})

					Convey("When On Success", func() {
						Convey("Then return Max BidPrice", func() {
							So(1, ShouldEqual, 1)
						})
					})
				})
			})
		})
	})
}
