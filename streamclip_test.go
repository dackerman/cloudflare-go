// File generated from our OpenAPI spec by Stainless.

package cloudflare_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/internal/testutil"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

func TestStreamClipNewWithOptionalParams(t *testing.T) {
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
	)
	_, err := client.Stream.Clips.New(context.TODO(), cloudflare.StreamClipNewParams{
		AccountID:             cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		ClippedFromVideoUid:   cloudflare.F("023e105f4ecef8ad9ca31a8372d0c353"),
		EndTimeSeconds:        cloudflare.F(int64(0)),
		StartTimeSeconds:      cloudflare.F(int64(0)),
		AllowedOrigins:        cloudflare.F([]string{"example.com"}),
		Creator:               cloudflare.F("creator-id_abcde12345"),
		MaxDurationSeconds:    cloudflare.F(int64(1)),
		RequireSignedURLs:     cloudflare.F(true),
		ThumbnailTimestampPct: cloudflare.F(0.529241),
		Watermark: cloudflare.F(cloudflare.StreamClipNewParamsWatermark{
			Uid: cloudflare.F("ea95132c15732412d22c1476fa83f27a"),
		}),
	})
	if err != nil {
		var apierr *cloudflare.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}