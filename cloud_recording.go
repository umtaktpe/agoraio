package agoraio

import (
	"fmt"
)

type AcquireParameters struct {
	Cname         string                         `json:"cname"`
	UID           string                         `json:"uid"`
	ClientRequest AcquireClientRequestParameters `json:"clientRequest"`
}

type AcquireClientRequestParameters struct {
	Region              string `json:"region"`
	ResourceExpiredHour int    `json:"resourceExpiredHour"`
	Scene               int    `json:"scene"`
}

type AcquireResponse struct {
	Cname      string `json:"cname"`
	UID        string `json:"uid"`
	ResourceID string `json:"resourceId"`
}

type StartRecordingClientRequestParameters struct {
	Token           string          `json:"token"`
	RecordingConfig RecordingConfig `json:"recordingConfig"`
	StorageConfig   StorageConfig   `json:"storageConfig"`
}

type RecordingConfig struct {
	ChannelType        int      `json:"channelType"` // 0: communication, 1: live broadcast
	StreamTypes        int      `json:"streamTypes"` // 0: audio, 1: video, 2: audio and video
	StreamMode         string   `json:"streamMode"`  // default, standard, original
	DecryptionMode     int      `json:"decryptionMode"`
	MaxIdleTime        int      `json:"maxIdleTime"`
	SubscribeAudioUids []string `json:"subscribeAudioUids"`
	SubscribeUidGroup  int      `json:"subscribeUidGroup"`
}

type TranscodingConfig struct {
	Width                      int                `json:"width"`
	Height                     int                `json:"height"`
	Fps                        int                `json:"fps"`
	Bitrate                    int                `json:"bitrate"`
	MaxResolutionUid           string             `json:"maxResolutionUid"`
	MixedVideoLayout           int                `json:"mixedVideoLayout"`
	BackgroundColor            string             `json:"backgroundColor"`
	BackgroundImage            string             `json:"backgroundImage"`
	DefaultUserBackgroundImage string             `json:"defaultUserBackgroundImage"`
	LayoutConfig               []LayoutConfig     `json:"layoutConfig"`
	BackgroundConfig           []BackgroundConfig `json:"backgroundConfig"`
}

type LayoutConfig struct {
	Uid        string  `json:"uid"`
	XAxis      float64 `json:"x_axis"`
	YAxis      float64 `json:"y_axis"`
	Width      float64 `json:"width"`
	Height     float64 `json:"height"`
	Alpha      float64 `json:"alpha"`
	RenderMode int     `json:"render_mode"`
}

type BackgroundConfig struct {
	Uid        string `json:"uid"`
	ImageUrl   string `json:"image_url"`
	RenderMode int    `json:"render_mode"`
}

type StorageConfig struct {
	Vendor          int             `json:"vendor"`
	Region          int             `json:"region"`
	Bucket          string          `json:"bucket"`
	AccessKey       string          `json:"accessKey"`
	SecretKey       string          `json:"secretKey"`
	FileNamePrefix  []string        `json:"fileNamePrefix"`
	ExtensionParams ExtensionParams `json:"extensionParams"`
}

type ExtensionParams struct {
	Sse string `json:"sse"`
	Tag string `json:"tag"`
}

type StartRecordingParameters struct {
	ResourceID    string                                `json:"resourceId"`
	Mode          string                                `json:"mode"`
	Cname         string                                `json:"cname"`
	UID           string                                `json:"uid"`
	ClientRequest StartRecordingClientRequestParameters `json:"clientRequest"`
}

type StartRecordingResponse struct {
	Sid        string `json:"sid"`
	ResourceID string `json:"resourceId"`
}

type RecordingStatusParameters struct {
	ResourceID string `json:"resourceId"`
	Sid        string `json:"sid"`
	Mode       string `json:"mode"`
}

type RecordingStatusResponse struct {
	ResourceID     string `json:"resourceId"`
	Sid            string `json:"sid"`
	ServerResponse struct {
		FileListMode   string      `json:"fileListMode"`
		FileList       interface{} `json:"fileList"`
		Status         int         `json:"status"`
		SliceStartTime int64       `json:"sliceStartTime"`
	} `json:"serverResponse"`
}

type StopRecordingParameters struct {
	AppID      string `json:"appid"`
	ResourceID string `json:"resourceId"`
	Sid        string `json:"sid"`
	Mode       string `json:"mode"`
	Cname      string `json:"cname"`
	Uid        string `json:"uid"`
}

func (c *Client) Acquire(params *AcquireParameters) (*AcquireResponse, error) {
	response := &AcquireResponse{}
	c.baseURL = "https://api.agora.io"
	url := fmt.Sprintf("/v1/apps/%s/cloud_recording/acquire", c.appID)
	err := c.request("POST", url, params, response)

	return response, err
}

func (c *Client) StartRecording(params *StartRecordingParameters) (*StartRecordingResponse, error) {
	response := &StartRecordingResponse{}
	c.baseURL = "https://api.agora.io"
	url := fmt.Sprintf("/v1/apps/%s/cloud_recording/resourceid/%s/mode/%s/start", c.appID, params.ResourceID, params.Mode)
	err := c.request("POST", url, params, response)

	return response, err
}

func (c *Client) RecordingStatus(params *RecordingStatusParameters) (*RecordingStatusResponse, error) {
	response := &RecordingStatusResponse{}
	c.baseURL = "https://api.agora.io"
	url := fmt.Sprintf("/v1/apps/%s/cloud_recording/resourceid/%s/sid/%s/mode/%s/query", c.appID, params.ResourceID, params.Sid, params.Mode)
	err := c.request("GET", url, nil, response)

	return response, err
}

func (c *Client) StopRecording(params *StopRecordingParameters) (interface{}, error) {
	var response interface{}
	c.baseURL = "https://api.agora.io"
	url := fmt.Sprintf("/v1/apps/%s/cloud_recording/resourceid/%s/sid/%s/mode/%s/stop", params.AppID, params.ResourceID, params.Sid, params.Mode)
	err := c.request("POST", url, params, response)

	return response, err
}
