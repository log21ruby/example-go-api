package model

import "time"

// User model
type User struct {
	ID                               int       `json:"id,omitempty"`
	ExtID                            int       `json:"ext_id,omitempty"`
	UUID                             string    `json:"uuid,omitempty"`
	EmailVerified                    uint8     `json:"email_verified"`
	Email                            string    `json:"email,omitempty"`
	OtpEnabled                       uint8     `json:"otp_enabled"`
	Tel                              string    `json:"tel,omitempty"`
	RiskId                           int       `json:"risk_id,omitempty"`
	Password                         string    `json:"password,omitempty"`
	AccountCode                      string    `json:"account_code,omitempty"`
	FirstName                        string    `json:"first_name"`
	LastName                         string    `json:"last_name"`
	FirstNameTh                      string    `json:"first_name_th"`
	LastNameTh                       string    `json:"last_name_th"`
	LastLogin                        string    `json:"last_login,omitempty"`
	Status                           uint8     `json:"status,omitempty"`
	ReferralID                       int       `json:"referral_id"`
	MobilePin                        string    `json:"mobile_pin,omitempty"`
	Freeze                           uint8     `json:"freeze,omitempty"`
	FreezeCryptoWithdraw             uint8     `json:"freeze_crypto_withdraw"`
	TradingCredit                    float64   `json:"trading_credit,omitempty"`
	LastPWUpdated                    time.Time `json:"last_pw_updated"`
	LastPWStatus                     uint8     `json:"last_pw_status"`
	GoogleAuthenticator              string    `json:"google_authenticator"`
	GoogleAuthenticatorVerified      int       `json:"google_authenticator_verified"`
	GoogleAuthenticatorTradeVerified int       `json:"google_authenticator_trade_verified"`
	LastIP                           string    `json:"last_ip,omitempty"`
	RankVip                          int       `json:"rank_vip,omitempty"`
	UpdatedAt                        string    `json:"updated_at,omitempty"`
	CreatedAt                        string    `json:"created_at,omitempty"`
}
