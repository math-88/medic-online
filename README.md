# Medical Service Example

This repository contains an example of an online medical service developed using a clean architecture approach, domain-driven design, and dependency injection. The application leverages various technologies to input patient medical history, request diagnoses from OpenAI's API, perform intermediate calculations, and return responses.

## Technologies Used

- Clean Architecture: The application is structured according to clean architecture principles, separating it into distinct layers: domain, ports, and interfaces.

- Domain-Driven Design: The domain logic of the application is organized using domain-driven design (DDD) principles, enabling a clear understanding of the core business logic.

- Dependency Injection: Dependency injection is employed to manage dependencies and promote decoupling between components, making the code more testable and maintainable.

- GRPC: The application utilizes GRPC for efficient communication between components, enhancing performance and enabling seamless integration of various services.

- HTTP over GRPC: HTTP is used over GRPC to provide a user-friendly interface for client interactions, allowing external services and clients to communicate with the application effectively.

- SQL Database: The SQL schema is designed externally and then imported into the application. The [sqlc](https://github.com/kyleconroy/sqlc) library is used to generate Go objects for the database, simplifying database interactions.

- In-Memory Database for Testing: For testing purposes, the application employs a virtual in-memory database, which is connected through interfaces, ensuring a reliable and isolated environment for testing.

## Project Structure

The application is structured as follows:

- Domain Layer: Contains the core business logic, models, and use cases.
- Ports and Interfaces: Define the input and output ports and interfaces for interacting with external services and databases.
- Delivery: Handles the HTTP and GRPC server configurations, acting as an entry point for clients.

## Getting Started

To set up and run the application, follow these steps:

1. Clone this repository to your local machine.
2. Set up the necessary dependencies, including Go, GRPC, and any other project-specific dependencies.
3. Configure your SQL database according to the schema provided in an external service.
4. Use the [sqlc](https://github.com/kyleconroy/sqlc) library to generate Go objects for your database.
5. Modify the configuration files for database connections and API keys as needed.
6. Build and run the application.

## Usage

The application allows clients to input patient medical history, request diagnoses, and perform various medical-related calculations. Clients can interact with the service using HTTP and GRPC, following the defined interfaces and protocols.

## Contributing

We welcome contributions from the open-source community to enhance and improve this example application. Please follow the established contribution guidelines and best practices when submitting pull requests.

## License

This project is licensed under the [MIT License](LICENSE), which means you are free to use and modify the code for your own projects.

Feel free to explore the code and experiment with the example online medical service. If you have any questions or encounter issues, please open a GitHub issue, and we'll be happy to assist you.

Thank you for using this example repository to explore clean architecture, domain-driven design, and dependency injection in the context of an online medical service. We hope it serves as a valuable resource for your software development endeavors.
