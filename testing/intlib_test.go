/*
* the file name has to end with _test.go to be picked up as a set of tests by go test
* the package name has to be the same as in the source file that has to be tested
* you have to import the package testing
* all test functions should start with Test to be run as a test
* the tests will be executed in the same order that they are appear in the source
* the test function TestXxx functions take a pointer to the type testing.T. You use it to record the test status and also for logging.
* the signature of the test function should always be func TestXxx ( *testing.T). You can have any combination of alphanumeric characters and the hyphen for the Xxx part, the only constraint that it should not begin with a small alphabet, [a-z].
* a call to any of the following functions of testing.T within the test code Error, Errorf, FailNow, Fatal, FatalIf will indicate to go test that the test has failed.
 */
package intpkg // same package name as source file

import (
	"testing" // import go package for testing related functionality
)

func Test_Add2Ints_1(t *testing.T) { //test function starts with "Test" and takes a pointer to type testing.T
	if Add2Ints(3, 4) != 7 { //try a unit test on function
		t.Error("Add2Ints did not work as expected.") // log error if it did not work as expected
	} else {
		t.Log("one test passed.") // log some info if you want
	}
}

func Test_Add2Ints_2(t *testing.T) { //test function starts with "Test" and takes a pointer to type testing.T
	t.Error("this is just hardcoded as an error.") //Indicate that this test failed and log the string as info
}
