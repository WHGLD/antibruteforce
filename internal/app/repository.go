package app

import (
	"context"
	"fmt"

	_ "github.com/jackc/pgx/stdlib" //nolint
	"github.com/jmoiron/sqlx"
)

type IPNetDb struct {
	Ip   string `db:"ip"`
	Mask string `db:"mask"`
}

type DBRepository struct {
	db *sqlx.DB
}

func NewRepository(dbConnection string) (*DBRepository, error) {
	db, err := sqlx.Open("pgx", dbConnection)
	if err != nil {
		return nil, err
	}

	return &DBRepository{db: db}, nil
}

func (repository *DBRepository) Close() error {
	return repository.db.Close()
}

func (repository *DBRepository) AddToBlackList(ctx context.Context, ip, mask string) error {
	query := `INSERT INTO blacklist (ip, mask) VALUES($1, $2)`
	_, err := repository.db.ExecContext(ctx, query, ip, mask)

	return err
}

func (repository *DBRepository) AddToWhiteList(ctx context.Context, ip, mask string) error {
	query := `INSERT INTO whitelist (ip, mask) VALUES($1, $2)`
	_, err := repository.db.ExecContext(ctx, query, ip, mask)

	return err
}

func (repository *DBRepository) RemoveFromBlackList(ctx context.Context, ip, mask string) error {
	_, err := repository.db.ExecContext(ctx, `DELETE FROM blacklist WHERE ip = $1 AND mask = $2`, ip, mask)

	return err
}

func (repository *DBRepository) RemoveFromWhiteList(ctx context.Context, ip, mask string) error {
	_, err := repository.db.ExecContext(ctx, `DELETE FROM whitelist WHERE ip = $1 AND mask = $2`, ip, mask)

	return err
}

func (repository *DBRepository) GetBlackList(ctx context.Context) ([]*IPNetDb, error) {
	return repository.getIPNetDbList(ctx, "SELECT ip, mask FROM blacklist")
}

func (repository *DBRepository) GetWhiteList(ctx context.Context) ([]*IPNetDb, error) {
	return repository.getIPNetDbList(ctx, "SELECT ip, mask FROM whitelist")
}

func (repository *DBRepository) getIPNetDbList(ctx context.Context, query string) ([]*IPNetDb, error) {
	res := make([]*IPNetDb, 0)
	rows, err := repository.db.QueryxContext(ctx, query)
	if err != nil {
		return res, err
	}
	defer func() {
		if err = rows.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	for rows.Next() {
		var n IPNetDb
		if err := rows.StructScan(&n); err != nil {
			return res, err
		}
		res = append(res, &n)
	}

	return res, err
}
