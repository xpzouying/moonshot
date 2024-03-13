package moonshot

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

const (
	apiFiles = "/v1/files"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

type FileDetail struct {
	ID           string `json:"id,omitempty"`
	Object       string `json:"object,omitempty"`
	Bytes        int    `json:"bytes,omitempty"`
	CreatedAt    int64  `json:"created_at,omitempty"`
	Filename     string `json:"filename,omitempty"`
	Purpose      string `json:"purpose,omitempty"`
	Status       string `json:"status,omitempty"`
	StatusDetail string `json:"status_details,omitempty"`
}

func (c *Client) GetFileInfo(ctx context.Context, fid string) (*FileDetail, error) {

	req, err := c.newGetFileRequest(ctx, fid)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result FileDetail

	if err := decodeResponse(resp.Body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) newGetFileRequest(ctx context.Context, fid string) (*http.Request, error) {

	targetURL := moonshotBaseURL + apiFiles + "/" + fid

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, targetURL, nil)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Authorization", "Bearer "+c.apiKey)

	httpReq = httpReq.WithContext(ctx)

	return httpReq, nil
}

func (c *Client) ListFiles(ctx context.Context) ([]FileDetail, error) {

	httpReq, err := c.newListFilesRequest(ctx)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Object string       `json:"object"` // list
		Data   []FileDetail `json:"data"`
	}

	if err := decodeResponse(resp.Body, &result); err != nil {
		return nil, err
	}
	return result.Data, nil
}

func (c *Client) newListFilesRequest(ctx context.Context) (*http.Request, error) {

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, moonshotBaseURL+apiFiles, nil)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+c.apiKey)

	return httpReq, nil
}

func (c *Client) UploadFile(ctx context.Context, path string) (*FileDetail, error) {

	req, err := c.newUploadFileRequest(ctx, path)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result FileDetail
	if err := decodeResponse(resp.Body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) newUploadFileRequest(ctx context.Context, path string) (*http.Request, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var requestBody bytes.Buffer

	multipartWriter := multipart.NewWriter(&requestBody)
	fileWriter, err := multipartWriter.CreateFormFile("file", filepath.Base(file.Name()))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		return nil, err
	}
	_ = multipartWriter.WriteField("purpose", "file-extract")
	multipartWriter.Close()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, moonshotBaseURL+apiFiles, &requestBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	return req, nil
}

func (c *Client) DeleteFile(ctx context.Context, fid string) error {
	req, err := c.newDeleteFilesRequest(ctx, fid)
	if err != nil {
		return err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.Errorf("delete file request status code: %v", resp.StatusCode)
	}

	return nil
}

func (c *Client) newDeleteFilesRequest(ctx context.Context, fid string) (*http.Request, error) {

	targetURL := moonshotBaseURL + apiFiles + "/" + fid

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodDelete, targetURL, nil)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Authorization", "Bearer "+c.apiKey)

	return httpReq, nil
}

type FileContent struct {
	Content  string `json:"content"`
	FileType string `json:"file_type"`
	Filename string `json:"filename"`
	Title    string `json:"title"`
	Type     string `json:"type"`
}

func (c *Client) GetFileContent(ctx context.Context, fid string) (*FileContent, error) {
	req, err := c.newGetFileContentRequest(ctx, fid)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result FileContent
	if err := decodeResponse(resp.Body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) newGetFileContentRequest(ctx context.Context, fid string) (*http.Request, error) {

	targetURL := moonshotBaseURL + apiFiles + "/" + fid + "/content"

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, targetURL, nil)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Authorization", "Bearer "+c.apiKey)

	return httpReq, nil
}
