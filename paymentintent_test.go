package stripego_test

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/pilinux/stripego"
	"github.com/stripe/stripe-go/v74"
)

var PaymentIntentID string

func TestCreatePaymentIntent(t *testing.T) {
	err := stripego.Env()
	if err != nil {
		t.Errorf(
			"failed to load .env: %v", err,
		)
		return
	}
	sk := strings.TrimSpace(os.Getenv("STRIPE_SK"))
	currency := strings.TrimSpace(os.Getenv("CURRENCY"))

	piReq := stripe.PaymentIntentParams{}
	piReq.Amount = stripe.Int64(1000)
	piReq.Currency = stripe.String(currency)

	piRes, err := stripego.CreatePaymentIntent(sk, piReq)
	if err != nil {
		t.Errorf(
			"got error when creating payment intent: %v", err,
		)
		return
	}
	PaymentIntentID = piRes.ID

	res := &stripe.PaymentIntent{}
	err = json.Unmarshal(piRes.LastResponse.RawJSON, &res)
	if err != nil {
		t.Errorf("got error when unmarshalling payment intent: %v", err)
		return
	}

	expected := &stripe.PaymentIntent{}
	expected.Object = "payment_intent"
	expected.Status = stripe.PaymentIntentStatusRequiresPaymentMethod
	expected.Amount = *piReq.Amount
	expected.Currency = stripe.Currency(*piReq.Currency)

	if res.Object != expected.Object {
		t.Errorf("got: %v, want: %v", res.Object, expected.Object)
	}
	if res.Status != expected.Status {
		t.Errorf("got: %v, want: %v", res.Status, expected.Status)
	}
	if res.Amount != expected.Amount {
		t.Errorf("got: %v, want: %v", res.Amount, expected.Amount)
	}
	if res.Currency != expected.Currency {
		t.Errorf("got: %v, want: %v", res.Currency, expected.Currency)
	}
}

func TestCancelPaymentIntent(t *testing.T) {
	err := stripego.Env()
	if err != nil {
		t.Errorf(
			"failed to load .env: %v", err,
		)
		return
	}
	sk := strings.TrimSpace(os.Getenv("STRIPE_SK"))
	paymentIntentID := PaymentIntentID

	piRes, err := stripego.CancelPaymentIntent(sk, paymentIntentID)
	if err != nil {
		t.Errorf(
			"got error when canceling payment intent: %v", err,
		)
		return
	}

	res := &stripe.PaymentIntent{}
	err = json.Unmarshal(piRes.LastResponse.RawJSON, &res)
	if err != nil {
		t.Errorf("got error when unmarshalling payment intent: %v", err)
		return
	}

	expected := &stripe.PaymentIntent{}
	expected.ID = paymentIntentID
	expected.Object = "payment_intent"
	expected.Status = stripe.PaymentIntentStatusCanceled

	if res.ID != expected.ID {
		t.Errorf("got: %v, want: %v", res.ID, expected.ID)
	}
	if res.Object != expected.Object {
		t.Errorf("got: %v, want: %v", res.Object, expected.Object)
	}
	if res.Status != expected.Status {
		t.Errorf("got: %v, want: %v", res.Status, expected.Status)
	}
}
