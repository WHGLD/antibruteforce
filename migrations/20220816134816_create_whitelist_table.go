package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateWhitelistTable, downCreateWhitelistTable)
}

func upCreateWhitelistTable(tx *sql.Tx) error {
	_, err := tx.Exec(
		`CREATE TABLE whitelist (
			ip inet NOT NULL,
			mask inet NOT NULL,
			PRIMARY KEY (ip, mask));
	`)
	if err != nil {
		return err
	}

	return nil
}

func downCreateWhitelistTable(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE whitelist;")
	if err != nil {
		return err
	}

	return nil
}
