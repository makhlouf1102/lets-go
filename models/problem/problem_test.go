// models/problem/problem_test.go
package problem_model

import (
	"os"
	"testing"

	"lets-go/database"

	_ "github.com/mattn/go-sqlite3"
)


func TestMain(m *testing.M) {
	database.SetUpDBForTests("../../database/scripts/db_build.sql", "./problem_test_data_script.sql")

	defer database.DB.Close()

	// Run tests
	code := m.Run()
	os.Exit(code)
}

func TestGetAllProblems(t *testing.T) {
	problems, err := GetAllProblems()
	if err != nil {
		t.Fatalf("Failed to get all problems: %v", err)
	}

	if len(problems) != 3 {
		t.Errorf("Expected 3 problems, got %d", len(problems))
	}
	
}

func TestGetProblem(t *testing.T) {
	problem, err := GetProblem("uuid1")
	if err != nil {
		t.Fatalf("Failed to get problem: %v", err)
	}

	if problem.Title != "testproblem1" {
		t.Errorf("Expected title to be 'testproblem1', got %s", problem.Title)
	}

	if problem.ID != "uuid1" {
		t.Errorf("Expected ID to be 'uuid1', got %s", problem.ID)
	}

	if problem.Title != "testproblem1" {
		t.Errorf("Expected title to be 'testproblem1', got %s", problem.Title)
	}

	if problem.Description != "testproblem1" {
		t.Errorf("Expected description to be 'testproblem1', got %s", problem.Description)
	}

	if problem.Difficulty != "easy" {
		t.Errorf("Expected difficulty to be 'easy', got %s", problem.Difficulty)
	}
}

func TestCreateProblem(t *testing.T) {
	problem := &Problem{
		ID: "uuid4",
		Title: "testproblem4",
		Description: "testproblem4",
		Difficulty: "easy",
	}

	err := problem.Create()
	if err != nil {
		t.Fatalf("Failed to create problem: %v", err)
	}

	createdProblem, err := GetProblem("uuid4")
	if err != nil {
		t.Fatalf("Failed to get created problem: %v", err)
	}

	if createdProblem.Title != "testproblem4" {
		t.Errorf("Expected title to be 'testproblem4', got %s", createdProblem.Title)
	}

	if createdProblem.Description != "testproblem4" {
		t.Errorf("Expected description to be 'testproblem4', got %s", createdProblem.Description)
	}

	if createdProblem.Difficulty != "easy" {
		t.Errorf("Expected difficulty to be 'easy', got %s", createdProblem.Difficulty)
	}
	
}

func TestUpdateProblem(t *testing.T) {
	problem := &Problem{
		ID: "uuid1",
		Title: "updatedtestproblem1",
		Description: "updatedtestproblem1",
		Difficulty: "easy",
	}

	err := problem.Update()
	if err != nil {
		t.Fatalf("Failed to update problem: %v", err)
	}

	updatedProblem, err := GetProblem("uuid1")
	if err != nil {
		t.Fatalf("Failed to get updated problem: %v", err)
	}

	if updatedProblem.Title != "updatedtestproblem1" {
		t.Errorf("Expected title to be 'updatedtestproblem1', got %s", updatedProblem.Title)
	}

	if updatedProblem.Description != "updatedtestproblem1" {
		t.Errorf("Expected description to be 'updatedtestproblem1', got %s", updatedProblem.Description)
	}

	if updatedProblem.Difficulty != "easy" {
		t.Errorf("Expected difficulty to be 'easy', got %s", updatedProblem.Difficulty)
	}

}

func TestDeleteProblem(t *testing.T) {

	problems, err := GetAllProblems()
	if err != nil {
		t.Fatalf("Failed to get all problems: %v", err)
	}


	problem, err := GetProblem("uuid1")
	if err != nil {
		t.Fatalf("Failed to get problem: %v", err)
	}

	err = problem.Delete()
	if err != nil {
		t.Fatalf("Failed to delete problem: %v", err)
	}

	newListProblems, err := GetAllProblems()
	if err != nil {
		t.Fatalf("Failed to get all problems: %v", err)
	}

	if len(newListProblems) == len(problems) {
		t.Errorf("Expected length to be %d, got %d", len(problems)-1, len(newListProblems))
	}
}
