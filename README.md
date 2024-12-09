Overview
The Zocket Online Assessment is a web application designed for managing products in an e-commerce platform. It utilizes Redis and PostgreSQL for efficient data storage and retrieval, implementing microservices for image processing. This application is designed to handle product details, including the storage and compression of product images, and provides the capability for image processing asynchronously using message queues like RabbitMQ or Kafka.
Features
Product Management: Store and manage product details, including product name, description, price, and images.
Image Compression: Asynchronously process and compress product images.
Database Integration: Utilize PostgreSQL for persistent storage and Redis for caching and real-time operations.
Asynchronous Task Handling: Use message queues (RabbitMQ or Kafka) for background image processing tasks.
HTTP API: Exposes endpoints to interact with the product data and handle file uploads.
Technologies Used
Go: Backend programming language for building the web server and database logic.
PostgreSQL: Relational database for persistent storage of product and user information.
Redis: In-memory data store for caching and managing temporary data like product images.
Message Queue RabbitMQ: For processing image compression tasks asynchronously.
Docker: For containerizing services (if applicable).
Gin: Go web framework for building REST APIs.
GORM: ORM for interacting with PostgreSQL in Go.
Architecture
The system is designed using a microservices architecture with the following components:
API Server:
Handles incoming HTTP requests.
Interacts with PostgreSQL and Redis to manage product data.
Provides API endpoints to perform CRUD operations on products.
Image Processing Microservice:

Consumes messages from the message queue RabbitMQ.
Downloads product images, compresses them, and uploads to S3 .
Updates the PostgreSQL database with compressed image information.
Message Queue:
Handles the asynchronous processing of product images.
Ensures the image processing doesn't block the main API operations.
Setup Instructions
Prerequisites
Go: Install Go 
PostgreSQL: Install PostgreSQL on your system or use a cloud-based PostgreSQL database.
Redis: Install Redis locally or use a cloud-based Redis service.
Message Queue (Optional): RabbitMQ  for asynchronous image processing 
Local Setup
1. Clone the Repository
git clone https://github.com/Prathyusha1404/zocket_online_assessment.git
cd zocket_online_assessment

3. Set Up the Environment Variables
Create a .env file in the root of the project with the following configuration:
DATABASE_URL=postgres://username:5432@localhost:5432/prd_mngt
REDIS_URL=redis://localhost:6379
QUEUE_URL=rabbitmq://localhost:5672

4. Install Dependencies
go mod tidy(This command will install all the required dependencies for the Go project.)
5. Start the Database and Redis
For PostgreSQL, run:sudo service postgresql start
For Redis, run:redis-server
6. Run the Application
Start the API server by running:go run main.go
By default, the server will be available on http://localhost:8080
Assumptions
1.PostgreSQL is used for storing product details and user data. It is assumed that there is a working PostgreSQL instance.
2.Redis is used for caching purposes and quick data retrieval (especially product images).
3.Image Processing is handled asynchronously using a message queue RabbitMQ. The image processing microservice is responsible for downloading and compressing images, and the compressed images are then stored in cloud storage like S3.
The API server and image processing microservice are decoupled and communicate via the message queue.
