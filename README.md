# go-cli-app-template
Go CLI Application Template

# Introduction
This template implies the following conditions:
  * the project is version controlled using Git
  * the project used semantic versioning
  * the project uses Git tags to create releases
  * the git tag equals to the released semantic version (e.g.: 2.4.0)
  * all built binaries are put to the [dist](./dist) folder

## Build/Host System Requirements
  * Sh/Bash (use Git Bash for Windows)
  * GNU Make (available for Windows too)
  * Go compiler utils available in `$PATH`

# Initialise The Project
Execute the steps described in these subsections to initialise the project using this template.
## Create a `go.mod` file:
```bash
go mod init example.com/m       # to initialize a v0 or v1 module
go mod init example.com/m/v2    # to initialize a v2 module
```
## Configure the values in `Makefile`
  * Set the value of `BUILD_TYPE`. This is used to distinguish between build types (optional)
  * Set the value of `BINARY_NAME`. This is used when generating the binary output

## Configure the values in `versioninfo.json`
Update the respective values under `StringFileInfo`. Optionally update the values of `FixedFileInfo.FileVersion`.

## Configure the values in `.github/workflows/makefile-build-release.yml`
Update the following fields in [makefile-build-release.yml](.github/workflows/makefile-build-release.yml):
  * Set up Go / go-version
  * GH Release Artefacts / files (all binary files must be listed manually)

The `GH Release Artefacts / files` array shall contain all binary outputs which has to be uploaded to the GitHub releases page.

# Acknowledgements
  * github.com/josephspurrier/goversioninfo
