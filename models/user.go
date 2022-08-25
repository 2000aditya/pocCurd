package user

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type (
	UserModelImpl interface {
		PostEmploye(e Employee) (bool, error)
		PutEmploye(e Employee) (bool, error)
		DeleteEmploye(id string) (bool, error)
		FetchEmploye() ([]Employee, error)

		//FindAll() []User
	}

	Employee struct {
		Id     string `json:"id"`
		Name   string `json:"name"`
		Salary string `json:"salary"`
		Age    string `json:"age"`
	}
	//Employees struct {
	//	Employees []Employee `json:"employees"`
	//}

	UserModel struct {
		db *sql.DB
	}
)

func NewUserModel(db *sql.DB) *UserModel {
	return &UserModel{
		db: db,
	}
}

func (u *UserModel) PostEmploye(e Employee) (success bool, err error) {
	sqlStatement := "INSERT INTO employeedetails6 (name, salary, age)VALUES ($1, $2, $3) "

	_, err = u.db.Query(sqlStatement, e.Name, e.Salary, e.Age)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func (u *UserModel) PutEmploye(e Employee) (success bool, err error) {
	sqlStatement := "UPDATE employeedetails6 SET name=$1,salary=$2,age=$3 WHERE id=$4"

	_, err = u.db.Query(sqlStatement, e.Name, e.Salary, e.Age, e.Id)
	if err != nil {
		return true, nil
	} else {
		return true, nil
	}
}

func (u *UserModel) DeleteEmploye(id string) (success bool, err error) {
	sqlStatement := "DELETE FROM employeedetails6 WHERE id =$1"

	_, err = u.db.Query(sqlStatement, id)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func (u *UserModel) FetchEmploye() ([]Employee, error) {
	// An albums slice to hold data from returned rows.
	var employee []Employee
	sqlStatement := "SELECT id, name, salary, age FROM employeedetails6 order by id"

	rows, err := u.db.Query(sqlStatement)
	if err != nil {

		return nil, err
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var emp Employee
		if err := rows.Scan(&emp.Id, &emp.Name, &emp.Salary, &emp.Age); err != nil {

			return nil, err
		}

		employee = append(employee, emp)
	}
	if err := rows.Err(); err != nil {

		return nil, err
	}

	return employee, nil
}
