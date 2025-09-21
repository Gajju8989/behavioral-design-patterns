package paymentstrategy

type CreditCardDetail struct {
	CardNumber string
	Cvv        string
}

func (s *CreditCardDetail) Pay(amount int) string {
	return "Payment Processed with credit card"
}
