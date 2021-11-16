package model

type UserPromoteTagInfo struct {
	AppsflyerId string `json:"appsflyer_id,omitempty"`
	Ip          string `json:"ip,omitempty"`
	AppVersion  string `json:"app_version,omitempty"`
	AppId       string `json:"app_id,omitempty"`
	PdTid       string `json:"pdtid,omitempty"`
	UserTag     string `json:"user_tag,omitempty"`
	Idfv        string `json:"idfv,omitempty"`
	Idfa        string `json:"idfa,omitempty"`
	Gaid        string `json:"gaid,omitempty"`
}
