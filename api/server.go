package api

import (
	"fmt"
	"simple_bank/token"
	"simple_bank/util"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	config     util.Config
	handler    Handler
	tokenMaker token.Maker
	Router     *gin.Engine
}

func NewServer(config util.Config, handler Handler) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		handler:    handler,
		tokenMaker: tokenMaker}

	// Register Custom validation
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}
	// TODO: add routes to router
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users/login", server.loginUser)

	authorized := router.Group("/admin", basicAuth(server.tokenMaker))

	router.POST("/users", server.createUser)

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	authorized.GET("/accounts", server.listAccount)

	router.POST("/transfers", server.createTransfer)

	server.Router = router
}

func basicAuth(tokenMaker token.Maker) gin.HandlerFunc {

	return func(c *gin.Context) {
		authorizationHeader := c.Request.Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			respondWithError(401, "Unauthorized", c)
			return
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		payload, err := tokenMaker.VerifyToken(tokenString)

		if err != nil {
			respondWithError(401, "Unauthorized", c)
			return
		}

		if !authenticateToken(*payload, tokenMaker.GetPayload()) {
			respondWithError(401, "Unauthorized", c)
			return
		}

		c.Next()
	}
}

func authenticateToken(payload token.Payload, originalPayload token.Payload) bool {
	if strings.Compare(string(payload.ID.String()), originalPayload.ID.String()) != 0 {
		return false
	}

	if strings.Compare(string(payload.Username), originalPayload.Username) != 0 {
		return false
	}

	if strings.Compare(string(payload.Issuer), originalPayload.Issuer) != 0 {
		return false
	}

	if !payload.IssuedAt.Equal(originalPayload.IssuedAt) {
		return false
	}

	if !payload.ExpiredAt.Equal(originalPayload.ExpiredAt) {
		return false
	}
	return true
}

// func basicAuth2(c *gin.Context) {
// 	user, password, hasAuth := c.Request.BasicAuth()
// 	if hasAuth && user == "testuser" && password == "testpass" {
// 		log.WithFields(log.Fields{
// 			"user": user,
// 		}).Info("User authenticated")
// 	} else {
// 		c.Abort()
// 		c.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
// 		return
// 	}
// }

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func respondWithError(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}

	c.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
	c.JSON(code, resp)
	c.Abort()
}

func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}
