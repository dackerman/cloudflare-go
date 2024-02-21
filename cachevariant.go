// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// CacheVariantService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewCacheVariantService] method
// instead.
type CacheVariantService struct {
	Options []option.RequestOption
}

// NewCacheVariantService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCacheVariantService(opts ...option.RequestOption) (r *CacheVariantService) {
	r = &CacheVariantService{}
	r.Options = opts
	return
}

// Variant support enables caching variants of images with certain file extensions
// in addition to the original. This only applies when the origin server sends the
// 'Vary: Accept' response header. If the origin server sends 'Vary: Accept' but
// does not serve the variant requested, the response will not be cached. This will
// be indicated with BYPASS cache status in the response headers.
func (r *CacheVariantService) List(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *CacheVariantListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CacheVariantListResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/variants", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Variant support enables caching variants of images with certain file extensions
// in addition to the original. This only applies when the origin server sends the
// 'Vary: Accept' response header. If the origin server sends 'Vary: Accept' but
// does not serve the variant requested, the response will not be cached. This will
// be indicated with BYPASS cache status in the response headers.
func (r *CacheVariantService) Delete(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *CacheVariantDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CacheVariantDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/variants", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Variant support enables caching variants of images with certain file extensions
// in addition to the original. This only applies when the origin server sends the
// 'Vary: Accept' response header. If the origin server sends 'Vary: Accept' but
// does not serve the variant requested, the response will not be cached. This will
// be indicated with BYPASS cache status in the response headers.
func (r *CacheVariantService) Edit(ctx context.Context, zoneID string, body CacheVariantEditParams, opts ...option.RequestOption) (res *CacheVariantEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CacheVariantEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/cache/variants", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Variant support enables caching variants of images with certain file extensions
// in addition to the original. This only applies when the origin server sends the
// 'Vary: Accept' response header. If the origin server sends 'Vary: Accept' but
// does not serve the variant requested, the response will not be cached. This will
// be indicated with BYPASS cache status in the response headers.
type CacheVariantListResponse struct {
	// ID of the zone setting.
	ID CacheVariantListResponseID `json:"id,required"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,required,nullable" format:"date-time"`
	// Value of the zone setting.
	Value CacheVariantListResponseValue `json:"value,required"`
	JSON  cacheVariantListResponseJSON  `json:"-"`
}

// cacheVariantListResponseJSON contains the JSON metadata for the struct
// [CacheVariantListResponse]
type cacheVariantListResponseJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheVariantListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type CacheVariantListResponseID string

const (
	CacheVariantListResponseIDVariants CacheVariantListResponseID = "variants"
)

// Value of the zone setting.
type CacheVariantListResponseValue struct {
	// List of strings with the MIME types of all the variants that should be served
	// for avif.
	Avif []interface{} `json:"avif"`
	// List of strings with the MIME types of all the variants that should be served
	// for bmp.
	Bmp []interface{} `json:"bmp"`
	// List of strings with the MIME types of all the variants that should be served
	// for gif.
	Gif []interface{} `json:"gif"`
	// List of strings with the MIME types of all the variants that should be served
	// for jp2.
	Jp2 []interface{} `json:"jp2"`
	// List of strings with the MIME types of all the variants that should be served
	// for jpeg.
	Jpeg []interface{} `json:"jpeg"`
	// List of strings with the MIME types of all the variants that should be served
	// for jpg.
	Jpg []interface{} `json:"jpg"`
	// List of strings with the MIME types of all the variants that should be served
	// for jpg2.
	Jpg2 []interface{} `json:"jpg2"`
	// List of strings with the MIME types of all the variants that should be served
	// for png.
	Png []interface{} `json:"png"`
	// List of strings with the MIME types of all the variants that should be served
	// for tif.
	Tif []interface{} `json:"tif"`
	// List of strings with the MIME types of all the variants that should be served
	// for tiff.
	Tiff []interface{} `json:"tiff"`
	// List of strings with the MIME types of all the variants that should be served
	// for webp.
	Webp []interface{}                     `json:"webp"`
	JSON cacheVariantListResponseValueJSON `json:"-"`
}

// cacheVariantListResponseValueJSON contains the JSON metadata for the struct
// [CacheVariantListResponseValue]
type cacheVariantListResponseValueJSON struct {
	Avif        apijson.Field
	Bmp         apijson.Field
	Gif         apijson.Field
	Jp2         apijson.Field
	Jpeg        apijson.Field
	Jpg         apijson.Field
	Jpg2        apijson.Field
	Png         apijson.Field
	Tif         apijson.Field
	Tiff        apijson.Field
	Webp        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheVariantListResponseValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Variant support enables caching variants of images with certain file extensions
// in addition to the original. This only applies when the origin server sends the
// 'Vary: Accept' response header. If the origin server sends 'Vary: Accept' but
// does not serve the variant requested, the response will not be cached. This will
// be indicated with BYPASS cache status in the response headers.
type CacheVariantDeleteResponse struct {
	// ID of the zone setting.
	ID CacheVariantDeleteResponseID `json:"id,required"`
	// last time this setting was modified.
	ModifiedOn time.Time                      `json:"modified_on,required,nullable" format:"date-time"`
	JSON       cacheVariantDeleteResponseJSON `json:"-"`
}

// cacheVariantDeleteResponseJSON contains the JSON metadata for the struct
// [CacheVariantDeleteResponse]
type cacheVariantDeleteResponseJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheVariantDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type CacheVariantDeleteResponseID string

const (
	CacheVariantDeleteResponseIDVariants CacheVariantDeleteResponseID = "variants"
)

// Variant support enables caching variants of images with certain file extensions
// in addition to the original. This only applies when the origin server sends the
// 'Vary: Accept' response header. If the origin server sends 'Vary: Accept' but
// does not serve the variant requested, the response will not be cached. This will
// be indicated with BYPASS cache status in the response headers.
type CacheVariantEditResponse struct {
	// ID of the zone setting.
	ID CacheVariantEditResponseID `json:"id,required"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,required,nullable" format:"date-time"`
	// Value of the zone setting.
	Value CacheVariantEditResponseValue `json:"value,required"`
	JSON  cacheVariantEditResponseJSON  `json:"-"`
}

// cacheVariantEditResponseJSON contains the JSON metadata for the struct
// [CacheVariantEditResponse]
type cacheVariantEditResponseJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheVariantEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the zone setting.
type CacheVariantEditResponseID string

const (
	CacheVariantEditResponseIDVariants CacheVariantEditResponseID = "variants"
)

// Value of the zone setting.
type CacheVariantEditResponseValue struct {
	// List of strings with the MIME types of all the variants that should be served
	// for avif.
	Avif []interface{} `json:"avif"`
	// List of strings with the MIME types of all the variants that should be served
	// for bmp.
	Bmp []interface{} `json:"bmp"`
	// List of strings with the MIME types of all the variants that should be served
	// for gif.
	Gif []interface{} `json:"gif"`
	// List of strings with the MIME types of all the variants that should be served
	// for jp2.
	Jp2 []interface{} `json:"jp2"`
	// List of strings with the MIME types of all the variants that should be served
	// for jpeg.
	Jpeg []interface{} `json:"jpeg"`
	// List of strings with the MIME types of all the variants that should be served
	// for jpg.
	Jpg []interface{} `json:"jpg"`
	// List of strings with the MIME types of all the variants that should be served
	// for jpg2.
	Jpg2 []interface{} `json:"jpg2"`
	// List of strings with the MIME types of all the variants that should be served
	// for png.
	Png []interface{} `json:"png"`
	// List of strings with the MIME types of all the variants that should be served
	// for tif.
	Tif []interface{} `json:"tif"`
	// List of strings with the MIME types of all the variants that should be served
	// for tiff.
	Tiff []interface{} `json:"tiff"`
	// List of strings with the MIME types of all the variants that should be served
	// for webp.
	Webp []interface{}                     `json:"webp"`
	JSON cacheVariantEditResponseValueJSON `json:"-"`
}

// cacheVariantEditResponseValueJSON contains the JSON metadata for the struct
// [CacheVariantEditResponseValue]
type cacheVariantEditResponseValueJSON struct {
	Avif        apijson.Field
	Bmp         apijson.Field
	Gif         apijson.Field
	Jp2         apijson.Field
	Jpeg        apijson.Field
	Jpg         apijson.Field
	Jpg2        apijson.Field
	Png         apijson.Field
	Tif         apijson.Field
	Tiff        apijson.Field
	Webp        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheVariantEditResponseValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheVariantListResponseEnvelope struct {
	Errors   []CacheVariantListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CacheVariantListResponseEnvelopeMessages `json:"messages,required"`
	// Variant support enables caching variants of images with certain file extensions
	// in addition to the original. This only applies when the origin server sends the
	// 'Vary: Accept' response header. If the origin server sends 'Vary: Accept' but
	// does not serve the variant requested, the response will not be cached. This will
	// be indicated with BYPASS cache status in the response headers.
	Result CacheVariantListResponse `json:"result,required"`
	// Whether the API call was successful
	Success CacheVariantListResponseEnvelopeSuccess `json:"success,required"`
	JSON    cacheVariantListResponseEnvelopeJSON    `json:"-"`
}

// cacheVariantListResponseEnvelopeJSON contains the JSON metadata for the struct
// [CacheVariantListResponseEnvelope]
type cacheVariantListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheVariantListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheVariantListResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    cacheVariantListResponseEnvelopeErrorsJSON `json:"-"`
}

// cacheVariantListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CacheVariantListResponseEnvelopeErrors]
type cacheVariantListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheVariantListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheVariantListResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    cacheVariantListResponseEnvelopeMessagesJSON `json:"-"`
}

// cacheVariantListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CacheVariantListResponseEnvelopeMessages]
type cacheVariantListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheVariantListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CacheVariantListResponseEnvelopeSuccess bool

const (
	CacheVariantListResponseEnvelopeSuccessTrue CacheVariantListResponseEnvelopeSuccess = true
)

type CacheVariantDeleteResponseEnvelope struct {
	Errors   []CacheVariantDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CacheVariantDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Variant support enables caching variants of images with certain file extensions
	// in addition to the original. This only applies when the origin server sends the
	// 'Vary: Accept' response header. If the origin server sends 'Vary: Accept' but
	// does not serve the variant requested, the response will not be cached. This will
	// be indicated with BYPASS cache status in the response headers.
	Result CacheVariantDeleteResponse `json:"result,required"`
	// Whether the API call was successful
	Success CacheVariantDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    cacheVariantDeleteResponseEnvelopeJSON    `json:"-"`
}

// cacheVariantDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [CacheVariantDeleteResponseEnvelope]
type cacheVariantDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheVariantDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheVariantDeleteResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    cacheVariantDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// cacheVariantDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CacheVariantDeleteResponseEnvelopeErrors]
type cacheVariantDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheVariantDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheVariantDeleteResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    cacheVariantDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// cacheVariantDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CacheVariantDeleteResponseEnvelopeMessages]
type cacheVariantDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheVariantDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CacheVariantDeleteResponseEnvelopeSuccess bool

const (
	CacheVariantDeleteResponseEnvelopeSuccessTrue CacheVariantDeleteResponseEnvelopeSuccess = true
)

type CacheVariantEditParams struct {
	// Value of the zone setting.
	Value param.Field[CacheVariantEditParamsValue] `json:"value,required"`
}

func (r CacheVariantEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the zone setting.
type CacheVariantEditParamsValue struct {
	// List of strings with the MIME types of all the variants that should be served
	// for avif.
	Avif param.Field[[]interface{}] `json:"avif"`
	// List of strings with the MIME types of all the variants that should be served
	// for bmp.
	Bmp param.Field[[]interface{}] `json:"bmp"`
	// List of strings with the MIME types of all the variants that should be served
	// for gif.
	Gif param.Field[[]interface{}] `json:"gif"`
	// List of strings with the MIME types of all the variants that should be served
	// for jp2.
	Jp2 param.Field[[]interface{}] `json:"jp2"`
	// List of strings with the MIME types of all the variants that should be served
	// for jpeg.
	Jpeg param.Field[[]interface{}] `json:"jpeg"`
	// List of strings with the MIME types of all the variants that should be served
	// for jpg.
	Jpg param.Field[[]interface{}] `json:"jpg"`
	// List of strings with the MIME types of all the variants that should be served
	// for jpg2.
	Jpg2 param.Field[[]interface{}] `json:"jpg2"`
	// List of strings with the MIME types of all the variants that should be served
	// for png.
	Png param.Field[[]interface{}] `json:"png"`
	// List of strings with the MIME types of all the variants that should be served
	// for tif.
	Tif param.Field[[]interface{}] `json:"tif"`
	// List of strings with the MIME types of all the variants that should be served
	// for tiff.
	Tiff param.Field[[]interface{}] `json:"tiff"`
	// List of strings with the MIME types of all the variants that should be served
	// for webp.
	Webp param.Field[[]interface{}] `json:"webp"`
}

func (r CacheVariantEditParamsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CacheVariantEditResponseEnvelope struct {
	Errors   []CacheVariantEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CacheVariantEditResponseEnvelopeMessages `json:"messages,required"`
	// Variant support enables caching variants of images with certain file extensions
	// in addition to the original. This only applies when the origin server sends the
	// 'Vary: Accept' response header. If the origin server sends 'Vary: Accept' but
	// does not serve the variant requested, the response will not be cached. This will
	// be indicated with BYPASS cache status in the response headers.
	Result CacheVariantEditResponse `json:"result,required"`
	// Whether the API call was successful
	Success CacheVariantEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    cacheVariantEditResponseEnvelopeJSON    `json:"-"`
}

// cacheVariantEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [CacheVariantEditResponseEnvelope]
type cacheVariantEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheVariantEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheVariantEditResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    cacheVariantEditResponseEnvelopeErrorsJSON `json:"-"`
}

// cacheVariantEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CacheVariantEditResponseEnvelopeErrors]
type cacheVariantEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheVariantEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CacheVariantEditResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    cacheVariantEditResponseEnvelopeMessagesJSON `json:"-"`
}

// cacheVariantEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CacheVariantEditResponseEnvelopeMessages]
type cacheVariantEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CacheVariantEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CacheVariantEditResponseEnvelopeSuccess bool

const (
	CacheVariantEditResponseEnvelopeSuccessTrue CacheVariantEditResponseEnvelopeSuccess = true
)