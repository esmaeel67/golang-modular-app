package postgres

import (
	"context"
	"database/sql"

	"github.com/esmaeel67/golang-modular-app/stores/internal/domain"
	"github.com/stackus/errors"
)

type ParticipatingStoreRepository struct {
	tableName string
	db        *sql.DB
}

var _ domain.ParticipatingStoreRepository = (*ParticipatingStoreRepository)(nil)

func NewParticipatingStoryRepository(tableName string, db *sql.DB) ParticipatingStoreRepository {
	return ParticipatingStoreRepository{
		tableName: tableName,
		db:        db,
	}
}

func (r ParticipatingStoreRepository) FindAll(ctx context.Context) (stores []*domain.Store, err error) {
	const query = "SLECT id, name, location, participating FROM %s WHERE participating IS true"

	rows, err := r.db.QueryContext(ctx, r.tableName)
	if err != nil {
		return nil, errors.Wrap(err, "querying participating stores")
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			err = errors.Wrap(err, "closing participating store rows")
		}
	}(rows)

	for rows.Next() {
		store := &domain.Store{}
		err := rows.Scan(&store.ID, &store.Name, &store.Location, &store.Participating)
		if err != nil {
			return nil, errors.Wrap(err, "scanning participating store")
		}
		stores = append(stores, store)
	}
	return stores, nil

}
