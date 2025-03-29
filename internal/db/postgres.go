package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

// Student struct to map student data from the database
type Student struct {
	StudentID              string
	Name                   string
	Dob                    string
	Gender                 string
	Nationality            string
	AadharNumber           string
	ContactNumber          string
	Email                  string
	PermanentAddress       string
	CurrentAddress         string
	ParentGuardianName     string
	ParentGuardianContact  string
	PreviousInstitution    string
	PreviousGrade          string
	MarksObtained          string
	PassingYear            string
	BoardName              string
	TransferCertificate    string
	CourseAppliedFor       string
	Session                string
	ClassSemester         string
	RegistrationNumber     string
	AdmissionCategory      string
	Documents              []string
	PaymentMode            string
	ScholarshipEligibility string
	LoanAssistance         string
}

// ConnectDB connects to the PostgreSQL database
func ConnectDB() (*pgx.Conn, error) {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatalf("DATABASE_URL environment variable not set")
		return nil, fmt.Errorf("database URL not set")
	}

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
		return nil, err
	}
	return conn, nil
}

// GetStudentByID fetches a student by ID from the database
func GetStudentByID(studentID string) (*Student, error) {
	conn, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close(context.Background())

	var student Student
	query := `SELECT * FROM students WHERE student_id=$1`
	err = conn.QueryRow(context.Background(), query, studentID).Scan(
		&student.StudentID,
		&student.Name,
		&student.Dob,
		&student.Gender,
		&student.Nationality,
		&student.AadharNumber,
		&student.ContactNumber,
		&student.Email,
		&student.PermanentAddress,
		&student.CurrentAddress,
		&student.ParentGuardianName,
		&student.ParentGuardianContact,
		&student.PreviousInstitution,
		&student.PreviousGrade,
		&student.MarksObtained,
		&student.PassingYear,
		&student.BoardName,
		&student.TransferCertificate,
		&student.CourseAppliedFor,
		&student.Session,
		&student.ClassSemester,
		&student.RegistrationNumber,
		&student.AdmissionCategory,
		&student.Documents,
		&student.PaymentMode,
		&student.ScholarshipEligibility,
		&student.LoanAssistance,
	)
	if err != nil {
		return nil, fmt.Errorf("could not fetch student: %v", err)
	}

	return &student, nil
}
