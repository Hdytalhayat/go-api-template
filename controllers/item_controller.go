// File: controllers/item_controller.go
// Deskripsi: Berisi logika bisnis (handler) untuk entitas 'Item'.

package controllers

import (
	"go-api-template/config"
	"go-api-template/models"
	"go-api-template/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetItems adalah handler untuk endpoint GET /items.
func GetItems(c *gin.Context) {
	var items []models.Item

	rows, err := config.DB.Query("SELECT id, title, author FROM items")
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Gagal mengambil data item")
		return
	}
	defer rows.Close()

	for rows.Next() {
		var b models.Item
		if err := rows.Scan(&b.ID, &b.Title, &b.Author); err != nil {
			utils.RespondError(c, http.StatusInternalServerError, "Gagal memindai data item")
			return
		}
		items = append(items, b)
	}

	// Gunakan helper untuk mengirim respons sukses
	utils.RespondJSON(c, http.StatusOK, items)
}

// AddItem adalah handler untuk endpoint POST /items.
func AddItem(c *gin.Context) {
	var newItem models.Item

	if err := c.ShouldBindJSON(&newItem); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Input tidak valid: "+err.Error())
		return
	}

	result, err := config.DB.Exec("INSERT INTO items (title, author) VALUES (?, ?)", newItem.Title, newItem.Author)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Gagal menyimpan item")
		return
	}

	id, _ := result.LastInsertId()
	newItem.ID = int(id)

	// Gunakan helper untuk mengirim respons 'created'
	utils.RespondJSON(c, http.StatusCreated, newItem)
}
