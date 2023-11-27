package migrations

import (
	"context"

	"gofr.dev/pkg/datastore"
	"gofr.dev/pkg/log"
)

type K20231126153059 struct {
}

func (k K20231126153059) Up(d *datastore.DataStore, logger log.Logger) error {
	logger.Info("Running migration up: 20231126153059_create_table_products.go")

	_, err := d.DB().ExecContext(context.Background(), createTableProducts)
	if err != nil {
		return err
	}

	return nil
}

func (k K20231126153059) Down(d *datastore.DataStore, logger log.Logger) error {
	logger.Info("Running migration up: 20231126153059_drop_table_products.go")

	_, err := d.DB().ExecContext(context.Background(), dropTableProducts)
	if err != nil {
		return err
	}

	return nil
}
