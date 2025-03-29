

# **Student Service (gRPC)**
This project is a **gRPC-based Student Management Service** built using **Go**, **PostgreSQL**, and **Protocol Buffers**.

## **📌 Features**
- **CRUD operations** for managing students.
- Uses **gRPC** for communication.
- **PostgreSQL** as the database.

---

# **📦 Project Structure**
```
├── README.md
├── cmd
│   └── main.go            # Main entry point for gRPC server
├── go.mod
├── go.sum
├── internal
│   ├── db
│   │   └── postgres.go    # Database connection logic
│   └── server
│       └── server.go      # gRPC server implementation
├── logfile                # Log files directory
└── proto
    ├── student.pb.go
    ├── student.proto
    └── student_grpc.pb.go
```

---

# **🛠️ Prerequisites**
Ensure you have the following installed on your system:

| Dependency     | Windows                             | Mac                                |
|---------------|------------------------------------|-----------------------------------|
| **Go**        | [Download Go](https://go.dev/dl/) | `brew install go`                 |
| **PostgreSQL**| [Download Postgres](https://www.postgresql.org/download/) | `brew install postgresql` |
| **gRPC Tools**| Installed via `go mod tidy`       | Installed via `go mod tidy`       |
| **Postman** (for API testing) | [Download Postman](https://www.postman.com/downloads/) | [Download Postman](https://www.postman.com/downloads/) |

---

# **📌 Step 1: Clone the Repository**
```sh
git clone https://github.com/your-repo/student-service.git
cd student-service
```

---

# **📌 Step 2: Set Up PostgreSQL**

## **🔹 MacOS Setup**
1. Install PostgreSQL:
   ```sh
   brew install postgresql
   ```
2. Start PostgreSQL:
   ```sh
   brew services start postgresql
   ```
3. Create a database:
   ```sh
   createdb eduaierp
   ```

## **🔹 Windows Setup**
1. Download and install **PostgreSQL** from [here](https://www.postgresql.org/download/).
2. Open **pgAdmin** or **Command Prompt** and run:
   ```sql
   CREATE DATABASE eduaierp;
   ```

---

# **📌 Step 3: Set Up Environment Variables**
Create an `.env` file in the root directory:
```sh
touch .env
```
Edit the `.env` file:
```sh
nano .env
```
Add the following:
```
DATABASE_URL=postgres://ravikigf:root@localhost:5432/eduaierp
```
**For Windows (Command Prompt)**
```sh
set DATABASE_URL=postgres://ravikigf:root@localhost:5432/eduaierp
```

---

# **📌 Step 4: Set Up Database Schema**
Launch `psql` and execute:
```sql
CREATE TABLE students (
    student_id TEXT PRIMARY KEY DEFAULT 'S' || nextval('students_student_id_seq'),
    name TEXT,
    dob TEXT,
    gender TEXT,
    nationality TEXT,
    aadhar_number TEXT,
    contact_number TEXT,
    email TEXT,
    permanent_address TEXT,
    current_address TEXT,
    parent_guardian_name TEXT,
    parent_guardian_contact TEXT,
    previous_institution TEXT,
    previous_grade TEXT,
    marks_obtained TEXT,
    passing_year TEXT,
    board_name TEXT
);

ALTER TABLE students OWNER TO ravikigf;
GRANT ALL PRIVILEGES ON TABLE students TO ravikigf;
```

---

# **📌 Step 5: Install Dependencies**
```sh
go mod tidy
```

---

# **📌 Step 6: Run the gRPC Server**
```sh
go run cmd/main.go
```
If successful, you’ll see:
```
gRPC server listening on :50051
```

---

# **📌 Step 7: Test API with Postman**
1. Open **Postman**.
2. Click **"New" → "gRPC Request"**.
3. In the **server URL**, enter:
   ```
   localhost:50051
   ```
4. Click **"Connect"**.
5. Choose **"GetStudent"** and enter:
   ```json
   {
     "student_id": "S12345"
   }
   ```
6. Click **"Send"** and check the response.

---

# **📌 Step 8: Stop the Server**
To stop the server, press:
```sh
CTRL + C
```

---

# **📌 Troubleshooting**
### **🔹 Database Connection Issues**
- Check if PostgreSQL is running:
  ```sh
  pg_ctl status
  ```
- Restart the service:
  ```sh
  brew services restart postgresql  # Mac
  ```

### **🔹 Port Already in Use**
If you get an error that port **50051** is in use, change the port in `server.go`:
```go
lis, err := net.Listen("tcp", ":50052")
```
Then, re-run the server:
```sh
go run cmd/main.go
```

---

