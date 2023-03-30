package routes

import (
	controller "Final-Project-gin/controllers"
	"Final-Project-gin/middleware"

	"github.com/gin-gonic/gin"
)

func Book_keeperRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/get_book", controller.Get_books())
	incomingRoutes.POST("/post_book", controller.Post_book())
	incomingRoutes.PATCH("/update_book", controller.Update_book())
	incomingRoutes.DELETE("/delete_book", controller.Delete_book())
	incomingRoutes.GET("/search_book_id", controller.Search_book_id())
	incomingRoutes.GET("/search_by_name", controller.Search_book_name())
	incomingRoutes.GET("/search_book_publication", controller.Search_book_publication())
	incomingRoutes.GET("/publication_year", controller.Publication_Year())
	incomingRoutes.GET("/book_price", controller.Book_price())
	incomingRoutes.POST("/update_book_price", controller.Update_book_Price())
	incomingRoutes.POST("/book_awards", controller.Book_Awards())
	incomingRoutes.GET("/book_infor", controller.Book_Infor())
}
