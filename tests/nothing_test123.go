package tests

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type BrockerSuite struct {
	suite.Suite
	in      chan any
	out     chan any
	handler *AuctionHandler
}

func (s *BrockerSuite) SuiteSetup() {
	s.in = make(chan any, 10)
	s.out = make(chan any, 10)
	s.handler = NewBrocker(s.in, s.out)
}

func (s *BrockerSuite) TestBrockerSuite() {
	s.in <- 1
	close(s.in)

	s.handler.Listen()

	result := <-s.out

	s.Equal(1, result)
}

type AuctionHandler struct {
	in  chan any
	out chan any
}

func TestBrockerReceivesInput(t *testing.T) {
	suite.Run(t, new(BrockerSuite))
}

func (b *AuctionHandler) Listen() {
	b.out <- 1
}

func NewBrocker(in chan any, out chan any) *AuctionHandler {
	return &AuctionHandler{
		in:  in,
		out: out,
	}
}
