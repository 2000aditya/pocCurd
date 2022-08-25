package user

import (
	"testing"

	"pocCurd/test"

	"github.com/stretchr/testify/assert"
)

func TestPostEmploye(t *testing.T) {
	mockDB, mock := test.MockDB(t)
	defer mockDB.Close()
	u := Employee{Name: "aditya", Salary: "563456", Age: "67"}

	//mock.ExpectQuery("INSERT INTO employeedetails6 (name, salary, age)VALUES ($1, $2, $3)").WithArgs("aditya", "56", "567").WillReturnRows()
	mock.ExpectQuery("INSERT INTO employeedetails6 (name, salary, age)VALUES ($1, $2, $3)")
	um := NewUserModel(mockDB)
	sucess, err := um.PostEmploye(Employee{Name: "aditya", Salary: "56", Age: "567"})
	//assert.Equal(t, nil, err)
	assert.Equal(t, sucess, u)
}
