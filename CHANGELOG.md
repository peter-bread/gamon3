## [1.1.1](https://github.com/peter-bread/gamon3/compare/v1.1.0...v1.1.1) (2025-09-19)

### Documentation

* **README:** dummy commit for dummy release to fix homebrew tap ([87523b2](https://github.com/peter-bread/gamon3/commit/87523b2dc9c54b9caea558a7a9de4998ec8d71a4))

## [1.1.0](https://github.com/peter-bread/gamon3/compare/v1.0.10...v1.1.0) (2025-09-18)

### Features

* **cmd:** add source cmd ([#19](https://github.com/peter-bread/gamon3/issues/19)) ([b06752e](https://github.com/peter-bread/gamon3/commit/b06752ef1741a9ce5f800f30da0d924bfc2b666f))
* **source:** adds source command ([939e67b](https://github.com/peter-bread/gamon3/commit/939e67bbd9d35eb4ad3690a87561a4151a3c6d5c)), closes [#17](https://github.com/peter-bread/gamon3/issues/17)

### Documentation

* **README:** adds note about other shells ([861871e](https://github.com/peter-bread/gamon3/commit/861871e83a5b771336fe950bf63ebdefb11c4c49)), closes [#12](https://github.com/peter-bread/gamon3/issues/12)
* **README:** updates installation script instructions ([7093b43](https://github.com/peter-bread/gamon3/commit/7093b432500e545cc1d05ab2bd03150bece1f824))

### Other

* **run:** split filepath and account resolution from runCmd ([95d3312](https://github.com/peter-bread/gamon3/commit/95d3312f092e7df836f63afa2101969fb583121c))

## [1.0.10](https://github.com/peter-bread/gamon3/compare/v1.0.9...v1.0.10) (2025-09-13)

### Bug Fixes

* **config:** don't error if `accounts` field is empty ([2431839](https://github.com/peter-bread/gamon3/commit/2431839794df6fa9eb0d0fc933eb96dcb46633ed))

### Documentation

* **README:** adds documentation for install script options ([28d486a](https://github.com/peter-bread/gamon3/commit/28d486a8a2789ff115e6da855eb73a63ac621655))
* **README:** adds installation docs for install script ([e6382fe](https://github.com/peter-bread/gamon3/commit/e6382feb3fbc95935e5ed8985e75718a3f0a83db))
* **README:** adds links to install script ([b391355](https://github.com/peter-bread/gamon3/commit/b391355d78fce105610f821f699457f6461958da))
* **README:** adds OS requirements ([d12afe6](https://github.com/peter-bread/gamon3/commit/d12afe677b5c84b1b7a47b848689891ac7bb4b67))

## [1.0.9](https://github.com/peter-bread/gamon3/compare/v1.0.8...v1.0.9) (2025-09-12)

### Documentation

* **README:** adds link to homebrew tap ([9b4bdf7](https://github.com/peter-bread/gamon3/commit/9b4bdf75b26859d8448bbde8c4c235d192311895))

## [1.0.8](https://github.com/peter-bread/gamon3/compare/v1.0.7...v1.0.8) (2025-09-11)

### Documentation

* **README:** remove TODO ([aa54266](https://github.com/peter-bread/gamon3/commit/aa5426663b052ada8236931dd675311eba0b7ec0))

## [1.0.7](https://github.com/peter-bread/gamon3/compare/v1.0.6...v1.0.7) (2025-09-11)

### Documentation

* **README:** adds homebrew installation docs ([7c22741](https://github.com/peter-bread/gamon3/commit/7c2274106479952d1e5b8a83d2771d0e181f2299))

## [1.0.6](https://github.com/peter-bread/gamon3/compare/v1.0.5...v1.0.6) (2025-09-11)

### Documentation

* **README:** adds installation instructions for `go install` ([a4f07e6](https://github.com/peter-bread/gamon3/commit/a4f07e6594443b51f31c0f1f6e858eae4b29c0a6))
* **README:** start homebrew installation docs ([e3723ba](https://github.com/peter-bread/gamon3/commit/e3723baacb3edb15f4d53488d76958d09e51c228))

## [1.0.5](https://github.com/peter-bread/gamon3/compare/v1.0.4...v1.0.5) (2025-09-09)

### Bug Fixes

* **hook:** clearer description ([9fc7618](https://github.com/peter-bread/gamon3/commit/9fc7618e155a25170f274737f72084fcbffa9459))

## [1.0.4](https://github.com/peter-bread/gamon3/compare/v1.0.3...v1.0.4) (2025-09-09)

### Documentation

* **README:** make TODO bold ([480bf9d](https://github.com/peter-bread/gamon3/commit/480bf9d485d0a4b289686871d5c889709562e59b))
* **README:** remove italics from TODO ([e6280e0](https://github.com/peter-bread/gamon3/commit/e6280e03f6acb4154616b0ae6f6900cf1e3ea251))

## [1.0.3](https://github.com/peter-bread/gamon3/compare/v1.0.2...v1.0.3) (2025-09-09)

### Documentation

* **README:** specify go 1.25 or later ([c9c69cf](https://github.com/peter-bread/gamon3/commit/c9c69cfcbb0b3a53f7e70c6bc4d6e195f942e0dc))

## [1.0.2](https://github.com/peter-bread/gamon3/compare/v1.0.1...v1.0.2) (2025-09-08)

### Documentation

* **README:** update build from source instructions ([739751f](https://github.com/peter-bread/gamon3/commit/739751fa3947a5b9de8dc28fa84035d350fc12cf))

## [1.0.1](https://github.com/peter-bread/gamon3/compare/v1.0.0...v1.0.1) (2025-09-08)

### Bug Fixes

* **root:** correct `--version` output ([abda223](https://github.com/peter-bread/gamon3/commit/abda223bb47375ac9070250add9ca394f37d16a8))

## 1.0.0 (2025-09-08)

### ⚠ BREAKING CHANGES

* remove map command and json mapping
* use cobra
* **paths:** use actual config and state filepaths

### Features

* better error messages ([316c773](https://github.com/peter-bread/gamon3/commit/316c77373624102876e8c4c6238f125efb322920))
* compare with current account in gh/hosts.yml ([953d11c](https://github.com/peter-bread/gamon3/commit/953d11c55b979b483d08edd6db6798a56176ee88))
* **config:** allow .yaml and .yml for config file ([98724f7](https://github.com/peter-bread/gamon3/commit/98724f7360519d5c50f0124566216b61468058ef))
* **config:** better validation ([eaf2a8d](https://github.com/peter-bread/gamon3/commit/eaf2a8d185707cf402b51c0c48b924761529c301))
* **config:** config can now contain `~` ([ff4a150](https://github.com/peter-bread/gamon3/commit/ff4a150b9ac2fb6847887af7874126ecb4fe6722)), closes [#5](https://github.com/peter-bread/gamon3/issues/5)
* **help:** adds config section to help ([2168a5b](https://github.com/peter-bread/gamon3/commit/2168a5b7047958ecac41c9f072fa78e8b4273e6d))
* **hook:** adds hook for fish shell ([75dd5f7](https://github.com/peter-bread/gamon3/commit/75dd5f779f0eca9d85d36cf70a2388f8d318b0fd))
* parse args ([804fb9c](https://github.com/peter-bread/gamon3/commit/804fb9c3dae75f4b39d7d8f421951aba9f6a3bf9))
* **paths:** use actual config and state filepaths ([667f3e2](https://github.com/peter-bread/gamon3/commit/667f3e23df1062b408faa81fd50d39d4acf3b45e))
* read yaml config and create json mapping ([3703f56](https://github.com/peter-bread/gamon3/commit/3703f563ccc2df0ad8882b004740b78b1955b8cb))
* remove map command and json mapping ([39689a1](https://github.com/peter-bread/gamon3/commit/39689a140527f57590f14d705dc623a4ed0099e7)), closes [#1](https://github.com/peter-bread/gamon3/issues/1)
* **root:** update version format ([afa6501](https://github.com/peter-bread/gamon3/commit/afa650160440b67d36cd9c76c850f010d045334f))
* **run:** better docs ([df60262](https://github.com/peter-bread/gamon3/commit/df602624c4028627041d5d1f5fa38c8a4a025c07))
* **run:** better error handling and validation ([01dbe99](https://github.com/peter-bread/gamon3/commit/01dbe9957730d225c5ab05a59444a22241edd489))
* **run:** check for .yaml and .yml local cfg files ([91f27f9](https://github.com/peter-bread/gamon3/commit/91f27f95d9dbce3e6cf0e4554008577a51b070ed))
* **run:** check GAMON3_ACCOUNT env var ([b098dab](https://github.com/peter-bread/gamon3/commit/b098dab4047d85ffdb9ec35021ca2b2febfc02c2))
* **run:** make it clear errors come from gamon3 ([10ae6f0](https://github.com/peter-bread/gamon3/commit/10ae6f0ce5c131204dee7ae9af056d04fc2c79ce)), closes [#4](https://github.com/peter-bread/gamon3/issues/4)
* **run:** walk up file tree to check for .gamon.yaml ([d4e0a18](https://github.com/peter-bread/gamon3/commit/d4e0a18787d999f180a3d808410bcdce324943b1))
* use cobra ([a191e5b](https://github.com/peter-bread/gamon3/commit/a191e5b953f3f180b51de0c6036e11b06f2b36c8))
* version flag + goreleaser config started ([1c48378](https://github.com/peter-bread/gamon3/commit/1c48378ecdfda672fc87ea748e97cc6b36d7d75c))

### Bug Fixes

* **run:** exit if fails to get config path ([dbf15e2](https://github.com/peter-bread/gamon3/commit/dbf15e213e53ed4d141a8bea97f2f59c7c5899c2))
* **run:** stop after finding first local config file ([d289dc1](https://github.com/peter-bread/gamon3/commit/d289dc14d1a2740fae77f53b4add0ac097dcaef7))
