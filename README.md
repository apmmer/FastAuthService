# FastAuthService
FastAuthService is an aspiring high-performance microservice project, designed for user authentication and management.

## Overview
As a part of a larger project, FastAuthService is intended to serve as a robust microservice that will be accessed via a gateway. Kubernetes is planned to be used for organizing these microservices. The service aims to provide a robust API for user registration, authentication, and account management, while utilizing advanced encryption and hashing methods for secure storage and processing of user credentials.

FastAuthService is designed with reusability in mind and can serve as a frame or template for authentication services in many microservices-based projects. This could be particularly useful for those who prefer not to rely on external services like Keycloak for authentication.

The project is currently under active development, and contributions are welcome.

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
- **Fully Documented Code**: All classes, methods, functions are documented.
- **Comprehensive Package Documentation**: All packages have *doc.go* file that explain their purpose, structure or usage.

### To Do :construction:

- DDoS attack mitigation
- Caching
- Monitoring
- More functionality for users administration
- Google auth
- Tests coverage at all levels
- Code review and refactor for improved maintainability
- Continuous integration and deployment setup
- Increase logging and observability
- Advanced CI and prod CD

Stay tuned for upcoming updates!
  
## Running the Service  
To run the service, follow these steps:  
 1. Clone the project using:
    ```bash
    git clone https://github.com/vik-backend/FastAuthService.git
    ```
 2. Navigate to the project directory.  
 3. Ensure [**Docker**](https://www.docker.com/) engine is running in your system.  
 4. Run the following command to start the service:
    ```bash
    docker-compose up --build
    ```

## Development and Deployment Workflow :arrow_forward:

The workflow for feature development and deployment in this project follows these steps:

1. **Feature Branch Creation**: A separate branch is created for each feature.
2. **Pull Request (PR) to Develop**: After the development is completed, a PR is opened into the `develop` branch. Upon PR creation, automated tests are run.
3. **Release and Staging**: When a new release is ready, it is tagged in the format "**v*.*.***". The tagged code is then automatically pushed to the staging environment on AWS.
4. **Pull Request to Main and Production Deployment**: If everything works fine on the staging environment, a PR is opened into the `main` branch. After merging into `main`, the project is automatically pushed to the production environment with the same tag that was last in the `develop` branch.

This process ensures consistent, test-driven development and seamless transitions between development, staging, and production environments.
