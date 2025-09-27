# Task Manager API
A robust, secure, and scalable task management backend built with Go, Gin, and PostgreSQL, following clean architecture principles.

## Features
### User Authentication & Authorization
1. Secure registration and login with email validation and bcrypt password hashing
2. JWT-based authentication and Gin middleware for protected endpoint
   
### Task Management (CRUD)
1. Create, view, update, and soft-delete tasks
2. Filter tasks by status and ensure users can only manage their own tasks

### Database & Migrations
1. PostgreSQL schema with users and tasks tables, including soft-delete support
2. Managed with migration files for reproducible database state

### API Design
1. RESTful endpoints: /api/v1/auth/register, /api/v1/auth/login, /api/v1/tasks, etc.
2. Consistent error handling with meaningful JSON responses

### Development & Deployment
1. Docker Compose for local development (API + database)
2. Makefile for common tasks (run, test, migrate)
3. Unit and integration tests for business logic and API endpoints

## Architecture
1. Clean separation of domain, service, repository, and API layers for maintainability and scalability
2. Custom error handling and validation throughout
