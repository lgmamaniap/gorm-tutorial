package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type DemoTestSuite struct {
	suite.Suite
}

type DataTest struct {
	A        int `json:"a"`
	B        int `json:"b"`
	Expected int `json:"expected"`
}

func (s *DemoTestSuite) TestSum() {
	tests := []DataTest{
		{3, 5, 8},
		{0, 0, 0},
		{-1, 1, 0},
		{-2, -3, -5},
	}

	for _, test := range tests {
		result := sum(test.A, test.B)
		s.Equal(test.Expected, result, "sum(%d, %d) should equal %d", test.A, test.B, test.Expected)
	}
}

func (s *DemoTestSuite) TestMultiply() {
	tests := []DataTest{
		{3, 5, 15},
		{0, 5, 0},
		{-1, 1, -1},
		{-2, -3, 6},
	}

	for _, test := range tests {
		result := multiply(test.A, test.B)
		s.Equal(test.Expected, result, "multiply(%d, %d) should equal %d", test.A, test.B, test.Expected)
	}
}

func (s *DemoTestSuite) TestDivide() {
	tests := []DataTest{
		{10, 2, 5},
		{0, 1, 0},
		{-6, -3, 2},
		{5, 0, 0}, // Division by zero case
	}

	for _, test := range tests {
		result := divide(test.A, test.B)
		if test.B == 0 {
			s.Equal(0, result, "divide(%d, %d) should handle division by zero", test.A, test.B)
		} else {
			s.Equal(test.Expected, result, "divide(%d, %d) should equal %d", test.A, test.B, test.Expected)
		}
	}
}

func (s *DemoTestSuite) SetupTest() {
	// This method is called before each test in the suite
}

func (s *DemoTestSuite) TearDownTest() {
	// This method is called after each test in the suite
}

func TestSumTestSuite(t *testing.T) {
	suite.Run(t, new(DemoTestSuite))
}
