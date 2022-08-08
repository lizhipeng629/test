package user

import (
	"context"
	"github.com/golang/mock/gomock"
	"test01/gomock/mocks"

	"testing"
)

func TestUse(t *testing.T) {
	ctx := context.Background()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDoer := mocks.NewMockDoer(mockCtrl)
	testUser := &User{Doer: mockDoer}
	user := User{}
	// 写好需要mock的函数的传参和返回值，当被测试的函数调用到这个方法时，就用这些值来代替。
	mockDoer.EXPECT().DoSomething(gomock.Any(), "Hello GoMock").Return(nil).Times(1)
	testUser.Use(ctx)
	testUser.testLow()
	user.testLow()

}
