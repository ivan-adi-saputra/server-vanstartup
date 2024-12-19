package campaign

type GetCampaignDetailByID struct {
	ID int `uri:"id" binding:"required"`
}