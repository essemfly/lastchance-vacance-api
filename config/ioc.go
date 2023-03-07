package config

import (
	"github.com/1000king/handover/internal/repository"
	"github.com/1000king/handover/internal/storage"
	"github.com/spf13/viper"
)

type iocRepo struct {
	Products        repository.ProductsRepository
	CrawlThreads    repository.CrawlThreadsRepository
	CrawlKeywords   repository.CrawlKeywordsRepository
	CrawlProducts   repository.CrawlProductsRepository
	Users           repository.UsersRepository
	UserLikes       repository.UserLikesRepository
	Orders          repository.OrdersRepository
	Keywords        repository.KeywordsRepository
	Notifications   repository.NotificationsRepository
	KeywordProducts repository.KeywordProductsRepository
}

var Repo iocRepo

func InitRegisterRepo() {
	env := viper.GetString("ENVIRONMENT")

	switch env {
	case "prod":
		registerProdEnv()
	case "test":
		registerTestEnv()
	case "local":
		registerLocalEnv()
	default:
		registerLocalEnv()
	}
}

func registerProdEnv() {
	mongoconn := storage.NewMongoDB()
	registerRepos(mongoconn)
}

func registerTestEnv() {
	mongoconn := storage.NewMongoDB()
	registerRepos(mongoconn)
}

func registerLocalEnv() {
	mongoconn := storage.NewMongoDB()
	registerRepos(mongoconn)
}

func registerRepos(conn *storage.MongoDB) {
	Repo.Products = storage.MongoProductsRepo(conn)
	Repo.CrawlThreads = storage.MongoCrawlThreadsRepo(conn)
	Repo.CrawlKeywords = storage.MongoCrawlKeywordsRepo(conn)
	Repo.CrawlProducts = storage.MongoCrawlProductsRepo(conn)
	Repo.Users = storage.MongoUsersRepo(conn)
	Repo.UserLikes = storage.MongoUserLikesRepo(conn)
	Repo.Orders = storage.MongoOrdersRepo(conn)
	Repo.Keywords = storage.MongoKeywordRepo(conn)
	Repo.KeywordProducts = storage.MongoKeywordProductsRepo(conn)
	Repo.Notifications = storage.MongoNotificationsRepo(conn)
}
