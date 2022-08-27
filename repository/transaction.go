package repository

import (
	"challenge-q2pay/db"
	"challenge-q2pay/models"
	"fmt"
)

func CreateTransaction(transaction *models.Transaction) error {
	_, err := db.Conn.NamedExec(`INSERT INTO transactions (value, payer, payee, status)
		VALUES (:value, :payer, :payee, :status)`, transaction)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_ = db.Conn.MustBegin().Commit()
	return nil
}
