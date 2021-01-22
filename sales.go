package metrc

import (
	"encoding/json"
	"fmt"
)

// SalesReceiptGet represents a receipt in the responses to `GET sales/v1/receipt` endpoints.
// See: https://api-ca.metrc.com/Documentation/#Sales.get_sales_v1_receipts_active
type SalesReceiptGet struct {
	Id                   int                   `json:"Id"`
	ReceiptNumber        *string               `json:"ReceiptNumber"`
	SalesDateTime        string                `json:"SalesDateTime"`
	SalesCustomerType    string                `json:"SalesCustomerType"`
	PatientLicenseNumber *string               `json:"PatientLicenseNumber"`
	TotalPackages        int                   `json:"TotalPackages"`
	TotalPrice           float64               `json:"TotalPrice"`
	Transactions         []SalesTransactionGet `json:"Transactions"`
	IsFinal              bool                  `json:"IsFinal"`
	ArchivedDate         *string               `json:"ArchivedDate"`
	RecordedDateTime     string                `json:"RecordedDateTime"`
	RecordedByUserName   *string               `json:"RecordedByUserName"`
	LastModified         string                `json:"LastModified"`
}

// SalesReceiptPost represents a receipt in the request to `POST sales/v1/receipts`.
// See: https://api-ca.metrc.com/Documentation/#Sales.post_sales_v1_receipts
type SalesReceiptPost struct {
	Id                     *int                   `json:"Id,omitempty"`
	SalesDateTime          string                 `json:"SalesDateTime"`
	SalesCustomerType      string                 `json:"SalesCustomerType"`
	PatientLicenseNumber   *string                `json:"PatientLicenseNumber"`
	CaregiverLicenseNumber *string                `json:"CaregiverLicenseNumber"`
	IdentificationMethod   *string                `json:"IdentificationMethod"`
	Transactions           []SalesTransactionPost `json:"Transactions"`
}

// SalesTransactionGet represents a Transaction in the responses to `GET sales/v1/transactions`.
// See: https://api-ca.metrc.com/Documentation/#Sales.get_sales_v1_transactions
type SalesTransactionGet struct {
	SalesDate         string  `json:"SalesDate"`
	TotalTransactions int     `json:"TotalTransactions"`
	TotalPackages     int     `json:"TotalPackages"`
	TotalPrice        float64 `json:"TotalPrice"`
}

// SalesTransactionPost represents a Transaction in the requests to `POST sales/v1/transactions`.
// See: https://api-ca.metrc.com/Documentation/#Sales.post_sales_v1_transactions_{date}
type SalesTransactionPost struct {
	PackageLabel   string  `json:"PackageLabel"`
	Quantity       float64 `json:"Quantity"`
	UnitsOfMeasure string  `json:"UnitOfMeasure"`
	TotalAmount    float64 `json:"TotalAmount"`
}

// GetSalesCustomerTypes gets the customer types for Sales.
// See: https://api-ca.metrc.com/Documentation/#Sales.get_sales_v1_customertypes
func (m *Metrc) GetSalesCustomerTypes() ([]string, error) {
	endpoint := "sales/v1/customertypes"
	resp, err := m.client.Get(endpoint)
	if err != nil {
		return []string{}, fmt.Errorf("could not get customer types: %s", err)
	}

	var cr []string
	err = json.Unmarshal(resp, &cr)
	if err != nil {
		return []string{}, fmt.Errorf("could not unmarshal customer types: %s", err)
	}

	return cr, nil
}

// GetSalesReceiptsActive gets all active receipts.
// See: https://api-ca.metrc.com/Documentation/#Sales.get_sales_v1_receipts_active
func (m *Metrc) GetSalesReceiptsActive(licenseNumber string, salesDateStart *string, salesDateEnd *string, lastModifiedStart *string, lastModifiedEnd *string) ([]SalesReceiptGet, error) {
	// Metrc API specification says that sales date range and last modified range cannot both be specified.
	hasSalesDate := (salesDateStart != nil || salesDateEnd != nil)
	hasLastModified := (lastModifiedStart != nil || lastModifiedEnd != nil)
	if hasSalesDate && hasLastModified {
		return []SalesReceiptGet{}, fmt.Errorf("cannot specify both salesDate and lastModified params")
	}

	// Append non-nil optional parameters.
	endpoint := fmt.Sprintf("sales/v1/receipts/active?licenseNumber=%s", licenseNumber)
	if salesDateStart != nil {
		endpoint += fmt.Sprintf("&salesDateStart=%s", *lastModifiedStart)
	}

	if salesDateEnd != nil {
		endpoint += fmt.Sprintf("&salesDateEnd=%s", *lastModifiedEnd)
	}

	if lastModifiedStart != nil {
		endpoint += fmt.Sprintf("&lastModifiedStart=%s", *lastModifiedStart)
	}

	if lastModifiedEnd != nil {
		endpoint += fmt.Sprintf("&lastModifiedEnd=%s", *lastModifiedEnd)
	}

	resp, err := m.client.Get(endpoint)
	if err != nil {
		return []SalesReceiptGet{}, fmt.Errorf("could not get active receipts from metrc, license number %s: %s", licenseNumber, err)
	}

	var arr []SalesReceiptGet
	err = json.Unmarshal(resp, &arr)
	if err != nil {
		return arr, fmt.Errorf("could not unmarshal receipts, license number %s: %s", licenseNumber, err)
	}

	return arr, nil
}

