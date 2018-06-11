package auth

import (
	"io/ioutil"
	"net/http"

	"github.com/sunho/gorani-reader-server/pkg/util"
)

func (s *Services) FetchUser(serviceName string, token string) (User, error) {
	service, err := s.GetService(serviceName)
	if err != nil {
		return User{}, err
	}

	client := util.CreateClient()

	req, err := http.NewRequest("GET", service.BaseUrl+service.UserEndPoint, nil)
	if err != nil {
		return User{}, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		return User{}, err
	}
	defer resp.Body.Close()

	payload, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return User{}, err
	}

	user, err := service.GetUserFromPayload(payload)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
