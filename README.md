# OAuth User Login with React and Go
An example project demonstrating user authentication and authorization built with React.js and Go.  
It uses Google OAuth 2.0 for secure login and cookies for session management.

[Check out the live demo here](https://oauth.martishin.com/)!

<img src="https://i.giphy.com/media/v1.Y2lkPTc5MGI3NjExbnRwbDB6cmN2emtiaXhpY3hydWI3ZGJtbGM0cHZ2dzEzZXAxaHA5dCZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/KtKvOlylZtd9oOJQNF/giphy.gif" width="400"/>

## How OAuth 2.0 Works 
OAuth 2.0 is a protocol that allows apps to securely delegate authentication to providers like Google. The flow looks like this:
1. The app redirects you to the provider's login page
2. After you log in, the provider asks if you’re okay with sharing specific data (scopes) with the app
3. If you agree, it redirects you to the app's predefined callback URL, passing along the temporary authorization code as a query parameter
4. The app exchanges this code for a secure access token and optionally a refresh token, which allows the app to request new access tokens without requiring you to log in again
5. Using the access token, the app fetches your data (like your email and name) from the provider's APIs

## Prerequisites
1. **Install Dependencies**:
    - [Docker](https://www.docker.com/products/docker-desktop)
    - [Node.js and npm](https://nodejs.org/) for the React frontend.
    - [Go](https://golang.org/) for the backend.
2. **Set Up Google OAuth Credentials**:
    - Create a project in [Google Developer Console](https://console.developers.google.com/).
    - Generate a **Client ID** and **Client Secret**.
    - Set the redirect URI to `http://localhost:8100/auth/callback?provider=google`.

## Getting Started

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/martishin/react-golang-user-login-oauth
   cd react-golang-user-login-oauth
   ```

2. **Set Up Environment Variables**:
   Copy the .env.example file to .env in the root of the project and update the values as needed:
   ```env
   PORT=8100
   DB_HOST=localhost
   DB_PORT=5432
   DB_DATABASE=oauth
   DB_USERNAME=postgres
   DB_PASSWORD=postgres
   GOOGLE_CLIENT_ID=YOUR_GOOGLE_APP_CLIENT_ID
   GOOGLE_CLIENT_SECRET=YOUR_GOOGLE_CLIENT_SECRET
   GOOGLE_CALLBACK_URL=http://localhost:8100/auth/callback?provider=google
   REDIRECT_SECURE=http://localhost:5173/secure
   SESSION_COOKIE_DOMAIN=localhost
   ENV=local
   SESSION_SECRET=YOUR_SESSION_SECRET
   ```

3. **Start the Database**:
   ```bash
   docker compose up db pgadmin -d
   ```

4. **Start the Backend Server**:
   Navigate to the `server` directory:
   ```bash
   cd server
   make run
   ```

   The backend will be running at [http://localhost:8100](http://localhost:8100).

5. **Start the Frontend**:
   Navigate to the `client` directory:
   ```bash
   cd client
   npm install
   npm run dev
   ```

   The frontend will be available at [http://localhost:5173](http://localhost:5173).

6. **Veiw table in pgAdmin**:

   The portal will be available at [http://localhost:5050](http://localhost:5050).

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
