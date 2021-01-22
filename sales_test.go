package metrc

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSalesGetCustomerTypes_Integration(t *testing.T) {
	types, err := m.GetSalesCustomerTypes()
	assert.NoError(t, err)
	assert.Equal(t, 4, len(types))

	// Check that the type of desired name still exists.
	wantType := "Consumer"
	foundType := false
	for _, typ := range types {
		if typ == wantType {
			foundType = true
			break
		}
	}
	assert.True(t, foundType)
}

func TestSalesGetReceiptsActive_Integration(t *testing.T) {
	// TODO: Complete this once we have non-zero receipts.
	receipts, err := m.GetSalesReceiptsActive(licenseNumber, nil, nil, nil, nil)
	assert.NoError(t, err)
	fmt.Printf("%+v\n", receipts)

	// TODO: test the various optional parameters once we populate the receipts.
}

func TestSalesGetReceiptsInactive_Integration(t *testing.T) {
	// TODO: Complete this once we have non-zero receipts.
	receipts, err := m.GetSalesReceiptsInactive(licenseNumber, nil, nil, nil, nil)
	assert.NoError(t, err)
	fmt.Printf("%+v\n", receipts)

	// TODO: test the various optional parameters once we populate the receipts.
}

func TestSalesGetReceiptsById_Integration(t *testing.T) {
	id := 0 // TODO: Define this once the other GET endpoints are working.
	receipt, err := m.GetSalesReceiptsById(id, &licenseNumber)
	assert.NoError(t, err)
	assert.Equal(t, id, receipt.Id)
}

// TODO: Confirm that this test works after Package works.
func TestSalesPostDeleteReceipts_Integration(t *testing.T) {
	label := "ABCDEF012345670000010331" // TODO: Replace with a valid Package label once Package Post works.
	rand.Seed(time.Now().Unix())
	wantId := rand.Int()
	wantReceipt := SalesReceiptPost{
		Id:                &wantId,
		SalesDateTime:     "2020-10-04T16:44:53.000",
		SalesCustomerType: "Consumer",
		Transactions: []SalesTransactionPost{
			{
				PackageLabel:   label,
				Quantity:       1.0,
				UnitsOfMeasure: "Ounces",
				TotalAmount:    9.99,
			},
		},
	}
	inputReceipts := []SalesReceiptPost{wantReceipt}
	_, err := m.PostSalesReceipts(inputReceipts, licenseNumber)
	assert.NoError(t, err)

	var gotReceipt SalesReceiptGet
	gotReceipts, err := m.GetSalesReceiptsActive(licenseNumber, nil, nil, nil, nil)
	fmt.Printf("%+v\n", gotReceipts)
	for _, gr := range gotReceipts {
		if gr.Id == wantId {
			gotReceipt = gr
			break
		}
	}

	assert.Equal(t, gotReceipt.Id, wantId)

	fmt.Printf("%+v\n", gotReceipt)
	_, err = m.DeleteSalesReceiptsById(gotReceipt.Id, licenseNumber)
	assert.NoError(t, err)
}

// TODO: Confirm that this test works once Package works.
func TestSalesPutDeleteReceipts_Integration(t *testing.T) {

	label := "ABCDEF012345670000010331" // TODO: Replace with a valid Package label once Package Post works.
	salesDateTime := "2020-10-04T16:44:53.000"
	wantReceipt := SalesReceiptPost{
		SalesDateTime:     salesDateTime,
		SalesCustomerType: "Consumer",
		Transactions: []SalesTransactionPost{
			{
				PackageLabel:   label,
				Quantity:       1.0,
				UnitsOfMeasure: "Ounces",
				TotalAmount:    9.99,
			},
		},
	}
	inputReceipts := []SalesReceiptPost{wantReceipt}
	_, err := m.PostSalesReceipts(inputReceipts, licenseNumber)
	assert.NoError(t, err)

	var gotReceipt SalesReceiptGet
	gotReceipts, err := m.GetSalesReceiptsActive(licenseNumber, nil, nil, nil, nil)
	fmt.Printf("%+v\n", gotReceipts)
	for _, gr := range gotReceipts {
		if gr.SalesDateTime == salesDateTime {
			gotReceipt = gr
			break
		}
	}

	fmt.Printf("%+v\n", gotReceipt)
	deleteId := gotReceipt.Id
	_, err = m.DeleteSalesReceiptsById(deleteId, licenseNumber)
	assert.NoError(t, err)
}

func TestSalesGetTransactions_Integration(t *testing.T) {
	sts, err := m.GetSalesTransactions(licenseNumber)
	assert.NoError(t, err)
	fmt.Printf("%+v\n", sts)
	// TODO: Add checks once posts are made.
}

func TestSalesGetTransactionsByDates_Integration(t *testing.T) {
	// TODO: Define a date range for testing.
	start := ""
	end := ""
	sts, err := m.GetSalesTransactionsByDates(start, end, licenseNumber)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(sts)) // TODO: Add checks once posts are made.
}

func TestSalesPostTransactions_Integration(t *testing.T) {
	// TODO: Implement once Package works.
}

func TestSalesPutTransactions_Integration(t *testing.T) {
	// TODO: Implement once Package works.
}
