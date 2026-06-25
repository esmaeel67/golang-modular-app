# Go Modular Event-Driven App

A modular, event-driven application built with Go following best practices and clean architecture.

## 🚀 Features

- Modular architecture with clear separation of concerns
- Event-driven communication between modules
- Structured logging with configurable levels
- Environment-based configuration using `stackus/dotenv`
- Dependency injection ready
- Comprehensive testing strategy

## 📋 Prerequisites

- Go 1.21 or higher
- Git
- (Optional) Docker for containerized development

## 🔧 Installation

```bash
# Clone the repository
git clone git@github.com:esmaeel67/golang-modular-app.git
cd golang-modular-app

# Install dependencies
go mod download

# Install dotenv package
go get github.com/stackus/dotenv

# Copy environment variables file
cp .env.example .env

# Build the application
go build -o bin/app cmd/app/main.go