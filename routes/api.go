// File: routes/api.go
// Deskripsi: Mendaftarkan semua rute (endpoints) untuk API.

package routes

import (
	"go-api-template/controllers"
	"go-api-template/utils"

	"github.com/gin-gonic/gin"
)

// SetupRouter mengkonfigurasi dan mengembalikan instance dari Gin router.
func SetupRouter() *gin.Engine {
	// Membuat router Gin dengan middleware default (logger dan recovery)
	router := gin.Default()

	// Menambahkan rute untuk health check. Ini adalah praktik yang baik.
	router.GET("/health", func(c *gin.Context) {
		utils.RespondJSON(c, 200, gin.H{"status": "UP"})
	})

	// Mengelompokkan rute API di bawah prefiks /api/v1
	apiV1 := router.Group("/api/v1")
	{
		// Rute yang berhubungan dengan 'items'
		items := apiV1.Group("/items")
		{
			// GET /api/v1/items -> Memanggil fungsi GetItems dari item_controller
			items.GET("", controllers.GetItems)
			// POST /api/v1/items -> Memanggil fungsi AddItem dari item_controller
			items.POST("", controllers.AddItem)

			// Anda bisa menambahkan rute lain di sini, contoh:
			// items.GET("/:id", controllers.GetItemByID)
			// items.PUT("/:id", controllers.UpdateItem)
			// items.DELETE("/:id", controllers.DeleteItem)
		}
	}

	return router
}
