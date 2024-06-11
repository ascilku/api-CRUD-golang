package campaign

import "gorm.io/gorm"

type Repository interface {
	FindAllRepo() ([]Campaign, error)
	FindAllActiveImageRepo() ([]Campaign, error)

	FindAllUserByID(userID int) ([]Campaign, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAllRepo() ([]Campaign, error) {
	var keyCampaign []Campaign
	err := r.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&keyCampaign).Error
	if err != nil {
		return keyCampaign, err
	}
	return keyCampaign, nil
}

func (r *repository) FindAllActiveImageRepo() ([]Campaign, error) {
	var keyCampaign []Campaign
	err := r.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&keyCampaign).Error
	if err != nil {
		return keyCampaign, err
	}
	return keyCampaign, nil
}

func (r *repository) FindAllUserByID(userID int) ([]Campaign, error) {
	var keyCampaign []Campaign
	// err := r.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&keyCampaign).Error
	err := r.db.Where("user_id = ?", userID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&keyCampaign).Error
	if err != nil {
		return keyCampaign, err
	}
	return keyCampaign, nil
}
