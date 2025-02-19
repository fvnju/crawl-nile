# Crawl Nile

Crawl Nile is a Go-based application designed to interact with the Nile University SIS (Student Information System). It allows users to log in, scrape course information, and log out programmatically.

## Features

- Retrieve session tokens for authentication.
- Log in to the Nile SIS using provided credentials.
- Scrape course information including course code, name, grade, and credit.
- Log out from the SIS securely.

## Prerequisites

- Go 1.23.2 or later
- Internet connection

## Installation

1. Clone the repository:

```bash
git clone https://github.com/fvnju/crawl-nile.git
```

2. Navigate to the project directory:

```bash
cd crawl-nile
```

3. Install dependencies:

```bash
go mod tidy
```

## Usage

Run the application in CLI mode with your username and password:

```bash
go run main.go -c <username> <password>
```

Optional flags:

- `-t`: Measure execution time.
