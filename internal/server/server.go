package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"student-service/proto" // Ensure the import path is correct for your project structure

	"google.golang.org/grpc"
)

// server struct to implement the StudentServiceServer interface
type server struct {
	proto.UnimplementedStudentServiceServer
}

// GetStudent method to fetch student details
func (s *server) GetStudent(ctx context.Context, req *proto.StudentRequest) (*proto.Student, error) {
	// In a real-world app, you would fetch data from a database or another service here.
	// For now, we'll return hardcoded data based on the student_id.

	// Just a sample to match the student_id, you can replace it with real logic.
	if req.GetStudentId() == "" {
		return nil, fmt.Errorf("student ID is required")
	}

	// Simulating student data response
	student := &proto.Student{
		StudentId:              req.GetStudentId(),
		Name:                   "John Doe",
		Dob:                    "2000-01-01",
		Gender:                 "Male",
		Nationality:            "Indian",
		AadharNumber:           "1234-5678-9012",
		ContactNumber:          "+91-1234567890",
		Email:                  "johndoe@example.com",
		PermanentAddress:       "1234, Some Street, City, State, 123456",
		CurrentAddress:         "5678, Another Street, City, State, 123456",
		ParentGuardianName:     "Jane Doe",
		ParentGuardianContact:  "+91-9876543210",
		PreviousInstitution:    "ABC High School",
		PreviousGrade:          "12th Grade",
		MarksObtained:          "85%",
		PassingYear:            "2018",
		BoardName:              "CBSE",
		TransferCertificate:    "None",
		CourseAppliedFor:       "Computer Science",
		Session:                "2025",
		ClassSemester:          "1st Semester",
		RegistrationNumber:     "CS123456",
		AdmissionCategory:      "General",
		Documents:              []string{"birth_certificate", "12th_marksheet", "transfer_certificate"},
		PaymentMode:            "Online",
		ScholarshipEligibility: "Yes",
		LoanAssistance:         "No",
	}

	return student, nil
}

// NewServer initializes and returns a new gRPC server
func NewServer() *grpc.Server {
	s := grpc.NewServer()
	proto.RegisterStudentServiceServer(s, &server{}) // Register the service
	return s
}

// StartServer starts the gRPC server on the given port
func StartServer() {
	port := ":50051" // Choose your preferred port
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := NewServer()

	log.Printf("gRPC server listening on %v", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
