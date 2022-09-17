package valorant

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"valorant_agent/cmd/model"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

var (
	db  *sqlx.DB
	err error
)

func getEnv(key string, defaultValue string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultValue
}

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

func Run() {
	e := echo.New()
	e.Debug = true
	e.Logger.SetLevel(log.DEBUG)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/api/agents", getAgents)
	e.GET("/api/agent/:agent_id", getAgent)
	e.GET("/login", login)

	db, err = connectDB()
	if err != nil {
		e.Logger.Fatal("faild to connect db: %v", err)
		return
	}
	defer db.Close()

	e.Logger.Fatal(e.Start(":1323"))
}

func getAgents(c echo.Context) error {
	ctx := context.Background()
	agentRow := []model.Agent{}
	if err := db.SelectContext(
		ctx,
		&agentRow,
		"SELECT * FROM agents",
	); err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("error agents: %v", err)
	}

	return c.JSON(http.StatusOK, &model.SuccessResult{
		Status: true,
		Data:   agentRow,
	})
}

func getAgent(c echo.Context) error {
	ctx := context.Background()

	agentID := c.Param("agent_id")
	if agentID == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "agent_id required")
	}

	agentRow := []model.Agent{}
	if err := db.SelectContext(
		ctx,
		&agentRow,
		"SELECT * FROM agents WHERE id = ?",
		agentID,
	); err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("error SELECT agents: %v", err)
	}

	return c.JSON(http.StatusOK, model.SuccessResult{
		Status: true,
		Data:   &agentRow,
	})
}

func login(c echo.Context) error {
	return nil
}
