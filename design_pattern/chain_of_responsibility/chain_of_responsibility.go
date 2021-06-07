package chain_of_responsibility

import (
	"container/list"
	"context"
	"errors"
)

type ChainResponsibility struct {
	*list.List
	ctx context.Context
}

func (c *ChainResponsibility) Run() error {
	handler := c.List.Front()
	for {
		if handler == nil {
			return nil
		}
		if err := handler.Value.(Handler).Handle(c.ctx); err != nil {
			return err
		}
	}
}

func NewChainResponsibility(ctx context.Context, handlers ...Handler) *ChainResponsibility {
	var c = &ChainResponsibility{
		list.New(),
		ctx,
	}
	for _, handler := range handlers {
		c.PushFront(handler)
	}
	return c
}

type Handler interface {
	Handle(ctx context.Context) error
}

type Handler1 struct{}

func (h Handler1) Handle(ctx context.Context) error {
	if !ctx.Value("switch").(bool) {
		return errors.New("switch should not false")
	}
	return nil
}

type Handler2 struct{}

func (h Handler2) Handle(ctx context.Context) error {
	panic("implement me")
}
