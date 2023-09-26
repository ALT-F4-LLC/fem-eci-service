## Build and Run

```bash
docker build -t test .

# Display env vars
docker run -it --rm --env=FEM_INSTRUCTOR=Erik --env=FEM_LOCATION=Minneapolis -p 8080:80 test

# Greet
curl http://localhost:8080

# Display env vars
curl http://localhost:8080/env
```
