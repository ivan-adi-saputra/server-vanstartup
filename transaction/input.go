package transaction

import "server-vanstartup/user"

type GetCampaignTransactionsInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}