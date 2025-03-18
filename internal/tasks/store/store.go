package store

import (
	"context"
	"sync"

	"gorm.io/gorm"
)

var (
	once sync.Once
	S    *datastore
)

type IStore interface {
	// DB 从传入的条件 对数据库示例进行筛选
	// 如果未传入任何条件 返回上下文中的数据库实例 - 事务实例 或者 核心实例.
	DB(ctx context.Context) *gorm.DB
	TX(ctx context.Context, fn func(ctx context.Context) error) error

	Task() TaskStore
	TaskScheduleCfg() TaskScheduleCfgStore
	TaskPos() TaskPosStore
}

// transactionKey is a value for context key.
type transactionKey struct{}

type datastore struct {
	coreDB *gorm.DB
}

func NewStore(db *gorm.DB) IStore {
	once.Do(func() {
		S = &datastore{db}
	})

	return S
}

// DB implements IStore.
func (ds *datastore) DB(ctx context.Context) *gorm.DB {
	db := ds.coreDB
	if tx, ok := ctx.Value(transactionKey{}).(*gorm.DB); ok {
		db = tx
	}

	return db
}

// TX implements IStore.
func (ds *datastore) TX(ctx context.Context, fn func(ctx context.Context) error) error {
	return ds.coreDB.WithContext(ctx).Transaction(
		func(tx *gorm.DB) error {
			ctx = context.WithValue(ctx, transactionKey{}, tx)
			return fn(ctx)
		},
	)
}

// Task implements IStore.
func (d *datastore) Task() TaskStore {
	panic("unimplemented")
}

// TaskPos implements IStore.
func (ds *datastore) TaskPos() TaskPosStore {
	return newTaskPosStore(ds)
}

// TaskScheduleCfg implements IStore.
func (d *datastore) TaskScheduleCfg() TaskScheduleCfgStore {
	panic("unimplemented")
}

var _ IStore = (*datastore)(nil)
