package campaign

import "gorm.io/gorm"

type Repository interface {
	FindAllRepo() ([]Campaign, error)
	FindAllActiveImageRepo() ([]Campaign, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAllRepo() ([]Campaign, error) {
	var keyCampaign []Campaign
	err := r.db.Preload("CampaignImages").Find(&keyCampaign).Error
	if err != nil {
		return keyCampaign, err
	}
	return keyCampaign, nil
}

func (r *repository) FindAllActiveImageRepo() ([]Campaign, error) {
	var keyCampaign []Campaign
	err := r.db.Preload("CampaignImages", "is_primary = 0").Find(&keyCampaign).Error
	if err != nil {
		return keyCampaign, err
	}
	return keyCampaign, nil
}
