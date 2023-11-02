# Obsidianer

A CLI tool to automatically sync files in a directory with a remote Git repository. Designed for use with Obsidian Vaults, but applicable to any folder you want to keep in sync.

## Features

- Watches for file changes in a specified directory and all its sub-directories
- Commits changes to a local Git repository whenever a file is modified
- Pushes commits to a remote repository
- Automatically initializes a Git repository if none exists in the specified directory
- Allows cloning from a remote repository if needed
- User-friendly setup for Git username and email

## Installation

1. Ensure you have Go installed on your machine.
2. Clone this repository.
3. Navigate to the repository and run `go build` to compile the application.
4. Optionally, move the compiled executable to a folder in your PATH for easier access.

## Usage

1. Run the compiled executable.
2. When prompted, enter the path to the directory you wish to watch.

## Tests

Run `go test ./...` from the root directory to run all tests.

## Contributing

Please read the [CONTRIBUTING.md](CONTRIBUTING.md) for details on how to contribute to this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.
