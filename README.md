--- 

Message Generator

This project is designed to generate messages and route them through various servers using gRPC. The architecture includes several components like Prometheus for monitoring, ELK stack for logging, and more. The project was built for fun and to learn technologies.
Project Architecture



Components
Message Generator (zg_generator)

This component generates messages at a specified interval and sends them to the router.

Docker Compose Configuration:

```yaml
version: '3.8'

networks:
    local-net:
      external: true

services:
    zg_generator:
        build:
            context: .
            dockerfile: ./Dockerfile
        container_name: zg_generator
        env_file:
        - .env-docker
        networks:
        - local-net
        ports:
        - "21122:21122"
        volumes:
        - ./zg_generator:/app
        restart: unless-stopped
```
Configuration File:

```yaml
prometheus:
  url: ${PROMETHEUS_URL}

generator:
    interval: 10
    count: 1

grpc_client:
  router_address: ${GRPC_SERVER_LISTEN_ADDRESS}

logstash:
  url: ${LOGSTASH_URL}
```
.env-docker File:

```env
GRPC_SERVER_LISTEN_ADDRESS=zg_router:50051
LOGSTASH_URL=http://logstash:5000
PROMETHEUS_URL=0.0.0.0:21122
```
Other Components:

    Router: Receives messages from the generator and routes them to processing servers.
    Processing Servers: Multiple servers that process the received messages.
    Prometheus: Monitors the application and collects metrics.
    ELK Stack: Collects and analyzes logs.
    Grafana: Visualizes the metrics collected by Prometheus.
    Kafka: A message broker that integrates with the backend.
    Databases: Includes MongoDB, MySQL, Redis for caching and indexing, and SQL/NoSQL repositories.

Getting Started
Prerequisites:

    Docker
    Docker Compose

Running the Project

Clone the repository:

```bash
git clone https://github.com/your-repo/message-generator.git
cd message-generator
```
Build and run the Docker containers:

```bash
docker-compose up --build
```

Environment Variables

Ensure to set the following environment variables in the .env-docker file:

    GRPC_SERVER_LISTEN_ADDRESS: Address of the gRPC server (e.g., zg_router:50051).
    LOGSTASH_URL: URL of the Logstash server (e.g., http://logstash:5000).
    PROMETHEUS_URL: URL of the Prometheus server (e.g., 0.0.0.0:21122).

This project is licensed under the MIT License.

--- 