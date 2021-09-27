package models

import "time"

type Revenue struct {
	PSOGeolocation                   string `csv:"PSO Geo" json:"psoGeo"`
	PSORegion                        string `csv:"PSO Region" json:"psoRegion"`
	PSOTerritory                     string `csv:"PSO Territory" json:"psoTerritory"`
	Name                             string `csv:"Engagement Name" json:"engagementName"`
	ID                               string `csv:"Engagement Id" json:"engagementId"`
	EngagementStatus                 string `csv:"Engagement Status" json:"engagementStatus"`
	EngagementType                   string `csv:"Engagement Type" json:"engagementType"`
	BusinessUnit                     string `csv:"Business Unit" json:"businessUnit"`
	BillingCustomer                  string `csv:"Bill To Customer" json:"billingCustomer"`
	EndCustomer                      string `csv:"End Customer" json:"endCustomer"`
	BudgetedHours                    string `csv:"Budgeted Hours Total" json:"budgetedHoursTotal"`
	PlannedHours                     string `csv:"Planned Hours Total" json:"plannedHoursTotal"`
	Created                          string `csv:"Created On" json:"createdOn"`
	PaymentType                      string `csv:"Payment Type" json:"paymentType"`
	RevenueType                      string `csv:"Rec Rev Type" json:"revenueType"`
	BillingSKU                       string `csv:"Billing SKU" json:"billingSKU"`
	EngagementAmount                 string `csv:"Engagement Amount" json:"engagementAmount"`
	Currency                         string `csv:"Billing Currency" json:"billingCurrency"`
	RedeemedAmount                   string `csv:"Total Billed-Redeemed" json:"totalBilled"`
	RemainingAmount                  string `csv:"Remaining Amount" json:"remainingAmount"`
	OperationsAnalyst                string `csv:"Operation Analysts" json:"operationsAnalysts"`
	TimeApprover                     string `csv:"Time Approver" json:"timeApprover"`
	RedeemedByDateRange              string `csv:"Billed-Redeemed by Date Range" json:"billedByDateRange"`
	RemainingByDateRange             string `csv:"Remaining Forecast\nby Date Range" json:"remainingForecast"`
	ForecastHours                    string `csv:"Forecast Hours this Qtr" json:"forecastHoursThisQtr"`
	ForecastRevenue                  string `csv:"Forecast Revenue  of Current  Qtr USD" json:"forecastRevenueThisQtrUSD"`
	BillableHoursAsOfPriorQtr        string `csv:"Total Billable hours as of  prior Qtr end" json:"billedAsOfPriorQtr"`
	BillableHoursThisQtr             string `csv:"Total Billable hours this Qtr" json:"billedThisQtr"`
	ManagementRevenueAsOfPriorQtr    string `csv:"Mgmt Rev as of Prior Qtr" json:"mgmtRevenueAsOfPriorQtr"`
	ManagementRevenueAsOfPriorQtrUSD string `csv:"Mgmt Rev as of Prior Qtr USD" json:"mgmtRevenueAsOfPriorQtrUSD"`
	ManagementRevenueThisQuarter     string `csv:"Qtr Management Revenue To-Date" json:"mgmtRevenueThisQtr"`
	ManagementRevenueThisQuarterUSD  string `csv:"Qtr Management Revenue To-date USD" json:"mgmtRevenueThisQtrUSD"`
	InternalHoursThisQtr             string `csv:"Internal Hours - QTD" json:"internalHoursThisQtr"`
	SubbingHoursThisQtr              string `csv:"Subbing Hours - QTD" json:"subbingHoursThisQtr"`
	ExpensesRedeemedToDate           string `csv:"Expenses Billed-Redeemed By Date Range" json:"expensesRedeemedByDateRange"`
	ContractStatus                   string `csv:"Contract Status" json:"contractStatus"`
	PurchaseOrder                    string `csv:"PO Number" json:"poNumber"`
	RedeemedByDateRangeUSD           string `csv:"Billed-Redeemed by Date Range (USD)" json:"billedByDateRangeUSD"`
	LastTimeEntry                    string `csv:"Last Time Entry Date" json:"lastTimeEntryDate"`
	StatusDuringDateRange            string `csv:"Status During Date Range" json:"statusDuringDateRange"`
	RedeemedAmountUSD                string `csv:"Total Billed -Redeemed (USD)" json:"totalBilledUSD"`
	EngagementAmountUSD              string `csv:"Engagement Amount in USD" json:"engagementAmountUSD"`
	TotalActualHours                 string `csv:"Total Actual Hours" json:"totalActualHours"`
	RevenueForecastForQtr            string `csv:"Hours Based Forecast Revenue of Current Qtr" json:"hoursBasedRevenueForecastThisQtr"`
	RevenueForecastForQtrUSD         string `csv:"Hours Based Forecast Revenue of Current Qtr(USD)" json:"hoursBasedRevenueForecastThisQtrUSD"`
	PlannedHoursForQtr               string `csv:"Plannned Hrs for Qtr" json:"plannedHoursThisQtr"`
	Program                          string `csv:"Program" json:"program"`
	ProjectProfile                   string `csv:"Project Profile" json:"projectProfile"`
	SalesExecutive                   string `csv:"SSM" json:"ssm"`
	ProjectManager                   string `csv:"Project\nManager" json:"projectManager"`
	ClosedData                       string `csv:"Engagement\nClosed Date" json:"dateClosed"`
	AgreementNumber                  string `csv:"Agreement Number" json:"agreementNumber"`
	OpportunityNumber                string `csv:"Opportunity Number" json:"opportunityNumber"`
}

