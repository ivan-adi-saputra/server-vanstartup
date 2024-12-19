package campaign

type Formatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

func CampaignFormatter(campaign Campaign) Formatter {
	imageURL := ""
	if len(campaign.CampaignImages) > 0 {
		imageURL = campaign.CampaignImages[0].FileName
	}

	return Formatter{
		ID:               campaign.ID,
		UserID:           campaign.UserID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		ImageURL:         imageURL,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		Slug:             campaign.Slug,
	}
}

func CampaignsFormatter(campaigns []Campaign) []Formatter {
	if len(campaigns) == 0 {
		return []Formatter{}
	}

	formattedCampaigns := make([]Formatter, len(campaigns))

	for i, campaign := range campaigns {
		formattedCampaigns[i] = CampaignFormatter(campaign)
	}

	return formattedCampaigns
}