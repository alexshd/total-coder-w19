package auction

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/suite"
)

type MultiBiddingRequestSuite struct {
	suite.Suite
	bidSerivces chan any
}

func (s *MultiBiddingRequestSuite) SetupSuite() {
	s.bidSerivces = make(chan any)
}

func (s *MultiBiddingRequestSuite) SetupSubTest() {
}

func (s *MultiBiddingRequestSuite) TestDataReceived() {
	s.Run("TheTest", func() {
		s.Run("SubTest", func() {
		})
	})
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

func (s *MultiBiddingRequestSuite) TestAuctionServiceClientExposedAPI() {
	Convey("Given publicly exposed API", s.T(), func() {
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

func TestMultiBidding(t *testing.T) {
	suite.Run(t, new(MultiBiddingRequestSuite))
}

func TestRacer(t *testing.T) {
	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)
		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		server := makeDelayedServer(25 * time.Millisecond)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
