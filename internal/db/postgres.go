package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"student-service/proto"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// ✅ Initialize Database Connection
func InitDB(dataSourceName string) error {
	var err error
	DB, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}
	return DB.Ping()
}

// ✅ Create Student
func CreateStudent(student *proto.Student) error {
	query := `
		INSERT INTO students (
			student_id, name, dob, gender, nationality, aadhar_number, 
			contact_number, email, permanent_address, current_address, 
			parent_guardian_name, parent_guardian_contact, previous_institution, 
			previous_grade, marks_obtained, passing_year, board_name, 
			transfer_certificate, course_applied_for, session, 
			class_semester, registration_number, admission_category, 
			payment_mode, scholarship_eligibility, loan_assistance
		) VALUES (
			$1, $2, $3, $4, $5, $6, 
			$7, $8, $9, $10, $11, 
			$12, $13, $14, $15, $16, 
			$17, $18, $19, $20, $21, 
			$22, $23, $24, $25, $26
		)
	`

	_, err := DB.Exec(query,
		student.StudentId, student.Name, student.Dob, student.Gender, student.Nationality, student.AadharNumber,
		student.ContactNumber, student.Email, student.PermanentAddress, student.CurrentAddress,
		student.ParentGuardianName, student.ParentGuardianContact, student.PreviousInstitution,
		student.PreviousGrade, student.MarksObtained, student.PassingYear, student.BoardName,
		student.TransferCertificate, student.CourseAppliedFor, student.Session,
		student.ClassSemester, student.RegistrationNumber, student.AdmissionCategory,
		student.PaymentMode, student.ScholarshipEligibility, student.LoanAssistance,
	)

	if err != nil {
		log.Printf("Failed to insert student: %v", err)
		return err
	}

	return nil
}

// Get Student by ID
func GetStudentByID(studentID string) (*proto.Student, error) {
	query := `SELECT * FROM students WHERE student_id = $1`
	row := DB.QueryRow(query, studentID)

	var student proto.Student
	var documents sql.NullString // ✅ Use sql.NullString to handle NULL values

	err := row.Scan(
		&student.StudentId, &student.Name, &student.Dob, &student.Gender, &student.Nationality, &student.AadharNumber,
		&student.ContactNumber, &student.Email, &student.PermanentAddress, &student.CurrentAddress,
		&student.ParentGuardianName, &student.ParentGuardianContact, &student.PreviousInstitution,
		&student.PreviousGrade, &student.MarksObtained, &student.PassingYear, &student.BoardName,
		&student.TransferCertificate, &student.CourseAppliedFor, &student.Session,
		&student.ClassSemester, &student.RegistrationNumber, &student.AdmissionCategory,
		&documents, // ✅ Handle NULL values correctly
		&student.PaymentMode, &student.ScholarshipEligibility, &student.LoanAssistance,
	)
	if err != nil {
		return nil, fmt.Errorf("could not fetch student: %v", err)
	}

	// ✅ Convert documents from JSON (if it's not NULL)
	student.Documents = []string{} // Default empty array
	if documents.Valid {           // Check if documents is NOT NULL
		err = json.Unmarshal([]byte(documents.String), &student.Documents)
		if err != nil {
			log.Printf("Warning: could not parse documents JSON: %v", err)
		}
	}

	return &student, nil
}

// ✅ Update Student
func UpdateStudent(student *proto.Student) error {
	query := `
		UPDATE students SET 
			name = $2, dob = $3, gender = $4, nationality = $5, aadhar_number = $6, 
			contact_number = $7, email = $8, permanent_address = $9, current_address = $10, 
			parent_guardian_name = $11, parent_guardian_contact = $12, previous_institution = $13, 
			previous_grade = $14, marks_obtained = $15, passing_year = $16, board_name = $17, 
			transfer_certificate = $18, course_applied_for = $19, session = $20, 
			class_semester = $21, registration_number = $22, admission_category = $23, 
			payment_mode = $24, scholarship_eligibility = $25, loan_assistance = $26
		WHERE student_id = $1
	`

	_, err := DB.Exec(query,
		student.StudentId, student.Name, student.Dob, student.Gender, student.Nationality, student.AadharNumber,
		student.ContactNumber, student.Email, student.PermanentAddress, student.CurrentAddress,
		student.ParentGuardianName, student.ParentGuardianContact, student.PreviousInstitution,
		student.PreviousGrade, student.MarksObtained, student.PassingYear, student.BoardName,
		student.TransferCertificate, student.CourseAppliedFor, student.Session,
		student.ClassSemester, student.RegistrationNumber, student.AdmissionCategory,
		student.PaymentMode, student.ScholarshipEligibility, student.LoanAssistance,
	)

	return err
}

// ✅ Delete Student
func DeleteStudent(studentID string) error {
	query := `DELETE FROM students WHERE student_id = $1`
	_, err := DB.Exec(query, studentID)
	return err
}
