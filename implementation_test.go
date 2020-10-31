package lab2

import (
	"fmt"
	"testing"

	gocheck "gopkg.in/check.v1"
)

func Test(t *testing.T) { gocheck.TestingT(t) }

type MySuite struct{}

var _ = gocheck.Suite(&MySuite{})

func (s *MySuite) TestSimple(c *gocheck.C) {
	res, err := CalculatePrefix("+ 2 2")
	c.Assert(res, gocheck.Equals, "4")
	c.Assert(err, gocheck.Equals, nil)
}

func (s *MySuite) TestComplex(c *gocheck.C) {
	res, err := CalculatePrefix("- * / 15 - 7 + 1 1 3 + 2 + 1 1")
	c.Assert(res, gocheck.Equals, "5")
	c.Assert(err, gocheck.Equals, nil)
}

func (s *MySuite) TestComplexPow(c *gocheck.C) {
	res, err := CalculatePrefix("^ 2 - * / 15 - 7 + 1 1 3 + 2 + 1 1")
	c.Assert(res, gocheck.Equals, "32")
	c.Assert(err, gocheck.Equals, nil)
}

func (s *MySuite) TestEmptyString(c *gocheck.C) {
	res, err := CalculatePrefix("")
	c.Assert(res, gocheck.Equals, "")
	c.Assert(err.Error(), gocheck.Equals, "Empty string")
}

func (s *MySuite) TestOddCharacters(c *gocheck.C) {
	res, err := CalculatePrefix("- * * / 15 - 7 + 1 1 3 + 2 + 1 1")
	c.Assert(res, gocheck.Equals, "")
	c.Assert(err.Error(), gocheck.Equals, "Odd characters")
}

func (s *MySuite) TestIncorrectCharacters(c *gocheck.C) {
	res, err := CalculatePrefix("-* / 15 - 7 + 1 1 3 + 2 + 1 1")
	c.Assert(res, gocheck.Equals, "")
	c.Assert(err.Error(), gocheck.Equals, "Incorrect characters")
}

func ExampleCalculatePrefix() {
	res, _ := CalculatePrefix("+ 2 2")
	fmt.Println(res)
	// Output:
	// 4
}
