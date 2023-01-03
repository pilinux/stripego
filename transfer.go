package stripego

import (
	"github.com/stripe/stripe-go/v74"
	transfer "github.com/stripe/stripe-go/v74/transfer"
)

// TransferBalance - transfer balance to a connected Stripe account
func TransferBalance(sk string, tp *stripe.TransferParams) (tRes *stripe.Transfer, err error) {
	// stripe secret key
	stripe.Key = sk

	tRes, err = transfer.New(tp)
	return
}
