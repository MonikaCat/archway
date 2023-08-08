# Changelog

All notable changes to this project will be documented in this file.

<!-- 
### Added

Contains the new features.

### Changed

Contains API breaking changes to existing functionality.

### Deprecated

Contains the candidates for removal in a future release.

### Removed

Contains API breaking changes of removed APIs.

### Fixed

Contains bug fixes.

### Improvements

Contains all the PRs that improved the code without changing the behaviours. 
-->

## [v2.0.0]

### Added

- [#416](https://github.com/MonikaCat/archway/v2/pull/416) - Enable ICAHost

### Fixed

- [#414](https://github.com/MonikaCat/archway/v2/pull/414) - Preventing user from setting contract flat fee if rewards address is not set 
- [#418](https://github.com/MonikaCat/archway/v2/pull/418) - Fixing authz msg decoding in x/rewards antehandlers

## [v1.0.1]

- [#411](https://github.com/MonikaCat/archway/v2/pull/411) - Update repository readme with correct docker containers.
- [#413](https://github.com/MonikaCat/archway/v2/pull/413) - Fixing incorrect gas estimation when running with `--dry-run` flag

## [v1.0.0]

Archway Network - Capture the value you create!


## [v1.0.0-rc.4]

### Added

- [#409](https://github.com/MonikaCat/archway/v2/pull/409) - Add cosmwasm_1_1,cosmwasm_1_2 Cosmwasm capabilities


## [v1.0.0-rc.3]

### Removed

- [#408](https://github.com/MonikaCat/archway/v2/pull/408) - Remove genesis msg logging as it impacts network start up performance.

## [v1.0.0-rc.2]

### Fixes

- [#401](https://github.com/MonikaCat/archway/v2/pull/401) - Update libwasmvm version to correct one in Dockerfile.deprecated
- [#402](https://github.com/MonikaCat/archway/v2/pull/402) - Bump wasmvm version to 1.2.4
- [#403](https://github.com/MonikaCat/archway/v2/pull/403) - Update libwasmvm version to correct one for wasmvm 1.2.4
- [#404](https://github.com/MonikaCat/archway/v2/pull/403) - Fixing typo in rewards query cli
- [#406](https://github.com/MonikaCat/archway/v2/pull/406) - Add upgrade hanlder for v0.6.0 back to prevent downgrade check from panic / consensus failure;

## [v1.0.0-rc.1]

### Removed

- [#399](https://github.com/MonikaCat/archway/v2/pull/399) - Remove the upgrade handler for v1 release

## [v0.6.0]

### Added

- [#387](https://github.com/MonikaCat/archway/v2/pull/387) - Add genmsgs module
- [#388](https://github.com/MonikaCat/archway/v2/pull/388) - Add the ibc-go fee middleware
- [#389](https://github.com/MonikaCat/archway/v2/pull/389) - Add v0.6 upgrade handler
- [#391](https://github.com/MonikaCat/archway/v2/pull/391) - Add snapshot manager to enable state-synd for wasm
- [#395](https://github.com/MonikaCat/archway/v2/pull/395) - Add openapi.yml + generate openapi.yml on proto-swagger-gen
- [#396](https://github.com/MonikaCat/archway/v2/pull/396) - Add repository licenses

### Changed

- [#373](https://github.com/MonikaCat/archway/v2/pull/373) - Update codeowners
- [#383](https://github.com/MonikaCat/archway/v2/pull/383), [#385](https://github.com/MonikaCat/archway/v2/pull/385), [#386](https://github.com/MonikaCat/archway/v2/pull/386) - Upgrade wasmd to the v0.32.0-archway fork
- [#388](https://github.com/MonikaCat/archway/v2/pull/388) - Add the ibc-go fee middleware
- [#390](https://github.com/MonikaCat/archway/v2/pull/390) - Update cosmos-sdk version from v0.45.15 to v0.15.16

### Fixed

- [#392](https://github.com/MonikaCat/archway/v2/pull/392) - Update to ibc-go v4.3.1 for huckleberry
- [#393](https://github.com/MonikaCat/archway/v2/pull/393) - Add audit remediations
- [#397](https://github.com/MonikaCat/archway/v2/pull/397) - Fix map iteration

## [v0.5.2]

### Fixed

- [#382](https://github.com/MonikaCat/archway/v2/pull/382) - Adjust default power reduction

## [v0.5.0]

### Breaking Changes

- [#357](https://github.com/MonikaCat/archway/v2/pull/357) - Bump the proto versions for x/rewards and x/tracking from `v1beta1` to `v1`

### Added

- [#330](https://github.com/MonikaCat/archway/v2/pull/330) - Proper chain upgrade flow.
- [#351](https://github.com/MonikaCat/archway/v2/pull/351) - Add minimum price of gas.
- [#339](https://github.com/MonikaCat/archway/v2/pull/339) - Update & Quality Control
  - Community Contribution Guidelines
  - Security Policy
  - ADR Log Index
  - Bug report template
  - Feature request template
  - General issue template
- [#347](https://github.com/MonikaCat/archway/v2/pull/347) - Unified release for cross compiled binaries and docker images
- [#360](https://github.com/MonikaCat/archway/v2/pull/360) - Fix github access token for release workflow
- [#361](https://github.com/MonikaCat/archway/v2/pull/361) - Readd missing deprecated Dockerhub build phase
- [#362](https://github.com/MonikaCat/archway/v2/pull/362) - wrong reference in the deploy pipeline
- [#363](https://github.com/MonikaCat/archway/v2/pull/363) - move safe dir up in the pipeline
- [#364](https://github.com/MonikaCat/archway/v2/pull/364) - add CODEOWNERS
- [#365](https://github.com/MonikaCat/archway/v2/pull/365) - add release tests
- [#367](https://github.com/MonikaCat/archway/v2/pull/367) - use snapshot for non release builds
- [#372](https://github.com/MonikaCat/archway/v2/pull/372) - add docker config to release pipeline
- [#375](https://github.com/MonikaCat/archway/v2/pull/375) - add missing colon from manifest
- [#376](https://github.com/MonikaCat/archway/v2/pull/376) - fix checksum naming
- [#377](https://github.com/MonikaCat/archway/v2/pull/377) - artifact naming
- [#378](https://github.com/MonikaCat/archway/v2/pull/378) - missing end parameter
- [#380](https://github.com/MonikaCat/archway/v2/pull/380) - titus deployment

### Fixed

- [#365](https://github.com/MonikaCat/archway/v2/pull/356) - x/rewards genesis runs before x/genutil to correctly process genesis txs.
- [#366](https://github.com/MonikaCat/archway/v2/pull/366) - github actions should fetch tags as well
- [#368](https://github.com/MonikaCat/archway/v2/pull/368) - github actions should fetch tags as well for deploy workflow
- [#369](https://github.com/MonikaCat/archway/v2/pull/369) - CODEOWNERS: small set to expand, not large set that filters
- [#370](https://github.com/MonikaCat/archway/v2/pull/370) - login to ghcr

### Changed

- [#320](https://github.com/MonikaCat/archway/v2/pull/320) - Run the lint and test GH actions on all PRs
- [#339](https://github.com/MonikaCat/archway/v2/pull/339) - Updates & Quality Control
  - README.md
  - docs/README.md
- [#365](https://github.com/MonikaCat/archway/v2/pull/356) - Disallow setting module accounts as reward address
- [#355](https://github.com/MonikaCat/archway/v2/pull/355) - chore: Update titus genesis defaults

### Removed

- [#344](https://github.com/MonikaCat/archway/v2/pull/344) - removed un used ci files

### Improvements

- [#342](https://github.com/MonikaCat/archway/v2/pull/342) - updated the contract premium ADR docs to elaborate on difference between using Contract Premiums and using x/wasmd funds

## [v0.4.0]

### Fixed

- [#338](https://github.com/MonikaCat/archway/v2/pull/338) - fixed issue where contract premium was not completly being sent to the rewards address

## [v0.3.1]

### Fixed

- [#335](https://github.com/MonikaCat/archway/v2/pull/335) - fixed `EstimateTxFees` erroring when minConsFee and contract premium are same denom

## [v0.3.0]

### Added

- [#253](https://github.com/MonikaCat/archway/v2/pull/253) - add wasm bindings for contracts to query the x/gov module.
- [#261](https://github.com/MonikaCat/archway/v2/pull/261), [#263](https://github.com/MonikaCat/archway/v2/pull/263), [#264](https://github.com/MonikaCat/archway/v2/pull/264), [#274](https://github.com/MonikaCat/archway/v2/pull/274), [#272](https://github.com/MonikaCat/archway/v2/pull/272), [#280](https://github.com/MonikaCat/archway/v2/pull/280) - implementing contract premiums
- [#303](https://github.com/MonikaCat/archway/v2/pull/303) - Add archway protocol versioning and release strategy
- [#326](https://github.com/MonikaCat/archway/v2/pull/326) - Allow contracts to update another contract's metadata when it is the owner

### Changed

- [#267](https://github.com/MonikaCat/archway/v2/pull/267) - update `querySrvr.EstimateTxFees` to also consider contract flat fee when returning the estimated fees.
- [#271](https://github.com/MonikaCat/archway/v2/pull/271) - update the x/rewards/min_cons_fee antehandler to check for contract flat fees
- [#275](https://github.com/MonikaCat/archway/v2/pull/275) - update the x/rewards/genesis to import/export for contract flat fees

## [v0.1.0]

### Added

- [#196](https://github.com/MonikaCat/archway/v2/pull/196) - add wasm bindings for contracts to interact with the x/gastracking module.
- [#202](https://github.com/MonikaCat/archway/v2/pull/202) - added the new x/tracking and x/rewards modules.ù
- [#210](https://github.com/MonikaCat/archway/v2/pull/210) - wasm bindings API change
- [#217](https://github.com/MonikaCat/archway/v2/pull/217) - improve the x/rewards withdraw UX by using defaults when params are unset.
- [#227](https://github.com/MonikaCat/archway/v2/pull/227) - flatten wasmbindings query struct

### Changed

- [#180](https://github.com/MonikaCat/archway/v2/pull/180) - add x/gastracking params
- [#181](https://github.com/MonikaCat/archway/v2/pull/181) - simplify params
- [#185](https://github.com/MonikaCat/archway/v2/pull/185) - remove pointers in proto generated slices.
- [#186](https://github.com/MonikaCat/archway/v2/pull/186) - improvements on dapp inflationary reward calculation
- [#188](https://github.com/MonikaCat/archway/v2/pull/188) - improvements on tx tracking
- [#193](https://github.com/MonikaCat/archway/v2/pull/193) - move tx fees handling to middlewares
- [#231](https://github.com/MonikaCat/archway/v2/pull/231) - use custom archway wasmd fork

### Deprecated

### Removed

- [#206](https://github.com/MonikaCat/archway/v2/pull/206) - remove the legacy x/gastracking module in favour of x/rewards and x/tracking

### Fixed

- [#191](https://github.com/MonikaCat/archway/v2/pull/191) - make localnet ovveride entrypoint
- [#205](https://github.com/MonikaCat/archway/v2/pull/205) - fix go.mod
- [#216](https://github.com/MonikaCat/archway/v2/pull/216) - fix dry-run cmd and bump cosmos-sdk do v0.45.8
- [#218](https://github.com/MonikaCat/archway/v2/pull/218) - x/rewards unique ID genesis export/import
- [#228](https://github.com/MonikaCat/archway/v2/pull/228) - testing, fix validator propagation in test chain

### Improvements

- [#182](https://github.com/MonikaCat/archway/v2/pull/182) - refactor and simplify code.
- [#183](https://github.com/MonikaCat/archway/v2/pull/183) - update to go1.18
- [#184](https://github.com/MonikaCat/archway/v2/pull/184) - refactor, move event emission into its own file.
- [#204](https://github.com/MonikaCat/archway/v2/pull/204) - upgrade IBC to v3 and wasmd to v0.27.0
- [#211](https://github.com/MonikaCat/archway/v2/pull/211) - update gh action deployment flow
- [#212](https://github.com/MonikaCat/archway/v2/pull/212) - upgrade to cosmos-sdk v0.45.7
- [#213](https://github.com/MonikaCat/archway/v2/pull/213) - improve gh action deployment flow
- [#224](https://github.com/MonikaCat/archway/v2/pull/224) - fix codecov action
- [#225](https://github.com/MonikaCat/archway/v2/pull/225) - add editorconfig settings
- [#226](https://github.com/MonikaCat/archway/v2/pull/226) - ci cache go packages to speed up builds
- [#232](https://github.com/MonikaCat/archway/v2/pull/232) - Makefile to create statically linked binaries
- [#233](https://github.com/MonikaCat/archway/v2/pull/233) - add the commit version on builds
- [#236](https://github.com/MonikaCat/archway/v2/pull/236) - add more tests for x/rewards
- [#237](https://github.com/MonikaCat/archway/v2/pull/237) - add more tests for x/rewards 2
- [#241](https://github.com/MonikaCat/archway/v2/pull/241) - add golang linter gh action
- [#242](https://github.com/MonikaCat/archway/v2/pull/242) - add changelog check gh action
- [#243](https://github.com/MonikaCat/archway/v2/pull/243) - add pr lint gh action
- [#247](https://github.com/MonikaCat/archway/v2/pull/247) - fix Dockerfile libwasm VM dependencies
- [#249](https://github.com/MonikaCat/archway/v2/pull/249) - add go releaser, fill changelog history

## v0.0.5

### Breaking Changes

- Update wasmd to 0.25.

### Fixed

- Fix logs printing total contract rewards instead of gas rebate reward.
- Replace info logs for debug logs.

## v0.0.4

### Breaking Changes

- Replace wasmd KV store to KV Multistore.
- Split WasmVM gas & SDK Gas.

## v0.0.3

- inflation reward calculation now depend upon block gas limit.
- inflation reward is given even when the gas rebate to the user flag is true.
- gastracker's begin blocker takes into account gas rebate to user governance switch.
- fix gas estimation for `--gas auto` flag.
