package cluster

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Launch gocheck tests
func Test(t *testing.T) {
	TestingT(t)
}

type PointSuite struct {
	points PointList
}

var _ = Suite(&PointSuite{})

func (s *PointSuite) SetUpTest(c *C) {
	s.points = PointList{Point{30.244759, 59.955982}, Point{30.24472, 59.955975}, Point{30.244358, 59.96698}}
}

func (s *PointSuite) TestCentroidAndBounds(c *C) {
	c1 := Cluster{C: 0, Points: []int{0, 1, 2}}

	center, min, max := c1.CentroidAndBounds(s.points)
	c.Check(center, DeepEquals, Point{30.244612333333333, 59.95964566666667})
	c.Check(min, DeepEquals, Point{30.244358, 59.955975})
	c.Check(max, DeepEquals, Point{30.244759, 59.96698})
}
