# Go Backend Test

This is a coding test to evaluate your skills in Go and backend development.

We use the [Echo framework](https://echo.labstack.com) for building REST services in Go.
If you prefer another framework, feel free to replace Echo.

## Preparation

You will need to install [Docker](https://docker.com).

This project contains Docker Compose configuration for starting a MongoDB database.

> We use Docker to run mongodb, which means there is no need for installing MongoDB directly on your machine.

To start the database, run the following command:

```bash
docker compose up -d
```

To stop the database, run the following command:

```bash
docker compose down
```

At first boot, the MongoDB database will be automatically populated with user data using the [`seed.js`](seed.js) script.
(if you need to start over, delete the `data` directory and restart the database)

The database will be available at `mongodb://localhost:27017`. See [`docker-compose.yml`](docker-compose.yml) file for default username and password.

> NOTE! If you are connecting to the database from inside a Docker container, you should use the service name `mongodb` as the hostname.

## Your tasks

You are expected to implement a REST service that provides specific endpoints.

Please make sure you handle errors and HTTP status codes in a meaningful way.  

The code should be well-structured and readable. Comments are fine, but the code should be as self-explanatory as possible.

> NOTE! The provided code is just a suggestion. You are free to change it as you see fit, including file names and project structure.

> IMPORTANT! Make sure to add instructions on how to run your code in the file [`INSTRUCTIONS.md`](INSTRUCTIONS.md).

### Get all users
Create a rest endpoint that returns all users from the database.

### Create new user
Create a rest endpoint that allows creating a new user.

### Deactivate user
We don't want to delete users entirely from our system (because of imaginary government regulations). Instead, we want to deactivate them.

Create a rest endpoint that allows deactivating a user.

(Of course, deactivated users should no longer be visible in the list of all users.)

"How do you want me to deactivate a user?", you might ask. It's up to you to decide.

## Bonus Tasks

If you want, you can also implement any or all of the following tasks:

### User sorting

Make it possible for the API client to sort the users by name, email, or age.

### Update user

Create a rest endpoint that allows updating a user.

### Simple Authentication

How would you implement a simple authentication mechanism for the REST service?

### Dockerize the application

Create a Dockerfile for the application and provide a Docker Compose configuration to start the application and the MongoDB database.

### Add a linter

Add a Makefile or similar containing a lint command that runs a linter on the code.
