# Backend of ST-JO NSI - Shifumi Project

## This Backend is made in GoLang with HTTP Server 

### Project Structure

- `main.go` : Entry point of the project
- `src/` : Contains all the source code of the project


- `src/handlers` : Contains all the handlers of the project
- `src/models` : Contains all the models of the project
- `src/utils` : Contains all the utils of the project


### HTTP Response

- `200` : Success
- `500` : Internal Server Error or Request Format Error

Example of a response : 

```json
{
    "status": 200,
    "message": "Request Success",
    "data": {
        "id": 1,
        "name": "John Doe",
        "email": "john@example.fr"
    }
}
```

### EndPoints 

- `GET /` : Return a welcome message

- `POST /auth/login` : Login a user with email and password
- `PUT /auth/register` : Register a user 
- `GET /user` : Get all user information

- `POST /game/play` : Create new Game 

- `GET /game/pub` : Get a pub for coins 
- `PATCH /game/pub` : Claim a pub for coins

### Data Structure

This Project use Mysql database (Docker Composer)
