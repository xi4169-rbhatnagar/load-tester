## Responsibilities of the Server:

#### A) Listen to the user requests
  * Port: 3000
  * endpoint: /submit
  * User Request Params (in json):
    - request_url: string
    - request_verb: one of {GET, POST, DELETE, PUT}
    - max_concurrent_users: integer >= 1

#### B)  Add a database entry for the incoming request

#### C) Add an entry in the task queue for the same request
  
#### D) Periodically poll the database's response table for an update
  * For each request, continuously ping the database for the latest updates
