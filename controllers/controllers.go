package controllers

import (
	"net/http"
	"quesiasts/digitalbank-api.git/database"
	"quesiasts/digitalbank-api.git/models"

	"github.com/gin-gonic/gin"
)

func Introduction(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"API says:": "Hi " + name + ", are you okay?",
	})
}

// GET /accounts - obtém a lista de contas
// GET /accounts/{account_id}/balance - obtém o saldo da conta
// POST /accounts - cria uma Account
// POST /login - autentica a usuaria
// GET /transfers - obtém a lista de transferencias da usuaria autenticada.
// POST /transfers - faz transferencia de uma Account para outra.

// Regras para esta rota

// Quem fizer a transferência precisa estar autenticada.
// O account_origin_id deve ser obtido no Token enviado.
// Caso Account de origem não tenha saldo, retornar um código de erro apropriado
// Atualizar o balance das contas

// POST /accounts - cria uma Account
func CreateAccount(c *gin.Context) {
	var account models.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := models.ValidateAccount(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&account)
	c.JSON(http.StatusOK, account)
}

// GET /accounts - obtém a lista de contas
func ListAllAccounts(c *gin.Context) {
	var account []models.Account
	database.DB.Find(&account)
	c.JSON(200, account)
}

// GET /accounts/{account_id}/balance
func SearchForBalance(c *gin.Context) {
	var account models.Account
	balance := c.GetFloat64("balance")
	database.DB.Where(&models.Account{Balance: balance}).First(&account)

	if account.Id == "" {
		c.JSON(http.StatusNotFound, gin.H{"Not Found": "Balance not found"})
		return
	}

	c.JSON(http.StatusOK, account)
}
