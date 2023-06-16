package experiment

import (
	"fmt"
	"gorm/internal"
	"gorm/models"
	"sync"
)

type Service struct {
	repository internal.UserRepository
}

var (
	svc     *Service
	svcOnce sync.Once
)

func NewService(repository internal.UserRepository) *Service {
	svcOnce.Do(func() {
		svc = &Service{repository: repository}
	})
	return svc
}

func (s *Service) CreateUser(name, age, phone string) string {
	user := models.User{
		Name:    name,
		DOB:     age,
		Contact: phone,
	}

	response := s.repository.CreateUser(&user)
	if response {
		return "User Created"
	} else {
		return "Error Occured"
	}
}

func (s *Service) CreateUserAll(name, age, phone []string) string {
	var users []*models.User
	n := len(name)
	for i := 0; i < n; i++ {
		users = append(users, &models.User{
			Name:    name[i],
			DOB:     age[i],
			Contact: phone[i],
		})
	}
	response := s.repository.CreateUserAll(users)
	if response {
		fmt.Println("Users created Successfully.")
		return "Users created Successfully."
	} else {
		fmt.Println("Error Occurred")
		return "Error Occurred"
	}
}

func prepareCondition(name, age, contact string) map[string]interface{} {
	mp := make(map[string]interface{})
	if len(name) > 0 {
		mp["name"] = name
	}
	if len(age) > 0 {
		mp["dob"] = age
	}
	if len(contact) > 0 {
		mp["contact"] = contact
	}
	return mp
}

func (s *Service) GetUser(name, age, contact string) []*models.User {
	mp := prepareCondition(name, age, contact)
	response := s.repository.GetUser(mp)
	for _, user := range response {
		fmt.Println(user)
	}
	return response
}

func (s *Service) UpdateUser(nameOrg, ageOrg, phoneOrg, name, age, phone string) string {
	org := prepareCondition(nameOrg, ageOrg, phoneOrg)
	upd := prepareCondition(name, age, phone)
	response := s.repository.UpdateUser(org, upd)
	if response {
		fmt.Println("Database Updated")
		return "Database Updated"
	} else {
		fmt.Println("Some problem Occurred")
		return "Some problem Occurred"
	}
}

func (s *Service) DeleteUser(name, age, contact string) string {
	mp := prepareCondition(name, age, contact)
	response := s.repository.DeleteUser(mp)
	if response {
		fmt.Println("Deletion successful")
		return "Deletion successful"
	} else {
		fmt.Println("Some error occurred")
		return "Some error occurred"
	}
}

func (s *Service) UpsertUser(name, age, contact string) string {
	response := s.repository.UpsertUser(&models.User{
		Name:    name,
		DOB:     age,
		Contact: contact,
	})
	if response {
		fmt.Println("User Upserted")
		return "User Upserted"
	} else {
		fmt.Println("Error Occurred")
		return "Error Occurred"
	}
}
