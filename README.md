# dockerfile-image-tags

[![version](https://img.shields.io/github/v/release/shivjm/dockerfile-image-tags?include_prereleases&sort=semver)](https://github.com/shivjm/dockerfile-image-tags/releases)
[![semantic versioning](https://img.shields.io/badge/semantic%20versioning-2.0.0-informational)](https://semver.org/spec/v2.0.0.html)

[![test](https://github.com/shivjm/dockerfile-image-tags/workflows/test/badge.svg)](https://github.com/shivjm/dockerfile-image-tags/actions?query=workflow%3Atest+branch%3Amain)

A tool to list images and tags used in a Dockerfile.

## Usage

Pass path to Dockerfile:

```sh
dockerfile-image-tags path/to/Dockerfile
```

Or pass Dockerfile as input:

```sh
cat path/to/Dockerfile | dockerfile-image-tags
```

The output is a JSON array listing all images and tags found.
