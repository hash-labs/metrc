package metrc

import (
	"encoding/json"
	"fmt"
)

// Transfer represents an incoming Transfer in Metrc.
// See: https://testing-api-ca.metrc.com/Documentation#Transfers.get_transfers_v1_incoming.GET
type Transfer struct {
	Id                                 int     `json:"Id"`
	ManifestNumber                     *string `json:"ManifestNumber"`
	ShipmentLicenseType                int     `json:"ShipmentLicenseType"`
	ShipperFacilityLicenseNumber       string  `json:"ShipperFacilityLicenseNumber"`
	ShipperFacilityName                string  `json:"ShipperFacilityName"`
	Name                               *string `json:"Name"`
	TransporterFacilityLicenseNumber   string  `json:"TransporterFacilityLicenseNumber"`
	TransporterFacilityName            string  `json:"TransporterFacilityName"`
	DriverName                         string  `json:"DriverName"`
	DriverOccupationalLicenseNumber    string  `json:"DriverOccupationalLicenseNumber"`
	DriverVehicleLicenseNumber         string  `json:"DriverVehicleLicenseNumber"`
	VehicleMake                        string  `json:"VehicleMake"`
	VehicleModel                       string  `json:"VehicleModel"`
	VehicleLicensePlateNumber          string  `json:"VehicleLicensePlateNumber"`
	DeliveryCount                      int     `json:"DeliveryCount"`
	ReceivedDeliveryCount              int     `json:"ReceivedDeliveryCount"`
	PackageCount                       int     `json:"PackageCount"`
	ReceivedPackageCount               int     `json:"ReceivedPackageCount"`
	ContainsPlantPackage               bool    `json:"ContainsPlantPackage"`
	ContainsProductPackage             bool    `json:"ContainsProductPackage"`
	ContainsTradeSample                bool    `json:"ContainsTradeSample"`
	ContainsDonation                   bool    `json:"ContainsDonation"`
	ContainsTestingSample              bool    `json:"ContainsTestingSample"`
	ContainsProductRequiresRemediation bool    `json:"ContainsProductRequiresRemediation"`
	ContainsRemediatedProductPackage   bool    `json:"ContainsRemediatedProductPackage"`
	CreatedDateTime                    string  `json:"CreatedDateTime"`
	CreatedByUserName                  *string `json:"CreatedByUserName"`
	LastModified                       string  `json:"LastModified"`
	DeliveryId                         int     `json:"DeliveryId"`
	RecipientFacilityLicenseNumber     *string `json:"RecipientFacilityLicenseNumber"`
	RecipientFacilityName              *string `json:"RecipientFacilityName"`
	ShipmentTypeName                   *string `json:"ShipmentTypeName"`
	ShipmentTransactionType            *string `json:"ShipmentTransactionType"`
	EstimatedDepartureDateTime         string  `json:"EstimatedDepartureDateTime"`
	ActualDepartureDateTime            *string `json:"ActualDepartureDateTime"`
	EstimatedArrivalDateTime           string  `json:"EstimatedArrivalDateTime"`
	ActualArrivalDateTime              *string `json:"ActualArrivalDateTime"`
	DeliveryPackageCount               int     `json:"DeliveryPackageCount"`
	DeliveryReceivedPackageCount       int     `json:"DeliveryReceivedPackageCount"`
	ReceivedDateTime                   *string `json:"ReceivedDateTime"`
	EstimatedReturnDepartureDateTime   *string `json:"EstimatedReturnDepartureDateTime"`
	ActualReturnDepartureDateTime      *string `json:"ActualReturnDepartureDateTime"`
	EstimatedReturnArrivalDateTime     *string `json:"EstimatedReturnArrivalDateTime"`
	ActualReturnArrivalDateTime        *string `json:"ActualReturnArrivalDateTime"`
}

func (m *Metrc) getTransfersByStatus(status string, licenseNumber string, lastModifiedStart *string, lastModifiedEnd *string) ([]Transfer, error) {
	endpoint := fmt.Sprintf("transfers/v1/%s", status)
	endpoint += fmt.Sprintf("?licenseNumber=%s", licenseNumber)
	if lastModifiedStart != nil {
		endpoint += fmt.Sprintf("?lastModifiedStart=%s", *lastModifiedStart)
	}
	if lastModifiedEnd != nil {
		endpoint += fmt.Sprintf("?lastModifiedEnd=%s", *lastModifiedEnd)
	}

	var tr []Transfer
	responseBody, err := m.client.Get(endpoint)
	if err != nil {
		return tr, fmt.Errorf("could not get endpoint: %s", err)
	}

	err = json.Unmarshal(responseBody, &tr)
	if err != nil {
		return tr, fmt.Errorf("could not unmarshal body: %s", err)
	}

	return tr, nil
}

