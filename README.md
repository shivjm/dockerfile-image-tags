# dockerfile-image-tags

[![version](https://img.shields.io/github/v/release/shivjm/dockerfile-image-tags?include_prereleases&sort=semver)](https://github.com/shivjm/dockerfile-image-tags/releases)
[![semantic versioning](https://img.shields.io/badge/semantic%20versioning-2.0.0-informational)](https://semver.org/spec/v2.0.0.html)

[![test](https://github.com/shivjm/dockerfile-image-tags/workflows/test/badge.svg)](https://github.com/shivjm/dockerfile-image-tags/actions?query=workflow%3Atest)

List or query images and tags used in a Dockerfile.

## Usage

### List all images and tags

Pass path to Dockerfile:

```sh
dockerfile-image-tags list path/to/Dockerfile
```

Or pass Dockerfile as input:

```sh
cat path/to/Dockerfile | dockerfile-image-tags list
```

Sample output (JSON):

```json
[{"name":"golang","tag":"1.17.0-alpine"},{"name":"common","tag":"?"},{"name":"common","tag":"?"},{"name":"common","tag":"?"},{"name":"viaductoss/ksops","tag":"v3.0.0"},{"name":"quay.io/argoproj/argocd","tag":"$ARGOCD_VERSION"}]
```

### Find single image tag

Use <kbd>query</kbd> to return tag for first occurrence of image with specified
name:

```sh
dockerfile-image-tags query path/to/Dockerfile golang
```

Pass <kbd>-n</kbd> (<kbd>--occurrence</kbd>) to return specified
occurrence instead of first occurrence. For example, to return tag for
second `FROM golang`:

```sh
dockerfile-image-tags query -n 2 path/to/Dockerfile golang
```

Sample output:

```output
1.17.0-alpine
```

## Docker image

See [shivjm/dockerfile-image-tags
packages](https://github.com/shivjm/dockerfile-image-tags/pkgs/container/dockerfile-image-tags/).
