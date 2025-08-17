# AWS Cognito with Go

This repository implements **end-to-end authentication** using **AWS Cognito** and **Go**.

---

## Directory Structure

- **config/**  
  Loads environment variables and injects them into the application.

- **models/**  
  Contains all the data structures and models representing entities and request/response payloads. For example, user sign up and sign in data.

- **utils/**  
  Stateless helper functions used throughout the application. These can include token hashing or any initializing clients to interact with AWS services.

- **services/**  
  Implements the core business logic and interfaces with AWS Cognito SDK. This layer performs operations like user registration, authentication, password reset flows, and token management.

- **middleware/**  
  Enables endpoint protection by implementing token verification.

- **handlers/**  
  Responsible for handling HTTP requests and responses. These functions parse incoming requests, call the relevant services, and return JSON responses with appropriate status codes.

- **routes/**  
  Defines the HTTP routes/endpoints and maps them to corresponding handlers. This keeps routing organized and separated from business logic.

- **main.go**  
  The application entry point. Sets up the Gin router, middleware, and starts the server.

---

## Features Status

| Feature                  | Status         | Notes                               |
| ------------------------ | -------------- | ----------------------------------- |
| Sign Up                  | ✅ Done        | User registration works             |
| Login / Token Generation | ✅ Done        | JWT tokens issued on login          |
| Confirm Sign Up          | ✅ Done        | To confirm account and enable login |
| Resend Confirmation Code | ✅ Done        | To resend signup codes              |
| Token Verification       | ✅ Done        | Middleware to protect routes        |
| Forgot Password          | ✅ Done        | Trigger password reset code         |
| Confirm Forgot Password  | ✅ Done        | Verify reset code & new password    |
| Refresh Token Handling   | ✅ Done        | Keep user sessions alive            |
| Logout                   | ⏳ In Progress | Invalidate tokens                   |
