package models

import (
	"errors"
	orm "first_gin-gorm_proj/api/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       uint `gorm:"primarykey;AUTO_INCREAMENT"`
	Name     string
	Password string
}

var Users []User

// 添加
func (user User) Insert() (id uint, err error) {
	var existingUser User
	if err = orm.Db.Where("name = ?", user.Name).First(&existingUser).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}

	if existingUser.ID != 0 {
		// 如果找到了一个已存在的用户，返回一个错误
		err = errors.New("用户已存在")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	user.Password = string(hashedPassword)

	result := orm.Db.Create(&user)
	id = user.ID
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

// 登陆
func (user *User) Login() (loggedInUser User, err error) {
	// 先根据用户名查询用户
	if err = orm.Db.Where("name = ?", user.Name).First(&loggedInUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("用户名不存在")
		}
		return
	}

	// 比较哈希密码和明文密码
	err = bcrypt.CompareHashAndPassword([]byte(loggedInUser.Password), []byte(user.Password))
	if err != nil {
		// 如果密码不一致，返回错误
		err = errors.New("密码错误")
		return
	}

	return
}

// 修改
func (user *User) Update() (updateUser User, err error) {
	if err = orm.Db.Where("name = ?", user.Name).First(&updateUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("用户不存在")
		}
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	updateUser.Password = string(hashedPassword) // 更新密码

	if err = orm.Db.Model(&updateUser).Update("password", updateUser.Password).Error; err != nil { // 保存更新
		return
	}
	return
}

// 删除
func (user *User) Delete(name string) (Result User, err error) {
	var deleteUser User
	if err = orm.Db.Where("name = ?", name).First(&deleteUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("用户不存在")
		}
		return
	}

	Result = deleteUser // 保存待删除的用户信息以便返回

	if err = orm.Db.Delete(&deleteUser).Error; err != nil {
		return
	}

	return
}
