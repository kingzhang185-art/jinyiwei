package router

import (
	"sentinel-opinion-monitor/internal/handler"
	"sentinel-opinion-monitor/internal/middleware"
	"sentinel-opinion-monitor/internal/repository"
	"sentinel-opinion-monitor/internal/service"

	"github.com/gin-gonic/gin"
)

// SetupRouter 设置路由
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 初始化依赖
	// 认证相关
	userRepo := repository.NewUserRepository()
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	// 用户管理
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// 角色管理
	roleRepo := repository.NewRoleRepository()
	roleService := service.NewRoleService(roleRepo)
	roleHandler := handler.NewRoleHandler(roleService)

	// 权限管理
	permissionRepo := repository.NewPermissionRepository()
	permissionService := service.NewPermissionService(permissionRepo)
	permissionHandler := handler.NewPermissionHandler(permissionService)

	// 舆情相关
	opinionRepo := repository.NewOpinionRepository()
	opinionService := service.NewOpinionService(opinionRepo)
	opinionHandler := handler.NewOpinionHandler(opinionService)
	pingHandler := handler.NewPingHandler()

	// 公开路由（无需认证）
	public := r.Group("/api/v1")
	{
		// 健康检查
		public.GET("/ping", pingHandler.Ping)

		// 认证相关
		auth := public.Group("/auth")
		{
			auth.POST("/register", authHandler.Register) // 用户注册
			auth.POST("/login", authHandler.Login)       // 用户登录
		}
	}

	// 需要认证的路由
	protected := r.Group("/api/v1")
	protected.Use(middleware.AuthMiddleware())
	{
		// 当前用户信息
		protected.GET("/auth/me", authHandler.GetUserInfo)          // 获取当前用户信息
		protected.PUT("/auth/password", userHandler.ChangePassword) // 修改密码

		// 用户管理（需要管理员权限）
		users := protected.Group("/users")
		users.Use(middleware.RequireRole("admin"))
		{
			users.POST("", userHandler.CreateUser)            // 创建用户
			users.GET("", userHandler.GetUsers)               // 获取用户列表
			users.GET("/:id", userHandler.GetUser)            // 获取用户详情
			users.PUT("/:id", userHandler.UpdateUser)         // 更新用户
			users.DELETE("/:id", userHandler.DeleteUser)      // 删除用户
			users.POST("/:id/roles", userHandler.AssignRoles) // 分配角色
		}

		// 角色管理（需要管理员权限）
		roles := protected.Group("/roles")
		roles.Use(middleware.RequireRole("admin"))
		{
			roles.POST("", roleHandler.CreateRole)                        // 创建角色
			roles.GET("", roleHandler.GetRoles)                           // 获取角色列表
			roles.GET("/:id", roleHandler.GetRole)                        // 获取角色详情
			roles.PUT("/:id", roleHandler.UpdateRole)                     // 更新角色
			roles.DELETE("/:id", roleHandler.DeleteRole)                  // 删除角色
			roles.POST("/:id/permissions", roleHandler.AssignPermissions) // 分配权限
		}

		// 权限管理（需要管理员权限）
		permissions := protected.Group("/permissions")
		permissions.Use(middleware.RequireRole("admin"))
		{
			permissions.POST("", permissionHandler.CreatePermission)       // 创建权限
			permissions.GET("", permissionHandler.GetPermissions)          // 获取权限列表
			permissions.GET("/:id", permissionHandler.GetPermission)       // 获取权限详情
			permissions.PUT("/:id", permissionHandler.UpdatePermission)    // 更新权限
			permissions.DELETE("/:id", permissionHandler.DeletePermission) // 删除权限
		}

		// 舆情相关接口（需要认证）
		opinions := protected.Group("/opinions")
		{
			opinions.GET("", opinionHandler.GetAllOpinions) // 获取舆情列表
			opinions.POST("", opinionHandler.CreateOpinion) // 创建舆情
			opinions.GET("/:id", opinionHandler.GetOpinion) // 获取舆情详情
		}
	}

	// 兼容旧的路由格式
	r.GET("/ping", pingHandler.Ping)
	r.GET("/opinion/:id", opinionHandler.GetOpinion)

	return r
}
