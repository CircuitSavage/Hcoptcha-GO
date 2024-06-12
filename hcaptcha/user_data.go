package hcaptcha

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

type UserData struct {
	Data  UserDataInfo `json:"data"`
	Error bool         `json:"error"`
}

type UserDataInfo struct {
	APIKey         string  `json:"api_key"`
	Balance        float64 `json:"balance"`
	MaxThreads     int     `json:"max_threads"`
	Rank           string  `json:"rank"`
	RunningThreads int     `json:"running_threads"`
	Username       string  `json:"username"`
}

// Hcaptcha represents an hCaptcha client.
type Hcaptcha struct {
	apiKey string
	client *fasthttp.Client
}

// NewHcaptchaClient creates a new hCaptcha client.
func NewHcaptchaClient(apiKey string) *Hcaptcha {
	return &Hcaptcha{
		apiKey: apiKey,
		client: &fasthttp.Client{},
	}
}

func (h *Hcaptcha) GetUserData() (bool, UserData) {
	url := fmt.Sprintf("https://api.hcoptcha.online/api/getUserData?api_key=%s", h.apiKey)

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url)
	req.Header.SetMethod("GET")

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := h.client.Do(req, resp); err != nil {
		return false, UserData{Error: true}
	}

	var userData UserData
	if err := json.Unmarshal(resp.Body(), &userData); err != nil {
		return false, UserData{Error: true}
	}

	return true, userData
}
