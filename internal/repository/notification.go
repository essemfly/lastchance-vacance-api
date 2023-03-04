package repository

import "github.com/1000king/handover/internal/domain"

type NotificationsRepository interface {
	Insert(*domain.Notification) (*domain.Notification, error)
	Get(ID string) (*domain.Notification, error)
	List(offset, limit int, notiTypes []domain.NotificationType, onlyReady bool) ([]*domain.Notification, error)
	Update(*domain.Notification) (*domain.Notification, error)
}
