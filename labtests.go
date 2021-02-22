package metrc

import (
	"encoding/json"
	"fmt"
)

// GetLabTestsStates gets all states for lab tests in Metrc.
// See: https://testing-api-ca.metrc.com/Documentation#LabTests.get_labtests_v1_states.GET
func (m *Metrc) GetLabTestsStates() ([]string, error) {
	endpoint := "labtests/v1/states"

	var ltr []string
	responseBody, err := m.client.Get(endpoint)
	if err != nil {
		return ltr, fmt.Errorf("could not get lab test states response: %s", err)
	}

	err = json.Unmarshal(responseBody, &ltr)
	if err != nil {
		return ltr, fmt.Errorf("could not unmarshal response: %s", err)
	}

	return ltr, nil
}

// LabTestType represents a type of lab test in Metrc.
// See: https://testing-api-ca.metrc.com/Documentation#LabTests.get_labtests_v1_types.GET
type LabTestType struct {
	Id                   int      `json:"Id"`
	Name                 string   `json:"Name"`
	RequiresTestResult   bool     `json:"RequiresTestResult"`
	InformationalOnly    bool     `json:"InformationalOnly"`
	AlwaysPasses         bool     `json:"AlwaysPasses"`
	LabTestResultMode    int      `json:"LabTestResultMode"`
	LabTestResultMinimum *float64 `json:"LabTestResultMinimum"`
	LabTestResultMaximum *float64 `json:"LabTestResultMaximum"`
	DependencyMode       int      `json:"DependencyMode"`
}

// GetLabTestsTypes displays the types of lab tests in Metrc.
// See: https://testing-api-ca.metrc.com/Documentation#LabTests.get_labtests_v1_types.GET
func (m *Metrc) GetLabTestsTypes() ([]LabTestType, error) {
	endpoint := "labtests/v1/types"

	var ltr []LabTestType
	responseBody, err := m.client.Get(endpoint)
	if err != nil {
		return ltr, fmt.Errorf("could not get lab test types response: %s", err)
	}

	err = json.Unmarshal(responseBody, &ltr)
	if err != nil {
		return ltr, fmt.Errorf("could not unmarshal response: %s", err)
	}

	return ltr, nil
}

// LabTestResult represents a result of a lab test.
// See: https://testing-api-ca.metrc.com/Documentation#LabTests.get_labtests_v1_results.GET
type LabTestResult struct {
	PackageId                int     `json:"PackageId"`
	LabTestResultId          int     `json:"LabTestResultId"`
	LabFacilityLicenseNumber string  `json:"LabFacilityLicenseNumber"`
	LabFacilityName          string  `json:"LabFacilityName"`
	SourcePackageLabel       string  `json:"SourcePackageLabel"`
	ProductName              string  `json:"ProductName"`
	ProductCategoryName      string  `json:"ProductCategoryName"`
	TestPerformedDate        string  `json:"TestPerformedDate"`
	OverallPassed            bool    `json:"OverallPassed"`
	RevokedDate              *string `json:"RevokedDate"`
	ResultReleased           bool    `json:"ResultReleased"`
	ResultReleaseDateTime    string  `json:"ResultReleaseDateTime"`
	TestTypeName             string  `json:"TestTypeName"`
	TestPassed               bool    `json:"TestPassed"`
	TestResultLevel          float64 `json:"TestResultLevel"`
	TestComment              string  `json:"TestComment"`
	TestInformationalOnly    bool    `json:"TestInformationalOnly"`
	LabTestDetailRevokedDate *string `json:"LabTestDetailRevokedDate"`
}

