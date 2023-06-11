## Task

Write a program, that will check the list of domains for accessibility

Once a minute we need to check, if domains are available and get their ping.
We have many users, who want to know domains ping.
Users have three options:

1. Get the time for specific domain
2. Get the domain with min ping
3. Get the domain with max ping

Both we have administrators, who want to get statistics for this endpoints

## Usage

1. `go mod download`
2. `go run ./cmd/server/main.go`
3. You're done!

## Endpoints

[https://localhost:54](https://localhost:54)

| Endpoint         | Description                                    |
|------------------|------------------------------------------------|
| /stats           | Provides statistics for endpoints usage        |
| /special/:domain | Provides ping, status information about domain |
| /min             | Provides domain with min ping                  |
| /max             | Provides domain with max ping                  |
