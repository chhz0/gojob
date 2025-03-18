package bizv1

import "github.com/chhz0/gojob/internal/job/store"

type TaskBiz interface {
}

type taskBiz struct {
	store store.IStore
}

func NewTaskBiz(s store.IStore) TaskBiz {
	return &taskBiz{s}
}