// GetLabTestsResults gets the lab tests results.
// See: https://testing-api-ca.metrc.com/Documentation#LabTests.get_labtests_v1_results.GET
func (m *Metrc) GetLabTestsResults(packageId int, licenseNumber string) ([]LabTestResult, error) {
	endpoint := "labtests/v1/results"
	endpoint += fmt.Sprintf("?packageId=%d", packageId)
	endpoint += fmt.Sprintf("?licenseNumber=%s", licenseNumber)

	var ltr []LabTestResult
	responseBody, err := m.client.Get(endpoint)
	if err != nil {
		return ltr, fmt.Errorf("could not get lab test results response: %s", err)
	}

	err = json.Unmarshal(responseBody, &ltr)
	if err != nil {
		return ltr, fmt.Errorf("could not unmarshal response: %s", err)
	}

	return ltr, nil
}

// LabTestRecord represents a lab test record in the request body.
// See: https://testing-api-ca.metrc.com/Documentation#LabTests.post_labtests_v1_record.POST
type LabTestRecord struct {
	Label              string                `json:"Label"`
	ResultDate         string                `json:"ResultDate"`
	DocumentFileName   *string               `json:"DocumentFileName"`
	DocumentFileBase64 *string               `json:"DocumentFileBase64"`
	Results            []LabTestRecordResult `json:"Results"`
}

// LabTestRecordResult represents a result in a lab test when recording.
// See: https://testing-api-ca.metrc.com/Documentation#LabTests.post_labtests_v1_record.POST
type LabTestRecordResult struct {
	LabTestTypeName string  `json:"LabTestTypeName"`
	Quantity        float64 `json:"Quantity"`
	Passed          bool    `json:"Passed"`
	Notes           string  `json:"Notes"`
}

// PostLabTestsRecord records lab tests in Metrc.
// See: https://testing-api-ca.metrc.com/Documentation#LabTests.post_labtests_v1_record.POST
func (m *Metrc) PostLabTestsRecord(records []LabTestRecord, licenseNumber string) ([]byte, error) {
	endpoint := "labtests/v1/record"
	endpoint += fmt.Sprintf("?licenseNumber=%s", licenseNumber)

	body, err := json.Marshal(records)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal lab test records: %s", err)
	}

	resp, err := m.client.Post(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed posting lab test results: %s", err)
	}

	return resp, nil
}

// PutLabTestsDocument represents a document when putting a lab test.
// See: https://testing-api-ca.metrc.com/Documentation#LabTests.put_labtests_v1_labtestdocument.PUT
type LabTestDocument struct {
	LabTestResultId    int    `json:"LabTestResultId"`
	DocumentFileName   string `json:"DocumentFileName"`
	DocumentFileBase64 string `json:"DocumentFileBase64"`
}

// PutLabTestsDocument puts a document for lab tests.
// See: https://testing-api-ca.metrc.com/Documentation#LabTests.put_labtests_v1_labtestdocument.PUT
func (m *Metrc) PutLabTestsDocument(documents []LabTestDocument, licenseNumber string) ([]byte, error) {
	endpoint := "labtests/v1/labtestdocument"
	endpoint += fmt.Sprintf("?licenseNumber=%s", licenseNumber)

	body, err := json.Marshal(documents)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal lab test documents: %s", err)
	}

	resp, err := m.client.Put(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed putting lab test documents: %s", err)
	}

	return resp, nil
}

// LabTestResultRelease represents a lab test result to release.
// See: https://testing-api-ca.metrc.com/Documentation#LabTests.put_labtests_v1_results_release.PUT
type LabTestResultRelease struct {
	PackageLabel string `json:"PackageLabel"`
}

// PutLabTestsResultsRelease puts test lab results for release.
// See: https://testing-api-ca.metrc.com/Documentation#LabTests.put_labtests_v1_results_release.PUT
func (m *Metrc) PutLabTestsResultsRelease(packages []LabTestRecordResult, licenseNumber string) ([]byte, error) {
	endpoint := "labtests/v1/results/release"
	endpoint += fmt.Sprintf("?licenseNumber=%s", licenseNumber)

	body, err := json.Marshal(packages)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal lab tests results: %s", err)
	}

	resp, err := m.client.Put(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed putting lab test results release: %s", err)
	}

	return resp, nil
}
