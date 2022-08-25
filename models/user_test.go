package user

import (
	"testing"

	"pocCurd/test"

	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestFetchEmploye(t *testing.T) {
	mockDB, mock := test.MockDB(t)
	defer mockDB.Close()

	mock.ExpectQuery("SELECT id, name, salary, age FROM employeedetails6 order by id").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "salary", "age"}).
			AddRow("pepe", "guerra", "pepe@gmail.com", "34"))
	subject := UserModel{
		db: mockDB,
	}
	resp, err := subject.FetchEmploye()

	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 1, len(resp))
}

func TestUpdate(t *testing.T) {
	mockDB, mock := test.MockDB(t)
	defer mockDB.Close()

	subject := UserModel{
		db: mockDB,
	}

	u := Employee{Name: "aditya", Salary: "563456", Age: "67"}

	query := "UPDATE employeedetails6 SET name=$1,salary=$2,age=$3 WHERE id=$4"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(u.Id, u.Name, u.Salary, u.Age).WillReturnResult(sqlmock.NewResult(0, 1))

	_, err := subject.PutEmploye(u)
	assert.NoError(t, err)
}

func TestPost(t *testing.T) {
	mockDB, mock := test.MockDB(t)
	defer mockDB.Close()

	subject := UserModel{
		db: mockDB,
	}
	u := Employee{Name: "aditya", Salary: "563456", Age: "67"}

	query := "INSERT INTO employeedetails6 (name, salary, age)VALUES ($1, $2, $3)"

	mock.MatchExpectationsInOrder(false)
	mock.ExpectBegin()
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(u.Name, u.Salary, u.Age).WillReturnResult(sqlmock.NewResult(0, 1))

	_, err := subject.PostEmploye(u)
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	mockDB, mock := test.MockDB(t)
	defer mockDB.Close()

	subject := UserModel{
		db: mockDB,
	}
	u := Employee{Name: "aditya", Salary: "563456", Age: "67"}

	query := "DELETE FROM employeedetails6 WHERE id =$1"
	mock.MatchExpectationsInOrder(false)
	mock.ExpectBegin()
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(u.Id).WillReturnResult(sqlmock.NewResult(0, 1))

	_, err := subject.DeleteEmploye(u.Id)
	assert.NoError(t, err)
}
