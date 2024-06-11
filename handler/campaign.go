package handler

import (
	"api-satu/campaign"
	"api-satu/respons"
	"net/http"
	"strconv"

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

func (h *handlerCampign) FindActiveImageAllHand(g *gin.Context) {
	findAllSer, err := h.service.FindAllActiveImageAllSer()
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		errorResponsValue := respons.ResponsValue("Error Not Data campaign all active", http.StatusUnprocessableEntity, "Error", errorMessage)
		g.JSON(http.StatusUnprocessableEntity, errorResponsValue)
		return
	}
	var formatterCampaign []campaign.Formatter
	var formatterCampaignImage []campaign.FormatterCampaignImage
	for _, keyfindAllSer := range findAllSer {
		for _, keyCampaignImages := range keyfindAllSer.CampaignImages {
			formatterCampaignImage = append(formatterCampaignImage, campaign.CampaignImageFormatterCampaignImage(keyCampaignImages))
		}
		formatterCampaign = append(formatterCampaign, campaign.FormatterCampaign(keyfindAllSer, formatterCampaignImage))
	}
	responsValue := respons.ResponsValue("Success Show Data campaign all active", http.StatusOK, "Success", formatterCampaign)
	g.JSON(http.StatusOK, responsValue)
}

func (h *handlerCampign) FindCampaignUserHand(g *gin.Context) {
	userID, _ := strconv.Atoi(g.Query("user_id"))
	FindAllUserByIDSer, err := h.service.FindAllUserByIDSer(userID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		respons := respons.ResponsValue("Error get campaign", http.StatusUnprocessableEntity, "Error", errorMessage)
		g.JSON(http.StatusUnprocessableEntity, respons)
	} else {
		respons := respons.ResponsValue("Success get campaign", http.StatusOK, "Success", FindAllUserByIDSer)
		g.JSON(http.StatusOK, respons)
	}
}
