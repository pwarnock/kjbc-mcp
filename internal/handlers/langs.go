package handlers

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/pwarnock/kjbc-mcp/internal/kjbc"
)

func LangsTool() server.ServerTool {
	return server.ServerTool{
		Tool: mcp.Tool{
			Name:        "langs",
			Description: "List supported languages.",
			InputSchema: mcp.ToolInputSchema{
				Type:       "object",
				Properties: map[string]any{},
			},
			Annotations: mcp.ToolAnnotation{
				Title:           "Langs",
				ReadOnlyHint:    mcp.ToBoolPtr(true),
				DestructiveHint: mcp.ToBoolPtr(false),
				IdempotentHint:  mcp.ToBoolPtr(true),
				OpenWorldHint:   mcp.ToBoolPtr(false),
			},
		},
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return okResult(map[string][]string{"languages": kjbc.Languages()}), nil
		},
	}
}
