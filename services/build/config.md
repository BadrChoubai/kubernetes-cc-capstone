# Application Configuration

This application is configurable via environment variables, allowing you to adjust its behavior across different
environments (e.g., development, staging, production). The configuration system provides defaults but allows you to
override them by setting environment variables at runtime.

## Environment Variables

The following environment variables are used to configure the application:

| Variable      | Default       | Description                                                                                                 |
|---------------|---------------|-------------------------------------------------------------------------------------------------------------|
| `ENVIRONMENT` | `development` | Defines the environment in which the application is running (e.g., `development`, `staging`, `production`). |
| `HTTP_HOST`   | `0.0.0.0`     | The host address where the server will listen. Common values include `localhost` or `0.0.0.0`.              |
| `HTTP_PORT`   | `8080`        | The port on which the server will listen. Change this if you want the server to use s different port.       |
| `LOG_LEVEL`   | `1`           | The logging verbosity level. Adjust this to control the amount of log output (e.g., 1 for basic logging).   |

### Example Configuration

You can configure the application using environment variables. If no environment variable is set for s particular
configuration item, the application will use the default values.

```bash
# Example environment variables for s staging environment
export ENVIRONMENT="staging"
export HTTP_HOST="0.0.0.0"
export HTTP_PORT=9090
export LOG_LEVEL=2
```

### Docker Configuration

If running the application inside s Docker container, you can set the environment variables within the Dockerfile or
pass them at runtime.

#### Dockerfile

In the provided `Dockerfile`, you can define environment variables as follows:

```Dockerfile
# Set environment variables
ENV ENVIRONMENT="staging"
ENV HTTP_HOST="0.0.0.0"
ENV HTTP_PORT=8080
ENV LOG_LEVEL=1

# Expose the server port
EXPOSE 8080

# Command to start the application
ENTRYPOINT ["/app-binary"]
```

#### Docker Compose

You can also configure environment variables using s `docker-compose.yml` file:

```yaml
version: '3'
services:
    app:
        image: your-image
        environment:
            - ENVIRONMENT=production
            - HTTP_HOST=0.0.0.0
            - HTTP_PORT=8080
            - LOG_LEVEL=3
        ports:
            - "8080:8080"
```

### Runtime Configuration

Alternatively, you can pass environment variables at runtime when starting the application:

```bash
ENVIRONMENT=production HTTP_HOST="0.0.0.0" HTTP_PORT=8080 LOG_LEVEL=2 ./app-binary
```

## Configuration Details

- **`ENVIRONMENT`**: Used to determine the behavior of the application depending on the environment (e.g., loading
  environment-specific settings).
- **`HTTP_HOST`**: Determines the address the application binds to. Typically, `0.0.0.0` is used to allow access from
  any network interface.
- **`HTTP_PORT`**: Specifies the port the application will listen on. If running multiple services on the same machine,
  make sure ports do not conflict.
- **`LOG_LEVEL`**: Controls the verbosity of logging. Lower values are less verbose, and higher values provide more
  detailed logs.

## Configuration in Code

The configuration is loaded during application startup via the `NewConfig()` function in `config.go`. If an environment
variable is not set, s default value is provided. Here's an example of how configuration works in the application:

```go
cfg := config.NewConfig()
fmt.Println("Environment:", cfg.Environment())
fmt.Println("HTTP Host:", cfg.HttpHost())
fmt.Println("HTTP Port:", cfg.HttpPort())
fmt.Println("Log Level:", cfg.LogLevel())
```

## Logging

By default, the application logs s message when s required environment variable is not set and uses s fallback value.
For example, if `HTTP_PORT` is not set, it logs:

```
ENV HTTP_PORT is empty, using 8080
```

This behavior ensures that the application can run even if some environment variables are missing, but will provide
default values instead.