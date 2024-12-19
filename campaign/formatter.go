package campaign

import "strings"

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

type CampaigDetailFormatter struct {
	ID               int      `json:"id"`
	Name             string   `json:"name"`
	ShortDescription string   `json:"short_description"`
	Description      string   `json:"description"`
	ImageURL         string   `json:"image_url"`
	GoalAmount       int      `json:"goal"`
	CurrentAmount    int      `json:"current_amount"`
	UserID           int      `json:"user_id"`
	Slug             string   `json:"slug"`
	Perks            []string `json:"perks"`
	User 			 CampaignUserFormatter `json:"user"`
	Images 			 []CampaignImageFormatter `json:"images"`
}

type CampaignUserFormatter struct {
	Name 	 string `json:"name"`
	ImageURL string `json:"image_url"`
}

type CampaignImageFormatter struct {
	ImageURL string `json:"image_url"`
	IsPrimary bool `json:"is_primary"`
}

func CampaignDetailFormatter(campaign Campaign) CampaigDetailFormatter {
	imageURL := ""
	if len(campaign.CampaignImages) > 0 {
		imageURL = campaign.CampaignImages[0].FileName
	}

	var perks []string
	for _, p := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(p))
	}

	var images []CampaignImageFormatter
	for _, i := range campaign.CampaignImages {
		isPrimary := false
		if i.IsPrimary == 1 {
			isPrimary = true
		}

		images = append(images, CampaignImageFormatter{
			ImageURL: i.FileName,
			IsPrimary: isPrimary,
		})
	}

	return CampaigDetailFormatter{
		ID:               campaign.ID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		Description:      campaign.Description,
		ImageURL:         imageURL,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		UserID: campaign.UserID,
		Slug:  campaign.Slug,
		Perks: perks,
		User: CampaignUserFormatter{
			Name: campaign.User.Name,
			ImageURL: campaign.User.AvatarFileName,
		},
		Images: images,
	}
}