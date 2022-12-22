Steps:

1. Setup postgresql and create new .env file and copy all environment variables from .env.example and replace with yours.
2. User Table will be automatically created under public schema when server will be started. 
3. Created 2 apis for now. One for registration and other for login.
    ```
    curl --location --request POST 'http://localhost:8000/auth/register' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "username": "vikram",
        "password": "vikram"
    }'
    ```
    
    ```
    curl --location --request POST 'http://localhost:8000/auth/login' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "username": "vikram",
        "password": "vikram"
    }'
    ```