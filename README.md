# api-go-crud

Example of crud api using go. It has the default methods: GET, POST, PUT and DELETE and use a Event model as example.

## How to run

```bash
go run cmd/main.go
```

## Endpoints

#### Get all Events

```bash
GET: "/api/events"
```

#### Get Event by id

```bash
GET: "/api/events/{id}"
```

#### Create Event

Need a body with the following structure

```json
{
  "name": "Big Event",
  "description": "a very big event",
  "address": "downtown square"
}
```

```bash
POST: "/api/events"
```

#### Update Event

Need a body with the following structure (if some field is not present, it will be set to empty)

```json
{
  "name": "Big Event",
  "description": "a very big event",
  "address": "downtown square"
}
```

```bash
PUT: "/api/events/{id}"
```

#### Delete Event

```bash
DELETE: "/api/events/{id}"
```
