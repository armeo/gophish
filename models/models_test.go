package models

import (
	"os"
	"testing"

	"github.com/jordan-wright/gophish/config"
	"gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { check.TestingT(t) }

type ModelsSuite struct{}

var _ = check.Suite(&ModelsSuite{})

func (s *ModelsSuite) SetUpSuite(c *check.C) {
	config.Conf.DBPath = "../gophish_test.db"
	err := Setup()
	if err != nil {
		c.Fatalf("Failed creating database: %v", err)
	}
}

func (s *ModelsSuite) TestGetUser(c *check.C) {
	u, err := GetUser(1)
	c.Assert(err, check.Equals, nil)
	c.Assert(u.Username, check.Equals, "admin")
}

func (s *ModelsSuite) TestPostGroup(c *check.C) {
	g := Group{Name: "Test Group"}
	g.Targets = []Target{Target{Email: "test@example.com"}}
	g.UserId = 1
	err := PostGroup(&g)
	c.Assert(err, check.Equals, nil)
	c.Assert(g.Name, check.Equals, "Test Group")
	c.Assert(g.Targets[0].Email, check.Equals, "test@example.com")
}

func (s *ModelsSuite) TestPostGroupNoName(c *check.C) {
	g := Group{Name: ""}
	g.Targets = []Target{Target{Email: "test@example.com"}}
	g.UserId = 1
	err := PostGroup(&g)
	c.Assert(err, check.Equals, ErrGroupNameNotSpecified)
}

func (s *ModelsSuite) TestPostGroupNoTargets(c *check.C) {
	g := Group{Name: "No Target Group"}
	g.Targets = []Target{}
	g.UserId = 1
	err := PostGroup(&g)
	c.Assert(err, check.Equals, ErrNoTargetsSpecified)
}

func (s *ModelsSuite) TestPutUser(c *check.C) {
	u, err := GetUser(1)
	u.Username = "admin_changed"
	err = PutUser(&u)
	c.Assert(err, check.Equals, nil)
	u, err = GetUser(1)
	c.Assert(u.Username, check.Equals, "admin_changed")
}

func (s *ModelsSuite) TearDownSuite(c *check.C) {
	db.DB().Close()
	err := os.Remove(config.Conf.DBPath)
	if err != nil {
		c.Fatalf("Failed deleting test database: %v", err)
	}
}
