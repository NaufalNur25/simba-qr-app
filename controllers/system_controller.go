package controllers

import (
	"image"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/liyue201/goqr"
	"github.com/naufal/simba-qr-app/services"
	"github.com/naufal/simba-qr-app/services/requests"
	"github.com/naufal/simba-qr-app/utils"
	"github.com/skip2/go-qrcode"
)

func CreateSystem(c *gin.Context) {
	var input requests.PostSystemRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.CreateSystem(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create system"})
		return
	}

	png, err := qrcode.Encode(result.Key, qrcode.Medium, 256)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate QR code"})
		return
	}

	c.Data(http.StatusOK, "image/png", png)
}

func UploadQRCode(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}
	defer src.Close()

	img, _, err := image.Decode(src)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image format"})
		return
	}

	qrCodes, err := goqr.Recognize(img)
	if err != nil || len(qrCodes) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No QR code detected"})
		return
	}

	encrypted := string(qrCodes[0].Payload)

	decrypted, err := utils.DecryptAES(encrypted)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decrypt QR code"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"decrypted_data": decrypted,
	})
}

func DeleteSystem(c *gin.Context) {
	id := c.Param("id")

	err := services.DeleteSystem(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete system"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "System deleted"})
}
