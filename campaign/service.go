package campaign

type Service interface {
	FindCampaigns(userID int) ([]Campaign, error)
	FindCampaign(input GetCampaignDetailByID) (Campaign, error)
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