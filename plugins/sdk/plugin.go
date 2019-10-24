package sdk

// Plugin is will be the interface for all plugins development
type Plugin interface {
	// CheckMail will run the test itself: is a spam or not,
	// with its own business logic
	CheckMail([]byte)bool
}
