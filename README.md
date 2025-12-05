# How to run the project 

1. Clone the repository
2. thanks to dev containers and docker the project should run automatically by running the next commands:

```bash
docker compose up -d --build redis db
sleep 10
docker compose up -d --build
```