# MetricsObservabilitySystem

1. Prerequisites
Before you begin, ensure you have met the following requirements:

- Python 3.8+.
- Go 1.18+
- Docker
- InfluxDB
- Grafana

2. Set Up the Python Virtual Environment for the dummy application

$ cd dummy-metric-app
$ python3 -m venv dummyapp
$ source venv/bin/activate
$ python3 app.py

3. Install Go dependencies

$ cd ..
$ go mod tidy

4. If Docker is not already installed, you can install it by following the official Docker installation guide

To start the system:

$ docker-compose up --build -d

To stop the system:

$ docker-compose down




