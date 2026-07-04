package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/esmaeel67/golang-modular-app/internal/ddd"
	"github.com/esmaeel67/golang-modular-app/internal/es"
	"github.com/esmaeel67/golang-modular-app/internal/registry"
)

type (
	EventStore struct {
		tableName string
		db        *sql.DB
		registry  registry.Registry
	}

	aggregateEvent struct {
		id         string
		name       string
		payload    ddd.EventPayload
		occurredAt time.Time
		aggregate  es.EventSourcedAggregate
		version    int
	}
)

var _ es.AggregateStore = (*EventStore)(nil)
var _ ddd.AggregateEvent = (*aggregateEvent)(nil)

func (s EventStore) Load(ctx context.Context, aggregate es.EventSourcedAggregate) (err error) {
	err = nil

	return
}

func (s EventStore) Save(ctx context.Context, aggregate es.EventSourcedAggregate) (err error) {

	return
}

func (s EventStore) table(query string) string {
	return fmt.Sprintf(query, s.tableName)
}

func (e aggregateEvent) ID() string                { return e.id }
func (e aggregateEvent) EventName() string         { return e.name }
func (e aggregateEvent) Payload() ddd.EventPayload { return e.payload }
func (e aggregateEvent) Metadata() ddd.Metadata    { return ddd.Metadata{} }
func (e aggregateEvent) OccurredAt() time.Time     { return e.occurredAt }
func (e aggregateEvent) AggregateName() string     { return e.aggregate.AggregateName() }
func (e aggregateEvent) AggregateID() string       { return e.aggregate.ID() }
func (e aggregateEvent) AggregateVersion() int     { return e.version }
