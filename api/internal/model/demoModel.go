package model

import "github.com/jinzhu/gorm"

type defaultDemoUserModel struct {
	conn  *gorm.DB
	table string
	model DemoUser
}

type DemoUser struct {
	gorm.Model
	Name    string
	Address string
	Age     int
}

type DemoUserDo interface {
	Insert(user *DemoUser) (uint, error)
	DeleteById(id uint) error
	UpdateAddressById(addr string, id uint) error
}

func NewDemoUserDo(db *gorm.DB) DemoUserDo {
	return &defaultDemoUserModel{
		conn:  db,
		table: "demo_model",
		model: DemoUser{},
	}
}

func (d *defaultDemoUserModel) Insert(user *DemoUser) (uint, error) {
	if err := d.conn.Save(user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (d *defaultDemoUserModel) DeleteById(id uint) error {
	if err := d.conn.Delete(&d.model, id).Error; err != nil {
		return err
	}
	return nil
}

func (d *defaultDemoUserModel) UpdateAddressById(addr string, id uint) error {
	if err := d.conn.Model(&d.model).Where("id = ?", id).Update("address", addr).Error; err != nil {
		return err
	}
	return nil
}
