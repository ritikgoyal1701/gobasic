package experiment

import (
	"fmt"
	"gorm.io/gorm"
	"gorm/models"
	"sync"
)

type Repository struct {
	db *gorm.DB
}

var (
	repo     *Repository
	repoOnce sync.Once
)

func NewRepository(db *gorm.DB) *Repository {
	repoOnce.Do(func() {
		repo = &Repository{db: db}
	})
	return repo
}

func (rep *Repository) CreateUser(user *models.User) bool {
	result := rep.db.Create(user)
	if result.Error != nil {
		fmt.Println(result.Error)
		return false
	} else {
		fmt.Println(result.RowsAffected)
		return true
	}
}

func (rep *Repository) CreateUserAll(users []*models.User) bool {
	result := rep.db.Create(users)
	if result.Error != nil {
		fmt.Println(result.Error)
		return false
	} else {
		fmt.Println(result.RowsAffected)
		return true
	}
}

func applyScope(mp map[string]interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(mp)
	}
}

func (rep *Repository) GetUser(mp map[string]interface{}) []*models.User {
	var users []*models.User
	result := rep.db.Scopes(applyScope(mp)).Find(&users)
	if result.Error != nil {
		fmt.Println(result.Error)
	} else {
		fmt.Println(result.RowsAffected)
	}
	return users
}

func (rep *Repository) UpdateUser(org map[string]interface{}, update map[string]interface{}) bool {
	result := rep.db.Model(&models.User{}).Scopes(applyScope(org)).Updates(update)
	if result.Error != nil {
		fmt.Println(result.Error)
		return false
	} else {
		fmt.Println(result.RowsAffected)
		return true
	}
}

func (rep *Repository) DeleteUser(mp map[string]interface{}) bool {
	result := rep.db.Scopes(applyScope(mp)).Delete(&models.User{})
	if result.Error != nil {
		fmt.Println(result.Error)
		return false
	} else {
		fmt.Println(result.RowsAffected)
		return true
	}
}

func (rep *Repository) UpsertUser(user *models.User) bool {
	query := `INSERT INTO users (contact, name, dob)
              VALUES (?, ?, ?)
              ON CONFLICT (contact) DO UPDATE SET name = EXCLUDED.name, dob = EXCLUDED.dob`

	result := rep.db.Exec(query, user.Contact, user.Name, user.DOB)

	if result.Error != nil {
		fmt.Println(result.Error)
		return false
	} else {
		fmt.Println(result.RowsAffected)
		return true
	}
}
