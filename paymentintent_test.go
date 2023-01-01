package stripego_test

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/pilinux/stripego"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/paymentmethod"
)

var StripeSK string
var Currency string
var PaymentIntentObject string
var PaymentIntentID string
var PaymentIntentAmount int64

func TestCreatePaymentIntent(t *testing.T) {
	stripeSK, ok := os.LookupEnv("STRIPE_SK")
	if !ok {
		err := stripego.Env()
		if err != nil {
			t.Errorf("failed to load .env: %v", err)
			return
		}

		StripeSK = strings.TrimSpace(os.Getenv("STRIPE_SK"))
		Currency = strings.TrimSpace(os.Getenv("CURRENCY"))
	} else {
		StripeSK = strings.TrimSpace(stripeSK)
		Currency = strings.TrimSpace(os.Getenv("CURRENCY"))
	}

	PaymentIntentObject = "payment_intent"
	PaymentIntentAmount = 1000

	piReq := stripe.PaymentIntentParams{}
	piReq.Amount = stripe.Int64(PaymentIntentAmount)
	piReq.Currency = stripe.String(Currency)

	piRes, err := stripego.CreatePaymentIntent(StripeSK, piReq)
	if err != nil {
		t.Errorf("got error when creating payment intent: %v", err)
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
	expected.Object = PaymentIntentObject
	expected.Status = stripe.PaymentIntentStatusRequiresPaymentMethod
	expected.Amount = PaymentIntentAmount
	expected.Currency = stripe.Currency(Currency)

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

func TestUpdateAmountPaymentIntent(t *testing.T) {
	PaymentIntentAmount = 500

	piRes, err := stripego.UpdateAmountPaymentIntent(StripeSK, PaymentIntentID, PaymentIntentAmount)
	if err != nil {
		t.Errorf("got error when updating payment intent: %v", err)
		return
	}

	res := &stripe.PaymentIntent{}
	err = json.Unmarshal(piRes.LastResponse.RawJSON, &res)
	if err != nil {
		t.Errorf("got error when unmarshalling payment intent: %v", err)
		return
	}

	expected := &stripe.PaymentIntent{}
	expected.ID = PaymentIntentID
	expected.Object = PaymentIntentObject
	expected.Status = stripe.PaymentIntentStatusRequiresPaymentMethod
	expected.Amount = PaymentIntentAmount
	expected.Currency = stripe.Currency(Currency)

	if res.ID != expected.ID {
		t.Errorf("got: %v, want: %v", res.ID, expected.ID)
	}
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

func TestUpdateMethodPaymentIntent(t *testing.T) {
	// test with card
	card := &stripe.PaymentMethodParams{
		Type: stripe.String(string(stripe.PaymentMethodTypeCard)),
		Card: &stripe.PaymentMethodCardParams{
			Number:   stripe.String("4242424242424242"),
			ExpMonth: stripe.Int64(12),
			ExpYear:  stripe.Int64(2023),
			CVC:      stripe.String("123"),
		},
	}

	// create a payment method
	pm, err := paymentmethod.New(card)
	if err != nil {
		t.Errorf("got error when creating a payment method: %v", err)
		return
	}

	piRes, err := stripego.UpdateMethodPaymentIntent(StripeSK, PaymentIntentID, pm)
	if err != nil {
		t.Errorf("got error when updating payment intent: %v", err)
		return
	}

	res := &stripe.PaymentIntent{}
	err = json.Unmarshal(piRes.LastResponse.RawJSON, &res)
	if err != nil {
		t.Errorf("got error when unmarshalling payment intent: %v", err)
		return
	}

	expected := &stripe.PaymentIntent{}
	expected.ID = PaymentIntentID
	expected.Object = PaymentIntentObject
	expected.Status = stripe.PaymentIntentStatusRequiresConfirmation
	expected.Amount = PaymentIntentAmount
	expected.Currency = stripe.Currency(Currency)

	if res.ID != expected.ID {
		t.Errorf("got: %v, want: %v", res.ID, expected.ID)
	}
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
	piRes, err := stripego.CancelPaymentIntent(StripeSK, PaymentIntentID)
	if err != nil {
		t.Errorf("got error when canceling payment intent: %v", err)
		return
	}

	res := &stripe.PaymentIntent{}
	err = json.Unmarshal(piRes.LastResponse.RawJSON, &res)
	if err != nil {
		t.Errorf("got error when unmarshalling payment intent: %v", err)
		return
	}

	expected := &stripe.PaymentIntent{}
	expected.ID = PaymentIntentID
	expected.Object = PaymentIntentObject
	expected.Status = stripe.PaymentIntentStatusCanceled
	expected.Amount = PaymentIntentAmount
	expected.Currency = stripe.Currency(Currency)

	if res.ID != expected.ID {
		t.Errorf("got: %v, want: %v", res.ID, expected.ID)
	}
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
