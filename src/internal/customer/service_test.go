package customer

import (
	"database/sql"
	"database/sql/driver"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// type Suite struct {
// 	suite.Suite
// 	DB   *gorm.DB
// 	mock sqlmock.Sqlmock

// 	person *customer.Customer
// }

// func (s *Suite) SetupSuite() {
// 	var (
// 		db  *sql.DB
// 		err error
// 	)

// 	db, s.mock, err = sqlmock.New()
// 	require.NoError(s.T(), err)

// 	dialector := postgres.New(postgres.Config{
// 		DSN:                  "sqlmock_db_0",
// 		DriverName:           "postgres",
// 		Conn:                 rs.conn,
// 		PreferSimpleProtocol: true,
// 	})

// 	s.DB, err = gorm.Open(dialector, db)
// 	require.NoError(s.T(), err)

// 	// s.DB.LogMode(true)

// 	customer.InitCustomerRepository(s.DB)
// }

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

type RepositorySuite struct {
	suite.Suite
	conn *sql.DB
	DB   *gorm.DB
	mock sqlmock.Sqlmock
}

func (rs *RepositorySuite) SetupSuite() {
	var (
		err error
	)
	rs.conn, rs.mock, err = sqlmock.New()
	assert.NoError(rs.T(), err)
	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 rs.conn,
		PreferSimpleProtocol: true,
	})
	rs.DB, err = gorm.Open(dialector, &gorm.Config{})
	assert.NoError(rs.T(), err)

	InitCustomerRepository(rs.DB)
}

func (rs *RepositorySuite) TestRepositoryGet() {

	rs.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "customers" WHERE "customers"."deleted_at" IS NULL`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "first_name", "last_name", "email"}).
			AddRow(1, nil, nil, nil, nil, nil, nil))

	service := CustomerService{}
	res, err := service.getCustumers()

	require.Nil(rs.T(), err)
	require.NotNil(rs.T(), res)
}

func (rs *RepositorySuite) TestRepositoryGetById() {

	rs.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "customers" WHERE "customers"."id" = $1 AND "customers"."deleted_at" IS NULL ORDER BY "customers"."id" LIMIT 1`)).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "first_name", "last_name", "email"}).
			AddRow(1, nil, nil, nil, nil, nil, nil))

	service := CustomerService{}
	res, err := service.getCustomerById(1)

	require.Nil(rs.T(), err)
	require.NotNil(rs.T(), res)
}

func (rs *RepositorySuite) TestRepositoryCreate() {

	entity := Customer{
		FirstName: "Jhon",
		LastName:  "Contoso",
		Email:     "jhon.c@email.com",
	}

	rs.mock.ExpectBegin()
	rs.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "customers" ("created_at","updated_at","deleted_at","first_name","last_name","email") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).
		WithArgs(AnyTime{}, AnyTime{}, nil, entity.FirstName, entity.LastName, entity.Email).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	rs.mock.ExpectCommit()

	service := CustomerService{}
	res, err := service.createCustomer(entity)

	require.Nil(rs.T(), err)
	require.NotNil(rs.T(), res)
}

func (rs *RepositorySuite) TestRepositoryUpdateById() {

	entity := Customer{
		FirstName: "Jhon",
		LastName:  "Contoso",
		Email:     "jhon.c@email.com",
	}

	rs.mock.ExpectBegin()
	rs.mock.ExpectExec(regexp.QuoteMeta(
		`UPDATE "customers" SET "updated_at"=$1,"first_name"=$2,"last_name"=$3,"email"=$4 WHERE id = $5 AND "customers"."deleted_at" IS NULL`)).
		WithArgs(AnyTime{}, entity.FirstName, entity.LastName, entity.Email, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	rs.mock.ExpectCommit()

	service := CustomerService{}
	err := service.updateCustomer(1, entity)

	require.Nil(rs.T(), err)
}

func (rs *RepositorySuite) TestRepositoryDeleteById() {

	rs.mock.ExpectBegin()
	rs.mock.ExpectExec(regexp.QuoteMeta(
		`UPDATE "customers" SET "deleted_at"=$1 WHERE "customers"."id" = $2 AND "customers"."deleted_at" IS NULL`)).
		WithArgs(AnyTime{}, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	rs.mock.ExpectCommit()

	service := CustomerService{}
	err := service.deleteCustomer(1)

	require.Nil(rs.T(), err)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(RepositorySuite))
}

func TestDummyTest(t *testing.T) {
	assert.Nil(t, nil)
	// fmt.Println("RRRRRRRRRRRRRR", res)
	// fmt.Println("EEEEEEEEEEEEEE", err)
}
