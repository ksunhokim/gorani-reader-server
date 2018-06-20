package auth

import (
	"errors"

	selector "github.com/sunho/json-selector"
	yaml "gopkg.in/yaml.v2"
)

var (
	ErrAlreadyExist = errors.New("auth: service already exists")
	ErrNotFound     = errors.New("auth: service not found")
)

type User struct {
	Service  string
	Username string
	Avator   string
	Id       string
}

type Service struct {
	Name    string `yaml:"name"`
	BaseUrl string `yaml:"base_url"`

	UserEndPoint     string `yaml:"user_end_point"`
	UsernameSelector string `yaml:"username_selector"`
	AvatorSelector   string `yaml:"avator_selector"`
	IdSelector       string `yaml:"id_selector"`
}

// is not thread-safe; shoul be initialized only once
type Services []Service

func NewServices(yamlBytes []byte) (Services, error) {
	conf := Services{}

	err := yaml.Unmarshal(yamlBytes, &conf)
	if err != nil {
		return Services{}, err
	}

	return conf, nil
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

func (s Services) GetService(name string) (Service, error) {
	for _, service := range s {
		if service.Name == name {
			return service, nil
		}
	}
	return Service{}, ErrNotFound
}

func (s *Services) AddService(service Service) error {
	_, err := s.GetService(service.Name)
	if err == nil {
		return ErrAlreadyExist
	}

	*s = append(*s, service)
	return nil
}
