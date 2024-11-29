# tinycode

This is a web application written in Go that allows users to sign up, write, and view code snippets.

- [Features](#features)
- [Requirements](#requirements)
- [Getting Started](#getting-started)

## Features

- User authentication system
- Create and view code snippets
- Clean and simple UI

## Requirements

- **Go**: [Download Go](https://go.dev/dl)
- **MySQL**: [Download MySQL](https://dev.mysql.com/doc/refman/8.0/en/installing.html)
- **Make**: A build tool (pre-installed on most Linux/macOS systems, can be installed via package managers)

## Getting Started

### 1. Clone the Repository
```sh
git clone https://github.com/alexandru-calin/tinycode
cd tinycode
```

### 2. Run the setup script
It sets up the database and generates a .envrc file.
```sh
chmod +x setup.sh
./setup.sh
```

### 3. Apply Database Migrations
```
make db/migrations/up
```

### 4. Start the application
```
make run/web
```
After starting the application, open your browser and navigate to
http://localhost:3000

You should see the home page of the web application. From here, users can sign up, log in, and interact with the platform.