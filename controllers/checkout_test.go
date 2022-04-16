package controllers

import (
	"context"
	d "fita-assignment/database"
	"fita-assignment/models"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(m *testing.M) {
	db := d.ConnectMySQL()
	defer db.Close()
	// Run tests
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestCheckout(t *testing.T) {
	ctx := context.Background()

	testCases := []struct {
		desc string
		req  []models.Item
		res  models.CheckoutResult
		err  error
	}{
		{
			desc: "Macbook Pro and Raspberry Pi",
			req: []models.Item{
				{
					SKU: "43N23P",
					Qty: 1,
				},
				{
					SKU: "234234",
					Qty: 1,
				},
			},
			res: models.CheckoutResult{
				FinalPrice: 5399.99,
			},
			err: nil,
		},
		{
			desc: "Raspberry Pi and Macbook Pro (Scanned Reverse)",
			req: []models.Item{
				{
					SKU: "234234",
					Qty: 1,
				},
				{
					SKU: "43N23P",
					Qty: 1,
				},
			},
			res: models.CheckoutResult{
				FinalPrice: 5399.99,
			},
			err: nil,
		},
		{
			desc: "Google Home x 3",
			req: []models.Item{
				{
					SKU: "120P90",
					Qty: 3,
				},
			},
			res: models.CheckoutResult{
				FinalPrice: 99.98,
			},
			err: nil,
		},
		{
			desc: "Alexa Speaker x3",
			req: []models.Item{
				{
					SKU: "A304SD",
					Qty: 3,
				},
			},
			res: models.CheckoutResult{
				FinalPrice: 295.65,
			},
			err: nil,
		},
	}

	Convey("Checkout", t, func() {
		for _, tc := range testCases {
			Convey(tc.desc, func() {
				res, err := Checkout(ctx, tc.req)
				if tc.err != nil {
					So(err, ShouldNotBeNil)
					return
				}
				So(err, ShouldBeNil)
				So(res.FinalPrice, ShouldEqual, tc.res.FinalPrice)
			})
		}
	})
}
