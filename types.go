package checkly

import (
	"io"
	"net/http"
	"time"
)

// Client represents a Checkly client. If the Debug field is set to an io.Writer
// (for example os.Stdout), then the client will dump API requests and responses
// to it.  To use a non-default HTTP client (for example, for testing, or to set
// a timeout), assign to the HTTPClient field. To set a non-default URL (for
// example, for testing), assign to the URL field.
type Client struct {
	apiKey     string
	URL        string
	HTTPClient *http.Client
	Debug      io.Writer
}

// TypeBrowser is used to identify a browser check.
const TypeBrowser = "BROWSER"

// TypeAPI is used to identify an API check.
const TypeAPI = "API"

// RunBased identifies a run-based escalation type, for use with an AlertSettings.
const RunBased = "RUN_BASED"

// TimeBased identifies a time-based escalation type, for use with an AlertSettings.
const TimeBased = "TIME_BASED"

// Check represents the parameters for an existing check.
type Check struct {
	ID                     string                `json:"id"`
	Name                   string                `json:"name"`
	Type                   string                `json:"checkType"`
	Frequency              int                   `json:"frequency"`
	Activated              bool                  `json:"activated"`
	Muted                  bool                  `json:"muted"`
	ShouldFail             bool                  `json:"shouldFail"`
	Locations              []string              `json:"locations"`
	Script                 string                `json:"script,omitempty"`
	CreatedAt              time.Time             `json:"created_at,omitempty"`
	UpdatedAt              time.Time             `json:"updated_at,omitempty"`
	EnvironmentVariables   []EnvironmentVariable `json:"environment_variables"`
	DoubleCheck            bool                  `json:"doubleCheck"`
	Tags                   []string              `json:"tags,omitempty"`
	SSLCheck               bool                  `json:"sslCheck,omitempty"`
	SSLCheckDomain         string                `json:"sslCheckDomain,omitempty"`
	SetupSnippetID         int64                 `json:"setupSnippetId,omitempty"`
	TearDownSnippetID      int64                 `json:"tearDownSnippetId,omitempty"`
	LocalSetupScript       string                `json:"localSetupScript,omitempty"`
	LocalTearDownScript    string                `json:"localTearDownScript,omitempty"`
	AlertChannels          AlertChannels         `json:"alertChannels, omitempty"`
	AlertSettings          AlertSettings         `json:"alertSettings,omitempty"`
	UseGlobalAlertSettings bool                  `jons:"useGlobalAlertSettings"`
	Request                Request               `json:"request"`
}

// Request represents the parameters for the request made by the check.
type Request struct {
	Method string `json:"method"`
	URL    string `json:"url"`
}

// EnvironmentVariable represents a key-value pair for setting environment
// values during check execution.
type EnvironmentVariable struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
	Locked bool   `json:"locked"`
}

// AlertChannels represents the possible ways an alert notification can be sent.
type AlertChannels struct {
	Email   []AlertEmail   `json:"email,omitempty"`
	Webhook []AlertWebhook `json:"webhook,omitempty"`
	Slack   []AlertSlack   `json:"slack,omitempty"`
	SMS     []AlertSMS     `json:"sms,omitempty"`
}

// AlertEmail represents an email contact for alerts.
type AlertEmail struct {
	Address string `json:"address"`
}

// AlertWebhook represents a webhook contact for alerts.
type AlertWebhook struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// AlertSlack represents a Slack contact for alerts.
type AlertSlack struct {
	URL string `json:"url"`
}

// AlertSMS represents an SMS contact for alerts.
type AlertSMS struct {
	Number string `json:"number"`
	Name   string `json:"name"`
}

// AlertSettings represents an alert configuration.
type AlertSettings struct {
	EscalationType      string              `json:"escalationType,omitempty"`
	RunBasedEscalation  RunBasedEscalation  `json:"runBasedEscalation,omitempty"`
	TimeBasedEscalation TimeBasedEscalation `json:"timeBasedEscalation,omitempty"`
	Reminders           Reminders           `json:"reminders,omitempty"`
	SSLCertificates     SSLCertificates     `json:"sslCertificates,omitempty"`
}

// RunBasedEscalation represents an alert escalation based on a number of failed
// check runs.
type RunBasedEscalation struct {
	FailedRunThreshold int `json:"failedRunThreshold,omitempty"`
}

// TimeBasedEscalation represents an alert escalation based on the number of
// minutes after a check first starts failing.
type TimeBasedEscalation struct {
	MinutesFailingThreshold int `json:"minutesFailingThreshold,omitempty"`
}

// Reminders represents the number of reminders to send after an alert
// notification, and the time interval between them.
type Reminders struct {
	Amount   int `json:"amount,omitempty"`
	Interval int `json:"interval,omitempty"`
}

// SSLCertificates represents alert settings for expiring SSL certificates.
type SSLCertificates struct {
	Enabled        bool `json:"enabled,omitempty"`
	AlertThreshold int  `json:"alertThreshold,omitempty"`
}