// GetTransfersIncoming gets all incoming Transfers from Metrc.
// See: https://testing-api-ca.metrc.com/Documentation#Transfers.get_transfers_v1_incoming.GET
func (m *Metrc) GetTransfersIncoming(licenseNumber string, lastModifiedStart *string, lastModifiedEnd *string) ([]Transfer, error) {
	transfers, err := m.getTransfersByStatus("incoming", licenseNumber, lastModifiedStart, lastModifiedEnd)
	if err != nil {
		return []Transfer{}, fmt.Errorf("could not get incoming transfers: %s", err)
	}

	return transfers, nil
}

// GetTransfersOutgoing gets all outgoing Transfers from Metrc.
// See: https://testing-api-ca.metrc.com/Documentation#Transfers.get_transfers_v1_outgoing.GET
func (m *Metrc) GetTransfersOutgoing(licenseNumber string, lastModifiedStart *string, lastModifiedEnd *string) ([]Transfer, error) {
	transfers, err := m.getTransfersByStatus("outgoing", licenseNumber, lastModifiedStart, lastModifiedEnd)
	if err != nil {
		return []Transfer{}, fmt.Errorf("could not get outgoing transfers: %s", err)
	}

	return transfers, nil
}

// GetTransfersRejected gets all rejected Transfers from Metrc.
// See: https://testing-api-ca.metrc.com/Documentation#Transfers.get_transfers_v1_rejected.GET
func (m *Metrc) GetTransfersRejected(licenseNumber string, lastModifiedStart *string, lastModifiedEnd *string) ([]Transfer, error) {
	transfers, err := m.getTransfersByStatus("rejected", licenseNumber, lastModifiedStart, lastModifiedEnd)
	if err != nil {
		return []Transfer{}, fmt.Errorf("could not get rejected transfers: %s", err)
	}

	return transfers, nil
}

// TransferDelivery represents a Delivery of a Transfer.
// See: https://testing-api-ca.metrc.com/Documentation#Transfers.get_transfers_v1_{id}_deliveries.GET
type TransferDelivery struct {
	Id                               int      `json:"Id"`
	RecipientFacilityLicenseNumber   string   `json:"RecipientFacilityLicenseNumber"`
	RecipientFacilityName            string   `json:"RecipientFacilityName"`
	ShipmentTypeName                 string   `json:"ShipmentTypeName"`
	ShipmentTransactionType          string   `json:"ShipmentTransactionType"`
	EstimatedDepartureDateTime       string   `json:"EstimatedDepartureDateTime"`
	ActualDepartureDateTime          *string  `json:"ActualDepartureDateTime"`
	GrossWeight                      *float64 `json:"GrossWeight"`
	GrossUnitOfWeightId              *int     `json:"GrossUnitOfWeightId"`
	GrossUnitOfWeightName            *string  `json:"GrossUnitOfWeightName"`
	PlannedRoute                     string   `json:"PlannedRoute"`
	DeliveryPackageCount             int      `json:"DeliveryPackageCount"`
	DeliveryReceivedPackageCount     int      `json:"DeliveryReceivedPackageCount"`
	ReceivedDateTime                 string   `json:"ReceivedDateTime"`
	EstimatedReturnDepartureDateTime *string  `json:"EstimatedReturnDepartureDateTime"`
	ActualReturnDepartureDateTime    *string  `json:"ActualReturnDepartureDateTime"`
	EstimatedReturnArrivalDateTime   *string  `json:"EstimatedReturnArrivalDateTime"`
	ActualReturnArrivalDateTime      *string  `json:"ActualReturnArrivalDateTime"`
	RejectedPackagesReturned         bool     `json:"RejectedPackagesReturned"`
}

// GetTransfersDeliveriesById retrieves deliveries with an id.
// See: https://testing-api-ca.metrc.com/Documentation#Transfers.get_transfers_v1_{id}_deliveries.GET
func (m *Metrc) GetTransfersDeliveriesById(id int) ([]TransferDelivery, error) {
	endpoint := fmt.Sprintf("transfers/v1/%d/deliveries", id)

	var dr []TransferDelivery
	responseBody, err := m.client.Get(endpoint)
	if err != nil {
		return dr, fmt.Errorf("could not get response: %s", err)
	}

	err = json.Unmarshal(responseBody, &dr)
	if err != nil {
		return dr, fmt.Errorf("could not unmarshal response: %s", err)
	}

	return dr, nil
}

// TransferTransporter represents a Transporter for a Transfer.
// See: https://testing-api-ca.metrc.com/Documentation#Transfers.get_transfers_v1_{id}_transporters.GET
type TransferTransporter struct {
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber"`
	TransporterFacilityName          string `json:"TransporterFacilityName"`
	TransporterDirection             int    `json:"TransporterDirection"`
}

// GetTransfersTransportersById retrieves transporters of transfers by id.
// See: https://testing-api-ca.metrc.com/Documentation#Transfers.get_transfers_v1_{id}_transporters.GET
func (m *Metrc) GetTransfersTransportersById(id int) ([]TransferTransporter, error) {
	endpoint := fmt.Sprintf("transfers/v1/%d/transporters", id)

	var tr []TransferTransporter
	responseBody, err := m.client.Get(endpoint)
	if err != nil {
		return tr, fmt.Errorf("could not get response: %s", err)
	}

	err = json.Unmarshal(responseBody, &tr)
	if err != nil {
		return tr, fmt.Errorf("could not unmarshal response: %s", err)
	}

	return tr, nil
}

