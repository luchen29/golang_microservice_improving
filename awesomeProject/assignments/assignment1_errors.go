package assignments

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
)

// assignment description: when we operate the database, like: when dao returns a sql.ErrNoRows,
// do we need to Wrap this error and return to the call stack? and why?

// Answer: I think no, since "no error fount" in real production environment is not an error, the metrics will send alarm
// if the number of error logs increases a lot in a short duration.

// implementation as follows:

type TestDbClient struct {
	ReadDb *sql.DB
	WriteDb *sql.DB
	TableName string
}

// call Init() in main.go/ server.go
func Init() {
	NewTestDbClient()
}

func NewDbTableClient(writeDb *sql.DB, readDb *sql.DB, tableName string) (*TestDbClient, error) {
	// .....
	// simplify db client creation and connection etc. config logic here ...
	testDbClient := &TestDbClient{
		ReadDb: readDb,
		WriteDb: writeDb,
		TableName: tableName,
	}
	return testDbClient, nil
}

func NewTestDbClient() (*TestDbClient, error) {
	testTableName := "test_table_name"
	testDbClient, err := NewDbTableClient(nil, nil,testTableName)
	if err != nil {
		return nil, errors.Wrapf(err, "new db table client error: %v", testTableName)
	}
	return testDbClient, nil
}

func (t *TestDbClient) GetSthBySth(ctx context.Context, name string) (string, error) {
	//readTable := t.ReadDb.Context(ctx)
	resultRow, err := t.ReadDb.Query("name = ?", name)
	if err != nil {
		if err == sql.ErrNoRows {
			//logs.CtxInfo(ctx, "no rows found for table name: %v", name)
			return "", nil
		}
		return "", errors.Wrapf(err, "err searching for name: %v", name)
	}
	colList, err := resultRow.Columns()
	if err != nil {
		return "", err
	}
	if len(colList) < 1 {
		return "", errors.New("all columns are empty for name")
	}
	return colList[0], nil
}