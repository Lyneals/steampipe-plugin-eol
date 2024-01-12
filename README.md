![image](https://hub.steampipe.io/images/plugins/turbot/cloudflare-social-graphic.png)

# endoflife.date Plugin for Steampipe

Use SQL to query EOL dates for major softwares and middlewares. 

## Quick start

Install the plugin locally (require go):

```shell
git clone git@git.sk5.io:skale-5/run/steampipe/steampipe-plugin-eol.git
cd steampipe-plugin-eol/
mkdir -p ~/.steampipe/plugins/local/steampipe-plugin-eol
go build -o ~/.steampipe/plugins/local/steampipe-plugin-eol/steampipe-plugin-eol.plugin 
cp config/* ~/.steampipe/config
```

Run a query:

```sql
select
    cycle,
    eol,
    days_to_eol
from
    eol_redis;
```

## Engines

This plugin is available for the following engines:

| Engine        | Description
|---------------|------------------------------------------
| [Steampipe](https://steampipe.io/docs) | The Steampipe CLI exposes APIs and services as a high-performance relational database, giving you the ability to write SQL-based queries to explore dynamic data. Mods extend Steampipe's capabilities with dashboards, reports, and controls built with simple HCL. The Steampipe CLI is a turnkey solution that includes its own Postgres database, plugin management, and mod support.

Try it!

```
steampipe query
> .inspect eol
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
