# wifi-sensor

A Viam sensor implementation that reads the system's wifi information. This is an example repo to show how to:
1. Make a Viam module with Go
1. Build it in CI and upload to the registry

## Running this with local exec

For iterative development, you can run a module locally (laptop or robot) and test it using a local instance of the RDK.

Create the binary with `make wifi`.

Your config will look something like this (replace /path/to/wifi-sensor with the actual path on your system):

```json
{
  "components": [
    {
      "name": "whatever",
      "model": "viam:sensor:linux-wifi",
      "type": "sensor"
    }
  ],
  "modules": [
    {
      "name": "wifi",
      "executable_path": "/path/to/wifi-sensor/wifi",
      "type": "local"
    }
  ]
}
```

Our docs for running local modules are [here](https://docs.viam.com/extend/modular-resources/configure/#local-modules).

## What's in this repo

- .github/workflows: CI and deploy logic
- Makefile: instructions for building the binary and bundling it into a tarball
- \*.go: the implementation
- meta.json: module configuration file

## Forking this repo

If you fork this and want to deploy to the registry, you'll need to update namespaces and CI secrets. Full fork instructions are in the [Python module example](https://github.com/viam-labs/python-example-module#forking-this-repo).
