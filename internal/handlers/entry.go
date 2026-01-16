package handlers

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/pwarnock/kjbc-mcp/internal/kjbc"
)

func EntryTool() server.ServerTool {
	return server.ServerTool{
		Tool: mcp.Tool{
			Name:        "entry",
			Description: "Get details for a Strong's entry number.",
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]any{
					"entry_number": map[string]any{
						"type":        "string",
						"description": "Strong's entry number to look up.",
					},
					"language": map[string]any{
						"type":        "string",
						"enum":        []string{kjbc.LanguageGreek, kjbc.LanguageHebrew},
						"default":     kjbc.LanguageGreek,
						"description": "Language of the entry (greek or hebrew).",
					},
				},
				Required: []string{"entry_number"},
			},
			Annotations: mcp.ToolAnnotation{
				Title:           "Entry",
				ReadOnlyHint:    mcp.ToBoolPtr(true),
				DestructiveHint: mcp.ToBoolPtr(false),
				IdempotentHint:  mcp.ToBoolPtr(true),
				OpenWorldHint:   mcp.ToBoolPtr(false),
			},
		},
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			number, err := request.RequireString("entry_number")
			if err != nil {
				return errorResult(err), nil
			}

			language := request.GetString("language", kjbc.LanguageGreek)
			result, err := kjbc.GetEntry(language, number)
			if err != nil {
				return errorResult(err), nil
			}

			return okResult(result), nil
		},
	}
}
