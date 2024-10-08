package tests

import (
	"fmt"
	"gotest/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

//test grade by grade function

// go test ./tests -run="TestCheckGradeA" -v
// func TestCheckGradeA(t *testing.T) {
// 	grade := services.CheckGrade(80)
// 	expected := "A"

// 	if grade != expected {
// 		t.Errorf("got %v expected %v", grade, expected)
// 	}

// }

// go test ./tests -run="TestCheckGradeB" -v
// func TestCheckGradeB(t *testing.T) {
// 	grade := services.CheckGrade(70)
// 	expected := "B"

// 	if grade != expected {
// 		t.Errorf("got %v expected %v", grade, expected)
// 	}

// }
// ------------------------------------------------------//

// test all grades in 1 function
// how to run function
// go test ./tests -run="TestCheckGrade" -v
// how to run sub function
// go test ./tests -run="TestCheckGrade/A" -v

// func TestCheckGrade(t *testing.T) {

// 	t.Run("A", func(t *testing.T) {
// 		grade := services.CheckGrade(80)
// 		expected := "A"

// 		if grade != expected {
// 			t.Errorf("got %v expected %v", grade, expected)
// 		}
// 	})

// 	t.Run("B", func(t *testing.T) {
// 		grade := services.CheckGrade(70)
// 		expected := "B"

// 		if grade != expected {
// 			t.Errorf("got %v expedted %v", grade, expected)
// 		}
// 	})
// }

// ------------------------------------------------------//
// loop check all grade
// how to run function
// go test gotest/tests -run="TestCheckGrade" -v
func TestCheckGrade(t *testing.T) {

	type testCase struct {
		grade    string
		score    int
		expected string
	}

	cases := []testCase{
		{grade: "A", score: 80, expected: "A"},
		{grade: "B", score: 70, expected: "B"},
		{grade: "C", score: 60, expected: "C"},
		{grade: "D", score: 50, expected: "D"},
		{grade: "F", score: 0, expected: "F"},
	}

	//c = case
	for _, c := range cases {
		t.Run(c.grade, func(t *testing.T) {
			grade := services.CheckGrade(c.score)

			// Refactor code from
			// if grade != c.expected {
			// 	t.Errorf("got %v expected %v", grade, c.expected)
			// }

			//Refactor code to
			assert.Equal(t, c.expected, grade)
		})
	}

}

// ------------------------------------------------------//
// benchmark check grade
// how to benchmark all
// go test ./tests -bench=.
// how to benchmark all and check memory
// go test ./tests -bench=. -benchmem

func BenchmarkCheckGrade(b *testing.B) {
	for i := 0; i < b.N; i++ {
		services.CheckGrade(80)
	}
}

// ------------------------------------------------------//
// document grade
// install godoc
//go get golang.org/x/tools/cmd/godoc
// how to run
// export PATH="$GOPATH/bin:$PATH"
//godoc -http=:6060

func ExampleCheckGrade() {
	grade := services.CheckGrade(80)
	fmt.Println(grade)
	// Output: A
}
