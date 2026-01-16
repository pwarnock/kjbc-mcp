package main

import (
	"log"

	"github.com/mark3labs/mcp-go/server"
	"github.com/pwarnock/kjbc-mcp/internal/handlers"
)

const version = "0.1.0"

func main() {
	mcpServer := server.NewMCPServer("kjbc-mcp", version)
	mcpServer.AddTools(
		handlers.WordTool(),
		handlers.EntryTool(),
		handlers.LangsTool(),
	)

	if err := server.ServeStdio(mcpServer); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
