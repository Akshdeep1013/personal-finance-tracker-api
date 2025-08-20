package repository

//1.what database, and its language(driver)
//2 Where is it local? or remote? port
//3.who am i for that, basically my identity?
//4.Which room(database) i need to go
//5.Be a good citizen by not destroying resources, means how will i use that connection? like, when to close, how many connections?

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"personal-finance-tracker-api/model"
)

type Operation interface {
	CreateTransaction(req *model.Transaction) error
	GetTransaction(id string) (*model.Transaction, error)
	GetAllTransaction() ([]*model.Transaction, error)
	UpdateTransaction(req *model.Transaction) error
	DeleteTransaction(id string) error

	CreateCategory(req model.Category) error
	GetCategory(id int) (*model.Category, error)
	GetAllCategory() (*[]model.Category, error)
	UpdateCategory(req *model.Category) error
	DeleteCategory(id int) error
}

type Repository struct {
	db *sql.DB
}

func NewRepository() (*Repository, error) {
	//dsn := "root:your_password@tcp(localhost:3306)/finance_tracker"
	//lets have a data source name
	dsn := "root:MyPass123!@tcp(localhost:3306)/"

	//make connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	//create database if not exist
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS finance_tracker")
	if err != nil {
		return nil, err
	}

	//lets close old connection becuase we can't create tables as it is utside any database
	db.Close()

	//lets reconnect to specifies database
	dsn = "root:MyPass123!@tcp(localhost:3306)/finance_tracker"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	r := &Repository{db: db}
	err = r.createTables()

	if err != nil {
		return nil, err
	}
	fmt.Printf("\n Tables has been created !!!\n")

	return r, nil
}

func (r *Repository) createTables() error {
	table := "CREATE TABLE IF NOT EXISTS transaction (" +
		"uuid VARCHAR(36) PRIMARY KEY," +
		"type VARCHAR(16) NOT NULL," +
		"amount INT," +
		"category_id INT," +
		"description TEXT," +
		"transaction_at TIMESTAMP," +
		"created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP," +
		"updated_at TIMESTAMP" +
		");"

	_, err := r.db.Exec(table)
	if err != nil {
		return err
	}

	table = "CREATE TABLE IF NOT EXISTS category (" +
		"id INT AUTO_INCREMENT PRIMARY KEY ," +
		"name VARCHAR(32) NOT NULL UNIQUE ," +
		"description TEXT ," +
		"created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ," +
		"updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP " +
		");"

	_, err = r.db.Exec(table)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) CreateTransaction(req *model.Transaction) error {
	query := "INSERT INTO transaction(uuid,type,amount,category_id,description,transaction_at,created_at,created_at,updated_at)" +
		"VALUES(?, ?, ?, ?, ?, ?, ?, ?)"
	param := []interface{
		req.UUID,
		req.Type,
		req.Amount,
		req.Category
	}

	return nil
}

func (r *Repository) GetTransaction(id string) (*model.Transaction, error) {

	return nil, nil
}

func (r *Repository) GetAllTransaction() ([]*model.Transaction, error) {

	return nil, nil
}

func (r *Repository) UpdateTransaction(req *model.Transaction) error {

	return nil
}

func (r *Repository) DeleteTransaction(id string) error {

	return nil
}

func (r *Repository) CreateCategory(req model.Category) error {

	return nil
}

func (r *Repository) GetCategory(id int) (*model.Category, error) {

	return nil, nil
}

func (r *Repository) GetAllCategory() (*[]model.Category, error) {

	return nil, nil
}

func (r *Repository) UpdateCategory(req *model.Category) error {

	return nil
}

func (r *Repository) DeleteCategory(id int) error {

	return nil
}
