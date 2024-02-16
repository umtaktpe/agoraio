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
	ResourceExpiredHour int `json:"resourceExpiredHour"`
	Scene               int `json:"scene"`
}

type AcquireResponse struct {
	Cname      string `json:"cname"`
	UID        string `json:"uid"`
	ResourceID string `json:"resourceId"`
}

type StartRecordingClientRequestParameters struct {
	Token           string          `json:"token,omitempty"`
	RecordingConfig RecordingConfig `json:"recordingConfig"`
	StorageConfig   StorageConfig   `json:"storageConfig"`
}

type RecordingConfig struct {
	ChannelType        int      `json:"channelType"` // 0: communication, 1: live broadcast
	StreamTypes        int      `json:"streamTypes"` // 0: audio, 1: video, 2: audio and video
	StreamMode         string   `json:"streamMode"`  // default, standard, original
	MaxIdleTime        int      `json:"maxIdleTime"`
	SubscribeAudioUids []string `json:"subscribeAudioUids"`
	SubscribeUidGroup  int      `json:"subscribeUidGroup"`
	AudioProfile       int      `json:"audioProfile"`
	VideoStreamType    int      `json:"videoStreamType"`
}

type StorageConfig struct {
	Vendor         int      `json:"vendor"`
	Region         int      `json:"region"`
	Bucket         string   `json:"bucket"`
	AccessKey      string   `json:"accessKey"`
	SecretKey      string   `json:"secretKey"`
	FileNamePrefix []string `json:"fileNamePrefix"`
}

type StartRecordingParameters struct {
	Cname         string                                `json:"cname"`
	UID           string                                `json:"uid"`
	ClientRequest StartRecordingClientRequestParameters `json:"clientRequest"`
}

type StartRecordingResponse struct {
	Sid        string `json:"sid"`
	ResourceID string `json:"resourceId"`
}

type RecordingStatusResponse struct {
	Cname          string `json:"cname"`
	Uid            string `json:"uid"`
	ResourceID     string `json:"resourceId"`
	Sid            string `json:"sid"`
	Code           int    `json:"code,omitempty"`
	ServerResponse struct {
		FileListMode   string      `json:"fileListMode"`
		FileList       interface{} `json:"fileList"`
		Status         int         `json:"status"`
		SliceStartTime int64       `json:"sliceStartTime"`
	} `json:"serverResponse,omitempty"`
}

type StopRecordingParameters struct {
	Cname         string   `json:"cname"`
	Uid           string   `json:"uid"`
	ClientRequest struct{} `json:"clientRequest"`
}

func (c *Client) Acquire(params *AcquireParameters) (*AcquireResponse, error) {
	response := &AcquireResponse{}
	c.baseURL = "https://api.agora.io"
	url := fmt.Sprintf("/v1/apps/%s/cloud_recording/acquire", c.appID)
	err := c.request("POST", url, params, response)

	return response, err
}

func (c *Client) StartRecording(resourceID, mode string, params *StartRecordingParameters) (*StartRecordingResponse, error) {
	response := &StartRecordingResponse{}
	c.baseURL = "https://api.agora.io"
	url := fmt.Sprintf("/v1/apps/%s/cloud_recording/resourceid/%s/mode/%s/start", c.appID, resourceID, mode)
	err := c.request("POST", url, params, response)

	return response, err
}

func (c *Client) RecordingStatus(resourceID, sid, mode string) (*RecordingStatusResponse, error) {
	response := &RecordingStatusResponse{}
	c.baseURL = "https://api.agora.io"
	url := fmt.Sprintf("/v1/apps/%s/cloud_recording/resourceid/%s/sid/%s/mode/%s/query", c.appID, resourceID, sid, mode)
	err := c.request("GET", url, nil, response)

	return response, err
}

func (c *Client) StopRecording(resourceID, sid, mode string, params *StopRecordingParameters) (interface{}, error) {
	var response interface{}
	c.baseURL = "https://api.agora.io"
	url := fmt.Sprintf("/v1/apps/%s/cloud_recording/resourceid/%s/sid/%s/mode/%s/stop", c.appID, resourceID, sid, mode)
	err := c.request("POST", url, params, &response)

	return response, err
}
