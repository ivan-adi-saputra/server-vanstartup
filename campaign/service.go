package campaign

type Service interface {
	FindCampaigns(userID int) ([]Campaign, error)
	FindCampaign(input GetCampaignDetailByID) (Campaign, error)
	SaveCampaign(input CreateCampaignInput) (Campaign, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) FindCampaigns(userID int) ([]Campaign, error) {
	if userID != 0 {
		campaigns, err := s.r.FindByUserID(userID)
		if err != nil {
			return campaigns, err
		}

		return campaigns, nil
	}

	campaigns, err := s.r.FindAll()
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (s *service) FindCampaign(input GetCampaignDetailByID) (Campaign, error) {
	campaign, err := s.r.FindByID(input.ID)
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (s *service) SaveCampaign(input CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{
		Name:             input.Name,
		ShortDescription: input.ShortDescription,
		Description:      input.Description,
		Perks:            input.Perks,
		GoalAmount:       input.GoalAmount,
		UserID:           input.User.ID,
	}

	newCampaign, err := s.r.Save(campaign)
	if err != nil {
		return Campaign{}, err
	}

	return newCampaign, nil
}