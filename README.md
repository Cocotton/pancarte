# Pancarte
This repository contains the Pancarte backend source code.

## Requirements
To run this code, you need to have
* Go installed
* MongoDB running

## Run
Once built, you'll need to provide the following environment variables to the application when running it
* `PANCARTE_DB_HOST`: The MongoDB host
* `PANCARTE_DB_NAME`: The MongoDB database to use
* `PANCARTE_PORT`: The port on which the application will listen
* `PANCARTE_JWT_SECRET`: The secret used to sign the JSON Web Token (JWT)

## Docker
To run Pancarte backend from a container, you'll first need to build the image. Make sure to update the previous environment variables in the Dockerfile.

You can also provide the environment variables while calling the `docker run` command.

You'll still need a MongoDB instance running somewhere.
