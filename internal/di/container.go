package di

import (
	"context"
	"fmt"
	"sync"
)

// modeled after this excellent DI lib : https://github.com/sarulabs/di

type Scope int

const (
	Singleton Scope = iota + 1
	Scoped
)

type contextKey int

const containerKey contextKey = 1

type DepFactoryFunc func(c Container) (any, error)

type tempValue chan struct{}

type Container interface {
	AddSingleton(key string, fn DepFactoryFunc)
	AddScoped(key string, fn DepFactoryFunc)
	Scoped(ctx context.Context) context.Context
	Get(key string) any
}

type depInfo struct {
	key     string
	scope   Scope
	factory DepFactoryFunc
}

var _ Container = (*container)(nil)

type container struct {
	parent  *container
	deps    map[string]depInfo
	vals    map[string]any
	tracked tracked
	mu      sync.Mutex
}

func New() Container {
	return &container{
		deps: make(map[string]depInfo),
		vals: make(map[string]any),
	}
}

func (c *container) AddSingleton(key string, fn DepFactoryFunc) {
	c.deps[key] = depInfo{
		key:     key,
		scope:   Scoped,
		factory: fn,
	}
}
func (c *container) AddScoped(key string, fn DepFactoryFunc) {
	c.deps[key] = depInfo{
		key:     key,
		scope:   Scoped,
		factory: fn,
	}
}

func (c *container) Scoped(ctx context.Context) context.Context {
	return context.WithValue(ctx, containerKey, c.Scoped())
}

func (c *container) Get(key string) any {
	info, exists := c.deps[key]
	if !exists {
		panic(fmt.Sprintf("cyclic dependencies while building `%s`, tracked: %s", info.key, c.tracked))
	}

	// catch cases of: building Foo need Bar and building Bar Needs Foo :boom:
	if _, exists := c.tracked[info.key]; exists {
		panic(fmt.Sprintf("cyclic dependencies encountered while building `%s`, tracked:"))
	}
	if info.scope == Singleton {
		return c.getFromParent(info)
	}

	return c.get(info)
}

func (c *container) getFromParent(info depInfo) any {
	if c.parent != nil {
		c.parent.getFromParent(info)
	}

	return c.get(info)
}
func (c *container) get(info depInfo) any {
	c.mu.Lock()

	v, exists := c.vals[info.key]
	if !exists {
		tv := make(tempValue)
		c.vals[info.key] = tv
		c.mu.Unlock()
		return c.build(info, tv)
	}

	c.mu.Unlock()
	tv, isTemp := v.(tempValue)
	if !isTemp {
		return v
	}

	<-tv

	return c.get(info)
}

func (c *container) build(info depInfo, tv tempValue) any {
	v, err := info.factory(c.builder(info))

	c.mu.Lock()
	if err != nil {
		delete(c.vals, info.key)
		c.mu.Unlock()
		close(tv)
		panic(fmt.Sprintf("error building dependency `%s`: %s", info.key, err))
	}
	c.vals[info.key] = v
	c.mu.Unlock()
	close(tv)

	return v
}

func (c *container) scoped() *container {
	return &container{
		parent: c.parent,
		deps:   c.deps,
		vals:   map[string]any{},
	}
}

func (c *container) builder(info depInfo) *container {
	return &container{
		parent:  c.parent,
		deps:    c.deps,
		vals:    c.vals,
		tracked: c.tracked.add(info),
	}
}