// TransferTransporterDetails represents the details of transporters for transfers by id.
// See: https://testing-api-ca.metrc.com/Documentation#Transfers.get_transfers_v1_{id}_transporters_details.GET
type TransferTransporterDetails struct {
	DriverName                      string  `json:"DriverName"`
	DriverOccupationalLicenseNumber string  `json:"DriverOccupationalLicenseNumber"`
	DriverVehicleLicenseNumber      string  `json:"DriverVehicleLicenseNumber"`
	VehicleMake                     string  `json:"VehicleMake"`
	VehicleModel                    string  `json:"VehicleModel"`
	VehicleLicensePlateNumber       string  `json:"VehicleLicensePlateNumber"`
	ActualDriverStartDateTime       *string `json:"ActualDriverStartDateTime"`
}

// GetTransfersTransportersDetails retrieves the details of transporters for transfers by ID.
// See: https://testing-api-ca.metrc.com/Documentation#Transfers.get_transfers_v1_{id}_transporters_details.GET
func (m *Metrc) GetTransfersTransportersDetailsById(id int) ([]TransferTransporterDetails, error) {
	endpoint := fmt.Sprintf("transfers/v1/%d/transporters/details", id)

	var tr []TransferTransporterDetails
	responseBody, err := m.client.Get(endpoint)
	if err != nil {
		return tr, fmt.Errorf("could not get response: %s", err)
	}

	err = json.Unmarshal(responseBody, &tr)
	if err != nil {
		return tr, fmt.Errorf("could not unmarshal response: %s", err)
	}

	return tr, nil
}

// TransferDeliveryPackage represents a Package in a Delivery within Transfers.
// See: https://testing-api-ca.metrc.com/Documentation#Transfers.get_transfers_v1_delivery_{id}_packages.GET
type TransferDeliveryPackage struct {
	PackageId                           int      `json:"PackageId"`
	PackageLabel                        string   `json:"PackageLabel"`
	PackageType                         string   `json:"PackageType"`
	SourceHarvestNames                  *string  `json:"SourceHarvestNames"`
	SourcePackageLabels                 *string  `json:"SourcePackageLabels"`
	ProductName                         string   `json:"ProductName"`
	ProductCategoryName                 string   `json:"ProductCategoryName"`
	ItemStrainName                      *string  `json:"ItemStrainName"`
	ItemUnitCbdPercent                  *float64 `json:"ItemUnitCbdPercent"`
	ItemUnitCbdContent                  *float64 `json:"ItemUnitCbdContent"`
	ItemUnitCbdContentUnitOfMeasureName *string  `json:"ItemUnitCbdContentUnitOfMeasureName"`
	ItemUnitThcPercent                  *float64 `json:"ItemUnitThcPercent"`
	ItemUnitThcContent                  *float64 `json:"ItemUnitThcContent"`
	ItemUnitThcContentUnitOfMeasureName *string  `json:"ItemUnitThcContentUnitOfMeasureName"`
	ItemUnitVolume                      *float64 `json:"ItemUnitVolume"`
	ItemUnitVolumeUnitOfMeasureName     *string  `json:"ItemUnitVolumeUnitOfMeasureName"`
	ItemUnitWeight                      *float64 `json:"ItemUnitWeight"`
	ItemUnitWeightUnitOfMeasureName     *string  `json:"ItemUnitWeightUnitOfMeasureName"`
	ItemServingSize                     *float64 `json:"ItemServingSize"`
	ItemSupplyDurationDays              *int     `json:"ItemSupplyDurationDays"`
	ItemUnitQuantity                    *float64 `json:"ItemUnitQuantity"`
	ItemUnitQuantityUnitOfMeasureName   *string  `json:"ItemUnitQuantityUnitOfMeasureName"`
	LabTestingState                     string   `json:"LabTestingState"`
	ProductionBatchNumber               *int     `json:"ProductionBatchNumber"`
	IsTradeSample                       bool     `json:"IsTradeSample"`
	IsTradeSamplePersistent             bool     `json:"IsTradeSamplePersistent"`
	SourcePackageIsTradeSample          bool     `json:"SourcePackageIsTradeSample"`
	IsDonation                          bool     `json:"IsDonation"`
	SourcePackageIsDonation             bool     `json:"SourcePackageIsDonation"`
	IsTestingSample                     bool     `json:"IsTestingSample"`
	ProductRequiresRemediation          bool     `json:"ProductRequiresRemediation"`
	ContainsRemediatedProduct           bool     `json:"ContainsRemediatedProduct"`
	RemediationDate                     *string  `json:"RemediationDate"`
	ShipmentPackageState                string   `json:"ShipmentPackageState"`
	ShippedQuantity                     float64  `json:"ShippedQuantity"`
	ShippedUnitOfMeasureName            string   `json:"ShippedUnitOfMeasureName"`
	GrossUnitOfWeightName               *string  `json:"GrossUnitOfWeightName"`
	ReceivedQuantity                    float64  `json:"ReceivedQuantity"`
	ReceivedUnitOfMeasureName           string   `json:"ReceivedUnitOfMeasureName"`
}

