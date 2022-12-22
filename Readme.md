# HCP CLI

[![Project Status: WIP – Initial development is in progress, but there has not yet been a stable, usable release suitable for the public.](https://www.repostatus.org/badges/latest/wip.svg)](https://www.repostatus.org/#wip)

CLI tooling for Hashicorp Cloud Platform.

Motivated by the web UI continuously logging me out, and wanting to change the iteration on a channel for a given bucket in HCP Packer Registry.

## Usage

*(NB: Unchecked things aren't yet working.)*

Set `HCP_ORGANISATION_ID` and `HCP_PROJECT_ID` in your shell to make the following work. It'll pop open a browser for you to auth the first time it's called.

Build tools with `go build ./...` then find the named binaries in the top level folder.

 `hcp-packer`

  - [x] `list-buckets` - list all buckets in the registry
  - [ ] `list-iterations --bucket ID|NAME` - list all iterations for a bucket
  - [ ] `list-channels --bucket ID|NAME` - list all channels for a bucket
  - [ ] `set-channel-iteration --bucket ID|NAME --channel ID|NAME --iteration ID` - set a channel to a specific iteration
