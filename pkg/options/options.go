package options

type Options struct {
	Quota        OptsQuota        `json:"quota,omitempty"`
	Conservation OptsConservation `json:"conservation,omitempty"`
}

type OptsQuota struct {
	TenLastMinutes   int `json:"tenlastminutes"`
	SixtyLastMinutes int `json:"sixtylastminutes"`
	LastDay          int `json:"lastday"`
	LastWeek         int `json:"lastweek"`
	LastMonth        int `json:"lastmonth"`
}

type OptsConservation struct {
	Sent   int  `json:"sent"`
	Unsent int  `json:"unsent"`
	Keep   bool `json:"keep,omitempty"`
}
