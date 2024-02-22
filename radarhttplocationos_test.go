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

func TestRadarHTTPLocationOSGetWithOptionalParams(t *testing.T) {
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
		option.WithAPIEmail("user@example.com"),
		option.WithAPIToken("Sn3lZJTBX6kkg7OdcBUAxOO963GEIyGQqnFTOFYY"),
		option.WithUserServiceKey("v1.0-144c9defac04969c7bfad8ef-631a41d003a32d25fe878081ef365c49503f7fada600da935e2851a1c7326084b85cbf6429c4b859de8475731dc92a9c329631e6d59e6c73da7b198497172b4cefe071d90d0f5d2719"),
	)
	_, err := client.Radar.HTTP.Locations.OS.Get(
		context.TODO(),
		cloudflare.RadarHTTPLocationOSGetParamsOSWindows,
		cloudflare.RadarHTTPLocationOSGetParams{
			ASN:          cloudflare.F([]string{"string", "string", "string"}),
			BotClass:     cloudflare.F([]cloudflare.RadarHTTPLocationOSGetParamsBotClass{cloudflare.RadarHTTPLocationOSGetParamsBotClassLikelyAutomated, cloudflare.RadarHTTPLocationOSGetParamsBotClassLikelyHuman}),
			DateEnd:      cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DateRange:    cloudflare.F([]cloudflare.RadarHTTPLocationOSGetParamsDateRange{cloudflare.RadarHTTPLocationOSGetParamsDateRange1d, cloudflare.RadarHTTPLocationOSGetParamsDateRange2d, cloudflare.RadarHTTPLocationOSGetParamsDateRange7d}),
			DateStart:    cloudflare.F([]time.Time{time.Now(), time.Now(), time.Now()}),
			DeviceType:   cloudflare.F([]cloudflare.RadarHTTPLocationOSGetParamsDeviceType{cloudflare.RadarHTTPLocationOSGetParamsDeviceTypeDesktop, cloudflare.RadarHTTPLocationOSGetParamsDeviceTypeMobile, cloudflare.RadarHTTPLocationOSGetParamsDeviceTypeOther}),
			Format:       cloudflare.F(cloudflare.RadarHTTPLocationOSGetParamsFormatJson),
			HTTPProtocol: cloudflare.F([]cloudflare.RadarHTTPLocationOSGetParamsHTTPProtocol{cloudflare.RadarHTTPLocationOSGetParamsHTTPProtocolHTTP, cloudflare.RadarHTTPLocationOSGetParamsHTTPProtocolHTTPS}),
			HTTPVersion:  cloudflare.F([]cloudflare.RadarHTTPLocationOSGetParamsHTTPVersion{cloudflare.RadarHTTPLocationOSGetParamsHTTPVersionHttPv1, cloudflare.RadarHTTPLocationOSGetParamsHTTPVersionHttPv2, cloudflare.RadarHTTPLocationOSGetParamsHTTPVersionHttPv3}),
			IPVersion:    cloudflare.F([]cloudflare.RadarHTTPLocationOSGetParamsIPVersion{cloudflare.RadarHTTPLocationOSGetParamsIPVersionIPv4, cloudflare.RadarHTTPLocationOSGetParamsIPVersionIPv6}),
			Limit:        cloudflare.F(int64(5)),
			Location:     cloudflare.F([]string{"string", "string", "string"}),
			Name:         cloudflare.F([]string{"string", "string", "string"}),
			TLSVersion:   cloudflare.F([]cloudflare.RadarHTTPLocationOSGetParamsTLSVersion{cloudflare.RadarHTTPLocationOSGetParamsTLSVersionTlSv1_0, cloudflare.RadarHTTPLocationOSGetParamsTLSVersionTlSv1_1, cloudflare.RadarHTTPLocationOSGetParamsTLSVersionTlSv1_2}),
		},
	)
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}