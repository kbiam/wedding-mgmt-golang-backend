package guest

import (
	"fmt"
	"gorm/db"
	"gorm/models/entity/guest"
	"gorm/services"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

func AddGuest(c *gin.Context) {
	var newGuest Guest
	if err := c.BindJSON(&newGuest); err != nil{
		return
	}
	exist := db.DB.Where("phone = ?", newGuest.Phone).First(&Guest{}).Error
	if exist == nil {
		c.JSON(400, gin.H{"message": "Guest with this phone number already exists"})
		return
	}

	result := db.DB.Create(&newGuest)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return

	}
	c.JSON(201, gin.H{"message": "Guest added successfully", "guest": newGuest})
}

func GetGuests(c *gin.Context) {
	relation := c.Query("relation")
	side := c.Query("side")
	phone := c.Query("phone")

	var guests []Guest

	query := db.DB.Model(&Guest{})
	query = query.Order("created_at DESC")
	if relation != ""{
		query = query.Where("relation = ?", relation)
	}
	if side != "" {
		query = query.Where("side = ?", side)
	}
	if phone != "" {
		query = query.Where("phone = ?", phone)
	}

	if err := query.Find(&guests).Error; err != nil{
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(200, guests)

}

type UpdateInvitationStatusRequest struct {
	IsInvited bool `json:"is_invited" binding:"required"`
}

func UpdateInvitationStatus(c *gin.Context){
	id := c.Param("id")
	
	var data UpdateInvitationStatusRequest
	if err := c.BindJSON(&data); err != nil{
		c.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}

	query := db.DB.Model(&Guest{})
	query = query.Where("id = ?", id)
	if err := query.Update("is_invited", data.IsInvited).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Invitation status updated successfully"})

}

func DeleteGuest(c *gin.Context){
	id := c.Param("id")
	query := db.DB.Model(&Guest{})

	query = query.Where("id = ?", id)
	if err := query.Delete(&Guest{}).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Guest deleted successfully"})
}

func UploadExcel(c *gin.Context){
	file, err := c.FormFile("file")

	if err !=nil {
		c.JSON(400, gin.H{"error": "File is required"})
		return
	}

	src, err := file.Open()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to open file"})
		return
	}

	defer src.Close()

	xlFile, err := excelize.OpenReader(src)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to read Excel file"})
		return
	}

	rows, err := xlFile.GetRows("Sheet1")
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get rows from Excel file"})
		return
	}

	// var guests []guest.Guest

	for i, row := range rows {
		if i == 0 {
			continue
		}
		fmt.Println(row)
	}
	// if err := services.AddGuestsFromExcel(rows); err != nil {
	// 	c.JSON(500, gin.H{"error": fmt.Sprintf("Failed to add guests from Excel: %s", err.Error())})
	// 	return
	// }
	c.JSON(200, gin.H{"message": "Guests uploaded successfully"})

}