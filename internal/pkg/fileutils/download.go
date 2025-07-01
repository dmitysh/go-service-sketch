package fileutils

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadFile(ctx context.Context, fp string, url string) error {
	f, err := os.OpenFile(fp, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0755) //nolint: gosec
	if err != nil {
		return fmt.Errorf("can't create file: %w", err)
	}
	defer f.Close()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return fmt.Errorf("can't prepare request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("can't do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("bad status: %s, body: %s", resp.Status, string(bodyBytes))
	}

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return fmt.Errorf("can't copy file to disk: %w", err)
	}

	return nil
}
