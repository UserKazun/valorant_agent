package valorant

import (
	"context"
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"valorant_agent/gen/openapi"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type Renderer struct {
	templates *template.Template
}

func (r *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return r.templates.ExecuteTemplate(w, name, data)
}

type ServiceInfo struct {
	Title string
}

type Agent struct {
	ID         int    `db:"id"`
	Name       string `db:"name"`
	Background string `db:"background"`
	CreatedAt  string `db:"created_at"`
	UpdatedAt  string `db:"updated_at"`
}

var (
	tr  = &Renderer{templates: template.Must(template.ParseGlob("../views/*.html"))}
	db  *sqlx.DB
	err error
)

var serviceInfo = ServiceInfo{
	"Valorant Agent",
}

func getEnv(key string, defaultValue string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultValue
}

// 管理用DBに接続する
func connectDB() (*sqlx.DB, error) {
	config := mysql.NewConfig()
	config.Net = "tcp"
	config.Addr = getEnv("DB_HOST", "127.0.0.1") + ":" + getEnv("DB_PORT", "3306")
	config.User = getEnv("DB_USER", "agent")
	config.Passwd = getEnv("DB_PASSWORD", "agent")
	config.DBName = getEnv("DB_NAME", "valorant")
	config.ParseTime = true
	dsn := config.FormatDSN()

	return sqlx.Open("mysql", dsn)
}

type getAgentImpl struct{}

func Run() {
	e := echo.New()
	e.Debug = true
	e.Logger.SetLevel(log.DEBUG)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Renderer = tr

	myApi := getAgentImpl{}
	openapi.RegisterHandlers(e, myApi)

	e.GET("/", example)
	e.GET("/home", home)

	// e.GET("/api/agents", getAgents)
	e.GET("/api/agent/:agent_id", getAgent)

	db, err = connectDB()
	if err != nil {
		e.Logger.Fatal("faild to connect db: %v", err)
		return
	}
	defer db.Close()

	e.Logger.Fatal(e.Start(":1323"))
}

type SuccessResult struct {
	Status bool `json:"status"`
	Data   any  `json:"data"`
}

// TODO: 後で消す
func example(c echo.Context) error {
	return c.String(http.StatusOK, "Hello")
}

func home(c echo.Context) error {
	data := struct {
		ServiceInfo
		Content string
	}{
		ServiceInfo: serviceInfo,
		Content:     "Jett",
	}

	return c.Render(http.StatusOK, "home", data)
}

func (h getAgentImpl) GetAgents(c echo.Context) error {
	ctx := context.Background()
	agentRow := []Agent{}
	if err := db.SelectContext(
		ctx,
		&agentRow,
		"SELECT * FROM agents",
	); err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("error agents: %v", err)
	}

	return c.JSON(http.StatusOK, &openapi.Agent{
		Status: true,
		Data:   &agentRow,
	})
}

func getAgent(c echo.Context) error {
	ctx := context.Background()

	agentID := c.Param("agent_id")
	if agentID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "agent_id required")
	}

	agentRow := []Agent{}
	if err := db.SelectContext(
		ctx,
		&agentRow,
		"SELECT * FROM agents WHERE id = ?",
		agentID,
	); err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("error SELECT agents: %v", err)
	}

	return c.JSON(http.StatusOK, SuccessResult{
		Status: true,
		Data:   &agentRow,
	})
}
