
# **Student Service (gRPC)**  
This project is a **gRPC-based Student Management Service** built using **Go**, **PostgreSQL**, and **Protocol Buffers**.

---

## **📦 Project Structure**
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
Ensure you have the following installed:

| Dependency     | Windows                             | Mac                                |
|---------------|------------------------------------|-----------------------------------|
| **Go**        | [Download Go](https://go.dev/dl/) | `brew install go`                 |
| **PostgreSQL**| [Download Postgres](https://www.postgresql.org/download/) | `brew install postgresql` |
| **gRPC Tools**| Installed via `go mod tidy`       | Installed via `go mod tidy`       |
| **Postman**   | [Download](https://www.postman.com/downloads/) | [Download](https://www.postman.com/downloads/) |

---

# **📌 Step 1: Clone the Repository**
```sh
git clone https://github.com/your-repo/student-service.git
cd student-service
```

---

# **📌 Step 2: Install & Start PostgreSQL**
## **🔹 MacOS Setup**
```sh
brew install postgresql
brew services start postgresql
createdb eduaierp
```
## **🔹 Windows Setup**
1. Download and install **PostgreSQL** from [here](https://www.postgresql.org/download/).
2. Open **pgAdmin** or **Command Prompt** and run:
   ```sql
   CREATE DATABASE eduaierp;
   ```

---

# **📌 Step 3: Create PostgreSQL User & Grant Permissions**
### **🔹 Open PostgreSQL CLI (`psql`)**
```sh
psql -U postgres -d eduaierp
```
### **🔹 Create a new user**
```sql
CREATE USER ravikigf WITH PASSWORD 'root';
```
### **🔹 Grant permissions**
```sql
ALTER DATABASE eduaierp OWNER TO ravikigf;
GRANT ALL PRIVILEGES ON DATABASE eduaierp TO ravikigf;
```
### **🔹 Connect as the new user**
```sh
psql -U ravikigf -d eduaierp
```
Now, you are logged in as `ravikigf`.

---

# **📌 Step 4: Set Up Database Schema**
Run the following SQL commands:
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
GRANT USAGE, SELECT, UPDATE, INSERT, DELETE ON SEQUENCE students_student_id_seq TO ravikigf;
```

---

# **📌 Step 5: Set Up Environment Variables**
Create an `.env` file in the root directory:
```sh
touch .env
```
Edit the `.env` file:
```sh
nano .env
```
Add:
```
DATABASE_URL=postgres://ravikigf:root@localhost:5432/eduaierp
```
**For Windows (Command Prompt)**
```sh
set DATABASE_URL=postgres://ravikigf:root@localhost:5432/eduaierp
```

---

# **📌 Step 6: Install Dependencies**
```sh
go mod tidy
```

---

# **📌 Step 7: Run the gRPC Server**
```sh
go run cmd/main.go
```
If successful, you’ll see:
```
gRPC server listening on :50051
```

---

# **📌 Step 8: Test API with Postman**
1. Open **Postman**.
2. Click **"New" → "gRPC Request"**.
3. Enter server URL:
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

# **📌 Step 9: Stop the Server**
To stop the server:
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
- Restart:
  ```sh
  brew services restart postgresql  # Mac
  ```

### **🔹 Port Already in Use**
If **50051** is in use, update `server.go`:
```go
lis, err := net.Listen("tcp", ":50052")
```
Then, restart:
```sh
go run cmd/main.go
```
