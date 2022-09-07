# envcfg [![Build Status](https://travis-ci.org/tomazk/envcfg.svg?branch=master)](https://travis-ci.org/tomazk/envcfg)

Un-marshaling environment variables to Go structs

## Getting Started

Let's set a bunch of environment variables and then run your go app
```bash
#!/usr/bin/env bash
export DEBUG="false"
export DB_HOST="localhost"
export DB_PORT="8012"

./your_go_app 
```
Within your Go app do
```go
import "github.com/tomazk/envcfg"

// declare a type that will hold your env variables
type Cfg struct {
	DEBUG   bool
	DB_PORT int
	DB_HOST string
}

func main() {
	var config Cfg
	envcfg.Unmarshal(&config)
	// config is now set to Config{DEBUG: false, DB_PORT: 8012, DB_HOST: "localhost"}
	
	// optional: clear env variables listed in the Cfg struct
	envcfg.ClearEnvVars(&config)
}
```
## Installation

```
$ go get github.com/tomazk/envcfg
```

## Motivation

As per **[12 factor app manifesto](http://12factor.net/)** configuration of an app should be stored in the [environment](http://12factor.net/config) since it varies between environments by nature. This convention is also dicated and popularized by emerging technologies like docker and cloud platforms.

Instead of having a bunch of `os.Getenv("ENV_VAR")` buried deep in your code when configuring clients and services, **`envcfg`** encourages you to:

1. **define a struct type** that will hold your environment variables and serve as documentation which env variables must be configured
2. use **`envcfg.Unmarshal`** to read your env variables and unmarhsal them to an object that now holds your configuration of an app
3. use **`envcfg.ClearEnvVars`** to unset env variables, removing potential vulnerability of passing secrets to unsafe child processes or vendor libraries that assume you're not storing unsafe values in the environment

## Documentation

### `envcfg.Unmarshal`

`func Unmarshal(v interface{}) error` can recieve a reference to an object or even a reference to a pointer:

```go
var val2 StructType
envcfg.Unmarshal(&val2)

var val1 *StructType 
envcfg.Unmarshal(&val1) // val1 will be initialized
```

#### Supported Struct Field Types

`envcfg.Unmarshal` supports `int`, `string`, `bool` and `[]int`, `[]string`, `[]bool` types of fields wihin a struct. In addition, fields that satisfy the `encoding.TextUnmarshaler` interface are also supported. `envcfg.Unmarshal` will return nil if a valid struct was passed or return an error if not.

```go
type StructType struct {
	INT           int
	BOOL          bool
	STRING        string
	SLICE_STRING  []string
	SLICE_BOOL    []bool
	SLICE_INT     []int
	CUSTOM_TYPE   MyType
}

type MyType struct{}
func (mt *MyType) UnmarshalText(text []byte) error {
	...
}
```
#### Validation
`envcfg.Unmarshal` also spares you from writing type validation code:

```go
type StructType struct {
	SHOULD_BE_INT int
}
```
If you'll pass `export SHOULD_BE_INT="some_string_value"` to your application `envcfg.Unmarshal` will return an error.

#### Struct Tags for Custom Mapping of env Variables
You can also use struct field tags to map env variables to fields wihin a struct
```bash
export MY_ENV_VAR=1
```
```go
type StructType struct {
	Field int `envcfg:"MY_ENV_VAR"`
}
```
#### Slices Support
`envcfg.Unmarshal` also supports `[]int`, `[]string`, `[]bool` slices. Values of the slice are ordered in respect to env name suffix. See example below.
```bash
export CASSANDRA_HOST_1="192.168.0.20" # *_1 will come as the first element of the slice
export CASSANDRA_HOST_2="192.168.0.21"
export CASSANDRA_HOST_3="192.168.0.22"
```
```go
type StructType struct {
	CASSANDRA_HOST []string
}
func main() {
	var config StructType
	envcfg.Unmarshal(&config)
	// config.CASSANDRA_HOST is now set to []string{"192.168.0.20", "192.168.0.21", "192.168.0.22"} 
}
```
### `envcfg.ClearEnvVars`

`func ClearEnvVars(v interface{}) error` recieves a reference to the same struct you've passed to `envcfg.Unmarshal` and it will unset any environment variables listed in the struct. Except for those that you want to keep and are tagged with `envcfgkeep:""` struct field tag. It will throw an error on unsupported types.

```bash
export SECRET_AWS_KEY="foobar" 
export PORT="8080" 
```
```go
type StructType struct {
	SECRET_AWS_KEY string
	PORT           int    `envcfgkeep:""`
}
func main() {
	var config StructType
	envcfg.ClearEnvVars(&config)
	// it will unset SECRET_AWS_KEY but keep env variable PORT
}
```


## Contributing
Send me a pull request and make sure tests pass on [travis](https://travis-ci.org/tomazk/envcfg/).

## Tests

Package comes with an extensive test suite that's continuously run on travis against go versions: 1.3, 1.4, 1.5, 1.6, 1.7, 1.8 and the development tip.
```
$ go test github.com/tomazk/envcfg
```

## Licence

See LICENCE file


# JSON Config

// Copyright 2012 The Gorilla Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package config provides convenient access methods to configuration stored as
JSON or YAML.

Let's start with a simple YAML example:

    development:
      database:
        host: localhost
      users:
        - name: calvin
          password: yukon
        - name: hobbes
          password: tuna
    production:
      database:
        host: 192.168.1.1

We can parse it using ParseYaml(), which will return a *Config instance on
success:

    cfg, err := config.ParseYaml(yamlString)

An equivalent JSON configuration could be built using ParseJson():

    cfg, err := config.ParseJson(jsonString)

From now, we can retrieve configuration values using a path in dotted notation:

    // "localhost"
    host, err := cfg.String("development.database.host")

    // or...

    // "192.168.1.1"
    host, err := cfg.String("production.database.host")

Besides String(), other types can be fetched directly: Bool(), Float64(),
Int(), Map() and List(). All these methods will return an error if the path
doesn't exist, or the value doesn't match or can't be converted to the
requested type.

A nested configuration can be fetched using Get(). Here we get a new *Config
instance with a subset of the configuration:

    cfg, err := cfg.Get("development")

Then the inner values are fetched relatively to the subset:

    // "localhost"
    host, err := cfg.String("database.host")

For lists, the dotted path must use an index to refer to a specific value.
To retrieve the information from a user stored in the configuration above:

    // map[string]interface{}{ ... }
    user1, err := cfg.Map("development.users.0")
    // map[string]interface{}{ ... }
    user2, err := cfg.Map("development.users.1")

    // or...

    // "calvin"
    name1, err := cfg.String("development.users.0.name")
    // "hobbes"
    name2, err := cfg.String("development.users.1.name")

JSON or YAML strings can be created calling the appropriate Render*()
functions. Here's how we render a configuration like the one used in these
examples:

    cfg := map[string]interface{}{
        "development": map[string]interface{}{
            "database": map[string]interface{}{
                "host": "localhost",
            },
            "users": []interface{}{
                map[string]interface{}{
                    "name":     "calvin",
                    "password": "yukon",
                },
                map[string]interface{}{
                    "name":     "hobbes",
                    "password": "tuna",
                },
            },
        },
        "production": map[string]interface{}{
            "database": map[string]interface{}{
                "host": "192.168.1.1",
            },
        },
    }

    json, err := config.RenderJson(cfg)

    // or...

    yaml, err := config.RenderYaml(cfg)

This results in a configuration string to be stored in a file or database.
*/
package config