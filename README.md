API Endpoints
POST /v1/events - Create a new event.
PATCH /v1/events/{event_id} - Update an existing event.
DELETE /v1/events/{event_id} - Delete an event.
GET /v1/events/{event_id} - Get an event by ID.
GET /v1/events - Get all events with pagination and category filter.
API Documentation
To view the API documentation, you can access the Swagger UI at http://localhost:8080/swagger/index.html after starting the application. The Swagger UI provides detailed information about the API endpoints, request parameters, and response data.

Testing with Postman
If you have Postman installed, you can use it to test the API endpoints.

Open Postman and import the collection provided in the repository https://www.postman.com/galactic-capsule-151700/workspace/mini-project/request/13711908-7503a816-5c34-4331-9dba-f2c7cf82277d

Make sure the Go application is running.

Use the imported collection to test each API endpoint. You can send requests to http://localhost:8080/v1/events and specify the required parameters.

Check the responses from the server to verify that the API is working as expected.

Deployment
Add any deployment instructions here if applicable.

Built With
Go (Golang) - The programming language used.
Chi - Lightweight and flexible HTTP router for Go.
Swagger - API documentation and specification.
