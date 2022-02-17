package NoCap

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	HCaptchaRegular = 1
	HCaptchaBackup = 2
	HCaptchaEnterprise = 3
)

type HCaptcha struct {
	ApiKey string
	Mode int
	SiteKey string
	Proxy string
}

type createTaskResp struct {
	Id   string `json:"task"`
}

type statusResp struct {
	Status   string `json:"status"`
	Answer   string `json:"answer"`
}

func (hCaptcha *HCaptcha) SolveHCaptcha() (string, error) {
	task, err := hCaptcha.createTask()
	if err != nil {
		return "", err
	}
	solution, err := hCaptcha.getSolution(task)
	if err != nil {
		return "", err
	}
	return solution, nil
}

// createTask creates a new HCaptcha task and
// returns createTaskResp.Id if the request was
// successful. An error is returned if the request
// was unable to be fulfilled.
func (hCaptcha *HCaptcha) createTask() (string, error) {
	req, _ := http.NewRequest("POST", fmt.Sprintf(CapEndpoint+ "/create?api_key=%s", hCaptcha.ApiKey), nil)
	resp, err := request(req)
	if err != nil {
		return "", err
	}
	var task createTaskResp
	err = json.Unmarshal([]byte(resp), &task)
	if err != nil {
		return "", err
	}
	return task.Id, nil
}

// getSolution attempts to retrieve a token given a task.
// An error is returned if the request was unable to be fulfilled,
// or if a captcha was not solved.
func (hCaptcha *HCaptcha) getSolution(task string) (string, error) {
	var status statusResp
	req, _ := http.NewRequest("GET", fmt.Sprintf(CapEndpoint+ "/status?api_key=%s&task_id=%s", hCaptcha.ApiKey, task), nil)
	// wait for server to respond with a token or error
	for status.Status == "solving" || status.Status == "" {
		resp, err := request(req)
		if err != nil {
			return "", err
		}
		_ = json.Unmarshal([]byte(resp), &status)
	}
	// ensure status is successful
	if status.Status != "success" {
		return "", errors.New("error retrieving hCaptcha token")
	}

	return status.Answer, nil
}