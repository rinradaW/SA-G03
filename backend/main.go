package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rinradaW/SA-G03/controller"
	"github.com/rinradaW/SA-G03/entity"
	"github.com/rinradaW/SA-G03/middlewares"
)

func main() {

	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	api := r.Group("")
	{
		protected := api.Use(middlewares.Authorizes())
		{

			//Club Routes
			protected.GET("/clubs", controller.ListClubs)
			protected.GET("/club/:id", controller.GetClub)
			protected.POST("/clubs", controller.CreateClub)
			protected.PATCH("/clubs", controller.UpdateClub)
			protected.DELETE("/clubs/:id", controller.DeleteClub)

			//Activity Routes
			protected.GET("/activities", controller.ListActivities)
			protected.GET("/activitiy/:id", controller.GetActivity)
			protected.POST("/activities", controller.CreateActivity)
			protected.PATCH("/activities", controller.UpdateActivity)
			protected.DELETE("/activities/:id", controller.DeleteActivity)

			//Student Routes
			protected.GET("/students", controller.ListStudents)
			protected.GET("/student/:id", controller.GetStudent)
			protected.POST("/students", controller.CreateStudent)
			protected.PATCH("/students", controller.UpdateStudent)
			protected.DELETE("/students/:id", controller.DeleteStudent)

			//ClubCommittee Routes
			protected.GET("/club_committees", controller.ListClubCommittees)
			protected.GET("/club_committee/:id", controller.GetClubCommittee)
			protected.POST("/club_committees", controller.CreateClubCommittee)
			protected.PATCH("/club_committees", controller.UpdateClubCommittee)
			protected.DELETE("/club_committees/:id", controller.DeleteClubCommittee)

			//JoinActivityHistory Routes
			protected.GET("/join_activity_histories", controller.ListJoinActivityHistories)
			protected.GET("/join_activity_history/:id", controller.GetJoinActivityHistory)
			protected.POST("/join_activity_histories", controller.CreateJoinActivityHistory)
			protected.PATCH("/join_activity_histories", controller.UpdateJoinActivityHistory)
			protected.DELETE("/join_activity_histories/:id", controller.DeleteJoinActivityHistory)

		}
	}

	// Authentication Routes
	r.POST("/login", controller.LoginByClubCommittee)

	//Run server
	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
