package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/Aerilate/htn-backend/model"
	"github.com/gin-gonic/gin"
)

type Repository interface {
	GetUsers() ([]model.User, error)
	GetOneUser(id int) (model.User, error)
	UpdateUser(id int, updatedInfo model.User) error
	GetSkills(minFreq *int, maxFreq *int) ([]model.SkillRating, error)
}

func serve(repo Repository) {
	r := gin.Default()
	r.SetTrustedProxies(nil)
	registerRoutes(r, repo)
	r.Run()
}

func registerRoutes(r *gin.Engine, repo Repository) {
	r.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	r.GET("/users", getUsers(repo))
	r.GET("/users/:id", getOneUser(repo))
	r.PUT("/users/:id", updateUser(repo))
	r.GET("/skills", getSkills(repo))
}

func getUsers(repo Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := repo.GetUsers()
		if err != nil {
			c.Status(http.StatusNotFound)
		}
		c.JSON(http.StatusOK, users)
	}
}

func getOneUser(repo Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.Status(http.StatusNotFound)
		}
		user, err := repo.GetOneUser(id)
		if err != nil {
			c.Status(http.StatusNotFound)
		}
		c.JSON(http.StatusOK, user)
	}
}

func updateUser(repo Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		// validate id parameter
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.Status(http.StatusBadRequest)
		}

		data, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.Status(http.StatusBadRequest)
		}
		var updatedInfo model.User
		if err := json.Unmarshal(data, &updatedInfo); err != nil {
			c.Status(http.StatusBadRequest)
		}
		if err := repo.UpdateUser(id, updatedInfo); err != nil {
			c.Status(http.StatusBadRequest)
		}
		c.Status(http.StatusOK)
	}
}

func getSkills(repo Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			minFreq *int
			maxFreq *int
		)

		if minFreqParam, ok := c.GetQuery("min_frequency"); ok {
			if strConv(minFreqParam, minFreq) != nil {
				c.Status(http.StatusNotFound)
			}
		}
		if maxFreqParam, ok := c.GetQuery("max_frequency"); ok {
			if strConv(maxFreqParam, maxFreq) != nil {
				c.Status(http.StatusNotFound)
			}
		}

		skills, err := repo.GetSkills(minFreq, maxFreq)
		if err != nil {
			c.Status(http.StatusNotFound)
		}
		c.JSON(http.StatusOK, skills)
	}
}

func strConv(s string, res *int) error {
	i, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	*res = i
	return nil
}
