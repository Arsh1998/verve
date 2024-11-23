Objective : Extension 2
Ensure that ID deduplication works effectively when the service operates behind a load balancer with multiple instances running concurrently. This ensures that even if two instances receive the same id simultaneously, the deduplication logic remains intact and no duplicate counts are logged.

Key Challenges
Distributed Nature: Each instance of the service may maintain its own local state, which can result in duplicate counts if the same id is processed by multiple instances.
Consistency Across Instances: There needs to be a central, shared mechanism to determine the uniqueness of ids across all instances.

Solution Approach
To tackle this, I designed the system to use Redis as a centralized storage for tracking processed ids. Redis offers atomic operations, which are crucial for ensuring that the deduplication logic is consistent across all instances.

Load Balancer Behavior:

Since all instances share the same Redis cluster, the load balancer's distribution of traffic does not impact deduplication. Whether the same id reaches one or multiple instances, the global Redis ensures consistency.


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
