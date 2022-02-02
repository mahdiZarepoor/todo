package todoDB

import (
	"testing"
)

func TestInsertToDo( t *testing.T) {


	type input struct {
		user string 
		title string 
		description string 
		deadline string 
	}
	type testcase struct {
		testname string 
		args input
		expected error 
	}

	tests := []testcase{
		{
			testname: "test",
			args : input{"alex", "sleeping", "to the end of world", "2023-02-11 12:55:17"},
			expected: nil ,
		},
	}

	for _, test := range tests {
		t.Run(test.testname , func(t *testing.T) {
			if actual := InsertTodo(test.args.user, test.args.title, test.args.description, test.args.deadline); actual != test.expected {
				t.Errorf("actual:%v wanted:%v",actual, test.expected)
			}
		})
	}


}