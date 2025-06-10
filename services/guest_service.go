package services

import (
	"errors"
	"fmt"
	"gorm/models/entity/guest"
	guest_request "gorm/models/request/guest"
	"strconv"
	"strings"
)

func AddGuest(req guest_request.AddGuestRequest)(*guest.Guest, error){
	if guest.GuestExists(req.Phone){
		return nil, errors.New("guest already exists")
	}

	newGuest := &guest.Guest{
		Name:       req.Name,
		Phone:      req.Phone,
		Relation:   req.Relation,
		Side:       req.Side,
		GuestCount: req.GuestCount,

	}

	err := guest.CreateGuest(newGuest)
	return newGuest, err
}

func ListGuests(filters map[string]string)([]guest.Guest, error) {

	query := map[string]interface{}{}

	for key, value := range filters {
		if value != ""{
			query[key] = value
		}
	}

	return guest.GetGuests(query)
}

func UpdateInvitationStatus(id string, req guest_request.UpdateInvitationStatusRequest) error {
	return guest.UpdateInvitationStatus(id, req.IsInvited)
}

func DeleteGuest(id string) error {
	return guest.DeleteGuest(id)
}

func AddGuestsFromExcel(rows [][]string)([]*guest.Guest,error){

	var addedGuests []*guest.Guest

	for i, row := range rows {
		if i == 0 {
			continue
		}
		if len(row) < 5 {
			return nil,errors.New("invalid row format at row " + string(i+1))
		}
		count, err := strconv.Atoi(strings.TrimSpace(row[4]))
		if err != nil {
			return nil,fmt.Errorf("invalid GuestCount at row %d: %v", i+1, err)
		}

		newGuest := &guest.Guest{
			Name: 	 strings.TrimSpace(row[0]),
			Phone:    strings.TrimSpace(row[1]),
			Relation: strings.TrimSpace(row[2]),
			Side:     strings.TrimSpace(row[3]),
			GuestCount: count,
		}

		if guest.GuestExists(newGuest.Phone) {
			return nil,fmt.Errorf("guest with phone %s already exists at row %d", newGuest.Phone, i+1)
		}
		if err := guest.CreateGuest(newGuest); err != nil {
			return nil,fmt.Errorf("failed to add guest at row %d: %v", i+1, err)
		}
		addedGuests = append(addedGuests, newGuest)

	}
	return addedGuests,nil
}