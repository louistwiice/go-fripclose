package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	logger "github.com/rs/zerolog/log"
	
	"github.com/louistwiice/go/fripclose/api/middlewares"
	"github.com/louistwiice/go/fripclose/api/handlers/user"
	"github.com/louistwiice/go/fripclose/api/handlers/category"
	"github.com/louistwiice/go/fripclose/configs"
	"github.com/louistwiice/go/fripclose/ent/migrate"
	"github.com/louistwiice/go/fripclose/entity"
)

// To load .env file
func init() {
	configs.Initialize()
}

func main() {
	logger.Info().Msg("Server starting ...")
	conf := configs.LoadConfigEnv()

	// Start by connecting to database
	db := configs.NewDBConnection()
	defer db.Close()

	rdb := configs.NewRedisConnection()

	// Run the automatic migration tool to create all schema resources.
	ctx := context.Background()
	err := db.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}


	app := gin.Default()
	app.MaxMultipartMemory = 2 << 20  // 2 MiB max for files to upload
	app.Static("media", "./media") // Path to get media files

	api_v1 := app.Group("api/v1")
	api_restricted := app.Group("api/v1/in")

	router_base := &entity.RouterBase{
		Database: db,
		Redis: rdb,
		OpenApp: api_v1,
	}
	router := &entity.Routers{
		RouterBase: *router_base,
		RestrictedApp: api_restricted,
	}

	middlewareController := middlewares.NewMiddlewareRouters(router)
	api_restricted.Use(middlewareController.JwAuthtMiddleware())

	handler_user.NewUserRouters(router)
	handler_category.NewCategoryRouters(router)

	logger.Info().Msg("Server ready to go ...")
	app.Run(conf.ServerPort)
}
