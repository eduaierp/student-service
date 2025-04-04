// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.29.3
// source: proto/student.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StudentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentId string `protobuf:"bytes,1,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
}

func (x *StudentRequest) Reset() {
	*x = StudentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_student_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StudentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudentRequest) ProtoMessage() {}

func (x *StudentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_student_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudentRequest.ProtoReflect.Descriptor instead.
func (*StudentRequest) Descriptor() ([]byte, []int) {
	return file_proto_student_proto_rawDescGZIP(), []int{0}
}

func (x *StudentRequest) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

type EmptyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRequest) Reset() {
	*x = EmptyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_student_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRequest) ProtoMessage() {}

func (x *EmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_student_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRequest.ProtoReflect.Descriptor instead.
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return file_proto_student_proto_rawDescGZIP(), []int{1}
}

type StudentList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Students []*Student `protobuf:"bytes,1,rep,name=students,proto3" json:"students,omitempty"`
}

func (x *StudentList) Reset() {
	*x = StudentList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_student_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StudentList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudentList) ProtoMessage() {}

func (x *StudentList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_student_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudentList.ProtoReflect.Descriptor instead.
func (*StudentList) Descriptor() ([]byte, []int) {
	return file_proto_student_proto_rawDescGZIP(), []int{2}
}

func (x *StudentList) GetStudents() []*Student {
	if x != nil {
		return x.Students
	}
	return nil
}

type StudentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *StudentResponse) Reset() {
	*x = StudentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_student_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StudentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudentResponse) ProtoMessage() {}

