package controllers

func (s *Server) SetupRoutes() {
	// Define your routes here
	s.Router.GET("/", s.GetHelthCheck())
	s.Router.POST("/todos", s.CreateTodo())
}
