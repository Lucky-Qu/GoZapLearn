package dao

import "GoZapLearn/01_HelloZap/db/model"

func CreatePerson(person *model.Person) error {
	if err := db.Create(&person).Error; err != nil {
		return err
	}
	return nil
}
func UpdatePerson(id int, person *model.Person) error {
	if err := db.Model(&model.Person{}).Where("id = ?", id).Updates(person).Error; err != nil {
		return err
	}
	return nil
}
func DeletePerson(id int) error {
	if err := db.Unscoped().Delete(&model.Person{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
func GetPerson(id int) (*[]model.Person, error) {
	var persons []model.Person
	if err := db.Where("id = ?", id).Find(&persons).Error; err != nil {
		return nil, err
	}
	return &persons, nil
}
