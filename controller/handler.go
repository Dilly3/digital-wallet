package controller

import (
	"github.com/dilly3/digital-wallet/models"
	"github.com/dilly3/digital-wallet/repository"
	services "github.com/dilly3/digital-wallet/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	DB repository.DB
}

func (h *Handler) GetCustomer(c *gin.Context) {
	accountNos := c.Query("accountNos")
	if !services.ValidateAccountNumber(accountNos) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "account number should be 10 digits"})
		return
	}
	customer, err := h.DB.Getcustomer(accountNos)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "user not found",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": customer,
	})
}

func (h *Handler) CreditWallet(c *gin.Context) {

	credit := &models.Money{}
	if err := c.ShouldBindJSON(credit); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "unable to bind json"})
		return
	}
	if !services.ValidateAccountNumber(credit.AccountNos) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "account number should be 10 digits"})
		return
	}

	transaction, CreditErr := h.DB.Creditwallet(credit)
	if CreditErr != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "unable to credit wallet"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":     "wallet credited successfully",
		"transaction": transaction,
	})

}
func (h *Handler) DebitWallet(c *gin.Context) {

	debit := &models.Money{}
	if err := c.ShouldBindJSON(debit); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "unable to bind json"})
		return
	}

	customer, err := h.DB.Getcustomer(debit.AccountNos)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "unable to get customer"})
		return
	}

	if h.DB.InsufficientFunds(customer, debit) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "insufficient funds"})
		return
	}

	transaction, debitErr := h.DB.Debitwallet(debit)
	if debitErr != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "unable to debit wallet"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "wallet debited successfully",
		"transaction": transaction,
	})
}
func (h *Handler) GetTransaction(c *gin.Context) {
	accountNos := c.Query("accountNos")
	if !services.ValidateAccountNumber(accountNos) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "account number should be 10 digits"})
		return
	}

	transactions, err := h.DB.Gettransaction(accountNos)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "transaction not found",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": transactions,
	})
}

func (h *Handler) AddCustomer(c *gin.Context) {
	var customer *models.Customer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "unable to bind json"})
		return
	}

	if _, userErr := h.DB.Getcustomer(customer.AccountNos); userErr == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "customer exists"})
		return
	}

	if CreateErr := h.DB.Addcustomer(*customer); CreateErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "unable to create customer"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Customer added successfully"})
}
