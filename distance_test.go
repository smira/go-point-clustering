package cluster

import (
	"math"

	. "gopkg.in/check.v1"
)

type DistanceSuite struct {
	p1, p2 Point
}

var _ = Suite(&DistanceSuite{})

func (s *DistanceSuite) SetUpTest(c *C) {
	s.p1 = Point{30.244759, 59.955982}
	s.p2 = Point{30.24472, 59.955975}
}

func (s *DistanceSuite) TestFastCos(c *C) {
	c.Check(FastCos(0), Equals, math.Cos(0))
	c.Check(math.Abs(FastCos(0.1)-math.Cos(0.1)) < 0.001, Equals, true)
	c.Check(math.Abs(FastCos(-0.1)-math.Cos(-0.1)) < 0.001, Equals, true)
	c.Check(math.Abs(FastCos(1.0)-math.Cos(1.0)) < 0.001, Equals, true)
}

func (s *DistanceSuite) TestDistanceSpherical(c *C) {
	c.Check(DistanceSpherical(&s.p1, &s.p2), Equals, 0.0023064907653812116)
	c.Check(DistanceSpherical(&s.p2, &s.p1), Equals, 0.0023064907653812116)
	c.Check(DistanceSpherical(&s.p1, &s.p1), Equals, 0.0)
	c.Check(DistanceSpherical(&s.p2, &s.p2), Equals, 0.0)
}

func (s *DistanceSuite) TestDistanceSphericalFast(c *C) {
	c.Check(DistanceSphericalFast(&s.p1, &s.p2), Equals, 4.3026720164084415e-10)
	c.Check(DistanceSphericalFast(&s.p2, &s.p1), Equals, 4.3026720164084415e-10)
	c.Check(DistanceSphericalFast(&s.p1, &s.p1), Equals, 0.0)
	c.Check(DistanceSphericalFast(&s.p2, &s.p2), Equals, 0.0)

	c.Check(math.Abs(math.Sqrt(DistanceSphericalFast(&s.p1, &s.p2))*DegreeRad*EarthR-
		DistanceSpherical(&s.p1, &s.p2)) < 0.000001, Equals, true)
}

func (s *DistanceSuite) BenchmarkDistanceSpherical(c *C) {
	for i := 0; i < c.N; i++ {
		DistanceSpherical(&s.p1, &s.p2)
	}
}

func (s *DistanceSuite) BenchmarkDistanceSphericalFast(c *C) {
	for i := 0; i < c.N; i++ {
		DistanceSphericalFast(&s.p1, &s.p2)
	}
}