type Credit struct {
	ExpiredQtr            string `csv:"Expired Quarter" json:"expiredQtr"`
	ExpiredMonth          string `csv:"Expired Month" json:"expiredMonth"`
	ExpiredDate           string `csv:"Credits Expired Date" json:"expiredDate"`
	StartDate             string `csv:"Credit Start Date" json:"startDate"`
	Geolocation           string `csv:"Geo" json:"globalGeo"`
	Region                string `csv:"Region" json:"globalRegion"`
	SubRegion             string `csv:"Sub Region" json:"globalSubRegion"`
	Territory             string `csv:"Territory" json:"globalTerritory"`
	Segment               string `csv:"Segment" json:"globalSegment"`
	PSOGeolocation        string `csv:"PSO Geo" json:"psoGeo"`
	PSORegion             string `csv:"PSO Region" json:"psoRegion"`
	PSOTerritory          string `csv:"PSO Territory" json:"psoTerritory"`
	CSE                   string `csv:"CSE" json:"cse"`
	ESE                   string `csv:"ESE" json:"ese"`
	CustomerName          string `csv:"Customer Name" json:"customerName"`
	CustomerNumber        string `csv:"Customer Number SFDC" json:"customerNumber"`
	SalesOrder            string `csv:"Sales Order ID" json:"salesOrderId"`
	MylearnAccountID      string `csv:"Mylearn Account ID" json:"mylearnId"`
	RedemptionLOB         string `csv:"Redemption LOB" json:"redemptionLOB"`
	PlannedRedemptionDate string `csv:"Planned Redemption Date" json:"plannedRedemptionDate"`
	RedemptionStatus      string `csv:"Redemption Status" json:"redemptionStatus"`
	ID                    string `csv:"Class Name" json:"engagementId"`
	ProjectManager        string `csv:"Project Manager" json:"projectManager"`
	CreditStatusSummary   string `csv:"Credit Status Summary" json:"statusSummary"`
	BusinessUnit          string `csv:"Engagement BU" json:"businessUnit"`
	AllocatedAmount       string `csv:"Allocated Amount" json:"allocatedAmount"`
	AllocatedNumber       string `csv:"Allocated Number" json:"allocatedNumber"`
	AllocatedComments     string `csv:"Allocated Comments" json:"comments"`
}

