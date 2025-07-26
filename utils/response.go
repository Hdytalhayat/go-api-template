// File: utils/response.go
// Deskripsi: Menyediakan fungsi helper untuk standarisasi respons API.

package utils

import "github.com/gin-gonic/gin"

// RespondJSON adalah helper untuk mengirim respons JSON yang sukses.
// Ini memastikan format respons selalu konsisten.
func RespondJSON(c *gin.Context, status int, payload interface{}) {
	c.JSON(status, payload)
}

// RespondError adalah helper untuk mengirim respons error dalam format JSON yang standar.
// Ini memastikan semua pesan error memiliki struktur yang sama.
func RespondError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"error": message,
	})
}
