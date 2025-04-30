# Music CRUD API

A simple RESTful API for managing songs built with Go and Gin.

## Features

- Create, Read, Update, and Delete songs
- In-memory database (no external DB required)
- Simple and lightweight

## API Endpoints

| Method | Endpoint   | Description               |
| ------ | ---------- | ------------------------- |
| GET    | /songs     | Get all songs             |
| GET    | /songs/:id | Get a specific song by ID |
| POST   | /songs     | Create a new song         |
| PUT    | /songs/:id | Update an existing song   |
| DELETE | /songs/:id | Delete a song             |

## Song Structure

```json
{
  "id": 1,
  "title": "Song Title",
  "artist": "Artist Name",
  "genre": "Music Genre",
  "released": 2025
}
```
````

## Getting Started

1. Clone the repository

```bash
git clone https://github.com/KibuuleNoah/MusicCrudApi.git
```

2. Install dependencies: `go get -u github.com/gin-gonic/gin`

```bash
go mod tidy
```

3. Run the server

```bash
go run main.go
```

4. The API will be available at

```bash
http://localhost:8080
```

## Examples

1. **GET /songs** - Get all songs

```bash
curl http://localhost:8080/songs
```

2. **GET /songs/:id** - Get a specific song by ID

```bash
curl http://localhost:8080/songs/1
```

3. **POST /songs** - Create a new song

```bash
curl -X POST -H "Content-Type: application/json" -d '{
    "title": "Sijja",
    "artist": "Dokta Brain",
    "genre": "RnB",
    "released": 2025
}' http://localhost:8080/songs
```

4. **PUT /songs/:id** - Update an existing song

```bash
curl -X PUT -H "Content-Type: application/json" -d '{
    "title": "Mpeke",
    "artist": "Anknown ft Aroma",
    "genre": "RnB",
    "released": 2025
}' http://localhost:8080/songs/1
```

5. **DELETE /songs/:id** - Delete a song

```bash
curl -X DELETE http://localhost:8080/songs/2
```

**Error Cases**

**Song Not Found (GET, PUT, DELETE)**

```bash
curl http://localhost:8080/songs/999
```

**Response:**

```json
{
  "msg": "Song with ID 999 Not Found"
}
```
