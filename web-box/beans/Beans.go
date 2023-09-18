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
	UserId   string `json:"user_id" gorm:"column:USER_ID"`
	UUMRowId string `json:"uum_row_id" gorm:"column:UUM_ROW_ID"`
}

type ResponseStruct struct {
	Status string      `json:"status" `
	Data   interface{} `json:"data" `
}
