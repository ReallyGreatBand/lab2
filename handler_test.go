package lab2

import (
	"bytes"
	"strings"
	"testing"

	gocheck "gopkg.in/check.v1"
)

func TestHandler(t *testing.T) { gocheck.TestingT(t) }

type MySuiteHandler struct{}

var _ = gocheck.Suite(&MySuiteHandler{})

func (s *MySuiteHandler) TestOutputResult(c *gocheck.C) {
	mockReader := strings.NewReader("- * / 15 - 7 + 1 1 3 + 2 + 1 1")
	mockWriter := new(bytes.Buffer)
	handler := ComputeHandler{mockReader, mockWriter}
	err := handler.Compute()
	res, _ := mockWriter.ReadString(0)
	c.Assert(res, gocheck.Equals, "5")
	c.Assert(err, gocheck.Equals, nil)
}

func (s *MySuiteHandler) TestInputError(c *gocheck.C) {
	mockReader := strings.NewReader("- * * / 15 - 7 + 1 1 3 + 2 + 1 1")
	mockWriter := new(bytes.Buffer)
	handler := ComputeHandler{mockReader, mockWriter}
	err := handler.Compute()
	res, _ := mockWriter.ReadString(0)
	c.Assert(res, gocheck.Equals, "")
	c.Assert(err.Error(), gocheck.Equals, "Odd characters")
}
