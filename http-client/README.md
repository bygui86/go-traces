
# Go traces - HTTP client

## Endpoints

### Monitoring

URL: `localhost:9190/metrics`

### Products

Root URL: `localhost:8180/`

| Method | URL | Description
| --- | --- | --- |
| GET | /products | Fetch list of products |
| GET | /products/{id} | Fetch a product by ID |
| POST | /products | Create a new product |
| PUT | /products/{id} | Update an existing product retrieved by ID |
| DELETE | /products/{id} | Delete a product by ID |

---

## Run

`WARN: this application requires the relative HTTP server to work properly`

1. start client
    ```shell script
    make run
    ```

2. play a bit with [Postman](https://www.postman.com/) loading the [prepared collection](postman/postman_collection.json)
