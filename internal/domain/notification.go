package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NotificationType string

const (
	NOTIFICATION_NEW_PRODUCT_NOTIFICATION = NotificationType("NOTIFICATION_NEW_PRODUCT")
	NOTIFICATION_GENERAL_NOTIFICATION     = NotificationType("NOTIFICATION_GENERA")
)

type NotificationStatus string

const (
	NOTIFICATION_READY     = NotificationStatus("READY")
	NOTIFICATION_IN_QUEUE  = NotificationStatus("IN_QUEUE")
	NOTIFICATION_CANCELED  = NotificationStatus("CANCELED")
	NOTIFICATION_SUCCEEDED = NotificationStatus("SUCCEEDED")
	NOTIFICATION_FAILED    = NotificationStatus("FAILED")
)

type Notification struct {
	ID               primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Status           NotificationStatus
	NotificationType NotificationType
	Title            string
	Message          string
	DeviceIDs        []string
	NavigateTo       string
	ReferenceID      string
	NumUsersPushed   int
	NumUsersFailed   int
	CreatedAt        time.Time
	UpdatedAt        time.Time
	SendedAt         time.Time
}