// GetTransfersDeliveryPackages retrieves Packages for Delivery from Transfers.
// See: https://testing-api-ca.metrc.com/Documentation#Transfers.get_transfers_v1_delivery_{id}_packages.GET
func (m *Metrc) GetTransfersDeliveryPackages(id int) ([]TransferDeliveryPackage, error) {
	endpoint := fmt.Sprintf("transfers/v1/delivery/%d/packages", id)

	var tr []TransferDeliveryPackage
	responseBody, err := m.client.Get(endpoint)
	if err != nil {
		return tr, fmt.Errorf("could not get response: %s", err)
	}

	err = json.Unmarshal(responseBody, &tr)
	if err != nil {
		return tr, fmt.Errorf("could not unmarshal response: %s", err)
	}

	return tr, nil
}

// TransferDeliveryPackageWholesale represents a wholesale package in a delivery from Transfers.
// See: https://testing-api-ca.metrc.com/Documentation#Transfers.get_transfers_v1_delivery_{id}_packages_wholesale.GET
type TransferDeliveryPackageWholesale struct {
	PackageId              int      `json:"PackageId"`
	PackageLabel           string   `json:"PackageLabel"`
	ShipperWholesalePrice  *float64 `json:"ShipperWholesalePrice"`
	ReceiverWholesalePrice *float64 `json:"ReceiverWholesalePrice"`
}

// GetTransfersDeliveryPackagesWholesale retrieves wholesale packages in a delivery from Transfers.
// See:  https://testing-api-ca.metrc.com/Documentation#Transfers.get_transfers_v1_delivery_{id}_packages_wholesale.GET
func (m *Metrc) GetTransfersDeliveryPackagesWholesale(id int) ([]TransferDeliveryPackageWholesale, error) {
	endpoint := fmt.Sprintf("transfers/v1/delivery/%d/packages/wholesale", id)

	var tr []TransferDeliveryPackageWholesale
	responseBody, err := m.client.Get(endpoint)
	if err != nil {
		return tr, fmt.Errorf("could not get response: %s", err)
	}

	err = json.Unmarshal(responseBody, &tr)
	if err != nil {
		return tr, fmt.Errorf("could not unmarshal response: %s", err)
	}

	return tr, nil
}

// TransferDeliveryPackageRequiredLabTestBatches represents a required lab test batch for a delivery package in Transfers.
// See: https://testing-api-ca.metrc.com/Documentation#Transfers.get_transfers_v1_delivery_package_{id}_requiredlabtestbatches.GET
type TransferDeliveryPackageRequiredLabTestBatches struct {
	PackageId        int    `json:"PackageId"`
	LabTestBatchId   int    `json:"LabTestBatchId"`
	LabTestBatchName string `json:"LabTestBatchName"`
}

// GetTransfersDeliveryPackagesRequiredLabTestBatches retrieves required lab test batches for a delivery package.
// See: https://testing-api-ca.metrc.com/Documentation#Transfers.get_transfers_v1_delivery_package_{id}_requiredlabtestbatches.GET
func (m *Metrc) GetTransfersDeliveryPackagesRequiredLabTestBatches(id int) ([]TransferDeliveryPackageRequiredLabTestBatches, error) {
	endpoint := fmt.Sprintf("transfers/v1/delivery/package/%d/requiredlabtestbatches", id)

	var tr []TransferDeliveryPackageRequiredLabTestBatches
	responseBody, err := m.client.Get(endpoint)
	if err != nil {
		return tr, fmt.Errorf("could not get response: %s", err)
	}

	err = json.Unmarshal(responseBody, &tr)
	if err != nil {
		return tr, fmt.Errorf("could not unmarshal response: %s", err)
	}

	return tr, nil
}

// GetTransfersDeliveryPackagesStates retrieves states for delivery packages of Transfers.
// See: https://testing-api-ca.metrc.com/Documentation#Transfers.get_transfers_v1_delivery_packages_states.GET
func (m *Metrc) GetTransfersDeliveryPackagesStates() ([]string, error) {
	endpoint := "transfers/v1/delivery/packages/states"

	var sr []string
	responseBody, err := m.client.Get(endpoint)
	if err != nil {
		return sr, fmt.Errorf("could not get response: %s", err)
	}

	err = json.Unmarshal(responseBody, &sr)
	if err != nil {
		return sr, fmt.Errorf("could not unmarshal response: %s", err)
	}

	return sr, nil
}

