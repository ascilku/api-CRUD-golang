package campaign

type formatter struct {
	ID               int                      `json:"id"`
	Name             string                   `json:"name"`
	ShortDescription string                   `json:"shortDescription"`
	Description      string                   `json:"description"`
	Perks            string                   `json:"imgURl"`
	BackerCount      int                      `json:"backerCount"`
	GoalAmount       int                      `json:"goalAmount"`
	Slug             string                   `json:"slug"`
	CampaignImages   []formatterCampaignImage `json:"sampaignImages"`
}

func FormatterCampaign(compaign Campaign) formatter {
	formatter := formatter{
		ID:               compaign.ID,
		Name:             compaign.Name,
		ShortDescription: compaign.ShortDescription,
		Description:      compaign.Description,
		Perks:            compaign.Perks,
		BackerCount:      compaign.BackerCount,
		GoalAmount:       compaign.GoalAmount,
		Slug:             compaign.Slug,
	}
	if len(compaign.CampaignImages) == 0 {
		formatter.CampaignImages = []formatterCampaignImage{}
	}
	if len(compaign.CampaignImages) > 0 {
		formatter.Perks = compaign.CampaignImages[0].FileNamw
		var keyFormatterCampaignImage []formatterCampaignImage
		for _, keyCampaignImages := range compaign.CampaignImages {
			keyFormatterCampaignImage = append(keyFormatterCampaignImage, FormatterCampaignImage(keyCampaignImages))
		}
		formatter.CampaignImages = keyFormatterCampaignImage
	}

	return formatter
}

type formatterCampaignImage struct {
	ID         int    `json:"id"`
	CampaignID int    `json:"campaignID"`
	FileNamw   string `json:"fileName"`
	IsPrimary  int    `json:"isPrimary"`
}

func FormatterCampaignImage(campaignImage CampaignImage) formatterCampaignImage {
	FormatterCampaignImage := formatterCampaignImage{
		ID:         campaignImage.ID,
		CampaignID: campaignImage.CampaignID,
		FileNamw:   campaignImage.FileNamw,
		IsPrimary:  campaignImage.IsPrimary,
	}
	return FormatterCampaignImage
}

func FormatterCampaigns(campaign []Campaign) []formatter {
	// if len(campaign) == 0 {
	// 	return []formatter{}
	// }
	formatter := []formatter{}
	for _, keyCampaign := range campaign {
		formatter = append(formatter, FormatterCampaign(keyCampaign))
	}
	return formatter
}
