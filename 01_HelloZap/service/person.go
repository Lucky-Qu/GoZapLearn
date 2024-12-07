package service

import (
	"GoZapLearn/01_HelloZap/db/dao"
	"GoZapLearn/01_HelloZap/db/model"
)

func GetPerson(id int) (*[]model.Person, error) {
	if person, err := dao.GetPerson(id); err != nil {
		return nil, err
	} else {
		return person, nil
	}
}
func CreatePerson(person *model.Person) error {
	if err := dao.CreatePerson(person); err != nil {
		return err
	}
	return nil
}
func DeletePerson(id int) error {
	if err := dao.DeletePerson(id); err != nil {
		return err
	}
	return nil
}
func UpdatePerson(id int, person *model.Person) error {
	if err := dao.UpdatePerson(id, person); err != nil {
		return err
	}
	return nil
}
