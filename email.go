// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// EmailService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewEmailService] method instead.
type EmailService struct {
	Options  []option.RequestOption
	Routings *EmailRoutingService
}

// NewEmailService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEmailService(opts ...option.RequestOption) (r *EmailService) {
	r = &EmailService{}
	r.Options = opts
	r.Routings = NewEmailRoutingService(opts...)
	return
}