// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// R2Service contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewR2Service] method instead.
type R2Service struct {
	Options []option.RequestOption
	Buckets *R2BucketService
	Sippy   *R2SippyService
}

// NewR2Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewR2Service(opts ...option.RequestOption) (r *R2Service) {
	r = &R2Service{}
	r.Options = opts
	r.Buckets = NewR2BucketService(opts...)
	r.Sippy = NewR2SippyService(opts...)
	return
}