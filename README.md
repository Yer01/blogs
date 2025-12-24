# Personal Blog

A Go-based personal blog application with authentication, article management, and responsive design. Built with Go, PostgreSQL, and HTML templates.

## Screenshots

### Guest Section
![Homepage](assets/screenshots/photo_2025-12-25_01-14-04.jpg)
![All Posts](assets/screenshots/photo_2025-12-25_01-14-23.jpg)
![Article View](assets/screenshots/photo_2025-12-25_01-14-26.jpg)

### Admin Section
![Create Post](assets/screenshots/photo_2025-12-25_01-14-29.jpg)
![Edit Post](assets/screenshots/photo_2025-12-25_01-14-31.jpg)

## Features

### Guest Section (Public)
- **Home Page**: Displays latest blog posts with post count
- **All Posts**: View complete list of published articles
- **Article Page**: Read full article content with publication metadata

### Admin Section (Authenticated)
- **Create Post**: Add new blog articles with title and content
- **Edit Post**: Update existing article content (title is immutable)
- **Delete Post**: Remove articles from the blog
- **Basic HTTP Authentication**: Protected routes using standard HTTP Basic Auth

## Tech Stack

- **Language**: Go 1.24+
- **Framework**: chi/v5 (HTTP router)
- **Database**: PostgreSQL with SQL migrations
- **Frontend**: HTML templates with CSS styling
- **Authentication**: HTTP Basic Authentication (SHA256 hashing)

## Project Structure

```
.
├── cmd/                           # Application code
│   ├── main.go                   # Entry point and server setup
│   ├── routes.go                 # Route definitions
│   ├── handlers.go               # HTTP request handlers
│   ├── middleware.go             # Authentication middleware
│   ├── templates.go              # Template rendering utilities
│   ├── helpers.go                # Helper functions
│   └── middleware.go             # Auth middleware
├── internal/
│   └── model/
│       └── blogs.go              # Blog data model and database operations
├── db/
│   └── migrations/               # Database migrations
│       ├── 000001_create_blogs_table.up.sql
│       ├── 000001_create_blogs_table.down.sql
│       ├── 000002_add_userid.up.sql
│       └── 000002_add_userid.down.sql
├── templates/                    # HTML templates
│   ├── base.tmpl                # Layout wrapper
│   ├── home.tmpl                # Homepage
│   ├── blog.tmpl                # Single article view
│   ├── blogs.tmpl               # All articles list
│   ├── create.tmpl              # Create article form
│   └── edit.tmpl                # Edit article form
├── static/
│   └── styles.css               # Stylesheet
├── go.mod                        # Go module definition
├── go.sum                        # Go dependencies lock file
└── .env                          # Environment variables (not committed)
```

## Installation

### Prerequisites
- Go 1.24+
- PostgreSQL 12+

### Setup Steps

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd blogs
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Configure environment variables**
   Create a `.env` file in the project root:
   ```env
   DATABASE_URL=postgres://user:password@localhost:5432/blogs_db
   AUTH_USERNAME=admin
   AUTH_PASSWORD=your_secure_password
   ```

4. **Create database and run migrations**
   ```bash
   # Create database
   createdb blogs_db
   
   # Run migrations (using a migration tool like migrate)
   migrate -path db/migrations -database "$DATABASE_URL" up
   ```

5. **Run the server**
   ```bash
   go run ./cmd/main.go
   ```

   Server will start on `http://localhost:8081`

## API Routes

### Public Routes (No Authentication Required)
- `GET /` - Homepage
- `GET /blogs` - All articles list
- `GET /blogs/:id` - View single article

### Protected Routes (Authentication Required)
- `GET /blogs/create` - Create article form
- `POST /blogs/create` - Submit new article
- `GET /blogs/:id/edit` - Edit article form
- `POST /blogs/:id/edit` - Submit article update
- `DELETE /blogs/:id` - Delete article

## Authentication

The application uses **HTTP Basic Authentication**:
- Username and password are stored as environment variables (`AUTH_USERNAME`, `AUTH_PASSWORD`)
- Passwords are hashed using SHA256 for comparison
- Protected routes return `401 Unauthorized` if credentials are invalid

To access protected routes, include Basic Auth header:
```bash
curl -u "admin:password" http://localhost:8081/blogs/create
```

## Database Schema

### blogs table
```sql
CREATE TABLE blogs (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) UNIQUE NOT NULL,
  content TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  user_id INT
);
```

## Development

### Run Tests
```bash
go test ./...
```

### Format Code
```bash
go fmt ./...
```

### Lint Code
```bash
golangci-lint run
```

## Future Enhancements

- User accounts and session management
- Comments system
- Tags and categories
- Search functionality
- Markdown support
- Image uploads
- Social media sharing
- Analytics and view tracking

## License

MIT License - feel free to use this project for personal and commercial purposes.

## Contributing

Contributions are welcome! Please feel free to submit pull requests with improvements.
