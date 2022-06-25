# evolve-credit-test
golang test api

Create a Golang RESTful API that returns a dataset of your choosing (e.g users). You are required to set up a Postgres database on any free platform you are familiar with. The API should return data in with the following flexibility.
By email (or any other unique identifier).
By a date range.
The API should allow for flexible pagination;
It should allow the client to retrieve data by pages (10 records per page) only.
It should allow the client to specify the number of records per page.


BASE  URL :107.20.78.154:9090

POST /user - saves a user


GET /users - gets all user (query parameters are page and limit, default limit is 10)


GET /users/{email} - gets user with that particular email (query parameters are page and limit, default limit is 10)


GET /users/date/{start}/{end} get users in that range of dates (query parameters are page and limit, default limit is 10)
