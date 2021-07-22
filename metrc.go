package metrc

// Metrc exposes the Metrc API endpoints to an external developer.
// Implements MetrcInterface.
type Metrc struct {
	Client ClientInterface
}

// MetrcInterface specifies the methods through which an external developer can call the Metrc API.
type MetrcInterface interface {
	// facilities
	Facilities() ([]Facilities, error)

	// items
	GetItemsById(id int, licenseNumber *string) (ItemGet, error)
	GetItemsActive(licenseNumber *string) ([]ItemGet, error)
	GetItemsCategories(licenseNumber *string) ([]ItemCategory, error)
	PostItemsCreate(items []ItemPost, licenseNumber *string) ([]byte, error)
	PostItemsUpdate(items []ItemPost, licenseNumber *string) ([]byte, error)
	DeleteItemById(id int, licenseNumber *string) ([]byte, error)

	// locations
	GetLocationsById(id int, licenseNumber *string) (LocationGet, error)
	GetLocationsActive(licenseNumber *string) ([]LocationGet, error)
	GetLocationsTypes(licenseNumber *string) ([]LocationGet, error)
	PostLocationsCreate(locs []LocationPost, licenseNumber *string) ([]byte, error)
	PostLocationsUpdate(locs []LocationPost, licenseNumber *string) ([]byte, error)
	DeleteLocationById(id int, licenseNumber *string) ([]byte, error)

	// packages
	GetPackagesById(id int, licenseNumber *string) (PackageGet, error)
	GetPackagesByLabel(label string, licenseNumber *string) (PackageGet, error)
	GetPackagesActive(licenseNumber string, lastModifiedStart *string, lastModifiedEnd *string) ([]PackageGet, error)
	GetPackagesOnHold(licenseNumber string, lastModifiedStart *string, lastModifiedEnd *string) ([]PackageGet, error)
	GetPackagesInactive(licenseNumber string, lastModifiedStart *string, lastModifiedEnd *string) ([]PackageGet, error)
	GetPackagesTypes() ([]string, error)
	GetPackagesAdjustReasons(licenseNumber string) ([]PackageAdjustReasons, error)
	PostPackagesCreate(packages []PackagePost, licenseNumber string) ([]byte, error)
	PostPackagesCreateTesting(packages []PackagePost, licenseNumber string) ([]byte, error)
	PostPackagesCreatePlanting(packages []PackagePost, licenseNumber string) ([]byte, error)
	ChangePackagesItem(packageItems []PackageItem, licenseNumber string) ([]byte, error)
	ChangePackagesNote(packageNotes []PackageNote, licenseNumber string) ([]byte, error)
	ChangePackagesLocations(packageLocations []PackageLocation, licenseNumber string) ([]byte, error)
	AdjustPackages(packageAdjusts []PackageAdjust, licenseNumber string) ([]byte, error)
	FinishPackages(packageFinishes []PackageFinish, licenseNumber string) ([]byte, error)
	UnfinishPackages(packageUnfinishes []PackageUnfinish, licenseNumber string) ([]byte, error)
	RemediatePackages(packageRemediates []PackageRemediate, licenseNumber string) ([]byte, error)

	// sales
	GetSalesCustomerTypes() ([]string, error)
	GetSalesReceiptsActive(licenseNumber string, salesDateStart *string, salesDateEnd *string, lastModifiedStart *string, lastModifiedEnd *string) ([]SalesReceiptGet, error)
	GetSalesReceiptsInactive(licenseNumber string, salesDateStart *string, salesDateEnd *string, lastModifiedStart *string, lastModifiedEnd *string) ([]SalesReceiptGet, error)
	GetSalesReceiptsById(id int, licenseNumber *string) (SalesReceiptGet, error)
	PostSalesReceipts(receipts []SalesReceiptPost, licenseNumber string) ([]byte, error)
	PutSalesReceipts(receipts []SalesReceiptPost, licenseNumber string) ([]byte, error)
	DeleteSalesReceiptsById(id int, licenseNumber string) ([]byte, error)
	GetSalesTransactions(licenseNumber string) ([]SalesTransactionGet, error)
	GetSalesTransactionsByDates(salesDateStart string, salesDateEnd string, licenseNumber string) ([]SalesTransactionGet, error)
	PostSalesTransactions(transactions []SalesTransactionPost, date string, licenseNumber string) ([]byte, error)
	PutSalesTransactions(transactions []SalesTransactionPost, date string, licenseNumber string) ([]byte, error)

	// strains
	GetStrainsById(id int, licenseNumber *string) (Strain, error)
	GetStrainsActive(licenseNumber *string) ([]Strain, error)
	PostStrainsCreate(strains []Strain, licenseNumber *string) ([]byte, error)
	PostStrainsUpdate(strains []Strain, licenseNumber *string) ([]byte, error)
	DeleteStrainById(id int, licenseNumber *string) ([]byte, error)

	// units of measure
	GetUnitsOfMeasure() ([]UnitsOfMeasure, error)
}

func MakeIntegrationMetrc(vendorKey string, userKey string) *Metrc {
	var ci ClientInterface = MakeHttpClient(vendorKey, userKey)
	return &Metrc{
		Client: ci,
	}
}
