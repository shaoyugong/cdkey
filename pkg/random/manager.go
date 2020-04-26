package random

import (
	"errors"
	"sync"
)

var BuilderManagerInstance = &builderManager{}

type builderManager struct {
	builderMap sync.Map
}

func (b *builderManager) Get(source string) (*builder, error) {
	if source == "" {
		return nil, errors.New("source can not be empty")
	}

	bd, ok := b.builderMap.Load(source)
	if ok {
		return bd.(*builder), nil
	}

	builder, err := NewBuilder(source)
	if err != nil {
		return nil, err
	}

	b.builderMap.Store(source, builder)
	return builder, nil
}