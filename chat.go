package moonshot

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

const (
	apiV1ChatCompletions = "/v1/chat/completions"
)

type Message struct {
	Role    string `json:"role,omitempty"`
	Content string `json:"content,omitempty"`
}

type RequestCompletionChat struct {
	Model       string    `json:"model,omitempty"`
	Messages    []Message `json:"messages,omitempty"`
	Temperature float64   `json:"temperature,omitempty"`
}

type ChatCompletionChoice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type ResponseCompletionChat struct {
	ID      string                 `json:"id,omitempty"`
	Object  string                 `json:"object,omitempty"`
	Created int64                  `json:"created,omitempty"`
	Model   string                 `json:"model,omitempty"`
	Choices []ChatCompletionChoice `json:"choices,omitempty"`
	Usage   Usage                  `json:"usage,omitempty"`
}

func (c *Client) CreateChatCompletions(ctx context.Context, req *RequestCompletionChat) (*ResponseCompletionChat, error) {

	httpReq, err := c.newChatCompletionsHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result ResponseCompletionChat
	if err := decodeResponse(resp.Body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) newChatCompletionsHTTPRequest(ctx context.Context, req *RequestCompletionChat) (*http.Request, error) {

	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, moonshotBaseURL+apiV1ChatCompletions, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+c.apiKey)

	return httpReq, nil
}
