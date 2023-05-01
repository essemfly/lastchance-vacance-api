package product

import (
	"log"
	"strings"

	"github.com/1000king/handover/config"
	"github.com/1000king/handover/internal/domain"
	"github.com/1000king/handover/pkg/push"
)

func AddKeywordProduct(pd *domain.Product) {
	liveKeywords, err := config.Repo.Keywords.ListAll()
	if err != nil {
		log.Println("failed to list keywords", err)
		return
	}

	for _, keyword := range liveKeywords {
		if strings.Contains(pd.Name, keyword.Keyword) {
			config.Repo.KeywordProducts.Insert(pd, keyword.UserID.Hex(), keyword.Keyword)
			user, _ := config.Repo.Users.Get(keyword.UserID)
			newNoti, err := AddNewProductNotification(pd, keyword.Keyword, user)
			if err != nil {
				push.SendNotification(newNoti)
			}
		}
	}
}

func AddNewProductNotification(pd *domain.Product, keyword string, user *domain.User) (*domain.Notification, error) {
	newNotification := &domain.Notification{
		Status:           domain.NOTIFICATION_READY,
		NotificationType: domain.NOTIFICATION_NEW_PRODUCT_NOTIFICATION,
		Title:            keyword + " 키워드 알림",
		Message:          pd.Name,
		DeviceIDs: []string{
			user.DeviceUUID,
		},
		NavigateTo:     "",
		ReferenceID:    "",
		NumUsersPushed: 1,
		NumUsersFailed: 0,
	}
	return config.Repo.Notifications.Insert(newNotification)
}
