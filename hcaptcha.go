package NoCap

import (
	"bytes"
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
	ApiKey  string
	Mode    int    `json:"mode,omitempty"`
	SiteKey string `json:"sitekey,omitempty"`
	Proxy   string `json:"proxy,omitempty"`
	Host    string `json:"host,omitempty"`
	Version string `json:"version,omitempty"`
	Href    string `json:"href,omitempty"`
	RqData  string `json:"rqdata,omitempty"`
}

type createTaskResp struct {
	Id   string `json:"task"`
	Error string `json:"error"`
}

type statusResp struct {
	Status   string `json:"status"`
	Answer   string `json:"answer"`
	Reason   string `json:"reason"`
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
// returns createTaskResp.Id. An error is
// returned if the request wasn't successful.
func (hCaptcha *HCaptcha) createTask() (string, error) {
	data, err := json.Marshal(hCaptcha)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf(CapEndpoint+ "/hcaptcha/create?api_key=%s", hCaptcha.ApiKey), bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	req.Header.Add("content-type", "application/json")

	resp, err := request(req)
	if err != nil {
		return "", err
	}

	var task createTaskResp
	err = json.Unmarshal([]byte(resp), &task)
	if err != nil {
		return "", err
	}

	if task.Error != "" {
		return "", errors.New(task.Error)
	}

	return task.Id, nil
}

// getSolution attempts to retrieve a token given a task.
// An error is returned if the request was unable to be fulfilled,
// or if a captcha was not solved.
func (hCaptcha *HCaptcha) getSolution(task string) (string, error) {
	var status statusResp
	req, err := http.NewRequest("GET", fmt.Sprintf(CapEndpoint+ "/hcaptcha/status?api_key=%s&task_id=%s", hCaptcha.ApiKey, task), nil)
	if err != nil {
		return "", err
	}
	// wait for server to respond with a token or error
	for status.Status == "solving" || status.Status == "" {
		resp, err := request(req)
		if err != nil {
			return "", err
		}

		err = json.Unmarshal([]byte(resp), &status)
		if err != nil {
			return "", err
		}
	}
	// ensure status is successful
	if status.Status != "success" {
		if status.Reason != "" {
			return "", errors.New(status.Reason)
		}
		return "", errors.New("error retrieving hCaptcha token")
	}

	return status.Answer, nil
}
