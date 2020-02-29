package go21

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type TestSuite struct{}

func (s *TestSuite) TestPositive(c *C) {
	res, err := PostfixToInfix("2 2 +")
	c.Assert(err, IsNil)
	c.Check(res, Equals, "2 + 2")
}
func (s *TestSuite) TestPositive2(c *C) {
	res, err := PostfixToInfix("2 2 + 5 * 25 ^ 12 +")
	c.Assert(err, IsNil)
	c.Check(res, Equals, "2 + 2")
}
func (s *TestSuite) TestPositive3(c *C) {
	res, err := PostfixToInfix("2 5 ^ 8 * 123 -")
	c.Assert(err, IsNil)
	c.Check(res, Equals, "2 + 2")
}
func (s *TestSuite) TestPositive4(c *C) {
	res, err := PostfixToInfix("123 2 * 12 / 6 ^")
	c.Assert(err, IsNil)
	c.Check(res, Equals, "2 + 2")
}
func (s *TestSuite) TestPositive5(c *C) {
	res, err := PostfixToInfix("8 4 + 12 / 55 * 12 / 75 ^ 54 * 7 / 6")
	c.Assert(err, IsNil)
	c.Check(res, Equals, "2 + 2")
}

// func (s *TestSuite) TestPostfixToInfix6(c *C) {
// 	res, err := PostfixToInfix("2 2 5 + 4 7 *")
// 	c.Assert(err, IsNil)
// 	c.Check(res, Equals, "2 + 2")
// }
// func (s *TestSuite) TestPostfixToInfix7(c *C) {
// 	res, err := PostfixToInfix("2 2 * + 4 ^")
// 	c.Assert(err, IsNil)
// 	c.Check(res, Equals, "2 + 2")
// }
// func (s *TestSuite) TestPostfixToInfix8(c *C) {
// 	res, err := PostfixToInfix("2 2 * 5 $ asf")
// 	c.Assert(err, IsNil)
// 	c.Check(res, Equals, "2 + 2")
// }
// func (s *TestSuite) TestPostfixToInfix9(c *C) {
// 	res, err := PostfixToInfix("")
// 	c.Assert(err, IsNil)
// 	c.Check(res, Equals, "2 + 2")
// }

func ExamplePostfixToInfix() {
	res, err := PostfixToInfix("2 2 + 81 * 12 / 5 + 12 - 1 ^ 12 * 3 +")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
	// Output:
	// (2 + 2) * 81 / 12 + 5 - 12 ^ 1 * 12 + 3
}
