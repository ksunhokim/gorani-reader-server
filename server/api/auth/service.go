package auth

import (
	"fmt"

	selector "github.com/sunho/json-selector"
)

type User struct {
	Service  string
	Username string
	Avator   string
	Id       string
}

type Service struct {
	Name    string
	BaseUrl string

	UserEndPoint     string
	UsernameSelector string
	AvatorSelector   string
	IdSelector       string
}

func (s *Service) GetUserFromPayload(payload []byte) (User, error) {
	user := User{}
	avator, err := selector.Select(payload, s.AvatorSelector)
	if err != nil {
		return user, err
	}

	id, err := selector.Select(payload, s.IdSelector)
	if err != nil {
		return user, err
	}

	username, err := selector.Select(payload, s.UsernameSelector)
	if err != nil {
		return user, err
	}

	user.Avator = string(avator)
	user.Id = string(id)
	user.Username = string(username)
	user.Service = s.Name

	return user, nil
}

var services = []Service{}

func GetService(name string) (Service, error) {
	for _, service := range services {
		if service.Name == name {
			return service, nil
		}
	}
	return Service{}, fmt.Errorf("couldn't find such service")
}

func AddService(service Service) error {
	_, err := GetService(service.Name)
	if err == nil {
		return fmt.Errorf("already exist")
	}

	services = append(services, service)
	return nil
}
