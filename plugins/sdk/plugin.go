package sdk

// Plugin is will be the interface for all plugins development
type Plugin interface {
	// CheckMail will run the test itself: is a spam or not,
	// with its own business logic
	CheckMail([]byte) *Result
}

// Result is the "standard" way to return the spam test resultat
type Result struct {
	// Name of the plugin
	Name string `json:"name"`
	// Version of the plugin
	Version string `json:"version"`
	// Score of the spam test
	Score int `json:"score"`
	// Details are additional information for the check
	Details string `json:"details"`
	// TimeIn is the time when we send the email(in ms)
	TimeIn int64 `json:"time_in"`
	// TimeOut is the time when we get back the email(in ms)
	TimeOut int64 `json:"time_out"`
	// ExecTime is the execution time: the actual time to run the test
	ExecTime int64 `json:"exec_time"`
	// Recommended is the plugin recommended action
	Recommended string `json:"recommended"`
	// Components are the tested components
	Components []string `json:"components"`
	// Headers are the custom headers to add to the mail
	Headers []map[string]string `json:"headers"`
}
