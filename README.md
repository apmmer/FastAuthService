# FastAuthService
FastAuthService is a high-performance microservice designed for user authentication and management.

## Overview
FastAuthService provides a robust API for user registration, authentication, and account management. It employs advanced encryption and hashing methods to ensure secure storage and processing of user credentials.

## Project Status

### Done :heavy_check_mark:

- **Hybrid Authentication**: Utilizes a hybrid model of JWT and user sessions for authentication.
- **Performance Optimized**: No database requests are required for validating access tokens.
- **Dynamic Browser Data Usage**: Leverages IP and UserAgent data.
- **Advanced Token Management**: Provides robust mechanisms for token creation, verification, and update with multiple additional validators.
- **Secure Communication**: Employs TLS for API and database connections.
- **No Database Frameworks**: Does not rely on database frameworks, thus reducing potential attack vectors. Optimized DB requests.
- **Protection Against Web Vulnerabilities**: Safeguards in place to prevent XSS attacks and SQL injections.
- **Well-Structured Project**: The project adheres to best practices in code organization and module separation.
- **Fully Documented Code**: All classes, methods, and functions are documented.
- **Comprehensive Package Documentation**: All packages have *doc.go* file that explain their purpose, structure, and usage.

### To Do :construction:

- DDoS attack mitigation
- Caching
- Monitoring
- More functionality for users administration
- Tests coverage at all levels
- Code review and refactor for improved maintainability
- Continuous integration and deployment setup
- Increase logging and observability

Stay tuned for upcoming updates!
  
## Running the Service  
To run the service, follow these steps:  
 1. Clone the project using:
    ```bash
    git clone https://github.com/vik-backend/FastAuthService.git
    ```
 2. Navigate to the project directory:  
    ```bash
    cd ../AuthService
    ```
 3. Ensure [**Docker**](https://www.docker.com/) engine is running in your system.  
 4. Run the following command to start the service:
    ```bash
    docker-compose up --build
    ```