// TransferExternalIncoming represents an external/incoming Transfer.
// See: https://testing-api-ca.metrc.com/Documentation#Transfers.post_transfers_v1_external_incoming.POST
type TransferExternalIncoming struct {
	TransferId                       *int                                  `json:"TransferId,omitempty"`
	ShipperLicenseNumber             string                                `json:"ShipperLicenseNumber"`
	ShipperName                      string                                `json:"ShipperName"`
	ShipperMainPhoneNumber           string                                `json:"ShipperMainPhoneNumber"`
	ShipperAddress1                  *string                               `json:"ShipperAddress1"`
	ShipperAddress2                  *string                               `json:"ShipperAddress2"`
	ShipperAddressCity               *string                               `json:"ShipperAddressCity"`
	ShipperAddressState              *string                               `json:"ShipperAddressState"`
	ShipperAddressPostalCode         *string                               `json:"ShipperAddressPostalCode"`
	TransporterFacilityLicenseNumber *string                               `json:"TransporterFacilityLicenseNumber"`
	DriverOccupationalLicenseNumber  *string                               `json:"DriverOccupationalLicenseNumber"`
	DriverName                       *string                               `json:"DriverName"`
	DriverLicenseNumber              *string                               `json:"DriverLicenseNumber"`
	PhoneNumberForQuestions          *string                               `json:"PhoneNumberForQuestions"`
	VehicleMake                      *string                               `json:"VehicleMake"`
	VehicleModel                     *string                               `json:"VehicleModel"`
	VehicleLicensePlateNumber        *string                               `json:"VehicleLicensePlateNumber"`
	Destination                      []TransferExternalIncomingDestination `json:"Destinations"`
}

// TransferExternalIncomingDestination represents a Destination in TransferExternalIncoming struct.
// See: https://testing-api-ca.metrc.com/Documentation#Transfers.post_transfers_v1_external_incoming.POST
type TransferExternalIncomingDestination struct {
	TransferDestinationId      *int                                  `json:"TransferDestinationId,omitempty"`
	RecipientLicenseNumber     string                                `json:"RecipientLicenseNumber"`
	TransferTypeName           string                                `json:"TransferTypeName"`
	PlannedRoute               string                                `json:"PlannedRoute"`
	EstimatedDepartureDateTime string                                `json:"EstimatedDepartureDateTime"`
	EstimatedArrivalDateTime   string                                `json:"EstimatedArrivalDateTime"`
	GrossWeight                *float64                              `json:"GrossWeight,omitempty"`
	GrossUnitOfWeightId        *int                                  `json:"GrossUnitOfWeightId,omitempty"`
	Transporters               []TransferExternalIncomingTransporter `json:"Transporters"`
	Packages                   []TransferExternalIncomingPackage     `json:"Packages"`
}

// TransferExternalIncomingTransporter represents a Transporter within the TransferExternalIncoming struct.
// See: https://testing-api-ca.metrc.com/Documentation#Transfers.post_transfers_v1_external_incoming.POST
type TransferExternalIncomingTransporter struct {
	TransporterFacilityLicenseNumber string  `json:"TransporterFacilityLicenseNumber"`
	DriverOccupationalLicenseNumber  int     `json:"DriverOccupationalLicenseNumber"`
	DriverName                       string  `json:"DriverName"`
	DriverLicenseNumber              string  `json:"DriverLicenseNumber"`
	PhoneNumberForQuestions          string  `json:"PhoneNumberForQuestions"`
	VehicleMake                      string  `json:"VehicleMake"`
	VehicleModel                     string  `json:"VehicleModel"`
	VehicleLicensePlateNumber        string  `json:"VehicleLicensePlateNumber"`
	IsLayover                        bool    `json:"IsLayover"`
	EstimatedDepartureDateTime       string  `json:"EstimatedDepartureDateTime"`
	EstimatedArrivalDateTime         string  `json:"EstimatedArrivalDateTime"`
	TransporterDetails               *string `json:"TransporterDetails"`
}

// TransferExternalIncomingPackage represents a Package within the TransferExternalIncoming struct.
// See: https://testing-api-ca.metrc.com/Documentation#Transfers.post_transfers_v1_external_incoming.POST
type TransferExternalIncomingPackage struct {
	ItemName              string   `json:"ItemName"`
	Quantity              float64  `json:"Quantity"`
	UnitOfMeasureName     string   `json:"UnitOfMeasureName"`
	PackagedDate          string   `json:"PackagedDate"`
	GrossWeight           *float64 `json:"GrossWeight"`
	GrossUnitOfWeightName *string  `json:"GrossUnitOfWeightName"`
	WholesalePrice        *float64 `json:"WholesalePrice"`
}

// PostTransfersExternalIncoming posts external incoming Transfers to Metrc.
// See: https://testing-api-ca.metrc.com/Documentation#Transfers.post_transfers_v1_external_incoming.POST
func (m *Metrc) PostTransfersExternalIncoming(transfers []TransferExternalIncoming, licenseNumber string) ([]byte, error) {
	endpoint := "transfers/v1/external/incoming"
	endpoint += fmt.Sprintf("?licenseNumber=%s", licenseNumber)

	body, err := json.Marshal(transfers)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal body: %s", err)
	}

	resp, err := m.client.Post(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("could not post body: %s", err)
	}

	return resp, nil
}

