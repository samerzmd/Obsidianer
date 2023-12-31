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

1. Clone the repository:

    ```bash
    git clone https://github.com/samerzmd/obsidianer.git
    ```

2. Navigate to the project directory:

    ```bash
    cd obsidianer
    ```

3. Build the project:

    ```bash
    go build -o obsidianer ./cmd/
    ```

This will generate an executable named `obsidianer`.

## Usage

Run the executable and follow the prompts.

```bash
./obsidianer
```
You will be asked to specify the path of the directory you want to watch. If the directory is not a Git repository, you'll also be prompted to provide an empty Git repository URL.

## Add to PATH (optional)

To make the obsidianer command available globally, add the executable to your PATH.

### For Bash:
Add the following line to your ~/.bashrc or ~/.bash_profile file:

```bash
export PATH=$PATH:/path/to/obsidianer/directory
```
Then apply the changes:
```bash 
source ~/.bashrc  # or source ~/.bash_profile
```
### For Zsh:
Add the following line to your ~/.zshrc file:

```bash
export PATH=$PATH:/path/to/obsidianer/directory
```
Then apply the changes:

```bash
source ~/.zshrc
```
Replace /path/to/obsidianer/directory with the actual directory path where your obsidianer executable resides.

After doing this, you should be able to run obsidianer from any directory in the terminal.

## Tests

Run `go test ./...` from the root directory to run all tests.

## Contributing

Please read the [CONTRIBUTING.md](CONTRIBUTING.md) for details on how to contribute to this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.
