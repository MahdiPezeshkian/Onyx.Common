# Onyx.Common

[![Go Version](https://img.shields.io/badge/go-1.27-00ADD8.svg)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/MahdiPezeshkian/Onyx.Common.svg)](https://pkg.go.dev/github.com/MahdiPezeshkian/Onyx.Common)
[![Go Report Card](https://goreportcard.com/badge/github.com/MahdiPezeshkian/Onyx.Common)](https://goreportcard.com/report/github.com/MahdiPezeshkian/Onyx.Common)

A collection of reusable Go utility packages: configuration loading, error
stacks, REST response wrappers, validation helpers and ISP detection.

## Installation

```bash
go get github.com/MahdiPezeshkian/Onyx.Common@v0.1.0
```

The module requires Go 1.27 or later.

## Packages

| Package | Import path | Description |
|---|---|---|
| `configmaster` | `github.com/MahdiPezeshkian/Onyx.Common/configmaster` | Multi-format configuration loader (YAML, TOML, JSON, XML, ENV) |
| `errstack` | `github.com/MahdiPezeshkian/Onyx.Common/errstack` | Error creation and wrapping with stack traces |
| `responswrapper` | `github.com/MahdiPezeshkian/Onyx.Common/responswrapper` | Generic REST API response wrappers |
| `validation` | `github.com/MahdiPezeshkian/Onyx.Common/validation` | Reflection-based validation helpers |
| `isp` | `github.com/MahdiPezeshkian/Onyx.Common/isp` | Internet service provider detection for a client IP |

## Usage

### configmaster — configuration loading

```go
package main

import (
	"fmt"
	"time"

	"github.com/MahdiPezeshkian/Onyx.Common/configmaster"
)

func main() {
	cfg := configmaster.New()

	// 1. Defaults (lowest priority)
	cfg.SetDefault("app.port", 8080)
	cfg.SetDefault("db.host", "localhost")

	// 2. Load a config file (format auto-detected from the extension)
	if err := cfg.LoadFile("config.yaml", ""); err != nil {
		panic(err)
	}

	// 3. Environment variables with a prefix (highest priority)
	cfg.BindEnv("APP_")

	port := cfg.GetInt("app.port")
	host := cfg.GetString("db.host", "127.0.0.1")
	timeout := cfg.GetDuration("app.timeout", 10*time.Second)

	fmt.Println(port, host, timeout)
}
```

### errstack — errors with stack traces

```go
import "github.com/MahdiPezeshkian/Onyx.Common/errstack"

func load(id string) error {
	err := fetchFromDB(id)
	return errstack.Wrapf(err, "failed to load resource %s", id)
}
```

### responswrapper — REST API envelopes

```go
import "github.com/MahdiPezeshkian/Onyx.Common/responswrapper"

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getUsers(page, size, total int) *responswrapper.PagedRestApiResponse[User] {
	users := []User{{ID: 1, Name: "Ada"}}
	return responswrapper.NewPagedRestApiResponseV2(users, page, size, total, 200, "ok")
}

func notFound() *responswrapper.RestApiResponse[any] {
	return responswrapper.SetRestApiError("user not found")
}
```

### validation — input checks

```go
import "github.com/MahdiPezeshkian/Onyx.Common/validation"

func validateSignup(email, password string, age int) error {
	if err := validation.ValidateNotEmpty(email, password); err != nil {
		return err
	}
	if err := validation.ValidateEmail(email); err != nil {
		return err
	}
	return validation.ValidateRange(18, 120, age)
}
```

### isp — client ISP detection

```go
import "github.com/MahdiPezeshkian/Onyx.Common/isp"

func main() {
	detector := isp.NewIPInfoDetector(nil) // nil => default client, 3s timeout
	org := detector.GetISP("8.8.8.8")      // e.g. "AS15169 Google LLC"
	_ = org
}
```

## Versioning

This module follows [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

While the module is in `v0.x`, the public API may change between minor
versions. To stay stable, pin an explicit version in your `go.mod`:

```
require github.com/MahdiPezeshkian/Onyx.Common v0.1.0
```

or install a pinned version explicitly:

```bash
go get github.com/MahdiPezeshkian/Onyx.Common@v0.1.0
```

Never rely on `@latest` for a `v0.x` module: the newest version may contain
breaking changes. Once the module reaches `v1.0.0`, the API will be covered by
SemVer compatibility guarantees.

## Contributing

1. Fork the repository and create a feature branch from `main`.
2. Write tests for any change and make sure `go build ./...`,
   `go vet ./...` and `go test ./...` pass.
3. Use [Conventional Commits](https://www.conventionalcommits.org/) for commit
   messages.
4. Open a pull request describing the motivation and the change.

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for
details.
