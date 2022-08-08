package user_test

import (
	"context"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"test01/gomock/mocks"
	"test01/gomock/user"
)

var _ = Describe("UserGingo", func() {
	ctx := context.Background()
	// 获取mockcontroller
	ctrl := gomock.NewController(GinkgoT())

	It("Test1Use", func() {
		// 根据controller获取mockDoer,mockDoer可以进行调用mock的函数
		mockDoer := mocks.NewMockDoer(ctrl)
		mockDoer.EXPECT().DoSomething(123, gomock.Any()).Return(nil).AnyTimes()

		// 构造可以调用user方法的结构体对象 使用类似于多态进行赋值
		u := &user.User{Doer: mockDoer}
		err := u.Use(ctx)
		Expect(err).To(BeNil())
	})

	It("Test2Use", func() {
		// 根据controller获取mockDoer,mockDoer可以进行调用mock的函数
		mockDoer := mocks.NewMockDoer(ctrl)
		mockDoer.EXPECT().DoSomething(123, gomock.Any()).Return(nil).AnyTimes()

		// 构造可以调用user方法的结构体对象 使用类似于多态进行赋值
		u := &user.User{Doer: mockDoer}
		err := u.Use(ctx)
		Expect(err).To(BeNil())
	})
})
