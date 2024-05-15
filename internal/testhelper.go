package internal

import (
	"github.com/smartystreets/goconvey/convey"
)

// ShouldPass a way to integrate `testify.assertion` with goconvey
func ShouldPass(actual any, expected ...any) string {
	if actual == true {
		return ""
	}
	return "suite test failed"
}

// Then rapper around So() for readability
func Then(assertion any) {
	convey.So(assertion, ShouldPass)
}
