package test_util

import (
	"context"
	"fmt"
	"go_learning/db"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/mock"
)

// MockRows implements pgx.Rows for testing
type MockRows struct {
	mock.Mock
	Data         [][]interface{}
	currentIndex int
}

func (m *MockRows) Next() bool {
	// Call the mock system for verification
	m.Called()

	// Check if there are more rows
	hasMore := m.currentIndex < len(m.Data)

	// Move to next row if available
	if hasMore {
		m.currentIndex++
	}

	return hasMore
}

func (m *MockRows) Scan(dest ...interface{}) error {
	// Call the mock system for verification
	result := m.Called(mock.Anything)

	// If set to return error, do so
	if result.Error(0) != nil {
		return result.Error(0)
	}

	// If no data or out of bounds, return
	if len(m.Data) == 0 || m.currentIndex <= 0 || m.currentIndex > len(m.Data) {
		return nil
	}

	// Get the current row data
	rowData := m.Data[m.currentIndex-1]

	// Debug output
	fmt.Printf("Scanning row %d: %v into %d destinations\n", m.currentIndex, rowData, len(dest))

	// Copy data to destination pointers
	for i := 0; i < len(rowData) && i < len(dest); i++ {
		switch v := dest[i].(type) {
		case *string:
			if str, ok := rowData[i].(string); ok {
				*v = str
				fmt.Printf("  Copied string '%s' to destination %d\n", str, i)
			} else {
				fmt.Printf("  Failed to copy to string: %v is not a string\n", rowData[i])
			}
		case *float64:
			switch val := rowData[i].(type) {
			case float64:
				*v = val
				fmt.Printf("  Copied float64 %f to destination %d\n", val, i)
			case float32:
				*v = float64(val)
			case int:
				*v = float64(val)
			default:
				fmt.Printf("  Failed to copy to float64: %v is not a number\n", rowData[i])
			}
		default:
			fmt.Printf("  Unsupported type for destination %d: %T\n", i, dest[i])
		}
	}

	return nil
}

func (m *MockRows) Close() {
	m.Called()
}

func (m *MockRows) Err() error {
	return nil
}

func (m *MockRows) CommandTag() pgconn.CommandTag {
	return pgconn.CommandTag{}
}

func (m *MockRows) FieldDescriptions() []pgconn.FieldDescription {
	return nil
}

func (m *MockRows) RawValues() [][]byte {
	return nil
}

func (m *MockRows) Values() ([]interface{}, error) {
	return nil, nil
}

func (m *MockRows) Conn() *pgx.Conn {
	return nil
}

// MockRow implements pgx.Row for testing
type MockRow struct {
	mock.Mock
	Data []interface{}
}

func (m *MockRow) Scan(dest ...interface{}) error {
	result := m.Called(mock.Anything)

	// If set to return error, do so
	if result.Error(0) != nil {
		return result.Error(0)
	}

	// Copy data to destination pointers
	for i := 0; i < len(m.Data) && i < len(dest); i++ {
		switch v := dest[i].(type) {
		case *string:
			if str, ok := m.Data[i].(string); ok {
				*v = str
			}
		case *float64:
			switch val := m.Data[i].(type) {
			case float64:
				*v = val
			case float32:
				*v = float64(val)
			case int:
				*v = float64(val)
			}
		}
	}

	return nil
}

// MockDBConn implements db.Database for testing
type MockDBConn struct {
	mock.Mock
}

func (m *MockDBConn) Close(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockDBConn) Ping(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockDBConn) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	mockArgs := []interface{}{ctx, sql}
	mockArgs = append(mockArgs, args...)
	called := m.Called(mockArgs...)

	result := called.Get(0)

	if result == nil {
		emptyRows := &MockRows{}
		return emptyRows, called.Error(1)
	}

	return result.(pgx.Rows), called.Error(1)
}

func (m *MockDBConn) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	mockArgs := []interface{}{ctx, sql}
	mockArgs = append(mockArgs, args...)
	called := m.Called(mockArgs...)
	return called.Get(0).(pgx.Row)
}

// Setup and teardown helpers
func SetupTestDB() (*MockDBConn, db.Database) {
	mockDB := new(MockDBConn)
	originalDB := db.DBConn
	db.DBConn = mockDB
	return mockDB, originalDB
}

func TeardownTestDB(originalDB db.Database) {
	db.DBConn = originalDB
}
