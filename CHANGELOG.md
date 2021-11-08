# Changelog

## [2.1.0](https://www.github.com/shivjm/dockerfile-image-tags/compare/v2.0.0...v2.1.0) (2021-11-08)


### Features

* add version command ([5b0f84b](https://www.github.com/shivjm/dockerfile-image-tags/commit/5b0f84bf08c39f64868729ba914b2a0315ad5a9c))

## [2.0.0](https://www.github.com/shivjm/dockerfile-image-tags/compare/v1.3.2...v2.0.0) (2021-11-08)


### ⚠ BREAKING CHANGES

* create separate `list` and `query` sub-commands
* **input:** allow passing `-` as filename for STDIN

### Features

* create separate `list` and `query` sub-commands ([52e2059](https://www.github.com/shivjm/dockerfile-image-tags/commit/52e205902e93bbf6149971dd7d5bf7d9190dc460))
* **input:** allow passing `-` as filename for STDIN ([58e6c28](https://www.github.com/shivjm/dockerfile-image-tags/commit/58e6c289b9494bc6f5674518c7a9f59ff9bc7194))

### [1.3.2](https://www.github.com/shivjm/dockerfile-image-tags/compare/v1.3.1...v1.3.2) (2021-11-07)


### Bug Fixes

* force CI ([8f83fcf](https://www.github.com/shivjm/dockerfile-image-tags/commit/8f83fcff1d242563c52675e2d8a1f3a969254ee3))

### [1.3.1](https://www.github.com/shivjm/dockerfile-image-tags/compare/v1.3.0...v1.3.1) (2021-11-07)


### Bug Fixes

* fix typo in command description ([4a1840d](https://www.github.com/shivjm/dockerfile-image-tags/commit/4a1840d59efd8f107431abeb890ec7b55c72c8f7))

## [1.3.0](https://www.github.com/shivjm/dockerfile-image-tags/compare/v1.2.0...v1.3.0) (2021-11-07)


### Features

* add Dockerfile ([1686dd8](https://www.github.com/shivjm/dockerfile-image-tags/commit/1686dd8b12173dd9ac46a189af8ef3fcd4a38140))

## [1.2.0](https://www.github.com/shivjm/dockerfile-image-tags/compare/v1.1.0...v1.2.0) (2021-11-07)


### Features

* add CLI flag to specify which occurrence to return ([f1fe02c](https://www.github.com/shivjm/dockerfile-image-tags/commit/f1fe02ccc04f2dcf8f1f4ecd92fccd77c1bdd1f0))
* allow `getSingleTag` to return later occurrences ([c1ecb37](https://www.github.com/shivjm/dockerfile-image-tags/commit/c1ecb379cb64160b506eed5251feb3fe280f8298))

## [1.1.0](https://www.github.com/shivjm/dockerfile-image-tags/compare/v1.0.0...v1.1.0) (2021-11-07)


### Features

* allow querying single tag ([47d2db8](https://www.github.com/shivjm/dockerfile-image-tags/commit/47d2db8818b1bf20d91ed62a3d72b6976d042e8a))

## 1.0.0 (2021-11-07)


### ⚠ BREAKING CHANGES

* allow specify custom marker for unknown tags

### Features

* allow specify custom marker for unknown tags ([2e874eb](https://www.github.com/shivjm/dockerfile-image-tags/commit/2e874eb487c0308d8ae71e72c0a3cef141bbd0be))
* allow specifying unknown marker in CLI ([13afe26](https://www.github.com/shivjm/dockerfile-image-tags/commit/13afe2632c708af6bcff0e970c8418de49384ebc))
* implement parsing Dockerfile image versions ([bd48f8d](https://www.github.com/shivjm/dockerfile-image-tags/commit/bd48f8dec3859f5c8f56b9dbd94fc2d2fa941e41))
