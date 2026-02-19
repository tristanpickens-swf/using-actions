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
docker run -p 8080:8080 phonebook:latest
```

API:
- POST /contacts  {"name":"Alice","phone":"123"}
```bash
curl -X POST http://localhost:8080/contacts \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice","phone":"123-456-7890","email":"alice@example.com"}'
```
- GET /contacts
```bash
curl http://localhost:8080/contacts
```