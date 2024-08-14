# Event Booking App - API Documentation

## Overview
This application provides a RESTful API for managing events, user registrations, and user accounts. It uses JWT for authorization and SQLite for database storage.

## Database Schema

### Tables
- **Users**: Stores user information.
- **Events**: Stores event details.
- **Registrations**: Records user registrations for events.

### Data Models
- **Register**: Represents user-event registrations.
- **User**: Represents user accounts with email and password.
- **Event**: Represents events including name, description, location, and date/time.

## Database Setup
- **Initialization**: SQLite is used with the database file named `api.db`. Tables are created if they do not exist.

## Utilities
- **Password Hashing**: Utilizes bcrypt to hash and validate passwords.
- **JWT Utilities**: Includes functions for generating and verifying JWT tokens:

## Routes

### Public Endpoints
- `GET /events`: Lists all events.
- `GET /events/:id`: Retrieves a specific event by ID.
- `GET /users`: Lists all users.
- `GET /registers`: Lists all registrations.
- `POST /signup`: Creates a new user account.
- `POST /login`: Authenticates user and provides a JWT token.

### Protected Endpoints (Requires JWT)
- `POST /events`: Adds a new event.
- `PUT /events/:id`: Updates an event by ID.
- `DELETE /events/:id`: Deletes an event by ID.
- `POST /events/:id/register`: Registers a user for an event by ID.
- `DELETE /events/:id/register`: Cancels a user's registration for an event by ID.

## Testing with Postman

1. **Sign Up User**
   - **Method**: POST
   - **Endpoint**: `/signup`
   - **Body**: JSON with `email` and `password`

2. **Log In User**
   - **Method**: POST
   - **Endpoint**: `/login`
   - **Body**: JSON with `email` and `password`
   - **Response**: You'll receive a JWT token.

3. **Use Token for Protected Routes**
   - Include the JWT token in the Authorization header:
     ```
     Authorization: token <your_jwt_token>
     ```
   - **Test Protected Endpoints** using the token you received from the login step.

Use Postman to send requests to these endpoints and verify the responses.


## Running the Application

1. **Clone the Repository**
 ```bash
 git clone <repository-url>
 ```
2. **Change the Directory**
 ```bash
 cd <project-directory>
```

3. **Install Dependencies**
 ```bash
 go tidy
```

4. **Run App**
 ```bash
 go run .
```
