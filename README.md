# Notification Service with Kafka

A scalable, microservice-based notification system built with Go, Kafka, and Firebase Cloud Messaging (FCM). This service provides a robust solution for sending push notifications with asynchronous processing capabilities.

## 🏗️ Architecture

This project follows a **Clean Architecture** pattern with clear separation of concerns:

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Controller    │───▶│    UseCase      │───▶│   Repository    │
│   (HTTP Layer)  │    │ (Business Logic)│    │ (Data Access)   │
└─────────────────┘    └─────────────────┘    └─────────────────┘
                                                       │
                                                       ▼
                                              ┌─────────────────┐
                                              │   Kafka Topic   │
                                              │ "notifications" │
                                              └─────────────────┘
                                                       │
                                                       ▼
                                              ┌─────────────────┐
                                              │   Consumer      │
                                              │ (Background)    │
                                              └─────────────────┘
                                                       │
                                                       ▼
                                              ┌─────────────────┐
                                              │   Firebase FCM  │
                                              │ (Push Service)  │
                                              └─────────────────┘
```

### Components

- **Controller Layer**: Handles HTTP requests and responses
- **UseCase Layer**: Contains business logic and validation
- **Repository Layer**: Manages data persistence and external integrations
- **Kafka**: Message queue for asynchronous processing
- **Firebase FCM**: Push notification delivery service

## 📋 Prerequisites

Before running this project, ensure you have:

- **Go 1.19+** installed
- **Docker & Docker Compose** for Kafka setup
- **Firebase Project** with FCM credentials
- **Git** for version control

## 🛠️ Installation & Setup

### 1. Clone the Repository

```bash
git clone https://github.com/lakshya1goel/Notification-Service-Using-Kafka.git
cd Notification-Service-Using-Kafka
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Set Up Kafka

Start Kafka using Docker Compose:

```bash
docker-compose up -d
```

This will start:
- Zookeeper (port 2181)
- Kafka (port 9092)
- Kafka UI (port 8080) - for monitoring

### 4. Configure Firebase

1. Create a Firebase project at [Firebase Console](https://console.firebase.google.com/)
2. Enable Cloud Messaging
3. Download your credential JSON file and save it in anywhere in the project.
4. Replace the path in util/util.go file.

## 🏃‍♂️ Running the Application

### Start the Service

```bash
go run main.go
```

The service will start on `http://localhost:8000`

### Verify Setup

**Check Kafka**: Visit `http://localhost:8080` for Kafka UI

## 📡 API Endpoints

### Send Notification

**POST** `/api/notification/`

Send a push notification to users.

#### Request Body

```json
{
  "title": "Hello World",
  "message": "This is a test notification"
}
```

#### Response

```json
{
  "success": true,
  "message": "Notification published to Kafka"
}
```

#### Example Usage

```bash
curl -X POST http://localhost:8000/api/notification \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Welcome!",
    "message": "Thank you for using our service"
  }'
```

## 🔄 Message Flow

1. **HTTP Request** → Controller receives notification request
2. **Validation** → UseCase validates business rules
3. **Publish** → Repository publishes message to Kafka topic
4. **Consume** → Background consumer picks up message
5. **Process** → Consumer deserializes and processes notification
6. **Send** → Firebase FCM delivers push notification to devices

## 📊 Monitoring

### Kafka UI
- **URL**: `http://localhost:8080`
- **Features**: Topic monitoring, message inspection, consumer group status

## 📁 Project Structure

```
notification/
├── api/
│   ├── controller/          # HTTP request handlers
│   ├── router/             # Route definitions
│   ├── service/            # Business services
│   └── util/               # Utility functions
├── domain/
│   ├── dto/                # Data Transfer Objects
│   └── model/              # Domain models
├── kafka/
│   ├── consumer.go         # Kafka consumer implementation
│   └── producer.go         # Kafka producer implementation
├── repository/             # Data access layer
├── usecase/                # Business logic layer
├── main.go                 # Application entry point
├── docker-compose.yaml     # Kafka setup
├── go.mod                  # Go module file
└── README.md               # This file
```

## 🔧 Configuration

### Kafka Configuration

Update Kafka settings in `main.go`:

```go
consumer := kafka.NewKafkaConsumer(kafka.KafkaConsumerConfig{
    Brokers:    []string{"localhost:9092"},    // Kafka broker addresses
    Topic:      "notifications",               // Topic name
    GroupID:    "notification-consumer-group", // Consumer group ID
    PushSender: pushNotificationSender,
})
```

### Producer Configuration

Update producer settings in `repository/notification_repo.go`:

```go
producer: kafka.NewKafkaProducer(kafka.KafkaProducerConfig{
    Broker: "localhost:9092", // Kafka broker address
    Topic:  "notifications",  // Topic name
}),
```

## 🔑 FCM Token Configuration

Update the FCM token in `api/service/push_notification_service.go`:


## 🤝 Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature-name`
3. Commit changes: `git commit -am 'Add feature'`
4. Push to branch: `git push origin feature-name`
5. Submit a pull request

## 🙏 Acknowledgments

- [Kafka Go Client](https://github.com/segmentio/kafka-go)
- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [Firebase Admin SDK](https://firebase.google.com/docs/admin/setup)

## 📞 Support

For support and questions:
- Create an issue in the GitHub repository
- Contact: [lakshya1234goel@gmail.com]

---