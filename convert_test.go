package humannumbers_test

import (
	"testing"

	"github.com/paridab/humannumbers"
)

func TestOutOfRange(t *testing.T) {
	//result, err := convert.NumToWord(-1)
	_, err := humannumbers.NumToWord(-1)
	if err == nil{
		t.Fatal("Out of range value didn't return error")
	}
 }

 func TestInRange(t *testing.T) {
	/*result, err := convert.NumToWord(0)
	if err != nil{
		t.Fatalf("Valid value %v resulted in an error %v", 0, err)
	}
	if result !="zero"{
		t.Fatalf("expected %v got %v", "zero", result)
	}*/
	numbertest(t,0,"zero")
	//numbertest(t,20,"twenty")
 }

 func TestUnit(t *testing.T) {
	numbertest(t,53,"fifty three")
 }

 func TestHundredUnit(t *testing.T) {
	numbertest(t,100,"hundred")
 }

 func numbertest(t *testing.T, numberToTest int, expectedResult string) {
	result, err := humannumbers.NumToWord(numberToTest)
	if err != nil{
		t.Fatalf("Valid value %v resulted in an error %v", numberToTest, err)
	}
	if result != expectedResult {
		t.Fatalf("expected %v got %v", expectedResult, result)
	}
 }
