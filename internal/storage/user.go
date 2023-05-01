package storage

import (
	"context"
	"time"

	"github.com/1000king/handover/internal/domain"
	"github.com/1000king/handover/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepo struct {
	col *mongo.Collection
}

func (repo *userRepo) Get(ID primitive.ObjectID) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	mongoFilter := bson.M{
		"_id": ID,
	}

	var user *domain.User
	if err := repo.col.FindOne(ctx, mongoFilter).Decode(&user); err != nil {
		return nil, err
	}
	return user, nil

}

func (repo *userRepo) Upsert(user *domain.User) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var prevUser *domain.User
	err := repo.col.FindOne(ctx, bson.M{"deviceuuid": user.DeviceUUID}).Decode(&prevUser)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, err
		}
		user.ID = primitive.NewObjectID()
		user.CreatedAt = time.Now()
	} else {
		user.Mobile = prevUser.Mobile
		user.Address = prevUser.Address
	}

	user.UpdatedAt = time.Now()

	opts := options.Update().SetUpsert(true)
	filter := bson.M{"deviceuuid": user.DeviceUUID}
	if _, err := repo.col.UpdateOne(ctx, filter, bson.M{"$set": user}, opts); err != nil {
		return nil, err
	}

	return user, nil
}

func MongoUsersRepo(conn *MongoDB) repository.UsersRepository {
	return &userRepo{
		col: conn.userCol,
	}
}

type userLikeRepo struct {
	col *mongo.Collection
}

func (repo *userLikeRepo) Get(userID primitive.ObjectID, productID primitive.ObjectID) (*domain.UserLike, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var userLike domain.UserLike
	err := repo.col.FindOne(ctx, bson.M{"userid": userID, "productid": productID}).Decode(&userLike)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, err
		}
		return nil, nil
	}

	return &userLike, nil
}

func (repo *userLikeRepo) Upsert(userId primitive.ObjectID, pd *domain.Product) (*domain.UserLike, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var userLike domain.UserLike
	err := repo.col.FindOne(ctx, bson.M{"userid": userId, "productid": pd.ID}).Decode(&userLike)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, err
		}
		userLike.ID = primitive.NewObjectIDFromTimestamp(time.Now())
		userLike.UserId = userId
		userLike.ProductId = pd.ID
		userLike.IsLiked = true
		userLike.CreatedAt = time.Now()
	} else {
		userLike.IsLiked = !userLike.IsLiked
	}

	userLike.UpdatedAt = time.Now()

	opts := options.Update().SetUpsert(true)
	filter := bson.M{"_id": userLike.ID}
	if _, err := repo.col.UpdateOne(ctx, filter, bson.M{"$set": &userLike}, opts); err != nil {
		return nil, err
	}

	return &userLike, nil
}

func (repo *userLikeRepo) List(filter *domain.UserLikeFilter) ([]*domain.UserLike, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	mongoFilter := bson.M{
		"isliked": true,
		"userid":  filter.UserId,
	}

	cursor, err := repo.col.Find(ctx, mongoFilter)
	if err != nil {
		return nil, err
	}

	var userLike []*domain.UserLike
	err = cursor.All(ctx, &userLike)
	if err != nil {
		return nil, err
	}
	return userLike, nil
}

func MongoUserLikesRepo(conn *MongoDB) repository.UserLikesRepository {
	return &userLikeRepo{
		col: conn.userLikeCol,
	}
}
