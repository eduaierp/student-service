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

type server struct {
	proto.UnimplementedStudentServiceServer
}

// ✅ Create Student
func (s *server) CreateStudent(ctx context.Context, req *proto.Student) (*proto.StudentResponse, error) {
	log.Printf("Received new student: %+v", req)

	err := db.CreateStudent(req)
	if err != nil {
		return nil, fmt.Errorf("could not create student: %v", err)
	}

	message := fmt.Sprintf("Student %s created successfully!", req.Name)
	return &proto.StudentResponse{Message: message}, nil
}

// ✅ Get Student
func (s *server) GetStudent(ctx context.Context, req *proto.StudentRequest) (*proto.Student, error) {
	student, err := db.GetStudentByID(req.GetStudentId())
	if err != nil {
		return nil, fmt.Errorf("could not fetch student: %v", err)
	}

	return student, nil
}

// ✅ Update Student
func (s *server) UpdateStudent(ctx context.Context, req *proto.Student) (*proto.StudentResponse, error) {
	err := db.UpdateStudent(req)
	if err != nil {
		return nil, fmt.Errorf("could not update student: %v", err)
	}

	message := fmt.Sprintf("Student %s updated successfully!", req.Name)
	return &proto.StudentResponse{Message: message}, nil
}

// ✅ Delete Student
func (s *server) DeleteStudent(ctx context.Context, req *proto.StudentRequest) (*proto.StudentResponse, error) {
	err := db.DeleteStudent(req.GetStudentId())
	if err != nil {
		return nil, fmt.Errorf("could not delete student: %v", err)
	}

	message := fmt.Sprintf("Student %s deleted successfully!", req.StudentId)
	return &proto.StudentResponse{Message: message}, nil
}

// StartServer function
func StartServer() {
	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterStudentServiceServer(s, &server{})

	go func() {
		log.Printf("gRPC server listening on %v", port)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
	s.GracefulStop()
}
