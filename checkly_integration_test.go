// +build integration

package checkly

import (
	"net/http"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func getAPIKey(t *testing.T) string {
	key := os.Getenv("CHECKLY_API_KEY")
	if key == "" {
		t.Fatal("'CHECKLY_API_KEY' must be set for integration tests")
	}
	return key
}

func testCheck(name string) Check {
	return Check{
		Name:      name,
		Type:      TypeAPI,
		Activated: true,
		Frequency: 1,
		Locations: []string{"eu-west-1"},
		Request: Request{
			Method:          http.MethodGet,
			URL:             "http://example.com",
			Headers:         []KeyValue{},
			QueryParameters: []KeyValue{},
			Assertions: []Assertion{
				Assertion{
					Source:     StatusCode,
					Comparison: Equals,
					Target:     "200",
				},
			},
			BodyType: "NONE",
		},
		AlertSettings: AlertSettings{
			RunBasedEscalation: RunBasedEscalation{
				FailedRunThreshold: 1,
			},
			TimeBasedEscalation: TimeBasedEscalation{
				MinutesFailingThreshold: 5,
			},
			Reminders: Reminders{
				Interval: 5,
			},
			SSLCertificates: SSLCertificates{
				Enabled:        false,
				AlertThreshold: 3,
			},
		},
		Tags:                   []string{},
		SSLCheck:               false,
		SSLCheckDomain:         "example.com",
		UseGlobalAlertSettings: false,
	}
}

func TestCreateGetIntegration(t *testing.T) {
	t.Parallel()
	client := NewClient(getAPIKey(t))
	checkCreate := testCheck("integrationTestCreate")
	// client.Debug = os.Stdout
	ID, err := client.Create(checkCreate)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Delete(ID)
	check, err := client.Get(ID)
	checkCreate.ID = ID
	if !cmp.Equal(checkCreate, check, cmpopts.IgnoreFields(Check{}, "CreatedAt")) {
		t.Error(cmp.Diff(checkCreate, check))
	}
}

func TestDeleteIntegration(t *testing.T) {
	t.Parallel()
	client := NewClient(getAPIKey(t))
	checkDelete := testCheck("integrationTestDelete")
	ID, err := client.Create(checkDelete)
	if err != nil {
		t.Fatal(err)
	}
	if err := client.Delete(ID); err != nil {
		t.Error(err)
	}
	_, err = client.Get(ID)
	if err == nil {
		t.Error("want error getting deleted check, but got nil")
	}
}
