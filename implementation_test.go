package go21

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type TestSuite struct{}

func (s *TestSuite) TestPostfixToInfix(c *C) {
	res, err := PostfixToInfix("2 2 +")
	c.Assert(err, IsNil)
	c.Check(res, Equals, "2 + 2")
}

func ExamplePostfixToInfix() {
	res, _ := PostfixToInfix("2 2 +")
	fmt.Println(res)
}
