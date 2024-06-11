package campaign

type Formatter struct {
	ID               int                      `json:"id"`
	Name             string                   `json:"name"`
	ShortDescription string                   `json:"shortDescription"`
	Description      string                   `json:"description"`
	Perks            string                   `json:"perks"`
	BackerCount      int                      `json:"backerCount"`
	GoalAmount       int                      `json:"goalAmount"`
	Slug             string                   `json:"slug"`
	CampaignImages   []FormatterCampaignImage `json:"campaignImages"`
}

func FormatterCampaign(compaign Campaign, formatterCampaignImage []FormatterCampaignImage) Formatter {
	formatter := Formatter{
		ID:               compaign.ID,
		Name:             compaign.Name,
		ShortDescription: compaign.ShortDescription,
		Description:      compaign.Description,
		Perks:            compaign.Perks,
		BackerCount:      compaign.BackerCount,
		GoalAmount:       compaign.GoalAmount,
		Slug:             compaign.Slug,
		CampaignImages:   formatterCampaignImage,
	}
	return formatter
}

// type Formatterr struct {
// 	Name             string          `json:"name"`
// 	ShortDescription string          `json:"shortDescription"`
// 	Description      string          `json:"description"`
// 	Perks            string          `json:"perks"`
// 	BackerCount      int             `json:"backerCount"`
// 	GoalAmount       int             `json:"goalAmount"`
// 	Slug             string          `json:"slug"`
// 	CampaignImages   []CampaignImage `json:"campaignImages"`
// }

// func FormatterCampaignn(compaign Campaign) Formatter {
// 	formatter := Formatter{
// 		Name:             compaign.Name,
// 		ShortDescription: compaign.ShortDescription,
// 		Description:      compaign.Description,
// 		Perks:            compaign.Perks,
// 		BackerCount:      compaign.BackerCount,
// 		GoalAmount:       compaign.GoalAmount,
// 		Slug:             compaign.Slug,
// 		CampaignImages:   compaign.CampaignImages,
// 	}
// 	return formatter
// }

type FormatterCampaignImage struct {
	ID         int    `json:"id"`
	CampaignID int    `json:"campaignID"`
	FileNamw   string `json:"fileName"`
	IsPrimary  int    `json:"isPrimary"`
}

func CampaignImageFormatterCampaignImage(campaignImage CampaignImage) FormatterCampaignImage {
	FormatterCampaignImage := FormatterCampaignImage{
		ID:         campaignImage.ID,
		CampaignID: campaignImage.CampaignID,
		FileNamw:   campaignImage.FileNamw,
		IsPrimary:  campaignImage.IsPrimary,
	}
	return FormatterCampaignImage
}
