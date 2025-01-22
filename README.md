
# pyqqserver

A Go-based server that provides an API for managing and retrieving educational resources. It uses MongoDB as the database and Gorilla Mux for routing.

## Getting Started

### Prerequisites

- Go 1.23.1 or later
- MongoDB
- A `.env` file with the following environment variables:
  - `MONGO_URI`: The connection string for your MongoDB instance.

### Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/vishalvivekm/pyqqserver.git
   cd pyqqserver
   ```

2. Install the dependencies:

   ```sh
   go mod download
   ```

3. Create a `.env` file in the root directory of the project and add your MongoDB URI:

   ```sh
   MONGO_URI=your_mongo_db_uri
   ```

### Running the Application

To run the application, from the root of the project, use the following command:

```sh
go run main.go
```

The server will start and listen on the `Port` specified in the `.env` file or default to port 8080.

## API Endpoints

The application provides the following API endpoints:

### Get Resources

- **URL:** `/drive/{type}/{subject}`
- **Method:** `GET`
- **Description:** Retrieves resources based on type and subject.
- **Example:** `/drive/notes/math`

### Get Subjects

- **URL:** `/{course}/{semester}/{branch}`
- **Method:** `GET`
- **Description:** Retrieves subjects for a specific course, semester, and branch.
- **Example:** `/btech/secondsemesters/GEC`

### Get Subject Details

- **URL:** `/{course}/{semester}/{branch}/{subject}`
- **Method:** `GET`
- **Description:** Retrieves details for a specific subject.
- **Example:** `/btech/firstsemesters/GCS/math`

## Project Structure

- `app/`: Contains the main application code.
- `handler/`: Contains the HTTP handler functions.
- `db/`: Contains the database initialization code.
- `main.go`: The entry point for the application.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgements

- [Gorilla Mux](https://github.com/gorilla/mux)
- [godotenv](https://github.com/joho/godotenv)
- [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver)
- [rs/cors](https://github.com/rs/cors)

Feel free to contribute to this project by submitting issues or pull requests.
