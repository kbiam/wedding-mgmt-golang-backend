package guest

import "gorm/db"

func CreateGuest(guest *Guest) error {
	return db.DB.Create(guest).Error
}


func GuestExists(phone string)bool{
	var guest Guest

	err := db.DB.Where("phone = ?", phone).First(&guest).Error
	return err == nil
}

func GetGuests(filter map[string]interface{})([]Guest, error){
	var guests []Guest
	query := db.DB.Model(&Guest{}).Order("created_at DESC")

	for key, value := range filter {
		
		query = query.Where(key + " = ?", value)
	}
	err := query.Find(&guests).Error
	return guests, err
}

func UpdateInvitationStatus(id string, invited bool)error{
	return db.DB.Model(&Guest{}).Where("id = ?", id).Update("is_invited", invited).Error
}

func DeleteGuest(id string) error {
	return db.DB.Model(&Guest{}).Where("id = ?", id).Delete(&Guest{}).Error
}