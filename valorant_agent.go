package valorant

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"valorant_agent/cmd/model"

	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
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
	e.Use(session.Middleware(sessions.NewFilesystemStore("", []byte("secret"))))

	e.GET("/", sessionHandler)
	e.POST("/api/register", register)
	e.GET("/api/login", login)
	e.GET("/api/agents", getAgents)
	e.GET("/api/agent/:agent_id", getAgent)

	db, err = connectDB()
	if err != nil {
		e.Logger.Fatal("faild to connect db: %v", err)
		return
	}
	defer db.Close()

	e.Logger.Fatal(e.Start(":1323"))
}

// 中で cookie がない場合は生成したり、session を確認したりする。
func sessionHandler(c echo.Context) error {
	session, _ := session.Get("session-valorant_agent", c)

	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	session.Values["foo"] = "bar"
	session.Save(c.Request(), c.Response())

	return c.NoContent(http.StatusOK)
}

func register(c echo.Context) error {
	ctx := context.Background()
	userName := c.FormValue("user_name")
	email := c.FormValue("email")
	password := c.FormValue("password")
	if userName == "" || email == "" || password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "email and password required.")
	}

	// ハッシュ化
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed hash passowrd: %w", err)
	}

	result, err := db.ExecContext(
		ctx,
		"INSERT INTO users (user_name, email, password) VALUES (?,?,?)",
		userName, email, hash,
	)

	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("error SELECT user: %v", err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error get lastInsertId: %w", err)
	}

	return c.JSON(http.StatusOK, model.SuccessResult{
		Status: true,
		Data:   lastInsertId,
	})
}

func login(c echo.Context) error {
	ctx := context.Background()
	email := c.Param("email")
	password := c.Param("password")

	userRow := []model.User{}
	if email == "" || password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "email and password required.")
	}

	if err := db.SelectContext(
		ctx,
		&userRow,
		"SELECT * FROM user WHERE email = ? AND password = ?",
		email, password,
	); err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("error SELECT user: %v", err)
	}

	return c.JSON(http.StatusOK, model.SuccessResult{
		Status: true,
		Data:   &userRow,
	})
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
		return echo.NewHTTPError(http.StatusBadRequest, "agent_id required.")
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
