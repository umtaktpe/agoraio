package agoraio

import (
	"fmt"
	"time"
)

type GetAllProjectsResponse struct {
	Projects []struct {
		ID              string      `json:"id"`
		Name            string      `json:"name"`
		SignKey         string      `json:"sign_key"`
		VendorKey       string      `json:"vendor_key"`
		RecordingServer interface{} `json:"recording_server"`
		Status          int         `json:"status"`
		Created         int         `json:"created"`
	} `json:"projects"`
}

type GetProjectParameters struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GetProjectResponse struct {
	Projects []struct {
		ID              string      `json:"id"`
		Name            string      `json:"name"`
		SignKey         string      `json:"sign_key"`
		VendorKey       string      `json:"vendor_key"`
		RecordingServer interface{} `json:"recording_server"`
		Status          int         `json:"status"`
		Created         int         `json:"created"`
	} `json:"projects"`
}

type CreateProjectParameters struct {
	Name          string `json:"name"`
	EnableSignKey bool   `json:"enable_sign_key"`
}

type CreateProjectResponse struct {
	Project struct {
		ID              string      `json:"id"`
		Name            string      `json:"name"`
		VendorKey       string      `json:"vendor_key"`
		SignKey         string      `json:"sign_key"`
		RecordingServer interface{} `json:"recording_server"`
		Status          int         `json:"status"`
		Created         int         `json:"created"`
	} `json:"project"`
}

type ProjectStatusParameters struct {
	ID     string `json:"id"`
	Status int    `json:"status"`
}

type ProjectStatusResponse struct {
	Project struct {
		ID              string      `json:"id"`
		Name            string      `json:"name"`
		VendorKey       string      `json:"vendor_key"`
		SignKey         string      `json:"sign_key"`
		RecordingServer interface{} `json:"recording_server"`
		Status          int         `json:"status"`
		Created         int         `json:"created"`
	} `json:"project"`
}

type RecordingConfigParameters struct {
	ID              string `json:"id"`
	RecordingServer string `json:"recording_server"`
}

type RecordingServerResponse struct {
	Project struct {
		ID              string      `json:"id"`
		Name            string      `json:"name"`
		VendorKey       string      `json:"vendor_key"`
		SignKey         string      `json:"sign_key"`
		RecordingServer interface{} `json:"recording_server"`
		Status          int         `json:"status"`
		Created         int         `json:"created"`
	} `json:"project"`
}

type SignKeyParameters struct {
	ID     string `json:"id"`
	Enable bool   `json:"enable"`
}

type SignKeyResponse struct {
	Project struct {
		ID              string      `json:"id"`
		Name            string      `json:"name"`
		VendorKey       string      `json:"vendor_key"`
		SignKey         string      `json:"sign_key"`
		RecordingServer interface{} `json:"recording_server"`
		Status          int         `json:"status"`
		Created         int         `json:"created"`
	} `json:"project"`
}

type ResetSignKeyParameters struct {
	ID string `json:"id"`
}

type ResetSignKeyResponse struct {
	Project struct {
		ID              string      `json:"id"`
		Name            string      `json:"name"`
		VendorKey       string      `json:"vendor_key"`
		SignKey         string      `json:"sign_key"`
		RecordingServer interface{} `json:"recording_server"`
		Status          int         `json:"status"`
		Created         int         `json:"created"`
	} `json:"project"`
}

type GetProjectUsageParameters struct {
	ProjectID string `json:"project_id"`
	FromDate  string `json:"from_date"`
	ToDate    string `json:"to_date"`
	Business  string `json:"business"`
}

type GetProjectUsageResponse struct {
	Meta struct {
		DurationAudioAll struct {
			En   string `json:"en"`
			Unit string `json:"unit"`
		} `json:"durationAudioAll"`
		DurationVideo1080P struct {
			En   string `json:"en"`
			Unit string `json:"unit"`
		} `json:"durationVideo1080P"`
		DurationVideo2K struct {
			En   string `json:"en"`
			Unit string `json:"unit"`
		} `json:"durationVideo2K"`
		DurationVideo4K struct {
			En   string `json:"en"`
			Unit string `json:"unit"`
		} `json:"durationVideo4K"`
		DurationVideoHd struct {
			En   string `json:"en"`
			Unit string `json:"unit"`
		} `json:"durationVideoHd"`
		DurationVideoHdp struct {
			En   string `json:"en"`
			Unit string `json:"unit"`
		} `json:"durationVideoHdp"`
	} `json:"meta"`
	Usages []struct {
		Date  time.Time `json:"date"`
		Usage struct {
			DurationAudioAll   int `json:"durationAudioAll"`
			DurationVideo1080P int `json:"durationVideo1080P"`
			DurationVideo2K    int `json:"durationVideo2K"`
			DurationVideo4K    int `json:"durationVideo4K"`
			DurationVideoHd    int `json:"durationVideoHd"`
			DurationVideoHdp   int `json:"durationVideoHdp"`
		} `json:"usage"`
	} `json:"usages"`
}

func (c *client) GetAllProjects() (*GetAllProjectsResponse, error) {
	response := &GetAllProjectsResponse{}
	err := c.request("GET", "/v1/projects", nil, response)

	return response, err
}

func (c *client) GetProject(params *GetProjectParameters) (*GetProjectResponse, error) {
	response := &GetProjectResponse{}
	url := fmt.Sprintf("/v1/project?id=%s&name=%s", params.ID, params.Name)
	err := c.request("GET", url, nil, response)

	return response, err
}

func (c *client) CreateProject(params *CreateProjectParameters) (*CreateProjectResponse, error) {
	response := &CreateProjectResponse{}
	err := c.request("POST", "/v1/project", params, response)

	return response, err
}

func (c *client) ChangeProjectStatus(params *ProjectStatusParameters) (*ProjectStatusResponse, error) {
	response := &ProjectStatusResponse{}
	err := c.request("POST", "/v1/project_status", params, response)

	return response, err
}

func (c *client) SetIPAddress(params *RecordingConfigParameters) (*RecordingServerResponse, error) {
	response := &RecordingServerResponse{}
	err := c.request("POST", "/v1/recording_config", params, response)

	return response, err
}

func (c *client) ChangePrimaryCertificateStatus(params *SignKeyParameters) (*SignKeyResponse, error) {
	response := &SignKeyResponse{}
	err := c.request("POST", "/v1/signkey", params, response)

	return response, err
}

func (c *client) ResetPrimaryCertificate(params *ResetSignKeyParameters) (*ResetSignKeyResponse, error) {
	response := &ResetSignKeyResponse{}
	err := c.request("POST", "/v1/reset_signkey", params, response)

	return response, err
}

func (c *client) GetProjectUsage(params *GetProjectUsageParameters) (*GetProjectUsageResponse, error) {
	response := &GetProjectUsageResponse{}
	url := fmt.Sprintf("/v3/usage?project_id=%s&from_date=%s&to_date=%s&business=%s", params.ProjectID, params.FromDate, params.ToDate, params.Business)
	err := c.request("GET", url, nil, response)

	return response, err
}
