package main

import "fmt"

// Paymenter is an interface.
// It defines a contract: any type that has a `pay(float32)` method
// is considered a Paymenter (no need to explicitly "implement").
type Paymenter interface {
	pay(amount float32)
}

// payment is a struct that uses a Paymenter interface as a field.
// This allows us to "plug in" any payment gateway that implements Paymenter.
type payment struct {
	GateWay Paymenter
}

// MakePayment calls the gateway's pay method.
// Notice that we don't care *which* gateway is being used (Razorpay, Stripe, etc.).
func (p payment) MakePayment(amount float32) {
	p.GateWay.pay(amount)
}

// razorPay is a type that represents RazorPay gateway.
type razorPay struct{}

// Implementing the pay method for razorPay.
// Since it matches the interface, razorPay automatically satisfies Paymenter.
func (r razorPay) pay(amount float32) {
	fmt.Println("Payment of", amount, "done using RazorPay")
}

// strip is a type that represents Stripe gateway.
type strip struct{}

// Implementing the pay method for strip.
func (s strip) pay(amount float32) {
	fmt.Println("Payment of", amount, "done using Stripe")
}

func main() {
	// Why interfaces? 🔑
	// Interfaces allow writing flexible, reusable code.
	// Instead of hardcoding payment logic for each gateway,
	// we can depend on the interface (abstraction),
	// and plug in any gateway implementation.

	// Creating gateway instances
	testStrip := strip{}
	testRazorPay := razorPay{}

	// Using Stripe gateway
	p1 := payment{GateWay: testStrip}
	p1.MakePayment(500.0)

	// Using RazorPay gateway
	p2 := payment{GateWay: testRazorPay}
	p2.MakePayment(1000.0)
}
