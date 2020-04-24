
# Go traces - HTTP server

## Endpoints

### Monitoring

URL: `localhost:9090/metrics`

### Products

Root URL: `localhost:8080/`

| Method | URL | Description
| --- | --- | --- |
| GET | /products | Fetch list of products |
| GET | /products/{id} | Fetch a product by ID |
| POST | /products | Create a new product |
| PUT | /products/{id} | Update an existing product retrieved by ID |
| DELETE | /products/{id} | Delete a product by ID |

---

## Run

1. start PostgreSQL
    ```shell script
    make run-postgres
    ```

2. start server
    ```shell script
    make run
    ```

3. play a bit with [Postman](https://www.postman.com/) loading the [prepared collection](postman/postman_collection.json)
