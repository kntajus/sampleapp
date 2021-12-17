# Pre-requisites

These instructions assume you have the necessary software (e.g. Docker, Docker Compose) installed already.

# How to run

Run `docker compose up` from this directory.

# How to test

Run `curl -v http://localhost:8080/ports/ZACPT`. Note the 404 Not Found response as there is no data yet.

Run `curl -X POST -v http://localhost:8080/update` to request that the data store be populated from the sample JSON file. (In the interest of simplicity, the data file has been made available to the service via a docker volume and is in a known location.)

Run `curl -v http://localhost:8080/ports/ZACPT` again to retrieve the data for the requested port.
