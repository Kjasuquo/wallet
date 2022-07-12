package database

import (
	"database/sql"
	"fmt"
	"github.com/kjasuquo/preparation/model"
	_ "github.com/mattn/go-sqlite3"
)

type Database interface {
	CreateWallet(Wallet model.Wallet) error
	GetWallet(id uint) (*model.Wallet, error)
	WalletTransaction(wallet *model.Wallet, Money model.AmountPaid) error
}

type dataB struct {
	db *sql.DB
}

func NewDB() *dataB {
	return &dataB{}
}

func (d *dataB) SetupDB(dbName string) {
	db, err := sql.Open("sqlite3", fmt.Sprintf("./%s.db", dbName))
	if err != nil {
		fmt.Println(err)
	}

	err = createWallet(db)
	if err != nil {
		return
	}
	err = createTransactionTable(db)
	if err != nil {
		return
	}

	d.db = db
}

func createWallet(db *sql.DB) error {
	statement, err := db.Prepare(
		"CREATE TABLE IF NOT EXISTS wallet (id INTEGER PRIMARY KEY, customer_name VARCHAR (50), balance INTEGER )")
	if err != nil {
		return err
	}

	_, err = statement.Exec()
	if err != nil {
		return err
	}

	return nil
}

func createTransactionTable(db *sql.DB) error {
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS trans (id INTEGER PRIMARY KEY, walletId INTEGER, amount INTEGER, type VARCHAR (50), status VARCHAR (50))")
	if err != nil {
		return err
	}

	_, err = statement.Exec()
	if err != nil {
		return err
	}

	return nil
}
