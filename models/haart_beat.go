package models

type HeartbeatInfo struct {
	RadioStatus  []WiFiStatus `json:"radioStatus" binding:"required"`
	SystemStatus SystemInfo   `json:"SystemStatus" binding:"required"`
	DeviceStatus DeviceInfo   `json:"deviceStatus" binding:"required"`
	DID          string       `json:"did" binding:"required"`
}

type DeviceInfo struct {
	Kernel    string `json:"kernel"`
	Hostname  string `json:"hostname"`
	System    string `json:"system"`
	Model     string `json:"model"`
	BoardName string `json:"board_name"`
	Release   struct {
		Distribution string `json:"distribution"`
		Version      string `json:"version"`
		Revision     string `json:"revision"`
		Target       string `json:"target"`
		Description  string `json:"description"`
	} `json:"release"`
}

type SystemInfo struct {
	Localtime int   `json:"localtime"`
	Uptime    int   `json:"uptime"`
	Load      []int `json:"load"`
	Memory    struct {
		Total     int `json:"total"`
		Free      int `json:"free"`
		Shared    int `json:"shared"`
		Buffered  int `json:"buffered"`
		Available int `json:"available"`
		Cached    int `json:"cached"`
	} `json:"memory"`
	Swap struct {
		Total int `json:"total"`
		Free  int `json:"free"`
	} `json:"swap"`
}

type WiFiStatus struct {
	Disabled        bool   `json:"disabled"`
	Type            string `json:"type"` // openroaming | free
	Phy             string `json:"phy"`
	Ssid            string `json:"ssid"`
	Bssid           string `json:"bssid"`
	Country         string `json:"country"`
	Mode            string `json:"mode"`
	Channel         int    `json:"channel"`
	CenterChan1     int    `json:"center_chan1"`
	Frequency       int    `json:"frequency"`
	FrequencyOffset int    `json:"frequency_offset"`
	Txpower         int    `json:"txpower"`
	TxpowerOffset   int    `json:"txpower_offset"`
	Quality         int    `json:"quality"`
	QualityMax      int    `json:"quality_max"`
	Signal          int    `json:"signal"`
	Noise           int    `json:"noise"`
	Bitrate         int    `json:"bitrate"`
	Encryption      struct {
		Enabled        bool     `json:"enabled"`
		Wpa            []int    `json:"wpa,omitempty"`
		Authentication []string `json:"authentication,omitempty"`
		Ciphers        []string `json:"ciphers,omitempty"`
	} `json:"encryption"`
	Htmodes []string         `json:"htmodes"`
	Hwmodes []string         `json:"hwmodes"`
	Hwmode  string           `json:"hwmode"`
	Htmode  string           `json:"htmode"`
	Station []WifiClientInfo `json:"station"`
}

type WifiClientInfo struct {
	Hostname      string `json:"hostname"`
	Mac           string `json:"mac"`
	Signal        int    `json:"signal"`
	SignalAvg     int    `json:"signal_avg"`
	Noise         int    `json:"noise"`
	Inactive      int    `json:"inactive"`
	ConnectedTime int    `json:"connected_time"`
}
