# Backend Golang Coding Test

## Part 1: User Management API

### Part 1 Objective

Build a RESTful API in Golang to manage users, using MongoDB for persistence, JWT for authentication, and adhering to clean code practices.

---

### Part 1 Requirements

#### 1. User Model

Define a user entity with the following fields:

- `ID` (auto-generated)
- `Name` (string)
- `Email` (string, unique)
- `Password` (hashed)
- `CreatedAt` (timestamp)

#### 2. Authentication

##### Functions

- Register a new user.
- Authenticate a user and return a JWT token.

##### JWT Implementation

- Protect endpoints with JWT.
- Validate tokens via middleware.
- Sign tokens using HMAC (HS256) with a secret key.

#### 3. User Operations

Implement the following operations:

- Create a new user.
- Fetch a user by ID.
- List all users.
- Update a user's name or email.
- Delete a user.

#### 4. MongoDB Integration

- Use the official Go MongoDB driver.
- Persist and retrieve user data from MongoDB.

#### 5. Middleware

- Implement logging middleware to capture HTTP method, path, and execution time.

#### 6. Concurrency Task

- Run a background goroutine every 10 seconds to log the total number of users in the database.

#### 7. Testing

- Write unit tests using Go's standard `testing` package.
- Mock MongoDB interactions where appropriate.

---

### Bonus (Optional)

- **Containerization**: Add Docker and `docker-compose` support for the API and MongoDB.
- **Abstraction**: Use Go interfaces to abstract MongoDB operations for better testability.
- **Validation**: Implement input validation (e.g., required fields, email format).
- **Graceful Shutdown**: Handle system signals using `context.Context`.
- **gRPC Support**:
  - Define a `.proto` file for `CreateUser` and `GetUser`.
  - Implement a gRPC server (optionally secure with token metadata).
- **Hexagonal Architecture**:
  - Structure the project using ports and adapters.
  - Separate domain, application, and infrastructure layers.
  - Decouple business logic from frameworks and drivers.

---

## Part 2: Lottery Search System

### Part 2 Objective

Design a solution to search through a large dataset of lottery tickets based on specific pattern matching requirements.

### Part 2 Requirements

#### 1. Data Volume

- The system must handle a dataset of **1 million** lottery tickets.
- Each ticket consists of a 6-digit number.

#### 2. Search Pattern

- Allow users to search using a 6-character pattern containing digits and wildcards (`*`).
- **Examples**:
  - `****23` (matches numbers ending in 23)
  - `1****5` (matches numbers starting with 1 and ending with 5)
  - `123***` (matches numbers starting with 123)

#### 3. Result Distribution

- **Constraint**: "Each user should NOT see the same search result when searching with the same pattern."
- Implement a mechanism to distribute available tickets among users searching for the same pattern, ensuring no two users select the same ticket simultaneously.

#### 4. Performance

- Ensure the search is performant when dealing with 1M+ documents.
- Propose an efficient solution for handling these queries.

---

## Submission Guidelines

### Submission: User API

Provide a Git repository containing:

- `README.md` with setup and execution instructions.
- A guide on how to generate and use JWT tokens.
- Sample API requests and responses.
- Documentation of any assumptions or design decisions.

### Submission: Lottery System

- **Design Document**: Explain the approach, data structures, and algorithms used.
- **Code Implementation**: Provide the code for the search logic (can be a separate module or part of the main repo).
- **Performance Analysis**: Briefly explain the efficiency of the solution for large datasets.

---

## Evaluation Criteria

### Evaluation: User API

- Code quality, structure, and readability.
- Correctness and completeness of the REST API.
- Security and implementation of JWT.
- Proper usage and abstraction of MongoDB.
- Test coverage and effective mocking.
- Idiomatic Go usage.
- (Bonus) Implementation of gRPC, Docker, validation, and architecture.

### Evaluation: Lottery System

- **Feasibility**: Does the solution correctly address the problem?
- **Performance**: Is the search efficient for the specified data volume?
- **Correctness**: Does it properly handle the result distribution constraint?
- **Creativity**: Innovative use of data structures or algorithms.