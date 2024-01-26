# CRUD API

## Endpoints

    [GET] /v1/devices
    List all devices
    Example: curl -X GET http://localhost:8080/v1/devices
    Response: [{"id":"1","name":"test","deviceBrand":"test","createdAt":"2021-07-04T16:00:00Z"}]

    [GET] /v1/devices/:id
    Get a device by id
    Example: curl -X GET http://localhost:8080/v1/devices/1
    Response: {"id":"1","name":"test","deviceBrand":"test","createdAt":"2021-07-04T16:00:00Z"}

    [POST] /v1/devices
    Add a new devices
    Example: curl -X POST http://localhost:8080/v1/devices -d '{"name":"test","deviceBrand":"test"}'
    Response: {"uuid":"1"}

    [DELETE] /v1/devices/:id
    Delete a device by id
    Example: curl -X DELETE http://localhost:8080/v1/devices/1
    Response: {}

    [PUT] /v1/devices/:id
    Update whole device object by id
    Example: curl -X PUT http://localhost:8080/v1/devices/1 -d '{"id":"1","name":"test","deviceBrand":"test","createdAt":"2021-07-04T16:00:00Z"}'
    Response: {"id":"1"}

    [PATCH] /v1/devices/:id
    Update partial device object by id
    Example: curl -X PATCH http://localhost:8080/v1/devices/1 -d '{"name":"test","deviceBrand":"test"}'
    Response: {"id":"1"}

    [GET] /v1/devices/search?q=test
    Search devices by brand
    Example: curl -X GET http://localhost:8080/v1/devices/search?q=test
    Response: [{"id":"1","name":"test","deviceBrand":"test","createdAt":"2021-07-04T16:00:00Z"}]
