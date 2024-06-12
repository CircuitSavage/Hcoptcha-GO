package hcaptcha

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

// Client represents an hCaptcha client.
type Client struct {
	apiKey string
	client *fasthttp.Client
}

// NewClient creates a new hCaptcha client.
func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
		client: &fasthttp.Client{},
	}
}

// CreateTask creates a new captcha solving task.
func (h *Client) CreateTask(proxy, siteKey, rqdata, url string) (string, error) {
	createTaskURL := "https://api.hcoptcha.online/api/createTask"
	payload := map[string]interface{}{
		"api_key":   h.apiKey,
		"task_type": "hcaptchaEnterprise",
		"data": map[string]string{
			"url":     url,
			"sitekey": siteKey,
			"proxy":   proxy,
		},
	}
	if rqdata != "" {
		payload["data"].(map[string]string)["rqdata"] = rqdata
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(createTaskURL)
	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/json")
	req.SetBody(jsonData)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := h.client.Do(req, resp); err != nil {
		return "", err
	}

	if resp.StatusCode() != fasthttp.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	var responseData map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &responseData); err != nil {
		return "", err
	}

	if errorVal, ok := responseData["error"]; ok && errorVal.(bool) {
		return "", fmt.Errorf("task creation failed: %s", responseData["message"])
	}

	return responseData["task_id"].(string), nil
}

// GetTaskData retrieves the data of a task by its ID.
func (h *Client) GetTaskData(taskID string) (map[string]interface{}, error) {
	getTaskDataURL := "https://api.hcoptcha.online/api/getTaskData"

	// Prepare the request body
	requestBody := map[string]string{
		"api_key": h.apiKey,
		"task_id": taskID,
	}
	requestBodyJSON, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(getTaskDataURL)
	req.Header.SetMethod("POST")
	req.Header.SetContentType("application/json")
	req.SetBody(requestBodyJSON)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := h.client.Do(req, resp); err != nil {
		return nil, err
	}

	var responseData map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &responseData); err != nil {
		return nil, err
	}

	return responseData, nil
}
