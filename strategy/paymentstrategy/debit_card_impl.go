package paymentstrategy

type DebitCardDetail struct {
	CardNumber string
	Cvv        string
}

func (s *DebitCardDetail) Pay(amount int) string {
	return "Payment Processed with debit card"
}
