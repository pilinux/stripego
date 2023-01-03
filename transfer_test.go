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
var BTxID string

func TestTransferBalance(t *testing.T) {
	StripeSK = strings.TrimSpace(os.Getenv("STRIPE_SK"))
	Currency = strings.TrimSpace(os.Getenv("CURRENCY"))
	Destination = strings.TrimSpace(os.Getenv("DESTINATION"))

	TransferObject = "transfer"
	TransferAmount = 100

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
	BTxID = tRes.BalanceTransaction.ID

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

func TestGetBalanceTx(t *testing.T) {
	txn, err := stripego.GetBalanceTx(StripeSK, BTxID)
	if err != nil {
		t.Errorf("got error when retrieving a transaction details: %v", err)
		return
	}

	res := &stripe.BalanceTransaction{}
	err = json.Unmarshal(txn.LastResponse.RawJSON, &res)
	if err != nil {
		t.Errorf("got error when unmarshalling transaction details: %v", err)
		return
	}

	if res.ID != BTxID {
		t.Errorf("balance transaction IDs do not match")
	}
	if res.Currency != stripe.Currency(Currency) {
		t.Errorf("got: %v, want: %v", res.Currency, stripe.Currency(Currency))
	}
}
