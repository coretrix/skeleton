package tests

import (
	"bytes"
	"coretrix/skeleton/api/web-api/graphql/graph"
	"coretrix/skeleton/api/web-api/graphql/graph/generated"
	"coretrix/skeleton/pkg/entity"
	"coretrix/skeleton/pkg/ioc"
	"coretrix/skeleton/pkg/ioc/registry/mocks"
	"coretrix/skeleton/pkg/ioc/service/oss"
	graphqlParser "coretrix/skeleton/tests/graphql-parser"
	"coretrix/skeleton/tests/graphql-parser/jsonutil"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/coretrix/hitrix/service"
	"github.com/coretrix/hitrix/service/component/app"
	hitrixRegistry "github.com/coretrix/hitrix/service/registry"

	"github.com/coretrix/hitrix"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gin-gonic/gin"
	"github.com/summer-solutions/orm"
)

var dbService *orm.DB
var ormService *orm.Engine
var ginTestInstance *gin.Engine
var testSpringInstance *hitrix.Hitrix

type Ctx struct {
	t *testing.T
	g *gin.Engine
	c *gin.Context
	w *httptest.ResponseRecorder
}

func (ctx *Ctx) HandleQuery(query interface{}, variables map[string]interface{}) *graphqlParser.Errors {
	buff, err := graphqlParser.NewQueryParser().ParseQuery(query, variables)
	if err != nil {
		ctx.t.Fatal(err)
	}

	return ctx.handle(buff, query)
}

func (ctx *Ctx) HandleMutation(mutation interface{}, variables map[string]interface{}) *graphqlParser.Errors {
	buff, err := graphqlParser.NewQueryParser().ParseMutation(mutation, variables)
	if err != nil {
		ctx.t.Fatal(err)
	}

	return ctx.handle(buff, mutation)
}

func (ctx *Ctx) handle(buff bytes.Buffer, v interface{}) *graphqlParser.Errors {
	r, _ := http.NewRequestWithContext(ctx.c, http.MethodPost, "/query", &buff)
	r.Header = http.Header{"Content-Type": []string{"application/json"}}
	ctx.c.Request = r
	ctx.g.HandleContext(ctx.c)

	var out struct {
		Data   *json.RawMessage
		Errors *graphqlParser.Errors
	}
	if err := json.NewDecoder(ctx.w.Body).Decode(&out); err != nil {
		ctx.t.Fatal(err)
	}

	if out.Errors != nil {
		return out.Errors
	}

	if out.Data != nil {
		if err := jsonutil.UnmarshalGraphQL(*out.Data, v); err != nil {
			ctx.t.Fatal(err)
		}
	}

	return nil
}

func CreateContextWebAPI(t *testing.T, projectName string, resolvers graphql.ExecutableSchema, iocMocks *IoCMocks) *Ctx {
	mockServices := make([]*service.Definition, 0)
	if iocMocks != nil {
		mockServices = iocMocks.initMocks()
	}

	return createContext(t,
		projectName,
		resolvers,
		mockServices...,
	)
}

func createContext(t *testing.T, projectName string, resolvers graphql.ExecutableSchema, mockServices ...*service.Definition) *Ctx {
	_, filename, _, _ := runtime.Caller(0)
	var deferFunc func()

	defaultServices := []*service.Definition{
		hitrixRegistry.ServiceProviderConfigDirectory(filepath.Dir(filename) + "/../config"),
		hitrixRegistry.ServiceDefinitionOrmRegistry(entity.Init),
		hitrixRegistry.ServiceDefinitionOrmEngine(),
		hitrixRegistry.ServiceDefinitionOrmEngineForContext(),
	}

	if testSpringInstance == nil {
		err := os.Setenv("TZ", "UTC")
		if err != nil {
			t.Fatal(err)
		}
		err = os.Setenv("APP_MODE", app.ModeTest)
		if err != nil {
			t.Fatal(err)
		}

		testSpringInstance, deferFunc = hitrix.New(projectName, "").RegisterDIService(append(defaultServices, mockServices...)...).Build()
		defer deferFunc()
		ginTestInstance = hitrix.InitGin(resolvers, nil)

		ormService = ioc.GetOrmEngineGlobalService()
		dbService = ormService.GetMysql()

		err = dropTables()
		if err != nil {
			t.Fatal(err)
		}

		alters := ormService.GetAlters()

		for _, alter := range alters {
			dbService.Exec(alter.SQL)
		}
	}

	if len(mockServices) != 0 {
		testSpringInstance, deferFunc = hitrix.New(projectName, "").RegisterDIService(append(defaultServices, mockServices...)...).Build()
		defer deferFunc()
		ginTestInstance = hitrix.InitGin(resolvers, nil)

		// TODO: fix multiple connections to mysql
		ormService = ioc.GetOrmEngineGlobalService()
		dbService = ormService.GetMysql()
	}

	err := truncateTables()
	if err != nil {
		t.Fatal(err)
	}

	ormService.GetLocalCache().Clear()
	ormService.GetRedis().FlushDB()

	resp := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(resp)
	return &Ctx{t: t, g: ginTestInstance, c: c, w: resp}
}

func dropTables() error {
	dbService.Exec("SET GLOBAL FOREIGN_KEY_CHECKS=0")
	dbService.Exec("SET FOREIGN_KEY_CHECKS=0")

	var query string
	rows, deferF := dbService.Query(
		"SELECT CONCAT('DROP TABLE ',table_schema,'.',table_name,';') AS query " +
			"FROM information_schema.tables WHERE table_schema IN ('hymn_test')",
	)
	defer deferF()

	if rows != nil {
		for rows.Next() {
			rows.Scan(&query)
			dbService.Exec(query)
		}
	}

	dbService.Exec("SET GLOBAL FOREIGN_KEY_CHECKS=1")
	dbService.Exec("SET FOREIGN_KEY_CHECKS=1")
	return nil
}

func truncateTables() error {
	dbService.Exec("SET GLOBAL FOREIGN_KEY_CHECKS=0")
	dbService.Exec("SET FOREIGN_KEY_CHECKS=0")

	var query string
	rows, deferF := dbService.Query(
		"SELECT CONCAT('truncate table ',table_schema,'.',table_name,';') AS query " +
			"FROM information_schema.tables WHERE table_schema IN ('hymn_test');",
	)
	defer deferF()

	if rows != nil {
		for rows.Next() {
			rows.Scan(&query)
			dbService.Exec(query)
		}
	}

	dbService.Exec("SET GLOBAL FOREIGN_KEY_CHECKS=1")
	dbService.Exec("SET FOREIGN_KEY_CHECKS=1")

	return nil
}

type IoCMocks struct {
	OSSService oss.Client
}

func (m *IoCMocks) initMocks() []*service.Definition {
	mockServices := make([]*service.Definition, 0)
	if m.OSSService != nil {
		mockServices = append(mockServices, mocks.FakeOSSService(m.OSSService))
	}
	return mockServices
}

func GetWebAPIResolver() (string, graphql.ExecutableSchema) {
	config := generated.Config{Resolvers: &graph.Resolver{}, Directives: generated.DirectiveRoot{Validate: hitrix.ValidateDirective()}}

	return "web-api", generated.NewExecutableSchema(config)
}
