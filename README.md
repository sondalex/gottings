# gottings ![gottings](assets/logo.svg)

[![Go Reference](https://pkg.go.dev/badge/github.com/sondalex/gottings?status.svg)](https://pkg.go.dev/github.com/sondalex/gottings?tab=doc)

**gottings** is a Go library for loading configuration data from environment variables and JSON files.

## Usage

## Installation

```bash
go get github.com/sondalex/gottings
```

### Declare Your Configuration

Define your configuration structure by specifying fields with corresponding tags for JSON and environment variables:

```go
import "github.com/sondalex/gottings"


type Config struct {
    Host string `json:"host" env:"APP_HOST"`
    Port int    `json:"port" env:"APP_PORT"`
}
```

### Load Configuration from Environment Variables

Set environment variables

```bash
export APP_PORT=1312
export APP_HOST="127.0.0.1"
```

Initialize your configuration from environment variables:

```go
func NewConfig() (*Config, error) {
    config := &Config{}
    err := gottings.LoadEnv(config)
    if err != nil {
        return nil, nil
    }
    return config, nil
}
```

### Load Configuration from JSON and Environment Variables

You can load configuration from a JSON file, with environment variables taking precedence:

```go
func NewConfig() (*Config, error) {
    config := &Config{}
    filePath := "info.json"
    data, err := os.ReadFile(filePath)
    if err != nil {
        return nil, err
    }
    err = gottings.LoadConfiguration(data, config)
    if err != nil {
        return nil, err
    }
    return config, nil
}
```

### Unsupported Types

If the type associated to the environment value you are trying to unmarshal is unsupported, implement the `UnmarshalEnvironmentValue` interface:

```go
func (s *DataStructure) UnmarshalEnvironmentValue(data []byte) error {
    // Custom parsing logic here...
    return nil
}
```

### Supported Types

gottings supports the following types:

- `int`, `int8`, `int16`, `int32`, `int64`
- `float32`, `float64`
- `string`
- `bool`
- `gottings.NullInt`, `gottings.NullInt8`, `gottings.NullInt16`, `gottings.NullInt32`, `gottings.NullInt64`
- `gottings.NullFloat32`, `gottings.NullFloat64`
- `gottings.NullBool`
- `gottings.NullString`

### Nullable Fields

gottings also supports pointer fields, allowing for nullable configuration values:

```go
type Config struct {
    Host *string `env:"APP_HOST"`
    Port *int    `env:"APP_PORT"`
}
```

Usage example:

```go
func NewConfig() (*Config, error) {
    config := &Config{}
    err := gottings.LoadEnv(config)
    if err != nil {
        return nil, err
    }

    if config.Port == nil {
        return nil, errors.New("PORT not initialized")
    }
    return config, nil
}
```

Alternatively, use `Null*` types:

```go
type Config struct {
    Host gottings.NullString `env:"APP_HOST"`
    Port gottings.NullInt    `env:"APP_PORT"`
}


func NewConfig() (*Config, error) {
    config := &Config{}
    err := gottings.LoadEnv(config)
    if err != nil {
        return nil, err
    }

    if !config.Port.Valid {
        return nil, errors.New("PORT not initialized")
    }
    return config, nil 
}
```

### Mix configuration initialization between CLI flags environment variable and JSON

You may want to prepopulate your configuration file with CLI flags values.
Example with [flags std library package](https://pkg.go.dev/flag)

```go
import (
	"flag"
	"fmt"

	"github.com/sondalex/gottings"
)

type Config struct {
	Host string `json:"host" env:"APP_HOST"`
	Port int    `json:"port" env:"APP_PORT"`
}

var port int

var flags map[string]interface{} = map[string]interface{}{
	"Port": &port,
}

func init() {
	flag.IntVar(flags["Port"].(*int), "port", 1312, "Port to be used by the Server")
}

func NewConfig() (*Config, error) {
	config := Config{}
	err := gottings.LoadOptions(flags, &config)
	if err != nil {
		return nil, nil
	}
	err = gottings.LoadConfiguration([]byte(`{"host": "127.0.0.1"}`), &config)
	if err != nil {
		return nil, nil
	}
	return &config, nil
}

func main() {
    config, err := NewConfig()
    if err != nil {
        panic(err)
    }
    fmt.Println(config)
    // &{127.0.0.1 1312}
}

```
