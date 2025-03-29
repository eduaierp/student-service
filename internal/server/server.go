package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"student-service/internal/db"
	"student-service/proto"
	"syscall"

	"google.golang.org/grpc"
)

// server struct to implement the StudentServiceServer interface
type server struct {
	proto.UnimplementedStudentServiceServer
}

// GetStudent method to fetch student details from the database
func (s *server) GetStudent(ctx context.Context, req *proto.StudentRequest) (*proto.Student, error) {
	student, err := db.GetStudentByID(req.GetStudentId())
	if err != nil {
		return nil, fmt.Errorf("could not fetch student with ID %v: %v", req.GetStudentId(), err)
	}

	// Return student data as gRPC response
	return &proto.Student{
		StudentId:              student.StudentID,
		Name:                   student.Name,
		Dob:                    student.Dob,
		Gender:                 student.Gender,
		Nationality:            student.Nationality,
		AadharNumber:           student.AadharNumber,
		ContactNumber:          student.ContactNumber,
		Email:                  student.Email,
		PermanentAddress:       student.PermanentAddress,
		CurrentAddress:         student.CurrentAddress,
		ParentGuardianName:     student.ParentGuardianName,
		ParentGuardianContact:  student.ParentGuardianContact,
		PreviousInstitution:    student.PreviousInstitution,
		PreviousGrade:          student.PreviousGrade,
		MarksObtained:          student.MarksObtained,
		PassingYear:            student.PassingYear,
		BoardName:              student.BoardName,
		TransferCertificate:    student.TransferCertificate,
		CourseAppliedFor:       student.CourseAppliedFor,
		Session:                student.Session,
		ClassSemester:          student.ClassSemester,
		RegistrationNumber:     student.RegistrationNumber,
		AdmissionCategory:      student.AdmissionCategory,
		Documents:              student.Documents,
		PaymentMode:            student.PaymentMode,
		ScholarshipEligibility: student.ScholarshipEligibility,
		LoanAssistance:         student.LoanAssistance,
	}, nil
}

// NewServer initializes a new gRPC server
func NewServer() *grpc.Server {
	s := grpc.NewServer()
	proto.RegisterStudentServiceServer(s, &server{}) 
	return s
}

// StartServer starts the gRPC server
func StartServer() {
	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := NewServer()

	go func() {
		log.Printf("gRPC server listening on %v", port)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Handle graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
	s.GracefulStop()
}
