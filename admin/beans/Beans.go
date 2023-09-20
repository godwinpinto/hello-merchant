package beans

import "time"

type UserMasterStruct struct {
	UUMRowId  string    `json:"uum_row_id" gorm:"column:UUM_ROW_ID"`
	UserId    string    `json:"user_id" gorm:"column:USER_ID"`
	Role      string    `json:"role" gorm:"column:ROLE"`
	Active    string    `json:"active" gorm:"column:ACTIVE"`
	CreatedDt time.Time `json:"created_dt" gorm:"column:CREATED_DT"`
	CreatedBy string    `json:"created_by" gorm:"column:CREATED_BY"`
	UpdatedDt time.Time `json:"updated_dt" gorm:"column:UPDATED_DT"`
	UpdatedBy string    `json:"updated_by" gorm:"column:UPDATED_BY"`
}

type XrplMappingStruct struct {
	UxmRowId  string    `json:"uxm_row_id" gorm:"column:UXM_ROW_ID"`
	UUMRowId  string    `json:"uum_row_id" gorm:"column:UUM_ROW_ID"`
	XrplAcNo  string    `json:"xrpl_ac_no" gorm:"column:XRPL_AC_NO"`
	Active    string    `json:"active" gorm:"column:ACTIVE"`
	CreatedDt time.Time `json:"created_dt" gorm:"column:CREATED_DT"`
	CreatedBy string    `json:"created_by" gorm:"column:CREATED_BY"`
	UpdatedDt time.Time `json:"updated_dt" gorm:"column:UPDATED_DT"`
	UpdatedBy string    `json:"updated_by" gorm:"column:UPDATED_BY"`
}

type SquareUpMappingStruct struct {
	UsmRowId  string    `json:"usm_row_id" gorm:"column:USM_ROW_ID"`
	UUMRowId  string    `json:"uum_row_id" gorm:"column:UUM_ROW_ID"`
	UrlUuid   string    `json:"url_uuid" gorm:"column:URL_UUID"`
	Active    string    `json:"active" gorm:"column:ACTIVE"`
	CreatedDt time.Time `json:"created_dt" gorm:"column:CREATED_DT"`
	CreatedBy string    `json:"created_by" gorm:"column:CREATED_BY"`
	UpdatedDt time.Time `json:"updated_dt" gorm:"column:UPDATED_DT"`
	UpdatedBy string    `json:"updated_by" gorm:"column:UPDATED_BY"`
}

type TransactionMasterStruct struct {
	UTMRowId      string    `json:"utm_row_id" gorm:"column:UTM_ROW_ID"`
	UUMRowId      string    `json:"uum_row_id" gorm:"column:UUM_ROW_ID"`
	Amount        string    `json:"amount" gorm:"column:AMOUNT"`
	Currency      string    `json:"currency" gorm:"column:CURRENCY"`
	Channel       string    `json:"channel" gorm:"column:CHANNEL"`
	Active        string    `json:"active" gorm:"column:ACTIVE"`
	CreatedDt     time.Time `json:"created_dt" gorm:"column:CREATED_DT"`
	TransactionId string    `json:"transaction_id" gorm:"column:TRANSACTION_ID"`
}

type JwtBeanStruct struct {
	Sub string `json:"sub"`
	Id  string `json:"id"`
}

type ResponseStruct struct {
	Status string      `json:"status" `
	Data   interface{} `json:"data" `
}

type DashboardSummaryStruct struct {
	RippleCount   int `json:"ripple_count"  gorm:"column:RIPPLE_COUNT"`
	SquareUpCount int `json:"squareup_count"  gorm:"column:SQUAREUP_COUNT"`
}

type DashboardReportStruct struct {
	DateRange   string `json:"date_range"  gorm:"column:DATE_RANGE"`
	RecordCount int    `json:"record_count"  gorm:"column:RECORD_COUNT"`
	Channel     string `json:"channel"  gorm:"column:CHANNEL"`
}

type DashboardResponseStruct struct {
	RippleCount      int         `json:"ripple_count"`
	SquareUpCount    int         `json:"squareup_count"`
	LastTransactions interface{} `json:"transactions"`
}

type ConfigResponseStruct struct {
	AcNo string `json:"ac_no"`
}
