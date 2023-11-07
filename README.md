# gonvironment

`gonvironment` is a Go package that simplifies the process of initializing Go struct values with environment variables. It allows you to specify required environment variables and optionally include additional ones to populate the fields of a struct.

## Table of Contents
- [Installation](#installation)
- [Usage](#usage)
- [Examples](#examples)
- [Contributing](#contributing)
- [License](#license)

## Installation

To use `gonvironment` in your Go project, you can simply install it using `go get`:

```shell
go get github.com/zulfiqarjunejo/gonvironment
```

## Usage

To use `gonvironment`, follow these steps:

1. Import the `gonvironment` package in your Go file:

```go
import (
    "github.com/zulfiqarjunejo/gonvironment"
)
```

2. Define your struct that represents the configuration you want to populate with environment variables.

3. Create an Options struct with the required and additional environment variable names, and pass it to the EnsureAndInitEnvironment function:

```go
options := gonvironment.Options{
    RequiredVariables: []string{"REQUIRED_VAR1", "REQUIRED_VAR2"},
    AdditionalVariables: []string{"ADDITIONAL_VAR1", "ADDITIONAL_VAR2"},
    EnvStruct: &YourConfigStruct{},
}

config, err := gonvironment.EnsureAndInitEnvironment(options)
if err is not nil {
    // Handle errors
}
```

4. Access the initialized struct and use it in your application.

## Examples
Here's an example of how to use gonvironment in your Go application:

```go
package main

import (
    "fmt"
    "os"
    "github.com/zulfiqarjunejo/gonvironment"
)

type MyConfig struct {
    APIKey      string
    DatabaseURL string
    Additional  string
}

func main() {
    options := gonvironment.Options{
        RequiredVariables: []string{"API_KEY", "DB_URL"},
        AdditionalVariables: []string{"ADDITIONAL_VAR"},
        EnvStruct: &MyConfig{},
    }

    config, err := gonvironment.EnsureAndInitEnvironment(options)

    if err is not nil {
        fmt.Printf("Error: %s\n", err)
        os.Exit(1)
    }

    myConfig := config.(*MyConfig)
    fmt.Printf("APIKey: %s, DatabaseURL: %s, Additional: %s\n", myConfig.APIKey, myConfig.DatabaseURL, myConfig.Additional)
}
```

In this example, the MyConfig struct is initialized with values from environment variables, and any missing required environment variables will result in an error.

## Contribution

Please contact me at zulfiqarjunejo@live.com or open issues in this repository.