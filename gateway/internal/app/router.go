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
	"gateway/internal/transport/http_gin/handler"
	"gateway/internal/transport/http_gin/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, env *config.Config) {
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

	// =========================
	// SERVICES
	// =========================
	jwtService := jwt.NewJWTService(env.GetSecret(), env.Exp)
	refreshService := refresh.NewRefreshService(rdb, env.Ref_time)

	authService := auth.NewAuthService(refreshService, jwtService, userRepo)
	userService := user.NewUserService(userRepo, profileRepo)
	resumeService := resume.NewResumeService(resumeRepo)

	// =========================
	// HANDLERS
	// =========================
	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService)
	resumeHandler := handler.NewResumeHandler(resumeService)

	public_api := r.Group("/api/" + env.VersionAPI)

	// Public
	// Авторизация / Регистрация
	public_api.POST("/auth/register", authHandler.Register)
	public_api.POST("/auth/login", authHandler.Login)

	// Protected
	protected_api := public_api.Group("/") // ALLOW 401
	protected_api.Use(middleware.JWTAuthMiddleware(jwtService))

	// списки должны выводится, рекомедованные в начале
	protected_api.GET("/students", userHandler.ListStudents)   // Список студентов
	protected_api.GET("/employers", userHandler.ListEmployers) // Список работодателей

	protected_api.GET("/students/:id", userHandler.GetStudent)   // Профиль студента
	protected_api.GET("/employers/:id", userHandler.GetEmployer) // Профиль работодателя
	protected_api.GET("/users/me", userHandler.GetMe) // Зарос на user_id
	// Private
	private_api := protected_api.Group("/") // Проверка по user_id
	private_api.Use(middleware.OwnerMiddleware())

	// Если профиль твой, то можно PUT
	private_api.PUT("/students/:id", userHandler.UpdateStudent)   // Изменить профиль студента (user_id + role)
	private_api.PUT("/employers/:id", userHandler.UpdateEmployer) // Изменить профиль работодателя (user_id + role)

	// Операции с резюме
	private_api.GET("/students/:id/resume", resumeHandler.GetResume)       // Получить файл резюме
	private_api.POST("/students/:id/resume", resumeHandler.UploadResume)   // Изменить файл резюме
	private_api.DELETE("/students/:id/resume", resumeHandler.DeleteResume) // Удалить файл резюме
}
