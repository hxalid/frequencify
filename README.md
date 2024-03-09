# Wiki Crawler
A cli tool to scrape Wikipedia API and print top N words on a given page.

### Dependencies

* [Cobra](https://github.com/spf13/cobra) - a Go cli library
* [Testify](https://github.com/stretchr/testify) - A toolkit with common assertions and mocks that plays nicely with Go standard library.
* [Package errors](https://github.com/pkg/errors) - Package errors provides simple error handling primitives.
* [Go 1.12](https://golang.org/doc/go1.12) - required due to Go modules.

**Note:**
* This application could have been written without depending on any external libraries. Especially using Cobra may sound like an overkill for two parameters.
* It has been tested by using [Go 1.12](https://golang.org/doc/go1.12) on macOS Mojave and Linux Mint 18.
* It should work on [Go 1.11](https://blog.golang.org/go1.11) as it supports Go modules as well.
* If you don't have Go 1.11+ please use the Docker build

## How to run
The project includes a Makefile with the following rules:
```
all                            Build and run tests
clean                          Remove previous build
docker                         Containerise the application
help                           Display available commands
```
To build it for your desired architecture only, please specify the architecture name to the make command. For instance to build it only for MacOS run the following command:
```bash
make darwin
```
The following steps are needed in order to run the application from a Docker container:
* `make docker` or `docker build -t frequencify .`
* `docker run -it frequencify:latest -n <number> -p <pageid>`
## How to test
The default Makefile target runs unit tests after successful build.

To exacute the command on MacOS run the following command:
```
./frequencify_darwin_amd64 -n 5 -p 21721040
```
Here `-n` indicates the top n words and `-p` indicates the page id. The application provides `-h` help option as well. 