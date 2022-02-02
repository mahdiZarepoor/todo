package db

import "testing"

func TestIsValid(t *testing.T) {
	user, pass := "ali" , "223369"
	expectedBool := true 
	var expectedError error= nil 
	if actualBool, actualError:=IsValid(user,pass);expectedBool!= actualBool || expectedError!= actualError {
		t.Errorf("actualBool : %v, actualError:%v",actualBool, actualError)
	}


	user, pass = "john doe" , "13808"
	expectedBool = true 
	expectedError = nil 
	if actualBool, actualError:=IsValid(user,pass);expectedBool!= actualBool || expectedError!= actualError {
		t.Errorf("actualBool : %v, actualError:%v",actualBool, actualError)
	}
}