package storage

import (
	"context"
	"log"
	"time"

	"github.com/1000king/handover/internal/domain"
	"github.com/1000king/handover/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type notificationRepo struct {
	col *mongo.Collection
}

// Get implements repository.NotificationsRepository
func (repo *notificationRepo) Get(ID string) (*domain.Notification, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	notiObjID, _ := primitive.ObjectIDFromHex(ID)

	noti := &domain.Notification{}
	if err := repo.col.FindOne(ctx, bson.M{"_id": notiObjID}).Decode(noti); err != nil {
		return nil, err
	}

	return noti, nil
}

// List implements repository.NotificationsRepository
func (repo *notificationRepo) List(offset int, limit int, notiTypes []domain.NotificationType, onlyReady bool) ([]*domain.Notification, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	notiFilters := []bson.M{}
	for _, notiType := range notiTypes {
		notiFilters = append(notiFilters, bson.M{
			"notificationtype": notiType,
		})
	}
	filter := bson.M{
		"$or": notiFilters,
	}
	if onlyReady {
		filter["status"] = domain.NOTIFICATION_READY
	}

	cursor, err := repo.col.Find(ctx, filter)
	if err != nil {
		log.Println("err on notification ", err)
		return nil, err
	}

	var notis []*domain.Notification
	err = cursor.All(ctx, &notis)
	if err != nil {
		return nil, err
	}

	return notis, nil
}

// Update implements repository.NotificationsRepository
func (repo *notificationRepo) Update(noti *domain.Notification) (*domain.Notification, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	_, err := repo.col.UpdateOne(ctx, bson.M{"_id": noti.ID}, bson.M{"$set": &noti})
	if err != nil {
		return nil, err
	}

	return noti, nil
}

func (repo *notificationRepo) Insert(notification *domain.Notification) (*domain.Notification, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	notification.ID = primitive.NewObjectID()
	notification.CreatedAt = time.Now()
	notification.UpdatedAt = time.Now()

	if _, err := repo.col.InsertOne(ctx, notification); err != nil {
		return nil, err
	}

	return notification, nil
}

func MongoNotificationsRepo(conn *MongoDB) repository.NotificationsRepository {
	return &notificationRepo{
		col: conn.notificationCol,
	}
}
