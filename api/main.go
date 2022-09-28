package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	logger "github.com/rs/zerolog/log"

	"github.com/louistwiice/go/fripclose/api/middlewares"
	"github.com/louistwiice/go/fripclose/api/users"
	category_controller "github.com/louistwiice/go/fripclose/api/category"
	"github.com/louistwiice/go/fripclose/configs"
	"github.com/louistwiice/go/fripclose/ent/migrate"
	"github.com/louistwiice/go/fripclose/repository"
	"github.com/louistwiice/go/fripclose/usecase/authentication"
	"github.com/louistwiice/go/fripclose/usecase/category"
	"github.com/louistwiice/go/fripclose/usecase/user"
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

	userRepo := repository.NewUserClient(db, rdb)
	categeryRepo := repository.NewCategoryClient(db)

	userService := user.NewUserService(userRepo)
	authService := authentication.NewAuthService(userRepo)
	categoryService := category.NewCategoryService(categeryRepo)

	userController := users.NewUserController(userService)
	authController := users.NewAuthController(authService)
	categoryController := category_controller.NewCategoryController(categoryService, userService)
	middlwareController := middlewares.NewMiddlewareControllers(authService)

	app := gin.Default()
	app.MaxMultipartMemory = 2 << 20  // 2 MiB max for files to upload
	app.Static("media", "./media") // Path to get media files

	api_v1 := app.Group("api/v1")
	authController.MakeAuthHandlers(api_v1.Group("auth/"))
	userController.MakeUserHandlersWithoutAuth(api_v1.Group("user/"))
	categoryController.MakeCategoryHandlersWithoutAuth(api_v1.Group("category/"))

	api_auth := app.Group("api/v1/in")
	api_auth.Use(middlwareController.JwAuthtMiddleware())
	userController.MakeUserHandlersWithAuth(api_auth.Group("user/"))
	categoryController.MakeCategoryHandlersWithAuth(api_auth.Group("category/"))

	logger.Info().Msg("Server ready to go ...")
	app.Run(conf.ServerPort)
}
