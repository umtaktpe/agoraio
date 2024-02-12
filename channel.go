package agoraio

import (
	"fmt"
)

type GetUserStatusParameters struct {
	UID         string `json:"uid"`
	ChannelName string `json:"channelName"`
}

type GetUserStatusResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Join      int  `json:"join"`
		UID       int  `json:"uid"`
		InChannel bool `json:"in_channel"`
		Platform  int  `json:"platform"`
		Role      int  `json:"role"`
	} `json:"data"`
}

type GetUserListParameters struct {
	ChannelName string `json:"channelName"`
}

type GetChannelListParameters struct {
	PageNo   int `json:"page_no"`
	PageSize int `json:"page_size"`
}

type GetChannelListResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Channels []struct {
			ChannelName string `json:"channel_name"`
			UserCount   int    `json:"user_count"`
		} `json:"channels"`
		TotalSize int `json:"total_size"`
	} `json:"data"`
}

func (c *Client) GetUserStatus(params *GetUserStatusParameters) (*GetUserStatusResponse, error) {
	response := &GetUserStatusResponse{}
	url := fmt.Sprintf("/v1/channel/user/property/%s/%s/%s", c.appID, params.UID, params.ChannelName)
	err := c.request("GET", url, nil, response)

	return response, err
}

func (c *Client) GetUserList(params *GetUserListParameters) (interface{}, error) {
	var response interface{}
	err := c.request("GET", "/v1/channel/user/"+c.appID+"/"+params.ChannelName, nil, &response)

	return response, err
}

func (c *Client) GetChannelList(params *GetChannelListParameters) (*GetChannelListResponse, error) {
	response := &GetChannelListResponse{}

	var url string
	if params != nil {
		url = fmt.Sprintf("/v1/channel/%s?page_no=%d&page_size=%d", c.appID, params.PageNo, params.PageSize)
	} else {
		url = fmt.Sprintf("/v1/channel/%s", c.appID)
	}
	err := c.request("GET", url, nil, response)

	return response, err
}
