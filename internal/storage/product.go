package storage

import (
	"github.com/1000king/handover/internal/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type productRepo struct {
	col *mongo.Collection
}

func (repo *productRepo) Create() {
	// ...
}

func (repo *productRepo) List() {
	// ...
}

func (repo *productRepo) Update() {
	// ...
}

func MongoProductsRepo(conn *MongoDB) repository.ProductsRepository {
	return &productRepo{
		col: conn.productCol,
	}
}
