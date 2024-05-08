// Package bidding handles the AdBidding Server logic
package bidding

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBiddingService(t *testing.T) {
	Convey("Scenario - Bidding Sever receives HTTP request", t, func() {
		Convey("Given request handler", nil)

		Convey("Then service receives request", func() {
			Convey("When request contains AdPlacmentID", func() {
				Convey("Then service should respond with AdOpbject", func() {
					Convey("Then AdObject should contain AdID", nil)

					Convey("And  AdObject should contain BidPrice ( random for now )", func() {
						Convey("Then BidPrice should be in cents ( avoid floats for currency !!! )", nil)
					})

					Convey("And StatusOK (200)", nil)
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
}