type Gar struct {
	PSOGeolocation                string `csv:"PSO Geo" json:"psoGeo"`
	PSORegion                     string `csv:"PSO Region" json:"psoRegion"`
	PSOSubRegion                  string `csv:"PSO Subregion" json:"psoSubRegion"`
	ID                            string `csv:"Engagement Id" json:"engagementId"`
	Name                          string `csv:"Engagement Name" json:"engagementName"`
	ProjectManager                string `csv:"Field Project Manager" json:"projectManager"`
	ProjectProfile                string `csv:"Project Profile" json:"projectProfile"`
	BudgetStatus                  string `csv:"Budget Status (PMO)" json:"budgetStatus"`
	ProjectStatus                 string `csv:"Project Status (PMO)" json:"projectStatus"`
	CustomerSatisfaction          string `csv:"Customer Satisfaction (PMO)" json:"customerSatisfaction"`
	ScheduleStatus                string `csv:"Schedule Status (PMO)" json:"scheduleStatus"`
	Comments                      string `csv:"comments" json:"comments"`
	MitigationAction              string `csv:"Mitigation Action" json:"mitigationAction"`
	EngagementManager             string `csv:"Engagement Manager" json:"engagementManager"`
	PracticeManager               string `csv:"Practice Manager" json:"practiceManager"`
	BudgetStatusUpdated           string `csv:"Budget Status UpdateOn" json:"budgetStatusUpdatedOn"`
	BudgetStatusUpdatedBy         string `csv:"Budget Status UpdateBy" json:"budgetStatusUpdatedBy"`
	CustomerSatisfactionUpdated   string `csv:"Customer Satisfaction UpdatedOn" json:"customerSatisfactionUpdatedOn"`
	CustomerSatisfactionUpdatedBy string `csv:"Customer Satisfaction UpdatedBy" json:"customerSatisfactionUpdatedBy"`
	ScheduleStatusUpdated         string `csv:"Schedule Status UpdatedOn" json:"scheduleStatusUpdatedOn"`
	ScheduleStatusUpdateBy        string `csv:"Schedule Status UpdatedBy" json:"scheduleStatusUpdatedBy"`
	ProjectStatusUpdated          string `csv:"Project Status UpdatedOn" json:"projectStatusUpatedBy"`
	ProjectStatusUpdatedBy        string `csv:"Project Status UpdatedBy" json:"projectStatusUpdatedOn"`
}

type Roles struct {
	ID             string `json:"engagementId"`
	EAM            string `json:"eam"`
	CSE            string `json:"cse"`
	ProjectManager string `json:"projectManager"`
	LeadConsultant string `json:"consultant"`
	TAM            string `json:"tam"`
}

type Comments struct {
	ID   string   `json:"engagementId"`
	Text []string `json:"comments"`
}

type Statistics struct {
	ID              string  `json:"engagementId"`
	PercentComplete float64 `json:"percentComplete"`
}

type EngagementSummary struct {
	EngagementTerritory  string      `json:"territory"`
	EngagementID         string      `json:"id"`
	EngagementName       string      `json:"name"`
	CustomerName         string      `json:"customerName"`
	EngagementAmount     float64     `json:"engagementAmount"`
	RevenueType          string      `json:"revenueType"`
	Phase                string      `json:"phase"`
	ProjectStatus        string      `json:"projectStatus"`
	BudgetStatus         string      `json:"budgetStatus"`
	ScheduleStatus       string      `json:"scheduleStatus"`
	CustomerSatisfaction string      `json:"customerSatisfaction"`
	Statistics           *Statistics `json:"stats"`
	Roles                *Roles      `json:"roles"`
	CreditExpiry         string      `json:"creditExpiry"`
	Comments             *Comments   `json:"comments"`
}

type Portfolio struct {
	CustomerName string              `json:"customerName"`
	Summaries    []EngagementSummary `json:"engagements"`
	Comments     Comments            `json:"comments"`
}

type PortfolioReport struct {
	Date       time.Time   `json:"date"`
	Portfolios []Portfolio `json:"report"`
}
