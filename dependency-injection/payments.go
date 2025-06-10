package payments

import (
	"errors"
	"fmt"
)

// Payments interface defines the contract for all payment processors
type Payments interface {
	Pay(amount float64, method string) error
}

// --- Payment Processors ---

type PayPal struct{}

func (p PayPal) Pay(amount float64, method string) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}
	if amount < 1.00 {
		return errors.New("PayPal minimum charge is $1.00")
	}
	if method == "" {
		return errors.New("payment method required")
	}

	fmt.Printf("Processing $%.2f via PayPal (%s)\n", amount, method)
	return nil
}

type Stripe struct{}

func (s Stripe) Pay(amount float64, method string) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}
	if amount < 0.50 {
		return errors.New("Stripe minimum charge is $0.50")
	}
	if method == "" {
		return errors.New("payment method required")
	}

	fmt.Printf("Charging $%.2f via Stripe (%s)\n", amount, method)
	return nil
}

// --- Payment Service ---

type PaymentService struct {
	gateway Payments
}

func NewPaymentService(gateway Payments) *PaymentService {
	return &PaymentService{gateway: gateway}
}

func (ps *PaymentService) ProcessPayment(amount float64, method string) error {
	// Service-level validation (applies to all processors)
	if amount <= 0 {
		return errors.New("payment amount must be positive")
	}
	if method == "" {
		return errors.New("payment method must be specified")
	}

	// Delegate to the specific payment processor
	return ps.gateway.Pay(amount, method)
}
