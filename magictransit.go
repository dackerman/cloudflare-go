// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// MagicTransitService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewMagicTransitService] method
// instead.
type MagicTransitService struct {
	Options         []option.RequestOption
	CfInterconnects *MagicTransitCfInterconnectService
	GRETunnels      *MagicTransitGRETunnelService
	IPSECTunnels    *MagicTransitIPSECTunnelService
	Routes          *MagicTransitRouteService
}

// NewMagicTransitService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMagicTransitService(opts ...option.RequestOption) (r *MagicTransitService) {
	r = &MagicTransitService{}
	r.Options = opts
	r.CfInterconnects = NewMagicTransitCfInterconnectService(opts...)
	r.GRETunnels = NewMagicTransitGRETunnelService(opts...)
	r.IPSECTunnels = NewMagicTransitIPSECTunnelService(opts...)
	r.Routes = NewMagicTransitRouteService(opts...)
	return
}