package controller

import (
	"fmt"
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

	port := os.Getenv("ROUTER_PORT")

	return router, port
}

func Start() error {
	var PDB = new(repo.PostgresDb)
	h := &Handler{DB: PDB}

	err := PDB.SetupDb()
	if err != nil {
		log.Fatal(err)
	}
	routes, port := SetupRouter(h)
	fmt.Println("Server running on port", port)
	err = routes.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
