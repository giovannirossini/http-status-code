# HTTP Status Code API
This API is built with Go and Gin framework and provides HTTP status code descriptions hosted on Heroku. 
You can access the API at https://http-status-code.herokuapp.com.

## Endpoints
### Get all status codes

Endpoint: `/`

This endpoint returns a JSON object containing all HTTP status codes and their descriptions.

Example Request:

```shell
curl -XGET https://http-status-code.herokuapp.com/
```
Example Response:

```json
[
  {
    "id": "1xx",
    "description": "Informational",
    "codes": [
      {
        "id": "100",
        "description": "Continue"
      },
      {
        "id": "101",
        "description": "Switching Protocols"
      },
      {
        "id": "102",
        "description": "Processing"
      }
    ]
  },
  ...
]
```

Get status code description by ID
Endpoint: `/:id`

This endpoint returns the description for the specified HTTP status code.

Example Request:

```shell
curl -XGET https://http-status-code.herokuapp.com/404
````

Example Response:

```json
{
  "status_code": 404,
  "description": "Not Found"
}
```

Get status code descriptions by group `1xx`, `2xx`, `3xx`, `4xx`, `5xx`.

Endpoint: `/:group`

Example Request:

```shell
curl -XGET https://http-status-code.herokuapp.com/1xx
```

Example Response:

```json
{
  "id": "1xx",
  "description": "Informational",
  "codes": [
    {
      "id": "100",
      "description": "Continue"
    },
    {
      "id": "101",
      "description": "Switching Protocols"
    },
    {
      "id": "102",
      "description": "Processing"
    }
  ]
}
```

## Running the API locally

Clone the repository:

```bash
git clone https://github.com/giovannirossini/http-status-code.git
```

Navigate into the project directory:

```bash
cd http-status-code
```

Execute docker-compose:

```bash
docker-compose up
```

The server will be available at http://localhost.
