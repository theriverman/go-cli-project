# go-cli-app-template
Go CLI Application Template

# Introduction
This template implies the following conditions:
  * the project is version controlled using Git
  * the project uses semantic versioning
  * the project uses Git tags to create releases
  * git tags equal to the released semantic version (e.g.: 2.4.0)

All built binaries are put to the [dist](./dist) folder

## Build/Host System Requirements
  * sh/Bash (use Git Bash for Windows)
  * GNU Make (available for Windows too)
  * Go compiler utils are available in `$PATH`

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
Update the respective values under `StringFileInfo`. Optionally update the values of `FixedFileInfo.FileVersion` too.

## Configure the values in `.github/workflows/makefile-build-release.yml`
Update the following fields in [makefile-build-release.yml](.github/workflows/makefile-build-release.yml):
  * Set up Go / go-version
  * GH Release Artefacts / files (all binary files must be listed manually)

The `GH Release Artefacts / files` array shall contain all binary outputs which has to be uploaded to the GitHub releases page.

# Project Structure
See below an explanation for the key files of a typical Go project.
```
$ tree ./go-cli-app-template
.
├── LICENSE           # license for using this template. override for your project
├── Makefile          # build targets
├── README.md         # instructions for using this template. override for your project
├── cli.app.go        # the main CLI application. extend for your project
├── cli.commands.go   # additional commands underneath the main CLI application. extend for your project
├── dist              # folder containing your binaries
│   └── MY-APP.exe    # your typical binary output(s)
├── go.mod            # should be generated using go mod init ...
├── go.sum            # should be generated using go mod init ...
├── main.go           # the entrypoint to your Go application. typically no changes needed here
└── versioninfo.json  # executable file details (Windows-only)
```

# Acknowledgements
  * github.com/josephspurrier/goversioninfo
