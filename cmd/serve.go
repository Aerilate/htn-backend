package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/Aerilate/htn-backend/model"
	"github.com/Aerilate/htn-backend/repository"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the server",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := gorm.Open(sqlite.Open(SQLitePath), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}
		repository := repository.NewRepo(db)
		server := NewServer(repository)
		server.Run()
	},
}

type Server struct {
	*gin.Engine
	repo Repository
}

type Repository interface {
	UserRepository
	SkillRatingRepository
}

type UserRepository interface {
	InsertUsers(users []model.User) error
	GetAllUsers() ([]model.User, error)
	GetUser(id int) (model.User, error)
	UpdateUser(id int, updatedInfo model.User) error
}

type SkillRatingRepository interface {
	AggregateSkills(minFreq *int, maxFreq *int) ([]model.SkillAggregate, error)
}

func NewServer(repo Repository) Server {
	g := gin.Default()
	g.SetTrustedProxies(nil)
	server := Server{g, repo}
	server.registerRoutes()
	return server
}

func (s Server) registerRoutes() {
	s.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	s.GET("/users/", s.getUsers())
	s.GET("/users/:id", s.getOneUser())
	s.PUT("/users/:id", s.updateUser())
	s.GET("/skills/", s.getSkills())
}

func (s Server) getUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := s.repo.GetAllUsers()
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.JSON(http.StatusOK, users)
	}
}

func (s Server) getOneUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		user, err := s.repo.GetUser(id)
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.JSON(http.StatusOK, user)
	}
}

func (s Server) updateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		// validate id parameter
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		data, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		var updatedInfo model.User
		if err := json.Unmarshal(data, &updatedInfo); err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		if err := s.repo.UpdateUser(id, updatedInfo); err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		c.Status(http.StatusOK)
	}
}

func (s Server) getSkills() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			minFreq *int
			maxFreq *int
			err     error
		)

		if minFreqParam, ok := c.GetQuery("min_frequency"); ok {
			if minFreq, err = strConv(minFreqParam); err != nil {
				c.Status(http.StatusBadRequest)
				return
			}
		}
		if maxFreqParam, ok := c.GetQuery("max_frequency"); ok {
			if maxFreq, err = strConv(maxFreqParam); err != nil {
				c.Status(http.StatusBadRequest)
				return
			}
		}

		skills, err := s.repo.AggregateSkills(minFreq, maxFreq)
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		c.JSON(http.StatusOK, skills)
	}
}

func strConv(s string) (*int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return nil, err
	}
	return &i, nil
}
