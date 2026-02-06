package main

import (
	"context"
	"log"

	sdk "github.com/apito-io/go-apito-plugin-sdk"
)

func main() {
	log.Printf("🎯 [hc-search-plugin] Starting Search plugin...")
	plugin := sdk.Init("hc-search-plugin", "1.0.0", "apito-plugin-key")
	statusType := sdk.NewObjectType("SearchStatus", "Search plugin status").
		AddStringField("status", "Plugin status", false).
		AddStringField("version", "Plugin version", false).
		Build()
	plugin.RegisterQuery("getSearchStatus",
		sdk.ComplexObjectField("Get search plugin status", statusType),
		func(ctx context.Context, rawArgs map[string]interface{}) (interface{}, error) {
			return map[string]interface{}{"status": "ready", "version": "1.0.0"}, nil
		})
	log.Printf("🚀 [hc-search-plugin] Plugin ready")
	plugin.Serve()
}
