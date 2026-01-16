package handlers

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/pwarnock/kjbc-mcp/internal/kjbc"
)

func WordTool() server.ServerTool {
	return server.ServerTool{
		Tool: mcp.Tool{
			Name:        "word",
			Description: "Find Strong's entries for an English word.",
			InputSchema: mcp.ToolInputSchema{
				Type: "object",
				Properties: map[string]any{
					"word": map[string]any{
						"type":        "string",
						"description": "English word to search for.",
					},
					"language": map[string]any{
						"type":        "string",
						"enum":        []string{kjbc.LanguageGreek, kjbc.LanguageHebrew},
						"default":     kjbc.LanguageGreek,
						"description": "Language to search in (greek or hebrew).",
					},
					"show_verses": map[string]any{
						"type":        "boolean",
						"default":     false,
						"description": "Include verse references for each entry.",
					},
				},
				Required: []string{"word"},
			},
			Annotations: mcp.ToolAnnotation{
				Title:           "Word",
				ReadOnlyHint:    mcp.ToBoolPtr(true),
				DestructiveHint: mcp.ToBoolPtr(false),
				IdempotentHint:  mcp.ToBoolPtr(true),
				OpenWorldHint:   mcp.ToBoolPtr(false),
			},
		},
		Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			word, err := request.RequireString("word")
			if err != nil {
				return errorResult(err), nil
			}

			language := request.GetString("language", kjbc.LanguageGreek)
			showVerses := request.GetBool("show_verses", false)

			result, err := kjbc.SearchWord(language, word, showVerses)
			if err != nil {
				return errorResult(err), nil
			}

			return okResult(result), nil
		},
	}
}