// GetSalesReceiptsInactive gets all inactive receipts.
// See: https://api-ca.metrc.com/Documentation/#Sales.get_sales_v1_receipts_inactive
func (m *Metrc) GetSalesReceiptsInactive(licenseNumber string, salesDateStart *string, salesDateEnd *string, lastModifiedStart *string, lastModifiedEnd *string) ([]SalesReceiptGet, error) {
	// Metrc API specification says that sales date range and last modified range cannot both be specified.
	hasSalesDate := (salesDateStart != nil || salesDateEnd != nil)
	hasLastModified := (lastModifiedStart != nil || lastModifiedEnd != nil)
	if hasSalesDate && hasLastModified {
		return []SalesReceiptGet{}, fmt.Errorf("cannot specify both salesDate and lastModified params")
	}

	// Append non-nil optional parameters.
	endpoint := fmt.Sprintf("sales/v1/receipts/inactive?licenseNumber=%s", licenseNumber)
	if salesDateStart != nil {
		endpoint += fmt.Sprintf("&salesDateStart=%s", *lastModifiedStart)
	}

	if salesDateEnd != nil {
		endpoint += fmt.Sprintf("&salesDateEnd=%s", *lastModifiedEnd)
	}

	if lastModifiedStart != nil {
		endpoint += fmt.Sprintf("&lastModifiedStart=%s", *lastModifiedStart)
	}

	if lastModifiedEnd != nil {
		endpoint += fmt.Sprintf("&lastModifiedEnd=%s", *lastModifiedEnd)
	}

	var srr []SalesReceiptGet
	resp, err := m.client.Get(endpoint)
	if err != nil {
		return srr, fmt.Errorf("could not get active receipts from metrc: %s", err)
	}

	err = json.Unmarshal(resp, &srr)
	if err != nil {
		return srr, fmt.Errorf("could not unmarshal receipts: %s", err)
	}

	return srr, nil
}

// GetSlaesReceiptsById gets a sales receipt by ID.
// See: https://api-ca.metrc.com/Documentation/#Sales.get_sales_v1_receipts_{id}
func (m *Metrc) GetSalesReceiptsById(id int, licenseNumber *string) (SalesReceiptGet, error) {
	endpoint := fmt.Sprintf("sales/v1/receipts/%d", id)
	if licenseNumber != nil {
		endpoint += fmt.Sprintf("?licenseNumber=%s", *licenseNumber)
	}

	var sr SalesReceiptGet
	resp, err := m.client.Get(endpoint)
	if err != nil {
		return sr, fmt.Errorf("could not get receipt from metrc: %s", err)
	}

	err = json.Unmarshal(resp, &sr)
	if err != nil {
		return sr, fmt.Errorf("could not unmarshal receipt response: %s", err)
	}

	return sr, nil
}

// PostSalesReceipts posts the sales receipts.
// See: https://api-ca.metrc.com/Documentation/#Sales.post_sales_v1_receipts
func (m *Metrc) PostSalesReceipts(receipts []SalesReceiptPost, licenseNumber string) ([]byte, error) {
	endpoint := fmt.Sprintf("sales/v1/receipts?licenseNumber=%s", licenseNumber)

	body, err := json.Marshal(receipts)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal receipts: %s", err)
	}

	resp, err := m.client.Post(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("could not post receipts: %s", err)
	}

	return resp, nil
}

