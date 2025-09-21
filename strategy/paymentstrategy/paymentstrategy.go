package paymentstrategy

type PaymentStrategy interface {
	Pay(amount int) string
}
