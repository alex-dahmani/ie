package models

import (
	"bufio"
	"encoding/json"
	"github.com/intervention-engine/fhir/models"
	"github.com/pebbe/util"
	. "gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/dbtest"
	"os"
	"strings"
)

type PipelineSuite struct {
	DBServer *dbtest.DBServer
	Query    *models.Query
}

var _ = Suite(&PipelineSuite{})

func (p *PipelineSuite) SetUpSuite(c *C) {
	p.DBServer = &dbtest.DBServer{}
	p.DBServer.SetPath(c.MkDir())
}

func (p *PipelineSuite) SetUpTest(c *C) {
	file, err := os.Open("../fixtures/facts.json")
	defer file.Close()
	util.CheckErr(err)

	session := p.DBServer.Session()
	defer session.Close()
	factCollection := session.DB("ie-test").C("facts")
	factCollection.DropCollection()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		decoder := json.NewDecoder(strings.NewReader(scanner.Text()))
		fact := &Fact{}
		err = decoder.Decode(fact)
		util.CheckErr(err)
		i := bson.NewObjectId()
		fact.Id = i.Hex()
		factCollection.Insert(fact)
	}
	util.CheckErr(scanner.Err())
	p.Query = LoadQueryFromFixture("../fixtures/sample-query.json")
}

func (p *PipelineSuite) TearDownTest(c *C) {
	p.DBServer.Wipe()
}

func (p *PipelineSuite) TearDownSuite(c *C) {
	p.DBServer.Stop()
}

func (p *PipelineSuite) TestNewPersonPipeline(c *C) {
	pipeline := NewPipeline(p.Query)
	c.Assert(3, Equals, len(pipeline.MongoPipeline))
	session := p.DBServer.Session()
	defer session.Close()
	qr, err := pipeline.ExecuteCount(session.DB("ie-test"))
	util.CheckErr(err)
	c.Assert(qr.Total, Equals, 5)
}

func (p *PipelineSuite) TestNewPersonPipelineMultiCode(c *C) {
	pipeline := NewPipeline(LoadQueryFromFixture("../fixtures/multi-code-query.json"))
	session := p.DBServer.Session()
	defer session.Close()
	qr, err := pipeline.ExecuteCount(session.DB("ie-test"))
	util.CheckErr(err)
	c.Assert(qr.Total, Equals, 7)
}

func (p *PipelineSuite) TestNewPersonPipelineList(c *C) {
	pipeline := NewPipeline(p.Query)
	c.Assert(3, Equals, len(pipeline.MongoPipeline))
	session := p.DBServer.Session()
	defer session.Close()
	qpl, err := pipeline.ExecutePatientList(session.DB("ie-test"))
	util.CheckErr(err)
	c.Assert(qpl.PatientIds, HasLen, 5)
	c.Assert(qpl.PatientIds, Contains, "554bc3fda749dcd47b0002ce")
}

func (p *PipelineSuite) TestAgePipeline(c *C) {
	pipeline := NewPipeline(LoadQueryFromFixture("../fixtures/age-query.json"))
	session := p.DBServer.Session()
	defer session.Close()
	qr, err := pipeline.ExecuteCount(session.DB("ie-test"))
	util.CheckErr(err)
	c.Assert(qr.Total, Equals, 6)
}

func (p *PipelineSuite) TestUnderAnAgePipeline(c *C) {
	pipeline := NewPipeline(LoadQueryFromFixture("../fixtures/under-an-age-query.json"))
	session := p.DBServer.Session()
	defer session.Close()
	qr, err := pipeline.ExecuteCount(session.DB("ie-test"))
	util.CheckErr(err)
	c.Assert(qr.Total, Equals, 4)
}

func (p *PipelineSuite) TestOverAnAgePipeline(c *C) {
	pipeline := NewPipeline(LoadQueryFromFixture("../fixtures/over-an-age-query.json"))
	session := p.DBServer.Session()
	defer session.Close()
	qr, err := pipeline.ExecuteCount(session.DB("ie-test"))
	util.CheckErr(err)
	c.Assert(qr.Total, Equals, 36)
}

func (p *PipelineSuite) TestObservationValuePipeline(c *C) {
	pipeline := NewPipeline(LoadQueryFromFixture("../fixtures/risk-query.json"))
	session := p.DBServer.Session()
	defer session.Close()
	qr, err := pipeline.ExecuteCount(session.DB("ie-test"))
	util.CheckErr(err)
	c.Assert(qr.Total, Equals, 1)
}

func (p *PipelineSuite) TestEmptyQuery(c *C) {
	pipeline := NewPipeline(&models.Query{})
	session := p.DBServer.Session()
	defer session.Close()
	qr, err := pipeline.ExecuteCount(session.DB("ie-test"))
	util.CheckErr(err)
	c.Assert(qr.Total, Equals, 40)
}

func (p *PipelineSuite) TestCreateConditionPipeline(c *C) {
	pipeline := NewConditionPipeline(p.Query)
	session := p.DBServer.Session()
	defer session.Close()
	qr, err := pipeline.ExecuteCount(session.DB("ie-test"))
	util.CheckErr(err)
	c.Assert(qr.Total, Equals, 10)
}

func (p *PipelineSuite) TestCreateEncounterPipeline(c *C) {
	pipeline := NewEncounterPipeline(p.Query)
	session := p.DBServer.Session()
	defer session.Close()
	qr, err := pipeline.ExecuteCount(session.DB("ie-test"))
	util.CheckErr(err)
	c.Assert(qr.Total, Equals, 12)
}

func LoadQueryFromFixture(fileName string) *models.Query {
	data, err := os.Open(fileName)
	defer data.Close()
	util.CheckErr(err)
	decoder := json.NewDecoder(data)
	query := &models.Query{}
	err = decoder.Decode(query)
	util.CheckErr(err)
	return query
}

type containsChecker struct {
	*CheckerInfo
}

func (c *containsChecker) Check(params []interface{}, names []string) (result bool, error string) {
	var (
		ok       bool
		list     []string
		expected string
	)
	list, ok = params[0].([]string)
	if !ok {
		return false, "List value is not a []string"
	}
	expected, ok = params[1].(string)
	if !ok {
		return false, "Expected value is not a string"
	}
	for _, v := range list {
		if v == expected {
			return true, ""
		}
	}
	return false, "Expected value not found in list"
}

var Contains Checker = &containsChecker{&CheckerInfo{Name: "Contains", Params: []string{"list", "expected"}}}
