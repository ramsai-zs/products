// This is auto-generated file using 'gofr migrate' tool. DO NOT EDIT.
package migrations

import (
	dbmigration "gofr.dev/cmd/gofr/migration/dbMigration"
)

func All() map[string]dbmigration.Migrator {
	return map[string]dbmigration.Migrator{

		"20231126153059": K20231126153059{},
		"20231126162213": K20231126162213{},
	}
}
