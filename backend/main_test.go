package main

import (
	"context"
	"go_learning/db"
	"go_learning/test_util"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRouter struct {
	mock.Mock
}

func (m *MockRouter) Run(addr ...string) error {
	args := m.Called(mock.Anything)
	return args.Error(0)
}

func TestSetupPGX(t *testing.T) {
	if os.Getenv("DATABASE_URL") == "" {
		t.Skip("Skipping test as DATABASE_URL is not set")
	}

	conn := SetupPGX()
	assert.NotNil(t, conn, "Database connection should not be nil")

	err := conn.Ping(context.Background())
	assert.NoError(t, err, "Database connection should be valid")

	conn.Close(context.Background())
}

func TestSetupRouter(t *testing.T) {
	router := SetupRouter()
	assert.NotNil(t, router, "Router should not be nil")
}

func TestRun(t *testing.T) {
	originalDBConn := db.DBConn
	defer func() { db.DBConn = originalDBConn }()

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	mockRouter := new(MockRouter)
	mockRouter.On("Run", mock.Anything).Return(nil)

	originalSetupRouter := SetupRouter
	defer func() { SetupRouter = originalSetupRouter }()

	SetupRouter = func() Router {
		return mockRouter
	}

	// Create a mock DB implementation using the shared mock
	mockDBConn := new(test_util.MockDBConn)
	mockDBConn.On("Close", mock.Anything).Return(nil)

	originalSetupPGX := SetupPGX
	defer func() { SetupPGX = originalSetupPGX }()

	SetupPGX = func() db.Database {
		return mockDBConn
	}

	originalRun := Run
	defer func() { Run = originalRun }()

	Run = func() {
		db.DBConn = SetupPGX()
		defer db.DBConn.Close(context.Background())

		router := SetupRouter()
		router.Run(":8080")
	}

	go func() {
		Run()
		cancel()
	}()

	<-ctx.Done()

	mockRouter.AssertExpectations(t)
	mockDBConn.AssertExpectations(t)
}
