# DevPod Go Example

[![Open in DevPod!](https://devpod.sh/assets/open-in-devpod.svg)](https://devpod.sh/open#https://github.com/loft-sh/devpod-example-go)

## Quickstart

Either click on the 'Open in DevPod' link above or start via `devpod up github.com/loft-sh/devpod-example-go`. This will start a small devcontainer that has all the needed tools to start working on a go server.
It will also serve static files on port 4000 and forward it to your local machine.

After creating your workspace, run `go run main.go` to start the server.

Then run `go get tailscale.com/tsnet` and `go run tailscale.go`