func (x *StudentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_student_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudentResponse.ProtoReflect.Descriptor instead.
func (*StudentResponse) Descriptor() ([]byte, []int) {
	return file_proto_student_proto_rawDescGZIP(), []int{3}
}

func (x *StudentResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Student struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentId              string   `protobuf:"bytes,1,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	Name                   string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Dob                    string   `protobuf:"bytes,3,opt,name=dob,proto3" json:"dob,omitempty"`
	Gender                 string   `protobuf:"bytes,4,opt,name=gender,proto3" json:"gender,omitempty"`
	Nationality            string   `protobuf:"bytes,5,opt,name=nationality,proto3" json:"nationality,omitempty"`
	AadharNumber           string   `protobuf:"bytes,6,opt,name=aadhar_number,json=aadharNumber,proto3" json:"aadhar_number,omitempty"`
	ContactNumber          string   `protobuf:"bytes,7,opt,name=contact_number,json=contactNumber,proto3" json:"contact_number,omitempty"`
	Email                  string   `protobuf:"bytes,8,opt,name=email,proto3" json:"email,omitempty"`
	PermanentAddress       string   `protobuf:"bytes,9,opt,name=permanent_address,json=permanentAddress,proto3" json:"permanent_address,omitempty"`
	CurrentAddress         string   `protobuf:"bytes,10,opt,name=current_address,json=currentAddress,proto3" json:"current_address,omitempty"`
	ParentGuardianName     string   `protobuf:"bytes,11,opt,name=parent_guardian_name,json=parentGuardianName,proto3" json:"parent_guardian_name,omitempty"`
	ParentGuardianContact  string   `protobuf:"bytes,12,opt,name=parent_guardian_contact,json=parentGuardianContact,proto3" json:"parent_guardian_contact,omitempty"`
	PreviousInstitution    string   `protobuf:"bytes,13,opt,name=previous_institution,json=previousInstitution,proto3" json:"previous_institution,omitempty"`
	PreviousGrade          string   `protobuf:"bytes,14,opt,name=previous_grade,json=previousGrade,proto3" json:"previous_grade,omitempty"`
	MarksObtained          string   `protobuf:"bytes,15,opt,name=marks_obtained,json=marksObtained,proto3" json:"marks_obtained,omitempty"`
	PassingYear            string   `protobuf:"bytes,16,opt,name=passing_year,json=passingYear,proto3" json:"passing_year,omitempty"`
	BoardName              string   `protobuf:"bytes,17,opt,name=board_name,json=boardName,proto3" json:"board_name,omitempty"`
	TransferCertificate    string   `protobuf:"bytes,18,opt,name=transfer_certificate,json=transferCertificate,proto3" json:"transfer_certificate,omitempty"`
	CourseAppliedFor       string   `protobuf:"bytes,19,opt,name=course_applied_for,json=courseAppliedFor,proto3" json:"course_applied_for,omitempty"`
	Session                string   `protobuf:"bytes,20,opt,name=session,proto3" json:"session,omitempty"`
	ClassSemester          string   `protobuf:"bytes,21,opt,name=class_semester,json=classSemester,proto3" json:"class_semester,omitempty"`
	RegistrationNumber     string   `protobuf:"bytes,22,opt,name=registration_number,json=registrationNumber,proto3" json:"registration_number,omitempty"`
	AdmissionCategory      string   `protobuf:"bytes,23,opt,name=admission_category,json=admissionCategory,proto3" json:"admission_category,omitempty"`
	Documents              []string `protobuf:"bytes,24,rep,name=documents,proto3" json:"documents,omitempty"`
	PaymentMode            string   `protobuf:"bytes,25,opt,name=payment_mode,json=paymentMode,proto3" json:"payment_mode,omitempty"`
	ScholarshipEligibility string   `protobuf:"bytes,26,opt,name=scholarship_eligibility,json=scholarshipEligibility,proto3" json:"scholarship_eligibility,omitempty"`
	LoanAssistance         string   `protobuf:"bytes,27,opt,name=loan_assistance,json=loanAssistance,proto3" json:"loan_assistance,omitempty"`
}

func (x *Student) Reset() {
	*x = Student{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_student_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Student) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Student) ProtoMessage() {}

func (x *Student) ProtoReflect() protoreflect.Message {
	mi := &file_proto_student_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Student.ProtoReflect.Descriptor instead.
func (*Student) Descriptor() ([]byte, []int) {
	return file_proto_student_proto_rawDescGZIP(), []int{4}
}

func (x *Student) GetStudentId() string {
	if x != nil {
		return x.StudentId
	}
	return ""
}

func (x *Student) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Student) GetDob() string {
	if x != nil {
		return x.Dob
	}
	return ""
}

func (x *Student) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *Student) GetNationality() string {
	if x != nil {
		return x.Nationality
	}
	return ""
}

func (x *Student) GetAadharNumber() string {
	if x != nil {
		return x.AadharNumber
	}
	return ""
}

func (x *Student) GetContactNumber() string {
	if x != nil {
		return x.ContactNumber
	}
	return ""
}

func (x *Student) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Student) GetPermanentAddress() string {
	if x != nil {
		return x.PermanentAddress
	}
	return ""
}

func (x *Student) GetCurrentAddress() string {
	if x != nil {
		return x.CurrentAddress
	}
	return ""
}

func (x *Student) GetParentGuardianName() string {
	if x != nil {
		return x.ParentGuardianName
	}
	return ""
}

func (x *Student) GetParentGuardianContact() string {
	if x != nil {
		return x.ParentGuardianContact
	}
	return ""
}

func (x *Student) GetPreviousInstitution() string {
	if x != nil {
		return x.PreviousInstitution
	}
	return ""
}

func (x *Student) GetPreviousGrade() string {
	if x != nil {
		return x.PreviousGrade
	}
	return ""
}

func (x *Student) GetMarksObtained() string {
	if x != nil {
		return x.MarksObtained
	}
	return ""
}

func (x *Student) GetPassingYear() string {
	if x != nil {
		return x.PassingYear
	}
	return ""
}

func (x *Student) GetBoardName() string {
	if x != nil {
		return x.BoardName
	}
	return ""
}

func (x *Student) GetTransferCertificate() string {
	if x != nil {
		return x.TransferCertificate
	}
	return ""
}

func (x *Student) GetCourseAppliedFor() string {
	if x != nil {
		return x.CourseAppliedFor
	}
	return ""
}

func (x *Student) GetSession() string {
	if x != nil {
		return x.Session
	}
	return ""
}

func (x *Student) GetClassSemester() string {
	if x != nil {
		return x.ClassSemester
	}
	return ""
}

func (x *Student) GetRegistrationNumber() string {
	if x != nil {
		return x.RegistrationNumber
	}
	return ""
}

func (x *Student) GetAdmissionCategory() string {
	if x != nil {
		return x.AdmissionCategory
	}
	return ""
}

func (x *Student) GetDocuments() []string {
	if x != nil {
		return x.Documents
	}
	return nil
}

func (x *Student) GetPaymentMode() string {
	if x != nil {
		return x.PaymentMode
	}
	return ""
}

func (x *Student) GetScholarshipEligibility() string {
	if x != nil {
		return x.ScholarshipEligibility
	}
	return ""
}

func (x *Student) GetLoanAssistance() string {
	if x != nil {
		return x.LoanAssistance
	}
	return ""
}

var File_proto_student_proto protoreflect.FileDescriptor

var file_proto_student_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x22, 0x2f,
	0x0a, 0x0e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22,
	0x0e, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x3b, 0x0a, 0x0b, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2c,
	0x0a, 0x08, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x52, 0x08, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x2b, 0x0a, 0x0f,
	0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x92, 0x08, 0x0a, 0x07, 0x53, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x6f, 0x62, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x6f, 0x62, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74,
	0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x61, 0x64, 0x68, 0x61, 0x72, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x61, 0x64,
	0x68, 0x61, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x2b, 0x0a, 0x11, 0x70, 0x65, 0x72, 0x6d, 0x61, 0x6e,
	0x65, 0x6e, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x70, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x30, 0x0a, 0x14,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x67, 0x75, 0x61, 0x72, 0x64, 0x69, 0x61, 0x6e, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x47, 0x75, 0x61, 0x72, 0x64, 0x69, 0x61, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x36,
	0x0a, 0x17, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x67, 0x75, 0x61, 0x72, 0x64, 0x69, 0x61,
	0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x15, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x47, 0x75, 0x61, 0x72, 0x64, 0x69, 0x61, 0x6e, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x31, 0x0a, 0x14, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f,
	0x75, 0x73, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x49, 0x6e,
	0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72, 0x65,
	0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x67, 0x72, 0x61, 0x64, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x47, 0x72, 0x61, 0x64, 0x65,
	0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x5f, 0x6f, 0x62, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x4f,
	0x62, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x73, 0x73, 0x69,
	0x6e, 0x67, 0x5f, 0x79, 0x65, 0x61, 0x72, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x61, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x59, 0x65, 0x61, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x14, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x12,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x5f, 0x66,
	0x6f, 0x72, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x73, 0x65,
	0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x53, 0x65, 0x6d, 0x65, 0x73, 0x74, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x13, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x12,
	0x61, 0x64, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x61, 0x64, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x18, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09,
	0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x19, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x37, 0x0a, 0x17,
	0x73, 0x63, 0x68, 0x6f, 0x6c, 0x61, 0x72, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x65, 0x6c, 0x69, 0x67,
	0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x73,
	0x63, 0x68, 0x6f, 0x6c, 0x61, 0x72, 0x73, 0x68, 0x69, 0x70, 0x45, 0x6c, 0x69, 0x67, 0x69, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x6c, 0x6f, 0x61, 0x6e, 0x5f, 0x61, 0x73,
	0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x6c, 0x6f, 0x61, 0x6e, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x32, 0xc6,
	0x02, 0x0a, 0x0e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3d, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x15, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x73, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x37, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x17,
	0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x3b, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x2e, 0x73, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x1a, 0x18, 0x2e, 0x73,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x2e, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x53,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_student_proto_rawDescOnce sync.Once
	file_proto_student_proto_rawDescData = file_proto_student_proto_rawDesc
)

func file_proto_student_proto_rawDescGZIP() []byte {
	file_proto_student_proto_rawDescOnce.Do(func() {
		file_proto_student_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_student_proto_rawDescData)
	})
	return file_proto_student_proto_rawDescData
}

var file_proto_student_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_student_proto_goTypes = []interface{}{
	(*StudentRequest)(nil),  // 0: student.StudentRequest
	(*EmptyRequest)(nil),    // 1: student.EmptyRequest
	(*StudentList)(nil),     // 2: student.StudentList
	(*StudentResponse)(nil), // 3: student.StudentResponse
	(*Student)(nil),         // 4: student.Student
}
var file_proto_student_proto_depIdxs = []int32{
	4, // 0: student.StudentList.students:type_name -> student.Student
	1, // 1: student.StudentService.GetAllStudents:input_type -> student.EmptyRequest
	0, // 2: student.StudentService.GetStudent:input_type -> student.StudentRequest
	4, // 3: student.StudentService.CreateStudent:input_type -> student.Student
	4, // 4: student.StudentService.UpdateStudent:input_type -> student.Student
	0, // 5: student.StudentService.DeleteStudent:input_type -> student.StudentRequest
	2, // 6: student.StudentService.GetAllStudents:output_type -> student.StudentList
	4, // 7: student.StudentService.GetStudent:output_type -> student.Student
	3, // 8: student.StudentService.CreateStudent:output_type -> student.StudentResponse
	3, // 9: student.StudentService.UpdateStudent:output_type -> student.StudentResponse
	3, // 10: student.StudentService.DeleteStudent:output_type -> student.StudentResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_student_proto_init() }
func file_proto_student_proto_init() {
	if File_proto_student_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_student_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StudentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_student_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_student_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StudentList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_student_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StudentResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_student_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Student); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_student_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_student_proto_goTypes,
		DependencyIndexes: file_proto_student_proto_depIdxs,
		MessageInfos:      file_proto_student_proto_msgTypes,
	}.Build()
	File_proto_student_proto = out.File
	file_proto_student_proto_rawDesc = nil
	file_proto_student_proto_goTypes = nil
	file_proto_student_proto_depIdxs = nil
}
