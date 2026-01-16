package handlers

import (
	"encoding/json"

	"github.com/mark3labs/mcp-go/mcp"
)

func okResult(payload any) *mcp.CallToolResult {
	data, err := json.Marshal(payload)
	if err != nil {
		return errorResult(err)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: string(data),
			},
		},
		StructuredContent: payload,
	}
}

func errorResult(err error) *mcp.CallToolResult {
	payload := map[string]string{"error": err.Error()}
	data, _ := json.Marshal(payload)

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			mcp.TextContent{
				Type: "text",
				Text: string(data),
			},
		},
		StructuredContent: payload,
		IsError:           true,
	}
}
