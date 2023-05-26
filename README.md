# FastAuthService
AuthService is a microservice responsible for user authentication and user management.
  
## Description
AuthService provides an API for user registration, authentication, and management of user accounts. It ensures secure storage and processing of user credentials using advanced encryption and hashing methods.  
  
## Features  
coming soon
  
## How to run service  
 - Clone project by `git clone https://github.com/vik-backend/FastAuthService.git`.  
 - Create a file **"./AuthService/.env"** and use it to set up environment.  
 - Ensure **CERTIFICATE_KEY_LOC** and **CERTIFICATE_FILE_LOC** are valid in your environment. Create sertificate and key if necessary and put them somewhere, then provide those paths in env.  
 - Ensure [**Docker**](https://www.docker.com/) engine is running in your system.  
 - Execute command `docker-compose up --build` from project **"./AuthService/"** (root) directory, where **docker-compose.yml** is located.  
  