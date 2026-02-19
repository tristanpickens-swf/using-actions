# Phonebook API (Go)

Simple phonebook HTTP service with JSON file persistence.

Build locally:

```bash
go build -o phonebook .
./phonebook
```

Run with Docker:

```bash
docker build -t phonebook:latest .
docker run -p 8080:8080 -v $(pwd)/data.json:/app/data.json phonebook:latest
```

API:
- POST /contacts  {"name":"Alice","phone":"123"}
- GET /contacts
