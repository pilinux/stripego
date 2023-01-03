package stripego

import (
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/balancetransaction"
)

// GetBalanceTx - get details of a balance transaction in Stripe
func GetBalanceTx(sk, bTxID string) (txn *stripe.BalanceTransaction, err error) {
	// stripe secret key
	stripe.Key = sk

	txn, err = balancetransaction.Get(bTxID, nil)
	return
}
