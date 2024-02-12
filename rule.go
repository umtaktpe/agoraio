package agoraio

import (
	"time"
)

type CreateRuleParameters struct {
	AppID      string   `json:"appid"`
	Cname      string   `json:"cname"`
	UID        int      `json:"uid"`
	IP         string   `json:"ip"`
	Time       int      `json:"time"`
	Privileges []string `json:"privileges"`
}

type CreateRuleResponse struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

type GetRuleListResponse struct {
	Status string `json:"status"`
	Rules  []struct {
		ID         int       `json:"id"`
		AppID      string    `json:"appid"`
		UID        string    `json:"uid"`
		Opid       int       `json:"opid"`
		Cname      string    `json:"cname"`
		IP         string    `json:"ip"`
		Ts         time.Time `json:"ts"`
		Privileges []string  `json:"privileges"`
		CreateAt   time.Time `json:"createAt"`
		UpdateAt   time.Time `json:"updateAt"`
	} `json:"rules"`
}

type UpdateRuleParameters struct {
	ID    string `json:"id"`
	AppID string `json:"appid"`
	Time  int    `json:"time"`
}

type UpdateRuleResponse struct {
	Status string `json:"status"`
	Result struct {
		ID int       `json:"id"`
		Ts time.Time `json:"ts"`
	} `json:"result"`
}

type DeleteRuleParameters struct {
	ID    string `json:"id"`
	AppID string `json:"appid"`
}

type DeleteRuleResponse struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

func (c *Client) GetRuleList() (*GetRuleListResponse, error) {
	response := &GetRuleListResponse{}
	err := c.request("GET", "/v1/kicking-rule?appid="+c.appID, nil, response)

	return response, err
}

func (c *Client) CreateRule(params *CreateRuleParameters) (*CreateRuleResponse, error) {
	response := &CreateRuleResponse{}
	err := c.request("POST", "/v1/kicking-rule", params, response)

	return response, err
}

func (c *Client) UpdateRule(params *UpdateRuleParameters) (*UpdateRuleResponse, error) {
	response := &UpdateRuleResponse{}
	err := c.request("PUT", "/v1/kicking-rule", params, response)

	return response, err
}

func (c *Client) DeleteRule(params *DeleteRuleParameters) (*DeleteRuleResponse, error) {
	response := &DeleteRuleResponse{}
	err := c.request("DELETE", "/v1/kicking-rule", params, response)

	return response, err
}
