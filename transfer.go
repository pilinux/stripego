package stripego

import (
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/balancetransaction"
	transfer "github.com/stripe/stripe-go/v74/transfer"
)

// TransferBalance - transfer balance to a connected Stripe account
func TransferBalance(sk string, tp *stripe.TransferParams) (tRes *stripe.Transfer, err error) {
	// stripe secret key
	stripe.Key = sk

	tRes, err = transfer.New(tp)
	return
}

// GetBalanceTx - get details of a balance transaction in Stripe
func GetBalanceTx(sk, bTxID string) (txn *stripe.BalanceTransaction, err error) {
	// stripe secret key
	stripe.Key = sk

	txn, err = balancetransaction.Get(bTxID, nil)
	return
}
