package stripego

import (
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/paymentintent"
)

// CreatePaymentIntent - create a new PaymentIntent object in stripe
func CreatePaymentIntent(sk string, piReq stripe.PaymentIntentParams) (piRes *stripe.PaymentIntent, err error) {
	// stripe secret key
	stripe.Key = sk

	// create PaymentIntentParams with amount and currency
	params := &stripe.PaymentIntentParams{
		Amount:   piReq.Amount,
		Currency: piReq.Currency,
		// enable all payment methods
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	// create a PaymentIntent
	piRes, err = paymentintent.New(params)
	return
}

// CancelPaymentIntent - cancel an existing PaymentIntent object in stripe
func CancelPaymentIntent(sk, paymentIntentID string) (piRes *stripe.PaymentIntent, err error) {
	// stripe secret key
	stripe.Key = sk

	piRes, err = paymentintent.Cancel(paymentIntentID, nil)
	return
}
