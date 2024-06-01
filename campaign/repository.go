package campaign

import "gorm.io/gorm"

type Repository interface {
	GetAll() ([]Compaign, error)
	GetByID(compaignID int) ([]Compaign, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]Compaign, error) {
	var compaigns []Compaign
	err := r.db.Find(&compaigns).Error
	if err != nil {
		return compaigns, err
	}
	return compaigns, nil
}

func (r *repository) GetByID(compaignID int) ([]Compaign, error) {
	var compaign []Compaign
	err := r.db.Where("user_id = ?", compaignID).Preload("CampaignImage", "campaign_images.is_primary = 1").Find(&compaign).Error
	if err != nil {
		return compaign, err
	}
	return compaign, nil
}
