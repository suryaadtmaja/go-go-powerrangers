package payments

import (
	"errors"
	"testing"
)

// --- Test Helpers ---

type MockPaymentProcessor struct {
	ShouldFail bool
}

func (m MockPaymentProcessor) Pay(amount float64, method string) error {
	if m.ShouldFail {
		return errors.New("mock payment failure")
	}
	return nil
}

// --- PayPal Tests ---

func TestPayPalValidation(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		method   string
		expected error
	}{
		{"Valid Payment", 10.00, "card_123", nil},
		{"Zero Amount", 0.00, "card_123", errors.New("amount must be positive")},
		{"Negative Amount", -1.00, "card_123", errors.New("amount must be positive")},
		{"Below Minimum", 0.99, "card_123", errors.New("PayPal minimum charge is $1.00")},
		{"Missing Method", 1.00, "", errors.New("payment method required")},
	}

	processor := PayPal{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := processor.Pay(tt.amount, tt.method)
			if (err != nil) != (tt.expected != nil) {
				t.Errorf("unexpected error state: got %v, want %v", err, tt.expected)
			}
			if err != nil && err.Error() != tt.expected.Error() {
				t.Errorf("wrong error: got %q, want %q", err.Error(), tt.expected.Error())
			}
		})
	}
}

// --- Stripe Tests ---

func TestStripeValidation(t *testing.T) {
	tests := []struct {
		name     string
		amount   float64
		method   string
		expected error
	}{
		{"Valid Payment", 1.00, "card_123", nil},
		{"Minimum Charge", 0.50, "card_123", nil},
		{"Below Minimum", 0.49, "card_123", errors.New("Stripe minimum charge is $0.50")},
	}

	processor := Stripe{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := processor.Pay(tt.amount, tt.method)
			if (err != nil) != (tt.expected != nil) {
				t.Errorf("unexpected error state: got %v, want %v", err, tt.expected)
			}
			if err != nil && err.Error() != tt.expected.Error() {
				t.Errorf("wrong error: got %q, want %q", err.Error(), tt.expected.Error())
			}
		})
	}
}

// --- Payment Service Tests ---

func TestPaymentServiceValidation(t *testing.T) {
	// Create a mock processor that always succeeds when called
	mockProcessor := MockPaymentProcessor{ShouldFail: false}
	service := NewPaymentService(mockProcessor)

	tests := []struct {
		name     string
		amount   float64
		method   string
		expected error
	}{
		{"Valid Payment", 1.00, "card_123", nil},
		{"Zero Amount", 0.00, "card_123", errors.New("payment amount must be positive")},
		{"Negative Amount", -1.00, "card_123", errors.New("payment amount must be positive")},
		{"Missing Method", 1.00, "", errors.New("payment method must be specified")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := service.ProcessPayment(tt.amount, tt.method)
			if (err != nil) != (tt.expected != nil) {
				t.Errorf("unexpected error state: got %v, want %v", err, tt.expected)
			}
			if err != nil && err.Error() != tt.expected.Error() {
				t.Errorf("wrong error: got %q, want %q", err.Error(), tt.expected.Error())
			}
		})
	}
}

func TestPaymentServiceProcessorFailure(t *testing.T) {
	// Create a mock processor that always fails
	mockProcessor := MockPaymentProcessor{ShouldFail: true}
	service := NewPaymentService(mockProcessor)

	err := service.ProcessPayment(1.00, "card_123")
	if err == nil {
		t.Error("expected payment to fail, but it succeeded")
	}
	if err.Error() != "mock payment failure" {
		t.Errorf("unexpected error: got %q", err.Error())
	}
}
