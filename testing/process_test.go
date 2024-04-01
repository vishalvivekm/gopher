/*package main

import "testing"

func TestCheckEven(t *testing.T) {
i := 10
expected := "YES"

res := checkEven(i)
if expected != res {
	t.Errorf("expected: %v, got: %v", expected, res)
    }
}
*/
package main

import (
"testing"
"example.com/testing/process"
)

func TestCheckEven(t *testing.T) {
i := 10
expected := "YES"

res := process.CheckEven(i)
if expected != res {
	t.Errorf("expected: %v, got: %v", expected, res)
    }
}
