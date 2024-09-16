# Wallet Transaction API

## Description

The Wallet Transaction API provides endpoints for managing wallet transactions. It includes functionalities for:
- **Top-Up**: Add funds to a wallet.
- **Send Money**: Transfer money from one wallet to another.
- **Get Balance**: Retrieve the current balance of a wallet.
- **Get Transaction History**: Obtain the transaction history for a wallet.

## Arcitecture
The architecture of this application is built using Golang with the Fiber framework and Gorm ORM, following a layered pattern that includes four main layers: Repository, Service, Model, and Controller. The application uses PostgreSQL as its database, ensuring a well-organized structure and efficient data management and request processing.

## Performance
The application implements locking mechanisms during data modifications to ensure safety from race conditions, especially when handling a high volume of requests. Additionally, it optimizes database interactions through connection pooling, which improves efficiency and prevents excessive load on the database. Nginx is utilized as a load balancer to distribute traffic evenly across instances, further enhancing the application's performance and scalability.

## How to Use

1. **Install Docker**: Ensure Docker is installed on your machine.

2. **Clone the Repository**:
   ```bash
   git clone https://github.com/llassingan/transwallet.git
   ```
3. **Set Up the .env File:** Create a `.env` file in the root of the project with the following format:
    ```dotenv
   POSTGRES_USER=
   POSTGRES_PASSWORD=
   POSTGRES_DB=

   DB_HOST=
   DB_PORT=
   DB_SSLMODE=
   DB_MAX_OPEN_CONNS=
   DB_MAX_IDLE_CONNS=
   DB_CONN_MAX_LIFETIME=

   PORT1=
   PORT2=
   ```
   Example:
   ```dotenv
   POSTGRES_USER=postgres
   POSTGRES_PASSWORD=admin
   POSTGRES_DB=db_wallet

   DB_HOST=pgwallet
   DB_PORT=5432
   DB_SSLMODE=disable
   DB_MAX_OPEN_CONNS=100
   DB_MAX_IDLE_CONNS=20
   DB_CONN_MAX_LIFETIME=30

   PORT1=8001
   PORT2=8002
   ```
4. **Run the Application** Start the application with Docker Compose by running:
   ```bash
   docker-compose up
   ```
5. **The application is now up and running**

## Improvement
To further enhance the application's reliability and performance, consider the following improvements:

1. Optimize Gorm Settings: Fine-tune Gorm configurations to improve query performance and reduce overhead. Consider using advanced features such as preloading, caching, and batch processing to optimize database interactions.

2. Database Tuning: Perform database optimizations, such as indexing frequently queried fields, optimizing query execution plans, and adjusting PostgreSQL settings for better performance under load.

3. Scaling: Implement horizontal scaling by adding more instances of the application.

4. Caching: Introduce caching mechanisms, such as Redis or Memcached, to reduce the load on the database and improve response times for frequently accessed data. Focus on caching data that does not need to be real-time, such as configuration settings, user preferences, or static content, to further enhance performance without compromising data accuracy.

5. Asynchronous Processing: Offload long-running tasks to background workers or message queues to improve the responsiveness of the application and prevent bottlenecks during high traffic periods.

6. Monitoring and Logging: Enhance monitoring and logging to gain deeper insights into application performance and quickly identify and address issues.

7. Load Testing: Conduct regular load testing to identify performance bottlenecks and ensure the application can handle expected traffic volumes.

By implementing these improvements, the application can achieve greater reliability, efficiency, and scalability, ensuring a better user experience and smoother operation under varying loads.