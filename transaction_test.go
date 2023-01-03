package stripego_test

import (
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/pilinux/stripego"
	"github.com/stripe/stripe-go/v74"
)

var BTxID string

func TestGetBalanceTx(t *testing.T) {
	StripeSK = strings.TrimSpace(os.Getenv("STRIPE_SK"))
	Currency = strings.TrimSpace(os.Getenv("CURRENCY"))
	BTxID = strings.TrimSpace(os.Getenv("BALANCE_TRANSACTION_ID"))

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
