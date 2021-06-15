package models

type SubscriptionModel struct {
	SubscriptionID string `gorm:"column:subscription_id"`
	UserID         string `gorm:"column:user_id"`
	UserEmail      string `gorm:"column:user_email"`
	UserName       string `gorm:"column:user_name"`
	TitleID        string `gorm:"column:title_id"`
	TitleName      string `gorm:"column:title_name"`
}
