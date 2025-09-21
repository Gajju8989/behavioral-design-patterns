package main

import (
	"fmt"
	"strategypattern/paymentprocess"
	"strategypattern/service"
)

func main() {
	processPayment := paymentprocess.ProcessPayment{}
	PaymentService := service.NewPaymentService(processPayment)
	processedAmount, err := PaymentService.Pay("credit_card", 10)
	if err != nil {
		return
	}

	fmt.Println(processedAmount)
}
