package handler

import (
	"api-satu/campaign"
	"api-satu/respons"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handlerCampign struct {
	service campaign.Service
}

func NewHandlerCampign(service campaign.Service) *handlerCampign {
	return &handlerCampign{service}
}

func (h *handlerCampign) FindAllHand(g *gin.Context) {
	findAllSer, err := h.service.FindAllSer()
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		errorResponsValue := respons.ResponsValue("Error get all data campaign", http.StatusUnprocessableEntity, "Error", errorMessage)
		g.JSON(http.StatusUnprocessableEntity, errorResponsValue)
		return
	}

	var campaignFormatter []campaign.Formatter
	for _, keyFindAllSer := range findAllSer {
		var keyCampaignImage []campaign.FormatterCampaignImage
		for _, keyFindAllCampaignImages := range keyFindAllSer.CampaignImages {
			keyCampaignImage = append(keyCampaignImage, campaign.CampaignImageFormatterCampaignImage(keyFindAllCampaignImages))
		}
		formatter := campaign.Campaign{
			Name:             keyFindAllSer.Name,
			ShortDescription: keyFindAllSer.ShortDescription,
			Description:      keyFindAllSer.Description,
			Perks:            keyFindAllSer.Perks,
			BackerCount:      keyFindAllSer.BackerCount,
			GoalAmount:       keyFindAllSer.GoalAmount,
			Slug:             keyFindAllSer.Slug,
		}
		campaignFormatter = append(campaignFormatter, campaign.FormatterCampaign(formatter, keyCampaignImage))
	}
	responsValue := respons.ResponsValue("Succes get all data campaign", http.StatusOK, "Succes", campaignFormatter)
	g.JSON(http.StatusOK, responsValue)

}
