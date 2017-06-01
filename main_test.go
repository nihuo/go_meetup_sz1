package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetHello(t *testing.T) {
	Convey("TestGetHello\n", t, func() {
		So(getHello(), ShouldEqual, "Hello world")
	})
}
