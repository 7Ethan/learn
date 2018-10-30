package test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {
	Convey("Given some integer with a starting value", t, func() {
		x := 1
		Convey("Given some integer with a starting value", func() {
			x++
			Convey("The value shoud be greater by one", func() {
				So(x, ShouldEqual, 2)
			})
		})
	})
}
