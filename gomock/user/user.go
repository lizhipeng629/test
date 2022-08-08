package user

import (
	"context"
	"test01/gomock/doer"
)

type User struct {
	Doer doer.Doer
}

func (user *User) Use(ctx context.Context) error {
	user.testLow()
	return user.Doer.DoSomething(123, "Hello GoMock")
}