// PutTransfersExternalIncoming puts external incoming Transfers to Metrc.
// See: https://testing-api-ca.metrc.com/Documentation#Transfers.put_transfers_v1_external_incoming.PUT
func (m *Metrc) PutTransfersExternalIncoming(transfers []TransferExternalIncoming, licenseNumber string) ([]byte, error) {
	endpoint := "transfers/v1/external/incoming"
	endpoint += fmt.Sprintf("?licenseNumber=%s", licenseNumber)

	body, err := json.Marshal(transfers)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal body: %s", err)
	}

	resp, err := m.client.Put(endpoint, body)
	if err != nil {
		return []byte{}, fmt.Errorf("could not post body: %s", err)
	}

	return resp, nil
}

// DeleteTransfersExternalIncoming deletes an external incoming transfer from Metrc.
// See: https://testing-api-ca.metrc.com/Documentation#Transfers.delete_transfers_v1_external_incoming_{id}.DELETE
func (m *Metrc) DeleteTransfersExternalIncoming(id int, licenseNumber string) ([]byte, error) {
	endpoint := "transfers/v1/external/incoming/%d"
	endpoint += fmt.Sprintf("?licenseNumber=%s", licenseNumber)

	resp, err := m.client.Delete(endpoint)
	if err != nil {
		return []byte{}, fmt.Errorf("could not delete: %s", err)
	}

	return resp, nil
}

// GetTransfersTemplates gets transfers templates.
// See: https://testing-api-ca.metrc.com/Documentation#Transfers.get_transfers_v1_templates.GET
func (m *Metrc) GetTransfersTemplates(licenseNumber string, lastModifiedStart *string, lastModifiedEnd *string) ([]Transfer, error) {
	endpoint := "transfers/v1/templates"
	endpoint += fmt.Sprintf("?licenseNumber=%s", licenseNumber)
	if lastModifiedStart != nil {
		endpoint += fmt.Sprintf("?lastModifiedStart=%s", *lastModifiedStart)
	}
	if lastModifiedEnd != nil {
		endpoint += fmt.Sprintf("?lastModifiedEnd=%s", *lastModifiedEnd)
	}

	var tr []Transfer
	responseBody, err := m.client.Get(endpoint)
	if err != nil {
		return tr, fmt.Errorf("could not get response: %s", err)
	}

	err = json.Unmarshal(responseBody, &tr)
	if err != nil {
		return tr, fmt.Errorf("could not unmarshal body: %s", err)
	}

	return tr, nil
}

// TransferTemplateDelivery represents the template for a delivery transfer.
// See: https://testing-api-ca.metrc.com/Documentation#Transfers.get_transfers_v1_templates_{id}_deliveries.GET
type TransferTemplateDelivery struct {
	Id                               int      `json:"Id"`
	RecipientFacilityLicenseNumber   string   `json:"RecipientFacilityLicenseNumber"`
	RecipientFacilityName            string   `json:"RecipientFacilityName"`
	ShipmentTypeName                 string   `json:"ShipmentTypeName"`
	ShipmentTransactionType          string   `json:"ShipmentTransactionType"`
	EstimatedDepartureDateTime       string   `json:"EstimatedDepartureDateTime"`
	ActualDepartureDateTime          *string  `json:"ActualDepartureDateTime"`
	EstimatedArrivalDateTime         string   `json:"EstimatedArrivalDateTime"`
	ActualArrivalDateTime            *string  `json:"ActualArrivalDateTime"`
	GrossWeight                      *float64 `json:"GrossWeight"`
	GrossUnitOfWeightId              *int     `json:"GrossUnitOfWeightId"`
	GrossUnitOfWeightName            *string  `json:"GrossUnitOfWeightName"`
	PlannedRoute                     string   `json:"PlannedRoute"`
	DeliveryPackageCount             int      `json:"DeliveryPackageCount"`
	DeliveryReceivedPackageCount     int      `json:"DeliveryReceivedPackageCount"`
	ReceivedDateTime                 string   `json:"ReceivedDateTime"`
	EstimatedReturnDepartureDateTime *string  `json:"EstimatedReturnDepartureDateTime"`
	ActualReturnDepartureDateTime    *string  `json:"ActualReturnDepartureDateTime"`
	EstimatedReturnArrivalDateTime   *string  `json:"EstimatedReturnArrivalDateTime"`
	ActualReturnArrivalDateTime      *string  `json:"ActualReturnArrivalDateTime"`
	RejectedPackagesReturned         bool     `json:"RejectedPackagesReturned"`
}

// TransferTemplateDelivery represents a delivery template.
// See: https://testing-api-ca.metrc.com/Documentation#Transfers.get_transfers_v1_templates_{id}_deliveries.GET
func (m *Metrc) GetTransfersTemplatesDeliveries(id int) ([]TransferTemplateDelivery, error) {
	endpoint := fmt.Sprintf("transfers/v1/templates/%d/deliveries", id)

	var tr []TransferTemplateDelivery
	responseBody, err := m.client.Get(endpoint)
	if err != nil {
		return tr, fmt.Errorf("could not get response: %s", err)
	}

	err = json.Unmarshal(responseBody, &tr)
	if err != nil {
		return tr, fmt.Errorf("could not unmarshal body: %s", err)
	}

	return tr, nil
}

