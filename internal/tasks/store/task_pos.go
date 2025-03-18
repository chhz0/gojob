package store

import (
	"context"

	modelv1 "github.com/chhz0/gotasks/internal/tasks/model/v1"
	"github.com/chhz0/gotasks/internal/pkg/errcode"
	"github.com/chhz0/gokit/pkg/log"
)

type TaskPosStore interface {
	Create(ctx context.Context, obj *modelv1.TaskPos) error
	Update(ctx context.Context, obj *modelv1.TaskPos) error
	Delete(ctx context.Context, obj *modelv1.TaskPos) error
	Get(ctx context.Context, obj *modelv1.TaskPos) (*modelv1.TaskPos, error)
	List(ctx context.Context, obj *modelv1.TaskPos) (int64, []*modelv1.TaskPos, error)

	TaskExpansion
}

type TaskExpansion interface {
	GetNextPos(taskType string) error
}

type taskPosStore struct {
	store *datastore
}

func newTaskPosStore(store *datastore) TaskPosStore {
	return &taskPosStore{
		store: store,
	}
}

// Delete implements TaskPosStore.
func (t *taskPosStore) Delete(ctx context.Context, obj *modelv1.TaskPos) error {
	panic("unimplemented")
}

// Get implements TaskPosStore.
func (t *taskPosStore) Get(ctx context.Context, obj *modelv1.TaskPos) (*modelv1.TaskPos, error) {
	panic("unimplemented")
}

// List implements TaskPosStore.
func (t *taskPosStore) List(ctx context.Context, obj *modelv1.TaskPos) (int64, []*modelv1.TaskPos, error) {
	panic("unimplemented")
}

// Update implements TaskPosStore.
func (t *taskPosStore) Update(ctx context.Context, obj *modelv1.TaskPos) error {
	if err := t.store.DB(ctx).Save(obj).Error; err != nil {
		log.Errorw("Failed to update task-pos in database", "err", err, "taskPos", obj)
		return errcode.ErrDBWrite.WithMessage("failed to update task pos: %v", err)
	}
	return nil
}

// Create implements TaskPosStore.
func (t *taskPosStore) Create(ctx context.Context, obj *modelv1.TaskPos) error {
	if err := t.store.DB(ctx).Create(&obj).Error; err != nil {

		return errcode.ErrDBWrite.WithMessage("Failed to create task pos: %v", err)
	}
	return nil
}

// GetNextPos implements TaskPosStore.
func (t *taskPosStore) GetNextPos(taskType string) error {
	panic("unimplemented")
}
