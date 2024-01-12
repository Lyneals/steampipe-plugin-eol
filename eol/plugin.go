package eol

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             "steampipe-plugin-eol",
		DefaultTransform: transform.FromGo(),
		TableMapFunc:     PluginTables,
	}
	return p
}

func PluginTables(ctx context.Context, d *plugin.TableMapData) (map[string]*plugin.Table, error) {
	// Initialize tables
	tables := map[string]*plugin.Table{}

	// Request https://endoflife.date/api/all.json to get all supported products
	resp, err := http.Get("https://endoflife.date/api/all.json")

	if err != nil {
		plugin.Logger(ctx).Error("eol.PluginTables", "all_request", err)
		return nil, err
	}

	var products []string
	err = json.NewDecoder(resp.Body).Decode(&products)

	if err != nil {
		plugin.Logger(ctx).Error("eol.PluginTables", "all_decode", err)
		return nil, err
	}

	for _, key := range products {
		tableName := fmt.Sprintf("eol_%s", strings.ReplaceAll(key, "-", "_"))
		tables[tableName] = tableGeneric(key)
	}

	plugin.Logger(ctx).Info("eol.PluginTables", "table_builded", tables)

	return tables, nil
}
