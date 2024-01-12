package eol

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

type Eol struct {
	Cycle             string      `json:"cycle",omitempty`
	ReleaseDate       string      `json:"releaseDate",omitempty`
	Latest            string      `json:"latest",omitempty`
	LatestReleaseDate string      `json:"latestReleaseDate",omitempty`
	Lts               bool        `json:"lts,omitempty"`
	Eol               interface{} `json:"eol",omitempty`
	DaysToEol         int
}

func tableGeneric(tech string) *plugin.Table {
	return &plugin.Table{
		Name:             fmt.Sprintf("eol_%s", strings.ReplaceAll(tech, "_", "-")),
		Description:      fmt.Sprintf("Retrieve %s EOL information using endoflife.date", tech),
		DefaultTransform: transform.FromGo(),
		List: &plugin.ListConfig{
			Hydrate: listGeneric(tech),
		},
		Columns: []*plugin.Column{
			{Name: "cycle", Type: proto.ColumnType_STRING, Description: "Major version"},
			{Name: "release_date", Type: proto.ColumnType_STRING, Description: "Major version release date"},
			{Name: "eol", Type: proto.ColumnType_STRING, Description: "End of life date"},
			{Name: "latest", Type: proto.ColumnType_STRING, Description: "Latest patch release"},
			{Name: "latest_release_date", Type: proto.ColumnType_STRING, Description: "Latest patch release date"},
			{Name: "lts", Type: proto.ColumnType_BOOL, Description: "Is it an LTS release ?"},
			{Name: "days_to_eol", Type: proto.ColumnType_INT, Description: "Days remaining before EOL"},
		},
	}
}

func listGeneric(tech string) func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
		resp, err := http.Get(fmt.Sprintf("https://endoflife.date/api/%s.json", tech))

		now := time.Now()

		if err != nil {
			plugin.Logger(ctx).Error("eol.listGeneric", tech, err)
			return nil, err
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			plugin.Logger(ctx).Error("eol.listGeneric", tech, resp.StatusCode)
			return nil, err
		}

		var versions []Eol
		err = json.NewDecoder(resp.Body).Decode(&versions)

		if err != nil {
			plugin.Logger(ctx).Error("eol.listGeneric", tech, err)
			return nil, err
		}

		for _, v := range versions {
			// For latest release, EOL might be a boolean
			if reflect.TypeOf(v.Eol).Kind() == reflect.Bool {
				v.Eol = now.AddDate(1, 0, 0).Format("2006-01-02")
			}
			eol := v.Eol.(string)
			time, _ := time.Parse("2006-01-02", eol)
			v.DaysToEol = int(time.Sub(now).Hours() / 24)
			plugin.Logger(ctx).Debug("eol.listGeneric", tech, v)
			d.StreamListItem(ctx, v)
		}
		return nil, nil
	}
}
