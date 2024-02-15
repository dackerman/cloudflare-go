// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AnalyticsService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewAnalyticsService] method instead.
type AnalyticsService struct {
	Options    []option.RequestOption
	Colo       *AnalyticsColoService
	Dashboards *AnalyticsDashboardService
	Latencies  *AnalyticsLatencyService
}

// NewAnalyticsService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAnalyticsService(opts ...option.RequestOption) (r *AnalyticsService) {
	r = &AnalyticsService{}
	r.Options = opts
	r.Colo = NewAnalyticsColoService(opts...)
	r.Dashboards = NewAnalyticsDashboardService(opts...)
	r.Latencies = NewAnalyticsLatencyService(opts...)
	return
}