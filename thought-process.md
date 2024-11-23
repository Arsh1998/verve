Objective : Extension 3
Instead of writing the count of unique received ids to a log file, send the count of unique received
ids to a distributed streaming service of your choice.

Solution Approach
I have replaced the log file mechanism with an AWS Kinesis-based solution, it will send the unique request count for each minute to a distributed streaming service.

Kinesis Provides these advantages :
Scalability: Kinesis is designed to handle massive amounts of data and can process high throughput, making it ideal for a service targeting 10,000 requests per second.
Durability: Records sent to Kinesis are stored durably and can be replayed if needed.
Integration: It integrates seamlessly with analytics tools, databases, and AWS services like S3, Lambda, or Redshift.
Real-Time Processing: Data in Kinesis can be consumed by multiple downstream systems for real-time analytics or other use cases.


By using AWS Kinesis, the application achieves a scalable and reliable mechanism to publish unique request counts. This approach not only meets the functional requirements but also provides a foundation for real-time analytics and future enhancements.


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

