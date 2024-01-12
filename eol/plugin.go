package eol

import (
	"context"
	"fmt"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

var eolKeys = [...]string{
	"amazon-eks",
	"amazon-rds-mysql",
	"amazon-rds-postgresql",
	"google-kubernetes-engine",
	"debian",
	"ubuntu",
	"argo-cd",
	"ansible",
	"mariadb",
	"redis",
	"nginx",
	"memcached",
}

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

	for _, key := range eolKeys {
		tableName := fmt.Sprintf("eol_%s", strings.ReplaceAll(key, "-", "_"))
		tables[tableName] = tableGeneric(key)
	}

	plugin.Logger(ctx).Info("eol.PluginTables", "table_builded", tables)

	return tables, nil
}
