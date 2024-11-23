Objective : 

Build a Go application - REST service which is able to process at least 10K requests per second.
● The service has one GET endpoint - /api/verve/accept which is able to accept an integer id as a
mandatory query parameter and an optional string HTTP endpoint query parameter. It should return
String “ok” if there were no errors processing the request and “failed” in case of any errors.
● Every minute, the application should write the count of unique requests your application received in
that minute to a log file - please use a standard logger. Uniqueness of request is based on the id
parameter provided.
● When the endpoint is provided, the service should fire an HTTP GET request to the provided endpoint
with count of unique requests in the current minute as a query parameter. Also log the HTTP status
code of the response

Solution :

I have implemented simple service using gin framework for above requirements and considering below pointers.

API Design**:
   - Chose Gin framework for high performance and simplicity.
   - Designed `/api/verve/accept` endpoint to handle required and optional parameters.

Concurrency**:
   - Used `sync.Map` for concurrent-safe operations on unique request IDs.
   - Implemented a Goroutine for periodic logging.

Logging**:
   - Used Logrus for structured JSON logging.

Scalability**:
   - Designed the application to handle 10K RPS by using efficient data structures and connection pooling.

Objective Extension-1

Instead of firing an HTTP GET request to the endpoint, fire a POST request. The data structure
of the content can be freely decided

I have implemented SendHTTPPOST method with payload to provided endpoint, now i can use this method to implementing
the desired logic.

So by switching from a GET to a POST request, we enable the application to send more flexible and structured data.