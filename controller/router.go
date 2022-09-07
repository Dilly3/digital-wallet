package controller

import (
	"log"
	"os"

	repo "github.com/dilly3/digital-wallet/repository"
	"github.com/gin-gonic/gin"
)

func SetupRouter(h *Handler) (*gin.Engine, string) {
	router := gin.Default()
	router.GET("/customer", h.GetCustomer)
	router.GET("/transactions", h.GetTransaction)
	router.PATCH("/credit", h.CreditWallet)
	router.PATCH("/debit", h.DebitWallet)
	router.POST("/addcustomer", h.AddCustomer)

	port := os.Getenv("PORT")

	return router, port
}

func Start() error {
	values := repo.InitializeDbParameters()
	var PDB = new(repo.PostgresDb)
	h := &Handler{DB: PDB}

	err := PDB.SetupDb(values.Host, values.User, values.Password, values.DbName, values.Port)
	if err != nil {
		log.Fatal(err)
	}
	routes, port := SetupRouter(h)
	err = routes.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
