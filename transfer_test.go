package stripego_test

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/pilinux/stripego"
	"github.com/stripe/stripe-go/v74"
)

var Destination string
var TransferObject string
var TransferAmount int64

func TestTransferBalance(t *testing.T) {
	StripeSK = strings.TrimSpace(os.Getenv("STRIPE_SK"))
	Currency = strings.TrimSpace(os.Getenv("CURRENCY"))
	Destination = strings.TrimSpace(os.Getenv("DESTINATION"))

	TransferObject = "transfer"
	TransferAmount = 1

	tp := &stripe.TransferParams{
		Amount:      stripe.Int64(TransferAmount),
		Currency:    stripe.String(Currency),
		Destination: stripe.String(Destination),
		// SourceTransaction: stripe.String(Source),
	}

	tRes, err := stripego.TransferBalance(StripeSK, tp)
	if err != nil {
		t.Errorf("got error when initiating balance transfer: %v", err)
		return
	}

	res := &stripe.Transfer{}
	err = json.Unmarshal(tRes.LastResponse.RawJSON, &res)
	if err != nil {
		t.Errorf("got error when unmarshalling transfer response: %v", err)
		return
	}

	expected := &stripe.Transfer{}
	expected.Object = TransferObject
	expected.Amount = TransferAmount
	expected.Currency = stripe.Currency(Currency)

	if res.Object != expected.Object {
		t.Errorf("got: %v, want: %v", res.Object, expected.Object)
	}
	if res.Amount != expected.Amount {
		t.Errorf("got: %v, want: %v", res.Amount, expected.Amount)
	}
	if res.Currency != expected.Currency {
		t.Errorf("got: %v, want: %v", res.Currency, expected.Currency)
	}
	if res.Destination.ID != Destination {
		t.Errorf("destination IDs do not match")
	}
}
