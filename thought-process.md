Objective : Extension 2
Ensure that ID deduplication works effectively when the service operates behind a load balancer with multiple instances running concurrently. This ensures that even if two instances receive the same id simultaneously, the deduplication logic remains intact and no duplicate counts are logged.

Key Challenges
Distributed Nature: Each instance of the service may maintain its own local state, which can result in duplicate counts if the same id is processed by multiple instances.
Consistency Across Instances: There needs to be a central, shared mechanism to determine the uniqueness of ids across all instances.

Solution Approach
To tackle this, I designed the system to use Redis as a centralized storage for tracking processed ids. Redis offers atomic operations, which are crucial for ensuring that the deduplication logic is consistent across all instances.

Load Balancer Behavior:

Since all instances share the same Redis cluster, the load balancer's distribution of traffic does not impact deduplication. Whether the same id reaches one or multiple instances, the global Redis ensures consistency.