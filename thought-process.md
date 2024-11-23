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



