package main

import (
	"fmt"
	"time"
)

// AttendanceManager interface untuk mengelola absensi karyawan
type AttendanceManager interface {
	RecordAttendance(employeeID int)
}

// Employee struct untuk merepresentasikan data karyawan
type Employee struct {
	ID        int
	Name      string
	Position  string
	CreatedAt time.Time
}

// EmployeeRepository struct untuk mengelola data karyawan
type EmployeeRepository struct {
	employees []Employee
}

// AddEmployee metode untuk menambahkan karyawan baru
func (er *EmployeeRepository) AddEmployee(employee Employee) {
	er.employees = append(er.employees, employee)
}

// RemoveEmployee metode untuk menghapus karyawan berdasarkan ID
func (er *EmployeeRepository) RemoveEmployee(employeeID int) {
	for i, employee := range er.employees {
		if employee.ID == employeeID {
			er.employees = append(er.employees[:i], er.employees[i+1:]...)
			break
		}
	}
}

// FindEmployeeByID metode untuk mencari karyawan berdasarkan ID
func (er *EmployeeRepository) FindEmployeeByID(employeeID int) *Employee {
	for _, employee := range er.employees {
		if employee.ID == employeeID {
			return &employee
		}
	}
	return nil
}

// ClockInAttendanceManager struct untuk mengelola absensi karyawan dengan metode absensi masuk
type ClockInAttendanceManager struct {
	employeeRepository *EmployeeRepository
}

// RecordAttendance implementasi metode RecordAttendance untuk absensi masuk
func (ciam *ClockInAttendanceManager) RecordAttendance(employeeID int) {
	employee := ciam.employeeRepository.FindEmployeeByID(employeeID)
	if employee != nil {
		fmt.Printf("Absensi masuk berhasil: %s\n", employee.Name)
	} else {
		fmt.Println("Karyawan tidak ditemukan")
	}
}

// ClockOutAttendanceManager struct untuk mengelola absensi karyawan dengan metode absensi keluar
type ClockOutAttendanceManager struct {
	employeeRepository *EmployeeRepository
}

// RecordAttendance implementasi metode RecordAttendance untuk absensi keluar
func (coam *ClockOutAttendanceManager) RecordAttendance(employeeID int) {
	employee := coam.employeeRepository.FindEmployeeByID(employeeID)
	if employee != nil {
		fmt.Printf("Absensi keluar berhasil: %s\n", employee.Name)
	} else {
		fmt.Println("Karyawan tidak ditemukan")
	}
}

func main() {
	employeeRepo := &EmployeeRepository{}

	attendanceManagers := []AttendanceManager{
		&ClockInAttendanceManager{employeeRepository: employeeRepo},
		&ClockOutAttendanceManager{employeeRepository: employeeRepo},
	}

	employee1 := Employee{ID: 1, Name: "John Doe", Position: "Manager", CreatedAt: time.Now()}
	employee2 := Employee{ID: 2, Name: "Jane Smith", Position: "Staff", CreatedAt: time.Now()}

	employeeRepo.AddEmployee(employee1)
	employeeRepo.AddEmployee(employee2)

	// Rekam absensi masuk
	for _, manager := range attendanceManagers {
		manager.RecordAttendance(1)
	}

	// Rekam absensi keluar
	for _, manager := range attendanceManagers {
		manager.RecordAttendance(2)
	}

	// Rekam absensi dengan ID karyawan yang tidak
	for _, manager := range attendanceManagers {
		manager.RecordAttendance(2)
	}

	// Rekam absensi dengan ID karyawan yang tidak ada
	for _, manager := range attendanceManagers {
		manager.RecordAttendance(3)
	}
}
