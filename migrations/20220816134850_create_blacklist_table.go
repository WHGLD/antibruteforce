package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateBlacklistTable, downCreateBlacklistTable)
}

func upCreateBlacklistTable(tx *sql.Tx) error {
	_, err := tx.Exec(
		`CREATE TABLE blacklist (
			ip inet NOT NULL,
			mask inet NOT NULL,
			PRIMARY KEY (ip, mask));
	`)
	if err != nil {
		return err
	}

	return nil
}

func downCreateBlacklistTable(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE blacklist;")
	if err != nil {
		return err
	}

	return nil
}
