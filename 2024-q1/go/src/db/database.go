package db

import (
	"context"
	"database/sql"
	"database/sql/driver"

	"github.com/tailscale/sqlite"

	"backend-brawl/src/model"

	"log/slog"
)

type Store struct {
	DB *sql.DB
}

func (s *Store) InitDatabase(sqliteURI string) {

	connInitFunc := func(ctx context.Context, conn driver.ConnPrepareContext) error {
		return sqlite.ExecScript(conn.(sqlite.SQLConn), "PRAGMA journal_mode=WAL;")
	}
	connector := sqlite.Connector(sqliteURI, connInitFunc, nil)

	db := sql.OpenDB(connector)
	s.DB = db

	// defer db.Close()
}

func (s *Store) CreateSchemas() {
	slog.Info("Creating schemas")

	tx, err := s.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()

	_, err = tx.Exec("CREATE TABLE IF NOT EXISTS customers (id INTEGER PRIMARY KEY, acc_limit INTEGER, acc_balance INTEGER)")
	if err != nil {
		panic(err)
	}

	tx.Commit()
}

func (s *Store) PopulateDatabase() {
	customers := []model.Customer{
		*model.NewCustomer(1, 100000),
		*model.NewCustomer(2, 80000),
		*model.NewCustomer(3, 1000000),
		*model.NewCustomer(4, 10000000),
		*model.NewCustomer(5, 500000),
	}
	for _, c := range customers {
		s.StoreCustomer(&c)
	}
}

func (s *Store) StoreCustomer(c *model.Customer) {
	slog.Info("Storing customer", slog.Int("id", c.ID), slog.Int64("acc_limit", int64(c.Limit)), slog.Int64("acc_balance", int64(c.Balance)))

	tx, err := s.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()

	_, err = tx.Exec("INSERT INTO customers (id, acc_limit, acc_balance) VALUES (?, ?, ?)", c.ID, c.Limit, c.Balance)
	if err != nil {
		panic(err)
	}

	slog.Info("Stored customer. Now committing transaction")

	tx.Commit()
}

func (s *Store) GetCustomers() []model.Customer {
	slog.Info("Getting customers")

	rows, err := s.DB.Query("SELECT id, acc_limit, acc_balance FROM customers")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	customers := []model.Customer{}
	for rows.Next() {
		var c model.Customer
		err := rows.Scan(&c.ID, &c.Limit, &c.Balance)
		if err != nil {
			panic(err)
		}
		customers = append(customers, c)
	}

	return customers

}

func PingDatabase() {
	sqliteURI := "file:sqlite.db?cache=shared&mode=rwc"

	connInitFunc := func(ctx context.Context, conn driver.ConnPrepareContext) error {
		return sqlite.ExecScript(conn.(sqlite.SQLConn), "PRAGMA journal_mode=WAL;")
	}
	connector := sqlite.Connector(sqliteURI, connInitFunc, nil)

	db := sql.OpenDB(connector)
	defer db.Close()

	db.Ping()
}
