# Go Environment Variable Loader

This project provides a Go lib to simplify the process of loading environment variables into a struct. The env package contains a function called LoadEnv that takes a struct as input and populates its fields with values from corresponding environment variables.

## Installation

To use this lib, you need to add it to go.mod of your Go project using the following command:

```
go get github.com/sesaquecruz/go-env-loader
```

## Usage

1. Import the env package into your Go file:

```
import "github.com/sesaquecruz/go-env-loader/pkg/env"
```

2. Define a struct with the desired fields and env tags:

```
type Config struct {
    DbHost string `env:"DB_HOST"`
    DbPort int    `env:"DB_PORT"`
    DbName string `env:"DB_NAME"`
    DbUser string `env:"DB_USER"`
    DbPass string `env:"DB_PASS"`
    // Add more fields as needed
}
```

3. Use the LoadEnv function to populate the struct fields with environment variable values:

```
func main() {
    var config Config
    err := env.LoadEnv(&config)
    if err != nil {
        // Handle error loading environment variables
    }

    // Use the config struct with populated values
    // ...
}
```

Ensure that the corresponding environment variables (e.g. DB_HOST, DB_PORT, etc.) are set in your environment before running the application.

## Tag Format

The `env` tag follows the format ``` `env:"<ENV_VAR_NAME>"` ```, where `<ENV_VAR_NAME>` is the name of the environment variable to be loaded into the struct field.

## Error Handling

The LoadEnv function returns an error if an environment variable does not exist or its type is invalid. The supported field types are string and int.

## Contributing

Contributions to this project are welcome. If you encounter any issues or have ideas for enhancements, feel free to open an issue or submit a pull request.

## License
This project is licensed under the MIT License. Please see the [LICENSE](./LICENSE) file for more details.
