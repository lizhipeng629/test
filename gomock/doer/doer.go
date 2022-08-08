package doer

//go:generate mockgen -destination=../mocks/mock_doer.go -package=mocks test01/gomock/doer Doer

type Doer interface {
	DoSomething(int, string) error
}
