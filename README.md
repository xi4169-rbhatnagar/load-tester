## Architecture Diagram:
---
![Architecture Diagram](https://github.com/xi4169-rbhatnagar/load-tester/blob/master/load-tester-architecture.png)

## Entities:
---

### Client:
Sends a request with the following params:  
  * request_url: string
  * request_verb: {GET, POST, DELETE, PUT}
  * max_concurrent_users: integer >= 1

### Server:
Responsibilities of the server:
  * Listen to the user requests
  * Validate the user request
  * Add a database entry for the incoming request
  * Add an entry in the task queue for the same request
  * For each request, continuously ping the database for the latest updates


### Controller:
Responsibilities of the controller:
  * Listen to the task queue
  * read the request data from the db for the given entry in the task-queue
  * for each task, spawn a worker instance to process that request with the data received from the task-queue
  * maintain the request `max_concurrent_users` users in the system. A single go routine would be treated as a user. So, should be `max_concurrent_users` users in the system at any given time.
  * after the worker is done processing any request, aggregated metrics must be re-calculated and updated in the database.

### Worker:
Responsibilities of a worker:
  * Call the given url with correct parameter also given by the controller.
  * Abort the current request if it breaches the end-time the user might've specified. For example: if the user has asked at 09:10 to run the requests for 10min and there is a request that is an ongoing request at 09:20, it must be aborted.
  * Report back the request status (pass|fail), response code, and the time required to serve it.


## Database:
- request_data:
    1. id: autogenerated
    2. request_url
    3. http_verb
    4. max_concurrent_users
    5. status: (ONGOING|FAILED|SUCCEEDED)
- metrics:
    1. request_data::id
    2. start_time: Unix time when the request was sent  
    3. request_status: PASSED|FAILED
    4. response_time: (in ms)
    5. response_status_code

Note: The metrics table has a very high velocity, we could've also used a queue to buffer the incoming requests until it is written in the database.
