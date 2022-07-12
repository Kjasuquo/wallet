package database

import (
	"errors"
	"github.com/kjasuquo/preparation/model"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func (d *dataB) CreateWallet(Wallet model.Wallet) error {
	statement, err := d.db.Prepare("INSERT INTO wallet(customer_name, balance) VALUES (?, ?)")
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = statement.Exec(Wallet.CustomerName, Wallet.Balance)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (d *dataB) GetWallet(id uint) (*model.Wallet, error) {
	rows, err := d.db.Query("SELECT * FROM wallet WHERE id =?", id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var wallet model.Wallet
	for rows.Next() {
		err := rows.Scan(&wallet.Id, &wallet.CustomerName, &wallet.Balance)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	return &wallet, nil
}

func (d *dataB) WalletTransaction(wallet *model.Wallet, Money model.AmountPaid) error {
	var Err = errors.New("insufficient fund")
	var Transaction *model.Transaction
	if Money.Type == "debit" {
		if Money.Amount > wallet.Balance {
			Transaction = &model.Transaction{
				WalletId: wallet.Id,
				Amount:   Money.Amount,
				Type:     Money.Type,
				Status:   "Not Successful",
			}

			statement, err := d.db.Prepare("INSERT INTO trans (walletId, amount, type, status) VALUES (?, ?, ?, ?)")
			if err != nil {
				log.Println(err)
				return err
			}
			_, err = statement.Exec(Transaction.WalletId, Transaction.Amount, Transaction.Type, Transaction.Status)
			if err != nil {
				log.Println(err)
				return err
			}

			return Err

		} else if Money.Amount <= wallet.Balance {
			Transaction = &model.Transaction{
				WalletId: wallet.Id,
				Amount:   Money.Amount,
				Type:     Money.Type,
				Status:   "Successful",
			}
		}
		wallet.Balance = wallet.Balance - Transaction.Amount
	} else if Money.Type == "credit" {
		Transaction = &model.Transaction{
			WalletId: wallet.Id,
			Amount:   Money.Amount,
			Type:     Money.Type,
			Status:   "Successful",
		}
		wallet.Balance = wallet.Balance + Transaction.Amount
	}

	statement, err := d.db.Prepare("INSERT INTO trans (walletId, amount, type, status) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = statement.Exec(Transaction.WalletId, Transaction.Amount, Transaction.Type, Transaction.Status)
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = d.db.Exec("update wallet set balance = ? where id = ?", wallet.Balance, wallet.Id)
	if err != nil {
		return err
	}

	return nil
}
