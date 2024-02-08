# please

>This project is under active development

- [Description](#description)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)

## Description

An http client written in Go.

This project is inspired by [httpie-go](https://github.com/nojima/httpie-go) and [httpie](https://github.com/httpie/cli)
but with some more functionality.

## Requirements
- [Go](https://go.dev/)

## Installation

```bash
# Clone the repo
$ git clone https://github.com/utox39/please-http.git

# cd to the repo
$ cd path/to/please-http

# Build please
$ go build -v ./... 

# Then move it somewhere in your $PATH. Here is an example:
$ mv please ~/bin
```

## Usage

### Make a GET request
```bash
$ please get https://httpbin.org/get
```

### Make a POST request
```bash
$ please post https://httpbin.org/post foo=bar
```

### Make a PUT request
```bash
$ please put https://httpbin.org/put foo=bar
```

### Make a PATCH request
```bash
$ please patch https://httpbin.org/patch foo=bar
```

### Make a DELETE request
```bash
$ please delete https://httpbin.org/delete
```

### Make a HEAD request
```bash
$ please head https://httpbin.org/
```

### Make a OPTIONS request
```bash
$ please options https://httpbin.org/
```

### Repeat a request n times
The --repeat flag will repeat a request n times.


```bash
$ please --repeat=5 post https://httpbin.org/post foo=bar
```

### Create a log file of the http request response
The --log flag will create a log.json file. 

If the --repeat flag is used the --log flag will create a dir named "logs" and all the log.json files
will be created inside it.

```bash
$ please --log post https://httpbin.org/post foo=bar
```
### Generate a response time chart

The --gen-chart will generate a response time chart and must be called with the --repeat flag (--repeat=n, n>= 2).
You can visualize it opening the stats.html file in your browser.

```bash
$ please --repeat=5 --gen-chart post https://httpbin.org/post foo=bar
```

## Contributing

If you would like to contribute to this project just create a pull request which I will try to review as soon as
possible.