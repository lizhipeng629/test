package test_study

import "testing"
import . "github.com/smartystreets/goconvey/convey"

func TestAdd(t *testing.T) {
	if testing.Short() {
		t.Skip("Short模式下会跳过该测试用例")
	}
	Convey("将两数相加", t, func() {
		So(Add(1, 2), ShouldEqual, 3)
	})

}

func TestSub(t *testing.T) {
	Convey("将两数相减", t, func() {
		So(Sub(1, 2), ShouldEqual, -1)
	})
}

func TestMultiply(t *testing.T) {
	Convey("将两数相减乘", t, func() {
		So(Multiply(1, 2), ShouldEqual, 2)
	})
}

func TestDivision(t *testing.T) {
	Convey("将两数相除", t, func() {
		Convey("被除数为0", func() {
			_, err := Division(10, 0)
			So(err, ShouldNotBeNil)
		})
		Convey("被除数不为0", func() {
			num, err := Division(10, 2)
			So(num, ShouldEqual, 5)
			So(err, ShouldBeNil)
		})
	})
}