// TransferTemplateTransporter represents a transporter template for a transfer.
// See: https://testing-api-ca.metrc.com/Documentation#Transfers.get_transfers_v1_templates_{id}_transporters.GET
type TransferTemplateTransporter struct {
	TransporterFacilityLicenseNumber string `json:"TransporterFacilityLicenseNumber"`
	TransporterFacilityName          string `json:"TransporterFacilityName"`
	TransporterDirection             int    `json:"TransporterDirection"`
}

// GetTransfersTemplatesTransporters represents a transporter template for a Transfer.
// Note: the id parameter represents a Transfer Template Delivery ID, not a Manifest Number.
// See: https://testing-api-ca.metrc.com/Documentation#Transfers.get_transfers_v1_templates_{id}_transporters.GET
func (m *Metrc) GetTransfersTemplatesTransporters(id int) ([]TransferTemplateTransporter, error) {
	endpoint := fmt.Sprintf("transfers/v1/templates/%d/transporters", id)

	var tr []TransferTemplateTransporter
	responseBody, err := m.client.Get(endpoint)
	if err != nil {
		return tr, fmt.Errorf("could not get response: %s", err)
	}

	err = json.Unmarshal(responseBody, &tr)
	if err != nil {
		return tr, fmt.Errorf("could not unmarshal body: %s", err)
	}

	return tr, nil
}

// GetTransfersTemplatesTransportersDetails return the details of transporters for a template.
// Note: id represents a Transfer Template Delivery ID, not a Manifest Number.
// See: https://testing-api-ca.metrc.com/Documentation#Transfers.get_transfers_v1_templates_{id}_transporters_details.GET
func (m *Metrc) GetTransfersTemplatesTransportersDetails(id int) ([]TransferTransporterDetails, error) {
	endpoint := fmt.Sprintf("transfers/v1/templates/%d/transporters/details", id)

	var tr []TransferTransporterDetails
	responseBody, err := m.client.Get(endpoint)
	if err != nil {
		return tr, fmt.Errorf("could not get response: %s", err)
	}

	err = json.Unmarshal(responseBody, &tr)
	if err != nil {
		return tr, fmt.Errorf("could not unmarshal body: %s", err)
	}

	return tr, nil
}

// GetTransfersTemplatesDeliveryPackages gets a template for a package delivery in Transfers
// Note: id represents a Transfer Template Delivery ID, not a Manifest Number
// See: https://testing-api-ca.metrc.com/Documentation#Transfers.get_transfers_v1_templates_delivery_{id}_packages.GET
func (m *Metrc) GetTransfersTemplatesDeliveryPackages(id int) ([]TransferDeliveryPackage, error) {
	endpoint := fmt.Sprintf("transfers/v1/templates/delivery/%d/packages", id)

	var tr []TransferDeliveryPackage
	responseBody, err := m.client.Get(endpoint)
	if err != nil {
		return tr, fmt.Errorf("could not get response: %s", err)
	}

	err = json.Unmarshal(responseBody, &tr)
	if err != nil {
		return tr, fmt.Errorf("could not unmarshal body: %s", err)
	}

	return tr, nil
}

// TransferTemplate represents a Template in a Transfer.
// See: https://testing-api-ca.metrc.com/Documentation#Transfers.put_transfers_v1_templates.PUT
type TransferTemplate struct {
	TransferTemplateId               *int                          `json:"TransferTemplateId,omitempty"`
	Name                             string                        `json:"Name"`
	TransporterFacilityLicenseNumber *string                       `json:"TransporterFacilityLicenseNumber"`
	DriverOccupationalLicenseNumber  *string                       `json:"DriverOccupationalLicenseNumber"`
	DriverName                       *string                       `json:"DriverName"`
	DriverLicenseNumber              *string                       `json:"DriverLicenseNumber"`
	PhoneNumberForQuestions          *string                       `json:"PhoneNumberForQuestions"`
	VehicleMake                      *string                       `json:"VehicleMake"`
	VehicleModel                     *string                       `json:"VehicleModel"`
	VehicleLicensePlateNumber        *string                       `json:"VehicleLicensePlateNumber"`
	Destinations                     []TransferTemplateDestination `json:"TransferExternalIncomingDestination"`
}

// TransferTemplateDestination represents a Destination in a Template for a Transfer.
// See: https://testing-api-ca.metrc.com/Documentation#Transfers.post_transfers_v1_templates.POST
type TransferTemplateDestination struct {
	RecipientLicenseNumber     string                                   `json:"RecipientLicenseNumber"`
	TransferTypeName           string                                   `json:"TransferTypeName"`
	PlannedRoute               string                                   `json:"PlannedRoute"`
	EstimatedDepartureDateTime string                                   `json:"EstimatedDepartureDateTime"`
	EstimatedArrivalDateTime   string                                   `json:"EstimatedArrivalDateTime"`
	Transporters               []TransferTemplateDestinationTransporter `json:"Transporters"`
	Packages                   []TransferTemplateDestinationPackage     `json:"Packages"`
}

