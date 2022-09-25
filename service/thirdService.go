package service

import (
	"encoding/json"
	"github.com/SBEKzy/testTask/model"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (s *service) SendRequest(requestModel *model.Request) (*model.Response, error) {
	var responseModel model.Response
	client := &http.Client{Timeout: 90 * time.Second}
	request, err := http.NewRequest(
		requestModel.Method,
		requestModel.URL,
		nil,
	)

	if err != nil {
		return nil, err
	}

	for key, val := range requestModel.Headers {
		request.Header.Add(key, val)
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseModel.Status = resp.StatusCode
	responseModel.Length = resp.ContentLength
	responseModel.Headers = make(map[string][]string)
	responseModel.Headers = resp.Header
	responseModel.ID = uuid.New()

	err = s.saveRequest(requestModel, &responseModel)
	if err != nil {
		return nil, err
	}
	return &responseModel, nil
}

func (s *service) saveRequest(request *model.Request, response *model.Response) error {

	reqJSONByte, err := json.Marshal(request)
	if err != nil {
		return err
	}
	resJSONByte, err := json.Marshal(response)
	if err != nil {
		return err
	}

	err = s.Repo.SaveRequest(string(reqJSONByte), string(resJSONByte))

	if err != nil {
		return err

	}

	return nil
}
