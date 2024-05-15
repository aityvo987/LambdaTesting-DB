package repository

import (
	"database/sql"
	"errors"
	"fmt"
	pkg "lambda-test/internal/pkg"
	model "lambda-test/internal/repository/model"
	"log"
	"sync"

	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

const (
	host     = "database-employee.cby6yekiafvf.ap-southeast-1.rds.amazonaws.com"
	port     = 5432
	user     = "postgres"
	password = "4g5hYObaPwvI7NAFODXc"
	dbname   = "postgres"
)

var (
	database *sql.DB
	once     sync.Once
)

type EmployeeRepository interface {
	ReadEmployee(id string) (model.Employee, error)
	ReadAllEmployee(id int) (model.Employee, error)
	CreateEmployee(e model.Employee) error
	Update(id int, e model.Employee) error
	DeleteEmployee(id int) error
}

func ConnectDB() (*sql.DB, error) {
	once.Do(func() {
		db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
		if err != nil {
			log.Printf("Error connect to database: %s\n", err)
			return
		}
		database = db
	})
	return database, nil
}

func findEmployee(db *sql.DB, id string) error {
	row := db.QueryRow("SELECT id FROM Employee WHERE id = $1", id)
	var existingID int
	err := row.Scan(&existingID)
	if err != nil {
		log.Printf("Error in finding Emmployee ID %s: %s\n", id, err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New(pkg.RecNoF.Code())
			return err
		}
		err = errors.New(pkg.DbError.Code())
		return err
	}
	return nil
}

func ReadEmployee(id string) (model.Employee, error) {
	var e model.Employee
	e.Id = id
	db, err := ConnectDB()
	if err != nil {
		log.Printf("Error connect to database in Read Employee %s: %s\n", id, err)
		err = errors.New(pkg.DbError.Code())
		return e, err
	}
	row := db.QueryRow("SELECT name,dob, email,phone,citizenId,address FROM Employee WHERE id=$1", e.Id)
	err = row.Scan(&e.Name, &e.Dob, &e.Email, &e.Phone, &e.CitizenId, &e.Address)
	if err != nil {
		log.Printf("Error get Employee %s, %s\n", id, err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New(pkg.RecNoF.Code())
			return e, err
		}
		err = errors.New(pkg.DbError.Code())
		return e, err
	}
	return e, nil
}
func ReadAllEmployee() ([]model.Employee, error) {
	var employees []model.Employee
	db, err := ConnectDB()
	if err != nil {
		log.Printf("Error connect to database in Read Employees: %s\n", err)
		err = errors.New(pkg.DbError.Code())
		return employees, err
	}

	rows, err := db.Query("SELECT id, name,dob, email,phone,citizenId,address FROM Employee")
	if err != nil {
		log.Printf("Error get Employees %s\n", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New(pkg.RecNoF.Code())
			return nil, err
		}

		err = errors.New(pkg.DbError.Code())
		return nil, err
	}
	defer rows.Close()
	count := 0
	for rows.Next() {
		var e model.Employee
		err := rows.Scan(&e.Id, &e.Name, &e.Dob, &e.Email, &e.Phone, &e.CitizenId, &e.Address)
		if err != nil {
			log.Printf("Error scanning Employee number %v: %s\n", count, err)
			count++
			err = errors.New(pkg.DbError.Code())
			return nil, err
		}
		employees = append(employees, e)
	}
	return employees, nil
}

func CreateEmployee(e model.Employee) error {
	db, err := ConnectDB()
	if err != nil {
		log.Printf("Error connect to database in Create Employee: %s\n", err)
		err = errors.New(pkg.DbError.Code())
		return err
	}
	_, err = db.Exec("INSERT INTO Employee(name, dob, email, phone, citizenId, address) VALUES($1, $2,$3,$4,$5,$6)", e.Name, e.Dob, e.Email, e.Phone, e.CitizenId, e.Address)
	if err != nil {
		log.Printf("Error insert to database in Create Employee: %s\n", err)
		err = errors.New(pkg.DbError.Code())
		return err
	}
	return nil
}

func UpdateEmployee(id string, e model.Employee) error {
	db, err := ConnectDB()
	if err != nil {
		log.Printf("Error connect to database in Update Employee %s: %s\n", id, err)
		err = errors.New(pkg.DbError.Code())
		return err
	}
	err = findEmployee(db, e.Id)
	if err != nil {
		return err
	}
	_, err = db.Exec("UPDATE Employee SET name=$1, dob=$2, email=$3, phone=$4, citizenId=$5, address=$6 WHERE id=$7", e.Name, e.Dob, e.Email, e.Phone, e.CitizenId, e.Address, id)
	if err != nil {
		log.Printf("Error insert to database in Create Employee %s: %s\n", id, err)
		err = errors.New(pkg.DbError.Code())
		return err
	}
	return nil
}

func DeleteEmployee(id string) error {
	db, err := ConnectDB()
	if err != nil {
		log.Printf("Error connect to database in Delete Employee %s: %s\n", id, err)
		err = errors.New(pkg.DbError.Code())
		return err
	}
	err = findEmployee(db, id)
	if err != nil {
		return err
	}
	_, err = db.Exec("DELETE FROM EMPLOYEE WHERE id=$1", id)
	if err != nil {
		log.Printf("Error insert to database in Delete Employee %s: %s\n", id, err)
		err = errors.New(pkg.DbError.Code())
		return err
	}
	return nil
}
