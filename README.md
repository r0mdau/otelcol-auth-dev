# Otelcol-dev

My lab for developing a custom authenticator extension in Opentelemetry-Collector.

Official documentation: [Building a custom collector](https://opentelemetry.io/docs/collector/custom-collector/).

## Getting started

Clone the repo and the submodule

```bash
git clone --recurse-submodules git@github.com:r0mdau/otelcol-auth-dev.git
```

### Starting your dev journey

Build & start the custom collector

```bash
make run
```

Send a log line without author

```bash
curl -X POST -H "Content-Type: application/json" -d @testdata/logs.json -i localhost:4318/v1/logs
# response is 401
```

Send a log line with author

```bash
curl -X POST -H "Authorization: romain" -H "Content-Type: application/json" -d @testdata/logs.json -i localhost:4318/v1/logs
# response should be okay if good string passed
```

### Customize the otel modules

A replace instruction is set to use local code in the `go.mod` file: `replace github.com/r0mdau/customauthextension => ./customauthextension`.

Add the module to the good factory map in the `components.go` file.

## Misc

Vscode debug config `.vscode/launch.json`

```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceRoot}",
            "args": ["--config", "${workspaceRoot}/testdata/config.yaml"]
        }
    ]
}
```
