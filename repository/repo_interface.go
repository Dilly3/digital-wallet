package repository

import (
	"github.com/dilly3/digital-wallet/models"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type DB interface {
	Getcustomer(accountNos string) (*models.Customer, error)
	Addcustomer(customer models.Customer) error
	Creditwallet(money *models.Money) (*models.Transaction, error)
	Debitwallet(money *models.Money) (*models.Transaction, error)
	Gettransaction(id string) (*[]models.Transaction, error)
	CreateTransaction(transaction *models.Transaction)
	InsufficientFunds(customer *models.Customer, debit *models.Money) bool
}

type DbParameters struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
}

func InitializeDbParameters() DbParameters {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	return DbParameters{
		Host:     host,
		User:     user,
		Password: password,
		Name:     dbName,
		Port:     port,
	}
}
