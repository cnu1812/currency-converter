package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ConvertRequest struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

type FrankfurterResponse struct {
	Amount float64            `json:"amount"`
	Base   string             `json:"base"`
	Date   string             `json:"date"`
	Rates  map[string]float64 `json:"rates"`
}

func main() {
	r := gin.Default()

	r.POST("/convert", func(c *gin.Context) {
		var req ConvertRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		url := fmt.Sprintf("https://api.frankfurter.app/latest?from=%s&to=%s", req.From, req.To)

		resp, err := http.Get(url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reach exchange API"})
			return
		}
		defer resp.Body.Close()

		var ratesData FrankfurterResponse
		if err := json.NewDecoder(resp.Body).Decode(&ratesData); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse exchange rate data"})
			return
		}

		rate, exists := ratesData.Rates[req.To]
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Currency code not found"})
			return
		}

		convertedValue := req.Amount * rate

		c.JSON(http.StatusOK, gin.H{
			"original_currency": req.From,
			"target_currency":   req.To,
			"original_amount":   req.Amount,
			"exchange_rate":     rate,
			"converted_amount":  convertedValue,
			"message":           "Calculated using live rates!",
		})
	})

	r.Run(":8080")
}
