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

// UpdateAmountPaymentIntent - update the amount of an existing PaymentIntent object in stripe
func UpdateAmountPaymentIntent(sk, paymentIntentID string, newAmount int64) (piRes *stripe.PaymentIntent, err error) {
	// stripe secret key
	stripe.Key = sk

	// create PaymentIntentParams with amount
	params := &stripe.PaymentIntentParams{
		Amount: stripe.Int64(newAmount),
	}

	// update the PaymentIntent
	piRes, err = paymentintent.Update(paymentIntentID, params)
	return
}

// UpdateMethodPaymentIntent - update the payment method of an existing PaymentIntent object in stripe
func UpdateMethodPaymentIntent(sk, paymentIntentID string, pm *stripe.PaymentMethod) (piRes *stripe.PaymentIntent, err error) {
	// stripe secret key
	stripe.Key = sk

	// update the PaymentIntent
	piRes, err = paymentintent.Update(
		paymentIntentID,
		&stripe.PaymentIntentParams{
			PaymentMethod: stripe.String(pm.ID),
		},
	)
	return
}

// CancelPaymentIntent - cancel an existing PaymentIntent object in stripe
func CancelPaymentIntent(sk, paymentIntentID string) (piRes *stripe.PaymentIntent, err error) {
	// stripe secret key
	stripe.Key = sk

	piRes, err = paymentintent.Cancel(paymentIntentID, nil)
	return
}
