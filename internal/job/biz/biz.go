package biz

import (
	bizv1 "github.com/chhz0/gojob/internal/job/biz/v1"
	"github.com/chhz0/gojob/internal/job/store"
)

type IBiz interface {
	UserV1() bizv1.TaskBiz
}

type biz struct {
	store store.IStore
}

func NewBiz(store store.IStore) IBiz {
	return &biz{store: store}
}

// UserV1 implements IBiz.
func (b *biz) UserV1() bizv1.TaskBiz {
	return bizv1.NewTaskBiz(b.store)
}
