# GitHub OAuth Application Documentation

This documentation provides an overview of the GitHub OAuth-based application, detailing the backend and frontend aspects. This application enables users to connect to their GitHub account via OAuth, view their six most recent repositories (or fewer), display an "About" section, log out, and access a default page if not logged in using `/octocat`.

## Backend

### Technology Stack
- Language: Golang 1.20
- Router: Router-chi for routing
- Caching: Router-chi for caching
- Idempotency: Router-chi for ensuring idempotency
- GitHub API Integration: Go-based wrappers for efficient communication with GitHub
- Repository Pattern: Utilized for organized data access

### Installation and Setup
1. Clone the repository.
2. Install Golang 1.20.
3. Run the application using `go run cmd/app.go`.

### Endpoints
- `/login`: Initiates the OAuth flow to connect to a GitHub account.
- `/repo`: Fetches and displays the user's six most recent repositories.
- `/user`: Displays user-specific "About" section.
- `/logout`: Logs the user out.
- `/callback`: Handles code exchange against the code from the front-callback layer.

### Repository Pattern
- Data access and storage are organized using the repository pattern for maintainability and separation of concerns.

### Caching and Idempotency
- Router-chi is used to cache responses for improved performance.
- Idempotency is enforced to ensure the same request does not have adverse effects.

## Frontend

### Technology Stack
- Framework: Next.js 13
- Design Framework: Tailwind CSS
- State Management: React Context
- Custom Hooks: Implemented for efficient data fetching and state management
- No External REST API Client: Leverages built-in libraries to interact with the backend

### Installation and Setup
1. Clone the repository.
2. Install Node.js and npm.
3. Run `npm install` to install the required dependencies.
4. Start the frontend server with `npm run dev`.

### Project Structure
- Atomic pattern is used for structuring the application.
- Custom hooks are implemented for reusability.
- React Context is employed for state management.
- Tailwind CSS is used for styling.

## GitHub API Integration
- All interactions with the GitHub API are abstracted in the backend using Go-based wrappers.
- The frontend communicates with the backend to retrieve data efficiently without the need for additional client libraries.

## Running the Application
1. Start the backend as instructed in the Backend section.
2. Start the frontend as instructed in the Frontend section.
3. Access the application in your web browser at `http://localhost:3000`.

This documentation should provide you with a clear understanding of the GitHub OAuth application, its components, and how to set up and run it locally. For further details on the code structure and implementation, please refer to the source code and comments in the respective repository files.
