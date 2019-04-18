package add

//we have to import the testing package
import "testing"
//All test case functions begin with "Test".
//Test functions of the form TestXxx() take a testing.T argument
// we can use this type to record errors or to get the testing status.
//Test function names must begin with Test; 
//the optional sufﬁx Name must begin with a capital letter.
//The t parameter provides methods for reporting test failures 
//and logging additional information.
func TestAdd(t *testing.T){
  x, y := 4, 5
  res := Add(x, y)
//By calling one of the Error, Errorf, FailNow, Fatal or FatalIf methods of testing.T 
//on our testing functions, we can fail the test. 
  // if res != exp {
  //   t.Errorf("Test failed, expected: '%s', got:  '%s'", exp, act)
  // }
  if res != 9{
  		t.Errorf("Test failed, expected: '%d', got:  '%d'", 9, res)
  }
}


func TestFoo(t *testing.T) {
	tc := []struct {
		i, j, exp int
	}{
		{1, 2, 3},
		{3, 4, 7},
		{11, 21, 32},
		{30, 40, 70},
	}
 
	for _, tt := range tc {
		if add(tt.i, tt.j) != tt.exp {
			t.Fail()
		}
	}
}
 