// PutSalesReceipts puts the sales receipt.
// The primary difference between this and PostSalesReceipts is that this function requires a non-nil Id.
// See: https://api-ca.metrc.com/Documentation/#Sales.put_sales_v1_receipts
func (m *Metrc) PutSalesReceipts(receipts []SalesReceiptPost, licenseNumber string) ([]byte, error) {
	endpoint := fmt.Sprintf("sales/v1/receipts?licenseNumber=%s", licenseNumber)

	body, err := json.Marshal(receipts)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal receipts: %s", err)
	}

	resp, err := m.client.Put(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("could not post receipts: %s", err)
	}

	return resp, nil
}

// DeleteSalesReceiptsById deletes receipts with a specific ID.
// See: https://api-ca.metrc.com/Documentation/#Sales.delete_sales_v1_receipts_{id}
func (m *Metrc) DeleteSalesReceiptsById(id int, licenseNumber string) ([]byte, error) {
	endpoint := fmt.Sprintf("sales/v1/receipts/%d?licenseNumber=%s", id, licenseNumber)

	resp, err := m.client.Delete(endpoint)
	if err != nil {
		return []byte{}, fmt.Errorf("could not delete receipt with id %d, license number %s: %s", id, licenseNumber, err)
	}

	return resp, nil
}

// GetSalesTransactions gets the sales transactions for a license.
// See: https://api-ca.metrc.com/Documentation/#Sales.get_sales_v1_transactions
func (m *Metrc) GetSalesTransactions(licenseNumber string) ([]SalesTransactionGet, error) {
	endpoint := fmt.Sprintf("sales/v1/transactions?licenseNumber=%s", licenseNumber)

	var sts []SalesTransactionGet
	resp, err := m.client.Get(endpoint)
	if err != nil {
		return sts, fmt.Errorf("could not get sales transactions for license number %s: %s", licenseNumber, err)
	}

	err = json.Unmarshal(resp, &sts)
	if err != nil {
		return sts, fmt.Errorf("could not unmarshap transactions for license number %s: %s", licenseNumber, err)
	}

	return sts, nil
}

// GetSalesTransactionsByDates get the sales transactions for a license between dates.
// See: https://api-ca.metrc.com/Documentation/#Sales.get_sales_v1_transactions_{salesDateStart}_{salesDateEnd}
func (m *Metrc) GetSalesTransactionsByDates(salesDateStart string, salesDateEnd string, licenseNumber string) ([]SalesTransactionGet, error) {
	endpoint := fmt.Sprintf("sales/v1/transactions/%s/%s?licenseNumber=%s", salesDateStart, salesDateEnd, licenseNumber)

	var sts []SalesTransactionGet
	resp, err := m.client.Get(endpoint)
	if err != nil {
		return sts, fmt.Errorf("could not get sales transaction for license number %s, start date %s, end date %s: %s", licenseNumber, salesDateStart, salesDateEnd, err)
	}

	err = json.Unmarshal(resp, &sts)
	if err != nil {
		return sts, fmt.Errorf("could not unmarshal sales transactions response for license number %s, start date %s, end date %s: %s", licenseNumber, salesDateStart, salesDateEnd, err)
	}

	return sts, nil
}

// PostSalesTransaction posts new Sales transactions.
// See: https://api-ca.metrc.com/Documentation/#Sales.post_sales_v1_transactions_{date}
func (m *Metrc) PostSalesTransactions(transactions []SalesTransactionPost, date string, licenseNumber string) ([]byte, error) {
	endpoint := fmt.Sprintf("sales/v1/transactions/%s?licenseNumber=%s", date, licenseNumber)

	body, err := json.Marshal(transactions)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal sales transaction for license number %s, date %s: %s", licenseNumber, date, err)
	}

	_, err = m.client.Post(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("could not post sales transaction for license %s, date %s: %s", licenseNumber, date, err)
	}

	return []byte{}, nil
}

// PutSalesTransactions puts new Sales transactions.
// See: https://api-ca.metrc.com/Documentation/#Sales.put_sales_v1_transactions_{date}
func (m *Metrc) PutSalesTransactions(transactions []SalesTransactionPost, date string, licenseNumber string) ([]byte, error) {
	endpoint := fmt.Sprintf("sales/v1/transactions/%s?licenseNumber=%s", date, licenseNumber)

	body, err := json.Marshal(transactions)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal sales transaction for license number %s, date %s: %s", licenseNumber, date, err)
	}

	_, err = m.client.Put(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("could not put sales transactions for license %s, date %s: %s", licenseNumber, date, err)
	}

	return []byte{}, err
}
