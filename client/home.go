package client

import (
	"fmt"
	"golang.org/x/net/context"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
)

// DownloadStatic downloads /files with the given filename and writes it to the file dest.
// It returns the number of bytes downloaded in case of success.
func (c *Client) DownloadStatic(ctx context.Context, filename, dest string) (int64, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	p := path.Join("/static/", filename)
	u := url.URL{Host: c.Host, Scheme: scheme, Path: p}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return 0, err
	}
	resp, err := c.Client.Do(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 200 {
		var body string
		if b, err := ioutil.ReadAll(resp.Body); err != nil {
			if len(b) > 0 {
				body = ": " + string(b)
			}
		}
		return 0, fmt.Errorf("%s%s", resp.Status, body)
	}
	defer resp.Body.Close()
	out, err := os.Create(dest)
	if err != nil {
		return 0, err
	}
	defer out.Close()
	return io.Copy(out, resp.Body)
}

// DownloadBowerComponents downloads /files with the given filename and writes it to the file dest.
// It returns the number of bytes downloaded in case of success.
func (c *Client) DownloadBowerComponents(ctx context.Context, filename, dest string) (int64, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	p := path.Join("/bower_components/", filename)
	u := url.URL{Host: c.Host, Scheme: scheme, Path: p}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return 0, err
	}
	resp, err := c.Client.Do(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 200 {
		var body string
		if b, err := ioutil.ReadAll(resp.Body); err != nil {
			if len(b) > 0 {
				body = ": " + string(b)
			}
		}
		return 0, fmt.Errorf("%s%s", resp.Status, body)
	}
	defer resp.Body.Close()
	out, err := os.Create(dest)
	if err != nil {
		return 0, err
	}
	defer out.Close()
	return io.Copy(out, resp.Body)
}

// DownloadFaviconIco downloads favicon.ico and writes it to the file dest.
// It returns the number of bytes downloaded in case of success.
func (c *Client) DownloadFaviconIco(ctx context.Context, dest string) (int64, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: "/images/favicon.ico"}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return 0, err
	}
	resp, err := c.Client.Do(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 200 {
		var body string
		if b, err := ioutil.ReadAll(resp.Body); err != nil {
			if len(b) > 0 {
				body = ": " + string(b)
			}
		}
		return 0, fmt.Errorf("%s%s", resp.Status, body)
	}
	defer resp.Body.Close()
	out, err := os.Create(dest)
	if err != nil {
		return 0, err
	}
	defer out.Close()
	return io.Copy(out, resp.Body)
}

// DownloadIndexHTML downloads index.html and writes it to the file dest.
// It returns the number of bytes downloaded in case of success.
func (c *Client) DownloadIndexHTML(ctx context.Context, dest string) (int64, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: "/index.html"}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return 0, err
	}
	resp, err := c.Client.Do(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 200 {
		var body string
		if b, err := ioutil.ReadAll(resp.Body); err != nil {
			if len(b) > 0 {
				body = ": " + string(b)
			}
		}
		return 0, fmt.Errorf("%s%s", resp.Status, body)
	}
	defer resp.Body.Close()
	out, err := os.Create(dest)
	if err != nil {
		return 0, err
	}
	defer out.Close()
	return io.Copy(out, resp.Body)
}

// DownloadIndex downloads index.html and writes it to the file dest.
// It returns the number of bytes downloaded in case of success.
func (c *Client) DownloadIndex(ctx context.Context, dest string) (int64, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: "/index"}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return 0, err
	}
	resp, err := c.Client.Do(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 200 {
		var body string
		if b, err := ioutil.ReadAll(resp.Body); err != nil {
			if len(b) > 0 {
				body = ": " + string(b)
			}
		}
		return 0, fmt.Errorf("%s%s", resp.Status, body)
	}
	defer resp.Body.Close()
	out, err := os.Create(dest)
	if err != nil {
		return 0, err
	}
	defer out.Close()
	return io.Copy(out, resp.Body)
}

// DownloadSrc downloads /files with the given filename and writes it to the file dest.
// It returns the number of bytes downloaded in case of success.
func (c *Client) DownloadSrc(ctx context.Context, filename, dest string) (int64, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	p := path.Join("/src/", filename)
	u := url.URL{Host: c.Host, Scheme: scheme, Path: p}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return 0, err
	}
	resp, err := c.Client.Do(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 200 {
		var body string
		if b, err := ioutil.ReadAll(resp.Body); err != nil {
			if len(b) > 0 {
				body = ": " + string(b)
			}
		}
		return 0, fmt.Errorf("%s%s", resp.Status, body)
	}
	defer resp.Body.Close()
	out, err := os.Create(dest)
	if err != nil {
		return 0, err
	}
	defer out.Close()
	return io.Copy(out, resp.Body)
}

// DashboardHomePath computes a request path to the dashboard action of home.
func DashboardHomePath() string {
	return fmt.Sprintf("/")
}

// Show basic actions
func (c *Client) DashboardHome(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDashboardHomeRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDashboardHomeRequest create the request corresponding to the dashboard action endpoint of the home resource.
func (c *Client) NewDashboardHomeRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ProfileHomePath computes a request path to the profile action of home.
func ProfileHomePath() string {
	return fmt.Sprintf("/profile")
}

// Set Printer Type
func (c *Client) ProfileHome(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewProfileHomeRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewProfileHomeRequest create the request corresponding to the profile action endpoint of the home resource.
func (c *Client) NewProfileHomeRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
