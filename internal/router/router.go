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

	// 标签管理
	tagRepo := repository.NewTagRepository()
	tagService := service.NewTagService(tagRepo)
	tagHandler := handler.NewTagHandler(tagService)

	// 渠道管理
	channelRepo := repository.NewChannelRepository()
	channelService := service.NewChannelService(channelRepo)
	channelHandler := handler.NewChannelHandler(channelService)

	// 场景管理
	scenarioRepo := repository.NewScenarioRepository()
	scenarioService := service.NewScenarioService(scenarioRepo, tagRepo)
	scenarioHandler := handler.NewScenarioHandler(scenarioService)

	// 监测组管理
	groupRepo := repository.NewMonitoringGroupRepository()
	groupService := service.NewMonitoringGroupService(groupRepo, scenarioRepo)
	groupHandler := handler.NewMonitoringGroupHandler(groupService)

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

		// 标签管理（查看需要认证，增删改需要admin权限）
		tags := protected.Group("/tags")
		{
			tags.GET("", tagHandler.GetTags)    // 获取标签列表（支持type和status查询参数）
			tags.GET("/:id", tagHandler.GetTag) // 获取标签详情
		}

		// 标签管理（需要管理员权限）
		tagsAdmin := protected.Group("/tags")
		tagsAdmin.Use(middleware.RequireRole("admin"))
		{
			tagsAdmin.POST("", tagHandler.CreateTag)       // 创建标签
			tagsAdmin.PUT("/:id", tagHandler.UpdateTag)    // 更新标签
			tagsAdmin.DELETE("/:id", tagHandler.DeleteTag) // 删除标签
		}

		// 渠道管理（查看需要认证，增删改需要admin权限）
		channels := protected.Group("/channels")
		{
			channels.GET("", channelHandler.GetChannels)    // 获取渠道列表（支持status查询参数）
			channels.GET("/:id", channelHandler.GetChannel) // 获取渠道详情
		}

		// 渠道管理（需要管理员权限）
		channelsAdmin := protected.Group("/channels")
		channelsAdmin.Use(middleware.RequireRole("admin"))
		{
			channelsAdmin.POST("", channelHandler.CreateChannel)       // 创建渠道
			channelsAdmin.PUT("/:id", channelHandler.UpdateChannel)    // 更新渠道
			channelsAdmin.DELETE("/:id", channelHandler.DeleteChannel) // 删除渠道
		}

		// 场景管理（查看需要认证，增删改需要admin权限）
		scenarios := protected.Group("/scenarios")
		{
			scenarios.GET("", scenarioHandler.GetScenarios)                     // 获取场景列表
			scenarios.GET("/:id", scenarioHandler.GetScenario)                  // 获取场景详情
			scenarios.GET("/:id/groups", scenarioHandler.GetScenarioWithGroups) // 获取场景及其监测组
		}

		// 场景管理（需要管理员权限）
		scenariosAdmin := protected.Group("/scenarios")
		scenariosAdmin.Use(middleware.RequireRole("admin"))
		{
			scenariosAdmin.POST("", scenarioHandler.CreateScenario)       // 创建场景
			scenariosAdmin.PUT("/:id", scenarioHandler.UpdateScenario)    // 更新场景
			scenariosAdmin.DELETE("/:id", scenarioHandler.DeleteScenario) // 删除场景
		}

		// 监测组管理（查看需要认证，增删改需要admin权限）
		groups := protected.Group("/monitoring-groups")
		{
			groups.GET("/scenario/:scenario_id", groupHandler.GetGroupsByScenario) // 根据场景ID获取监测组列表
			groups.GET("/:id", groupHandler.GetGroup)                              // 获取监测组详情
			groups.GET("/:id/keywords", groupHandler.GetKeywords)                  // 获取关键词列表
			groups.GET("/:id/exclusion-words", groupHandler.GetExclusionWords)     // 获取排除词列表
		}

		// 监测组管理（需要管理员权限）
		groupsAdmin := protected.Group("/monitoring-groups")
		groupsAdmin.Use(middleware.RequireRole("user"))
		{
			groupsAdmin.POST("", groupHandler.CreateGroup)                                        // 创建监测组
			groupsAdmin.PUT("/:id", groupHandler.UpdateGroup)                                     // 更新监测组
			groupsAdmin.DELETE("/:id", groupHandler.DeleteGroup)                                  // 删除监测组
			groupsAdmin.POST("/:id/channels", groupHandler.AssignChannels)                        // 分配渠道
			groupsAdmin.POST("/:id/keywords", groupHandler.AddKeyword)                            // 添加关键词
			groupsAdmin.DELETE("/:id/keywords/:keyword_id", groupHandler.RemoveKeyword)           // 删除关键词
			groupsAdmin.POST("/:id/exclusion-words", groupHandler.AddExclusionWord)               // 添加排除词
			groupsAdmin.DELETE("/:id/exclusion-words/:word_id", groupHandler.RemoveExclusionWord) // 删除排除词
		}
	}

	// 兼容旧的路由格式
	r.GET("/ping", pingHandler.Ping)
	r.GET("/opinion/:id", opinionHandler.GetOpinion)

	return r
}
