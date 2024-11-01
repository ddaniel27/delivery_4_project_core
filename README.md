# Project Core
This microservice is responsible for managing projects. It provides the following functionalities:
- Create a project
- Consult a project

## How to run
To run this microservice, you need to have the following tools installed:
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/) (optional)

## API
This microservice provides the following endpoints:
- `POST /api/project`: Create a new project
- `GET /api/project/:id`: Get a project by id
- `GET /api/project/owner/:id`: Get all projects from an owner by id

## Database
This microservice uses a MongoDB database to store the projects. This is the database schema:
```json
{
  "_id": "string",
  "name": "string",
  "owner": {
    "id": "integer",
    "name": "string",
    "institution": "string"
  },
  "groups": [
    {
      "id": "integer",
      "name": "string"
    }
  ],
  "nominations": [
    {
      "name": "string",
      "url": "string"
    }
  ],
  "metadata": {
    "is_uploaded": "boolean",
    "is_public": "boolean",
    "filename": "string",
    "filetype": "string",
    "filesize": "integer",
    "object_key": "string",
    "duration": "integer",
    "created_at": "string",
    "updated_at": "string"
  }
}
```
