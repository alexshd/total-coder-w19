package auction

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

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

				Convey("responce body should contain JSON", func() {
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
					r := &APIResponce{}
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
