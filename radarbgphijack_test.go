// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestRadarBGPHijackEventsWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudflare.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("144c9defac04969c7bfad8efaa8ea194"),
		option.WithAPIEmail("dev@cloudflare.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("My User Service Key"),
	)
	_, err := client.Radar.BGP.Hijacks.Events(context.TODO(), cloudflare.RadarBGPHijackEventsParams{
		DateEnd:         cloudflare.F(time.Now()),
		DateRange:       cloudflare.F(cloudflare.RadarBGPHijackEventsParamsDateRange7d),
		DateStart:       cloudflare.F(time.Now()),
		EventID:         cloudflare.F(int64(0)),
		Format:          cloudflare.F(cloudflare.RadarBGPHijackEventsParamsFormatJson),
		HijackerAsn:     cloudflare.F(int64(0)),
		InvolvedAsn:     cloudflare.F(int64(0)),
		InvolvedCountry: cloudflare.F("string"),
		MaxConfidence:   cloudflare.F(int64(0)),
		MinConfidence:   cloudflare.F(int64(0)),
		Page:            cloudflare.F(int64(0)),
		PerPage:         cloudflare.F(int64(0)),
		Prefix:          cloudflare.F("string"),
		SortBy:          cloudflare.F(cloudflare.RadarBGPHijackEventsParamsSortByTime),
		SortOrder:       cloudflare.F(cloudflare.RadarBGPHijackEventsParamsSortOrderDesc),
		VictimAsn:       cloudflare.F(int64(0)),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}