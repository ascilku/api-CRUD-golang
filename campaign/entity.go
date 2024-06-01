package campaign

import "time"

type Compaign struct {
	ID               int
	UserID           int
	Name             string
	ShortDescription string
	Description      string
	Perks            string
	BackerCount      int
	GoalAmount       int
	Slug             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	CampaignImage    []CampaignImage //`gorm:"foreignkey:CampaignID"` = ini di pake pas beda field atau kata
}

type CampaignImage struct {
	ID         int
	CompaignID int
	FileNamw   string
	IsPrimary  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
