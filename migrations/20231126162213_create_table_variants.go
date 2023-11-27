package migrations

import (
	"context"

	"gofr.dev/pkg/datastore"
	"gofr.dev/pkg/log"
)

type K20231126162213 struct {
}

func (k K20231126162213) Up(d *datastore.DataStore, logger log.Logger) error {
	logger.Info("Running migration up: 20231126162213_create_table_variants.go")

	_, err := d.DB().ExecContext(context.Background(), createTableVariants)
	if err != nil {
		return err
	}

	return nil
}

func (k K20231126162213) Down(d *datastore.DataStore, logger log.Logger) error {
	logger.Info("Running migration up: 20231126162213_drop_table_variants.go")

	_, err := d.DB().ExecContext(context.Background(), dropTableVariants)
	if err != nil {
		return err
	}

	return nil
}
