package handler

import (
	"net/http"
	"server-vanstartup/campaign"
	"server-vanstartup/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	s campaign.Service
}

func NewCampaignHandler(s campaign.Service) *campaignHandler {
    return &campaignHandler{s}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.s.FindCampaigns(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ApiResponse("Get campaign failed", http.StatusInternalServerError, "FAILED", err.Error()))
		return
	}

	c.JSON(http.StatusOK, helper.ApiResponse("Get campaign successfully", http.StatusOK, "SUCCESS", campaign.CampaignsFormatter(campaigns)))
}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	var input campaign.GetCampaignDetailByID
	err := c.ShouldBindUri(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiResponse("Invalid request", http.StatusBadRequest, "FAILED", nil))
        return
	}

	data, err := h.s.FindCampaign(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiResponse("Get detail campaign failed", http.StatusBadRequest, "FAILED", nil))
        return
	}

	c.JSON(http.StatusOK, helper.ApiResponse("Get detail campaign", http.StatusOK, "SUCCESS", data))
}