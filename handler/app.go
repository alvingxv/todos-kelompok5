package handler

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

func StartApp() {

	database.HandleDatabaseConnection()
	db := database.GetDatabaseInstance()

	// Category Injection
	categoryRepo := category_pg.NewCategoryPG(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := NewCategoryHandler(categoryService)

	// port := os.Getenv("PORT")
	port := "4000"
	// port := helpers.GoDotEnvVariable("PORT")
	r := gin.Default()

	docs.SwaggerInfo.Title = "Kanban Board Kelompok 5"
	docs.SwaggerInfo.Description = "Final Project 3 Hactiv8 by Kelompok 5"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost"
	docs.SwaggerInfo.Schemes = []string{"http"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	categoryRoute := r.Group("/categories")
	{

		categoryRoute.GET("", categoryHandler.GetCategory)
		categoryRoute.POST("", categoryHandler.CreateCategory)
		categoryRoute.PATCH("/:id", categoryHandler.UpdateCategory)
		categoryRoute.DELETE("/:id", categoryHandler.DeleteCategory)
	}
	r.Run(":" + port)
}
