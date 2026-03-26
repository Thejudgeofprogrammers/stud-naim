package app

import (
	"gateway/internal/config"
	"gateway/internal/infra/redis"
	"gateway/internal/repository/memory"
	auth "gateway/internal/service/auth/impl"
	jwt "gateway/internal/service/jwt/impl"
	refresh "gateway/internal/service/refresh/impl"
	resume "gateway/internal/service/resume/impl"
	user "gateway/internal/service/user/impl"
	opportunity "gateway/internal/service/opportunity/impl"
	"gateway/internal/transport/http_gin/handler"
	"gateway/internal/transport/http_gin/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, env *config.Config) {
	r.Static("/uploads", "./uploads")

	// =========================
	// INFRA
	// =========================
	rdb := redis.NewRedisClient(
		env.RedisAddr,
		env.GetRedisPassword(),
		env.RedisDB,
	)

	// =========================
	// REPOSITORIES
	// =========================
	userRepo := memory.NewUserRepositoryMemory()
	profileRepo := memory.NewProfileRepositoryMemory()
	resumeRepo := memory.NewResumeRepositoryMemory()
	opportunityRepo := memory.NewOpportunityRepositoryMemory()

	// =========================
	// SERVICES
	// =========================
	jwtService := jwt.NewJWTService(env.GetSecret(), env.Exp)
	refreshService := refresh.NewRefreshService(rdb, env.Ref_time)

	authService := auth.NewAuthService(refreshService, jwtService, userRepo, profileRepo)
	userService := user.NewUserService(userRepo, profileRepo)
	resumeService := resume.NewResumeService(resumeRepo)
	opportunityService := opportunity.NewOpportunityService(opportunityRepo)

	// =========================
	// HANDLERS
	// =========================
	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService)
	resumeHandler := handler.NewResumeHandler(resumeService)
	opportunityHandler := handler.NewOpportunityHandler(opportunityService)

	// =========================
	// ROUTES
	// =========================
	publicAPI := r.Group("/api/" + env.VersionAPI)

	// ---------- PUBLIC ----------
	publicAPI.POST("/auth/register", authHandler.Register)
	publicAPI.POST("/auth/login", authHandler.Login)

	// ---------- PROTECTED ----------
	protectedAPI := publicAPI.Group("/")
	protectedAPI.Use(middleware.JWTAuthMiddleware(jwtService))

	// Users
	protectedAPI.GET("/students", userHandler.ListStudents)
	protectedAPI.GET("/employers", userHandler.ListEmployers)

	protectedAPI.GET("/students/:id", userHandler.GetStudent)
	protectedAPI.GET("/employers/:id", userHandler.GetEmployer)

	protectedAPI.GET("/users/me", userHandler.GetMe)

	// Opportunities (доступны всем авторизованным)
	protectedAPI.GET("/opportunities", opportunityHandler.List)
	protectedAPI.GET("/opportunities/filter", opportunityHandler.Filter)
	protectedAPI.GET("/opportunities/:id", opportunityHandler.Get)

	// ---------- PRIVATE ----------
	privateAPI := protectedAPI.Group("/")
	privateAPI.Use(middleware.OwnerMiddleware())

	// Profiles
	privateAPI.PUT("/students/:id", userHandler.UpdateStudent)
	privateAPI.PUT("/employers/:id", userHandler.UpdateEmployer)

	// Resume
	privateAPI.GET("/students/:id/resume", resumeHandler.GetResume)
	privateAPI.POST("/students/:id/resume", resumeHandler.UploadResume)
	privateAPI.DELETE("/students/:id/resume", resumeHandler.DeleteResume)

	// Opportunities (только владелец)
	privateAPI.POST("/opportunities", opportunityHandler.Create)
	privateAPI.PUT("/opportunities/:id", opportunityHandler.Update)
	privateAPI.DELETE("/opportunities/:id", opportunityHandler.Delete)
}
