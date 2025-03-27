Here is the full `README.md` content that you can directly copy and paste into your `README.md` file:

# Student Service - gRPC API

This is a simple gRPC-based server that provides an API for fetching student details using gRPC. The service simulates student data based on a `student_id`.

## Project Structure

```plaintext
student-service/
├── cmd/
│   └── main.go
├── proto/
│   ├── student.proto
│   ├── student.pb.go
│   └── student_grpc.pb.go
├── internal/
│   └── server/
│       └── server.go
├── go.mod
├── go.sum
└── README.md
```

- **cmd/main.go**: Main entry point for the server.
- **proto/student.proto**: The Protocol Buffers file that defines the service, request, and response types.
- **proto/student.pb.go & student_grpc.pb.go**: Generated Go code from the `.proto` file.
- **internal/server/server.go**: Server-side logic where gRPC service methods are implemented.

## Requirements

- **Go**: Ensure that Go is installed (version 1.18+).
- **Protocol Buffers**: Install `protoc` to generate Go code from the `.proto` file.
  - You can install it from [here](https://grpc.io/docs/protoc-installation/).
- **Postman** or **gRPC Client** to test the server.

## Setup Instructions

### Step 1: Install Dependencies

First, ensure Go and necessary gRPC dependencies are installed. 

1. Install the gRPC and Protobuf Go libraries:

   ```bash
   go get google.golang.org/grpc
   go get google.golang.org/protobuf
   ```

2. Install `protoc-gen-go` and `protoc-gen-go-grpc` for generating Go code from `.proto` files:

   ```bash
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   ```

### Step 2: Generate Go Code from Proto File

Once the dependencies are installed, generate the Go code for the gRPC service by running the following command:

```bash
protoc --go_out=. --go-grpc_out=. proto/student.proto
```

This will generate the necessary `student.pb.go` and `student_grpc.pb.go` files inside the `proto` folder.

### Step 3: Run the Server

Once the setup is complete, you can run the server with the following command:

```bash
go run cmd/main.go
```

By default, the server will start listening on port `50051`.

### Step 4: Testing the gRPC Server

You can test the gRPC server using **Postman** or any other gRPC client:

1. Open Postman and create a new request.
2. Set the **Request Type** to **gRPC**.
3. Set the **URL** to `localhost:50051`.
4. Click on **Import a .proto file** and import the `student.proto` file.
5. Select the **Service** `StudentService` and **Method** `GetStudent`.
6. In the **Request Payload** section, provide the required parameters (e.g., `student_id`):

   ```json
   {
     "student_id": "12345"
   }
   ```

7. Click **Send** to see the response.

### Example Response

If the `student_id` is valid, you will get a response like:

```json
{
   "student_id": "12345",
   "name": "John Doe",
   "dob": "2000-01-01",
   "gender": "Male",
   "nationality": "Indian",
   "aadhar_number": "1234-5678-9012",
   "contact_number": "+91-1234567890",
   "email": "johndoe@example.com",
   "permanent_address": "1234, Some Street, City, State, 123456",
   "current_address": "5678, Another Street, City, State, 123456",
   "parent_guardian_name": "Jane Doe",
   "parent_guardian_contact": "+91-9876543210",
   "previous_institution": "ABC High School",
   "previous_grade": "12th Grade",
   "marks_obtained": "85%",
   "passing_year": "2018",
   "board_name": "CBSE",
   "transfer_certificate": "None",
   "course_applied_for": "Computer Science",
   "session": "2025",
   "class_semester": "1st Semester",
   "registration_number": "CS123456",
   "admission_category": "General",
   "documents": ["birth_certificate", "12th_marksheet", "transfer_certificate"],
   "payment_mode": "Online",
   "scholarship_eligibility": "Yes",
   "loan_assistance": "No"
}
```

## Code Explanation

- **server.go**: This file contains the server implementation. The `GetStudent` method is where the student data is returned based on the `student_id` provided in the request.
  
- **student.proto**: This file defines the service `StudentService`, along with the `StudentRequest` and `Student` messages. The `StudentRequest` message is used to pass the `student_id`, while the `Student` message contains the student details.

### Example `GetStudent` Method (server.go):

```go
func (s *server) GetStudent(ctx context.Context, req *proto.StudentRequest) (*proto.Student, error) {
    // Fetch student data based on student_id (this is a hardcoded example)
    if req.GetStudentId() == "" {
        return nil, fmt.Errorf("student ID is required")
    }

    student := &proto.Student{
        StudentId: "12345",
        Name:      "John Doe",
        Dob:       "2000-01-01",
        Gender:    "Male",
        Nationality: "Indian",
        AadharNumber: "1234-5678-9012",
        ContactNumber: "+91-1234567890",
        Email: "johndoe@example.com",
        // More fields...
    }

    return student, nil
}
```
