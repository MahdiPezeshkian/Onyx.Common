# ConfigMaster — Configuration Library for Go

## Overview

ConfigMaster is a flexible configuration loader for Go applications. It supports multiple file formats (YAML, TOML, JSON, XML, ENV), environment variable binding, nested key access via dot notation, default values, merging, and type-safe getters with fallback support.

## Quick Start

```go
import "yourmodule/pkg/configmaster"

cfg := configmaster.New()

// Set defaults
cfg.SetDefault("app.port", 8080)
cfg.SetDefault("db.host", "localhost")

// Load config file (format auto-detected from extension)
cfg.LoadFile("config.yaml", "")

// Load environment variables with prefix
cfg.BindEnv("APP_")

// Read values
port := cfg.GetInt("app.port")
host := cfg.GetString("db.host")
```

## Initialization

### `configmaster.New() *Config`

Creates a new empty Config instance.

```go
cfg := configmaster.New()
```

## Loading Configuration

### `LoadFile(path string, format string) error`

Load a file. If format is `""`, it is detected from the file extension.

```go
cfg.LoadFile("config.yaml", "")       // YAML, auto-detected
cfg.LoadFile("config.toml", "")       // TOML, auto-detected
cfg.LoadFile("config.json", "")       // JSON, auto-detected
cfg.LoadFile("myconfig", "yaml")      // explicit format, no extension needed
```

Supported extensions: `.yaml`, `.yml`, `.toml`, `.json`, `.xml`, `.env`

### `LoadReader(r io.Reader, format string) error`

Load from any `io.Reader`.

```go
reader := strings.NewReader(`{"app": {"port": 8080}}`)
cfg.LoadReader(reader, "json")
```

### `LoadFiles(paths ...string) error`

Load multiple files in order. Later files override earlier ones.

```go
cfg.LoadFiles("defaults.yaml", "override.json", "local.toml")
```

## Environment Variables

### `BindEnv(prefix string)`

Loads environment variables with the given prefix. Variables are lowercased, underscores become dots. Does NOT overwrite existing keys.

```go
cfg.BindEnv("APP_")
// APP_DATABASE_HOST → "database.host"
// APP_DB_PORT        → "db.port"
```

To force overwrite, modify the `if c.Get(key) == nil` check inside the function.

## Defaults & Merging

### `SetDefault(key string, value interface{})`

Set a value only if the key does not already exist.

```go
cfg.SetDefault("app.port", 3000)
```

### `Set(key string, value interface{})`

Set or overwrite a value.

```go
cfg.Set("app.name", "MyApp")
```

### `Merge(other *Config)`

Merge another Config into this one. The other Config's values take precedence.

```go
other := configmaster.New()
other.LoadFile("override.yaml", "")
cfg.Merge(other)
```

## Reading Values

All getters support dot notation for nested keys and optional fallback values.

### `Get(key string) interface{}`

Raw value.

```go
val := cfg.Get("database.host")
```

### `GetString(key string, fallback ...string) string`

```go
host := cfg.GetString("database.host", "localhost")
```

### `GetInt(key string, fallback ...int64) int64`

```go
port := cfg.GetInt("app.port", 8080)
```

### `GetBool(key string, fallback ...bool) bool`

```go
debug := cfg.GetBool("app.debug", false)
```

### `GetFloat(key string, fallback ...float64) float64`

```go
rate := cfg.GetFloat("app.rate", 0.5)
```

### `GetDuration(key string, fallback ...time.Duration) time.Duration`

Accepts `"5s"`, `"10m"`, `"1h"` strings or numeric seconds.

```go
timeout := cfg.GetDuration("app.timeout", 10*time.Second)
```

### `GetStringSlice(key string, fallback ...[]string) []string`

```go
tags := cfg.GetStringSlice("app.tags", []string{"default"})
```

### `GetMap(key string) map[string]interface{}`

```go
settings := cfg.GetMap("app.settings")
```

### `All() map[string]interface{}`

Entire config map.

```go
all := cfg.All()
```

## Nested Key Convention

Use dots (`.`) for nested access:

```
"database.connection.pool.size"
```

## Supported File Formats

| Format   | Extensions         | Library                        |
|----------|-------------------|--------------------------------|
| YAML     | `.yaml`, `.yml`   | `gopkg.in/yaml.v3`             |
| TOML     | `.toml`           | `github.com/BurntSushi/toml`   |
| JSON     | `.json`           | `encoding/json` (stdlib)       |
| XML      | `.xml`            | `encoding/xml` (stdlib)        |
| ENV      | `.env`            | custom parser                  |

## Load Order (Recommended)

```go
cfg := configmaster.New()

// 1. Defaults (lowest priority)
cfg.SetDefault("app.port", 8080)

// 2. Base config file
cfg.LoadFile("config.yaml", "")

// 3. Local overrides (git-ignored)
cfg.LoadFile("local.yaml", "")

// 4. Environment variables (highest priority, but does NOT overwrite)
cfg.BindEnv("APP_")
```

## Dependency Installation

```bash
go get gopkg.in/yaml.v3
go get github.com/BurntSushi/toml
```

## Design Notes

- All configuration is stored in `map[string]interface{}` — no struct mapping.
- Thread-safety is NOT built-in. If concurrent access is needed, wrap with `sync.RWMutex`.
- XML parsing is simplistic (flat key-value). For complex XML, consider a dedicated library.
- ENV files support `#` comment lines, blank lines, and `KEY=VALUE` pairs. Values are always strings.
- The `errors` package import was removed as unused. All errors use `fmt.Errorf`.