// TransferTemplateDestinationTransporter represents a Transporter in a Destination of a Template for a Transfer.
// See: https://testing-api-ca.metrc.com/Documentation#Transfers.post_transfers_v1_templates.POST
type TransferTemplateDestinationTransporter struct {
	TransporterFacilityLicenseNumber string  `json:"TransporterFacilityLicenseNumber"`
	DriverOccupationalLicenseNumber  string  `json:"DriverOccupationalLicenseNumber"`
	DriverName                       string  `json:"DriverName"`
	DriverLicenseNumber              string  `json:"DriverLicenseNumber"`
	PhoneNumberForQuestions          string  `json:"PhoneNumberForQuestions"`
	VehicleMake                      string  `json:"VehicleMake"`
	VehicleModel                     string  `json:"VehicleModel"`
	VehicleLicensePlateNumber        string  `json:"VehicleLicensePlateNumber"`
	IsLayover                        bool    `json:"IsLayover"`
	EstimatedDepartureDateTime       string  `json:"EstimatedDepartureDateTime"`
	EstimatedArrivalDateTime         string  `json:"EstimatedArrivalDateTime"`
	TransporterDetails               *string `json:"TransporterDetails"`
}

// TransferTemplateDestinationPackage represents a Package in a Destination of a Template for a Transfer.
// See: https://testing-api-ca.metrc.com/Documentation#Transfers.post_transfers_v1_templates.POST
type TransferTemplateDestinationPackage struct {
	PackageLabel   string   `json:"PackageLabel"`
	WholesalePrice *float64 `json:"WholesalePrice"`
}

// PostTransfersTemplates posts templates for Transfers.
// See: https://testing-api-ca.metrc.com/Documentation#Transfers.post_transfers_v1_templates.POST
func (m *Metrc) PostTransfersTemplates(transfers []TransferTemplate, licenseNumber string) ([]byte, error) {
	endpoint := "transfers/v1/templates"
	endpoint += fmt.Sprintf("?licenseNumber=%s", licenseNumber)

	responseBody, err := json.Marshal(transfers)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal body: %s", err)
	}

	resp, err := m.client.Post(endpoint, responseBody)
	if err != nil {
		return []byte{}, fmt.Errorf("could not post body: %s", err)
	}

	return resp, nil
}

// PutTransfersTemplates puts templates for Transfers.
// These typically have a top level TransferTemplateId, unlike the initial POST.
// See: https://testing-api-ca.metrc.com/Documentation#Transfers.put_transfers_v1_templates.PUT
func (m *Metrc) PutTransfersTemplates(transfers []TransferTemplate, licenseNumber string) ([]byte, error) {
	endpoint := "transfers/v1/templates"
	endpoint += fmt.Sprintf("licenseNumber=%s", licenseNumber)

	responseBody, err := json.Marshal(transfers)
	if err != nil {
		return []byte{}, fmt.Errorf("could not marshal body: %s", err)
	}

	resp, err := m.client.Put(endpoint, responseBody)
	if err != nil {
		return []byte{}, fmt.Errorf("could not put body: %s", err)
	}

	return resp, nil
}

// DeleteTransfersTemplates deletes a Transfer template with the given id.
// See: https://testing-api-ca.metrc.com/Documentation#Transfers.delete_transfers_v1_templates_{id}.DELETE
func (m *Metrc) DeleteTransfersTemplates(id int, licenseNumber string) ([]byte, error) {
	endpoint := fmt.Sprintf("transfers/v1/templates/%d", id)
	endpoint += fmt.Sprintf("?licenseNumber=%s", licenseNumber)

	resp, err := m.client.Delete(endpoint)
	if err != nil {
		return []byte{}, fmt.Errorf("could not delete: %s", err)
	}

	return resp, nil
}

// TransferType represents the Type of a Transfer.
// See: https://testing-api-ca.metrc.com/Documentation#Transfers.get_transfers_v1_types.GET
type TransferType struct {
	Name                           string `json:"Name"`
	ForLicensedShipments           bool   `json:"ForLicensedShipments"`
	ForExternalIncomingShipments   bool   `json:"ForExternalIncomingShipments"`
	ForExternalOutgoingShipments   bool   `json:"ForExternalOutgoingShipments"`
	RequiresDestinationGrossWeight bool   `json:"RequiresDestinationGrossWeight"`
	RequiresPackagesGrossWeight    bool   `json:"RequiresPackagesGrossWeight"`
}

// GetTransfersTypes retrieves the types of Transfers.
// See: https://testing-api-ca.metrc.com/Documentation#Transfers.get_transfers_v1_types.GET
func (m *Metrc) GetTransfersTypes(licenseNumber string) ([]TransferType, error) {
	endpoint := "transfers/v1/types"
	endpoint += fmt.Sprintf("?licenseNumber=%s", licenseNumber)

	var tr []TransferType
	responseBody, err := m.client.Get(endpoint)
	if err != nil {
		return tr, fmt.Errorf("could not get response: %s", err)
	}

	err = json.Unmarshal(responseBody, &tr)
	if err != nil {
		return tr, fmt.Errorf("could not unmarshal body: %s", err)
	}

	return tr, nil
}
