package utils

import "github.com/reactivex/rxgo/v2"

type ObservableResult struct {
	Result rxgo.Observable
	Error  error
}
