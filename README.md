# Introduction

This is a web server in Go language for generating random values ​​and getting them by ID.

Implemented methods:
- `POST /api/generate/` — generate a random value and its ID

- `GET /api/retrieve/` — getting the value by id that was returned in the generate method

# Prerequisites

- You need to have [Go](https://golang.org/dl/) installed on your computer. The version used in this project is **1.21.4**.

# Setup

1. Clone this repository:

```bash
git clone https://github.com/gerstudent/generate-random && cd generate-random
```

2. Build docker image:

```bash
docker build -t generate-random .
```

3. Run docker image:

```bash
docker run -p 8081:8081 generate-random
```

4. Make a request to the server and get generated value:

```bash
curl -X POST http://localhost:8081/api/generate/
curl http://localhost:8081/api/retrieve/<id>
```