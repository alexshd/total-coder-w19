package auction

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAuctionService(t *testing.T) {
	Convey("Given multiple bidding services", t, func() {
		Convey("The Auction service sends concurrent requests", func() {
			Convey("When the response StatusCodeOK (200)", func() {
				Convey("Then accumulate valid bids", nil)
			})

			Convey("When response StatusCode != 200", func() {
				Convey("Or Bidding Service response time over 200ms", func() {
					Convey("Then Auction should  not accept the bid from that Bidding Service", nil)
				})
			})

			Convey("When FanOut finished", func() {
				Convey("Then return the Max BidPrice", nil)
			})
		})
	})
}

func TestAuctionServiceClientExposedAPI(t *testing.T) {
	Convey("Given publicly exposed API", t, func() {
		Convey("On external client request", func() {
			Convey("When contains AdPlacementID", func() {
				Convey("Then \"FanOut\" ( optimize ) client request", func() {
					Convey("When no eligible bids are received", func() {
						Convey("Then return StatusCodeNoContent ( 204 ) to the client", nil)
					})

					Convey("When On Success", func() {
						Convey("Then return Max BidPrice", nil)
					})
				})
			})
		})
	})
}
