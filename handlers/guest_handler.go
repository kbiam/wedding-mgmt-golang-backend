package handlers

import (
	"fmt"
	"gorm/models/request/guest"
	"gorm/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)


func AddGuest(c *gin.Context){
	var req guest_request.AddGuestRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	guest, err := services.AddGuest(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Guest added successfully", "guest": guest})
}

func GetGuests(c *gin.Context){
	filters := map[string]string{
		"relation": c.Query("relation"),
		"side":     c.Query("side"),
		"phone":    c.Query("phone"),
	}

	guests, err := services.ListGuests(filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, guests)
}

func UpdateInvitationStatus(c *gin.Context){
	var req guest_request.UpdateInvitationStatusRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	id := c.Param("id")
	
	if err := services.UpdateInvitationStatus(id, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Invitation status updated successfully"})

}

func DeleteGuest(c *gin.Context) {
	id := c.Param("id")

	if err := services.DeleteGuest(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Guest deleted successfully"})
}

func AddGuestsFromExcel(c *gin.Context){
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

	xlFile, err := excelize.OpenReader(src)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read Excel file"})
		return
	}

	rows, err := xlFile.GetRows("Sheet1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get rows from Excel file"})
		return
	}

	
	addedGuests, err := services.AddGuestsFromExcel(rows)
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Failed to add guests from Excel: %s", err.Error())})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Excel file processed successfully", "guests": addedGuests})
}