package service

// TwilioClient represents the Twilio client.
type TwilioClient struct {
	accountSID string
	authToken  string
}

// NewTwilioClient creates a new Twilio client.
func NewTwilioClient(accountSID, authToken string) *TwilioClient {
	return &TwilioClient{
		accountSID: accountSID,
		authToken:  authToken,
	}
}
