package auction

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

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

func TestAuctionServiceClientExposedAPI(t *testing.T) {
	Convey("Given publicly exposed API", t, func() {
		w := httptest.NewRecorder()
		Convey("Featureing client facing API for single request", func() {
			query := url.Values{"ad_placement_id": {"THE-COOLEST-ID-ETHER"}}
			req := httptest.NewRequest(http.MethodGet, "/bid?"+query.Encode(), nil)
			AuctionHandler(w, req)

			Convey("When handeling the client request", func() {
				res := w.Result()
				decod := new(AuctionResponce)

				So(json.NewDecoder(w.Body).Decode(&decod), ShouldBeNil)
				Convey("Then the responce is JSON", func() {
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
						expected := "THE-COOLEST-ID-ETHER"
						output := make(chan string, 4)
						list := Acum{}
						Convey("When running the function", func() {
							a := assert.New(t)
							for _, s := range []string{"1", "2", "3", "4", "5"} {
								input := httptest.NewRequest(http.MethodGet, "/bid?ad_placement_id=THE-COOLEST-ID-ETHER-"+s, nil)
								ProcessNumber(input, output)
							}
							for range 5 {
								select {
								case result := <-output:
									So(result, ShouldContainSubstring, expected)
									list.Lock()
									list.list = append(list.list, result)
									list.Unlock()
								case <-time.After(1 * time.Second):
									a.Fail("Test timed out")
								}
							}
							Convey("The max bid is", func() {
								So(max(list.list), ShouldContainSubstring, "THE-COOLEST-ID-ETHER-5")
							})
						})
					})
				})
			})
		})
	})
}

func ProcessNumber(r *http.Request, output chan<- string) {
	go func() {
		w := httptest.NewRecorder()
		AuctionHandler(w, r)
		output <- w.Body.String()
	}()
}
