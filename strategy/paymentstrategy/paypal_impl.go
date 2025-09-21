package paymentstrategy

type PaypalAccountDetail struct {
	UserName string
}

func (s *PaypalAccountDetail) Pay(amount int) string {
	return "Payment Processed with paypal account"
}
