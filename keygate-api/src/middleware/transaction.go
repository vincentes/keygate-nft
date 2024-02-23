package kgmiddleware

import (
	"database/sql"

	"github.com/labstack/echo"
)

const (
	TxKey = "Tx"
)

func TransactionHandler(db *sql.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			tx, _ := db.Begin()

			c.Set(TxKey, tx)

			if err := next(c); err != nil {
				tx.Rollback()
				c.Logger().Error("Rollback: ", err)
				return err
			}

			tx.Commit()
			return nil
		})
	}
}