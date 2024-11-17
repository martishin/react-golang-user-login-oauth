# React-Golang OAuth Example
An example project integrating React.js frontend with Go backend. This project demonstrates user authentication via Google OAuth 2.0
(using [goth](https://github.com/markbates/goth/tree/master) library) and session management using secure cookies.

## Demo
[![Demo](https://img.youtube.com/vi/OIEf-HuxvRY/0.jpg)](https://www.youtube.com/watch?v=OIEf-HuxvRY)

## Prerequisites
1. **Install Dependencies**:
    - [Docker](https://www.docker.com/products/docker-desktop)
    - [Node.js and npm](https://nodejs.org/) for the React frontend.
    - [Go](https://golang.org/) for the backend.
2. **Set Up Google OAuth Credentials**:
    - Create a project in [Google Developer Console](https://console.developers.google.com/).
    - Generate a **Client ID** and **Client Secret**.
    - Set the redirect URI to `http://localhost:3000/auth/callback?provider=google`.

## Getting Started

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/martishin/react-golang-goth-oauth
   cd react-golang-goth-auth
   ```

2. **Set Up Environment Variables**:
   Copy the .env.example file to .env in the root of the project and update the values as needed:
   ```env
    PORT=3000
    APP_ENV=local
    BLUEPRINT_DB_HOST=localhost
    BLUEPRINT_DB_PORT=5432
    BLUEPRINT_DB_DATABASE=oauth
    BLUEPRINT_DB_USERNAME=postgres
    BLUEPRINT_DB_PASSWORD=postgres
    BLUEPRINT_DB_SCHEMA=public
    GOOGLE_CLIENT_ID=YOUR_GOOGLE_APP_CLIENT_ID
    GOOGLE_CLIENT_SECRET=YOUR_GOOGLE_CLIENT_SECRET
    GOOGLE_CALLBACK_URL=http://localhost:3000/auth/callback?provider=google
    SESSION_SECRET=YOUR_SESSION_SECRET
   ```

3. **Start the Database**:
   ```bash
   docker-compose up
   ```

4. **Start the Backend Server**:
   Navigate to the `server` directory:
   ```bash
   cd server
   go run main.go
   ```

   The backend will be running at [http://localhost:3000](http://localhost:3000).

5. **Start the Frontend**:
   Navigate to the `client` directory:
   ```bash
   cd client
   npm install
   npm run dev
   ```

   The frontend will be available at [http://localhost:5173](http://localhost:5173).

## API Endpoints

| Method | Endpoint              | Description                      |
|--------|-----------------------|----------------------------------|
| GET    | `/auth`               | Initiate Google OAuth flow       |
| GET    | `/auth/callback`      | Google OAuth callback            |
| GET    | `/auth/logout`        | Log out and clear session        |
| GET    | `/api/user`           | Get authenticated user details   |

## Tech Stack

- **Frontend**: React.js, Tailwind CSS
- **Backend**: Go, Chi, goth
- **Database**: PostgreSQL
- **Auth Provider**: Google OAuth 2.0
- **Others**: Docker
