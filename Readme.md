# Snippetbox



Snippetbox is a web application written in Go that allows users to create and share code snippets. This project has been developed while following the book *Let's Go* by Alex Edwards.

## Features

- **User Authentication:** Signup and login functionality.
- **Create Snippets:** Logged-in users can create new code snippets.
- **View Snippets:** Anyone can view the created snippets.
- **Future Enhancements:** Additional features will be added beyond the book's implementation.

## Installation & Running

1. Clone the repository:
   ```sh
   git clone https://github.com/supersaint7780/snippetbox.git
   cd snippetbox
   ```
2. Set up the MySQL database:
   - Ensure MySQL is installed and running.
   - Create the required tables using the following SQL queries:
     ```sql
     -- Create a `snippets` table.
     CREATE TABLE snippets (
        id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
        title VARCHAR(100) NOT NULL,
        content TEXT NOT NULL,
        created DATETIME NOT NULL,
        expires DATETIME NOT NULL
     );

     -- Add an index on the created column.
     CREATE INDEX idx_snippets_created ON snippets(created);

     CREATE TABLE sessions (
        token CHAR(43) PRIMARY KEY,
        data BLOB NOT NULL,
        expiry TIMESTAMP(6) NOT NULL
     );

     CREATE INDEX sessions_expiry_idx ON sessions (expiry);

     CREATE TABLE users (
        id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
        name VARCHAR(255) NOT NULL,
        email VARCHAR(255) NOT NULL,
        hashed_password CHAR(60) NOT NULL,
        created DATETIME NOT NULL
     );

     ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);
     ```
3. Run the application:
   ```sh
   go run ./cmd/web -addr=:4000 -dsn="<user>:<password>@/snippetbox?parseTime=true"
   ```
   - `-addr` specifies the HTTP network address (default `:4000`).
   - `-dsn` specifies the MySQL data source name.

## Dependencies

Ensure you have Go installed and properly set up on your system.

## Acknowledgments

This project is built while learning from *Let's Go* by Alex Edwards.




