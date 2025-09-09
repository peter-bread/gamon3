## [1.0.3](https://github.com/peter-bread/gamon3/compare/v1.0.2...v1.0.3) (2025-09-09)

### Documentation

* **README:** specify go 1.25 or later ([c9c69cf](https://github.com/peter-bread/gamon3/commit/c9c69cfcbb0b3a53f7e70c6bc4d6e195f942e0dc))

## [1.0.2](https://github.com/peter-bread/gamon3/compare/v1.0.1...v1.0.2) (2025-09-08)

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
