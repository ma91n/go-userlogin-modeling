package model

import "time"

type User struct {
	UserID   string
	Name     string
	UserType int // '1:会員 2:管理者 3:クライアント'
}

type Member struct { // 会員
	UserID string
	Status string // 通常、プレミアム、仮会員、ゲスト
}

type AdminLog struct {
	UserID string
}

type ClientInfo struct {
	UserID      string
	CompanyName string
	PhoneNumber string
}

type UserPersonal struct {
	UserID      string
	ZipCD       string
	Addr        time.Time
	Birthday    time.Time
	PhoneNumber string
}

type UserEmail struct {
	UserID string
	Email  string
}

type UserLoginHashPassword struct {
	UserID         string
	HashedPassword string
}

type UserLoginLine struct {
	UserID    string
	LineID    string
	LineToken string
}

type UserCreditCard struct {
	UserID                string
	CardNumberLast4Digits string
	ExpirationDate        string
}

type UserDocomoPayment struct {
	UserID string
}

type UserLinePayment struct {
	UserID string
}

type UserDefaultPayment struct {
	UserID            string
	UserCreditCardID  int64
	UserLinePaymentID int64
	userDocomoPayment int64
}
