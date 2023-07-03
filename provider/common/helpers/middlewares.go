package helpers

import (
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"regexp"
)

func SetTracingDetails(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("requestID", uuid.New().String())
		projectManagerRegex, _ := regexp.Compile("/\\/projects\\//gm")
		kafkaManagerRegex, _ := regexp.Compile("/\\/kafka\\//gm")
		logManagerRegex, _ := regexp.Compile("/\\/logs\\//gm")
		gatewayManagerRegex, _ := regexp.Compile("/\\/gateway\\//gm")
		redisManagerRegex, _ := regexp.Compile("/\\/redis\\//gm")
		containerManagerRegex, _ := regexp.Compile("/\\/ecs\\//gm")
		bucketManagerRegex, _ := regexp.Compile("/\\/bucket\\//gm")
		if projectManagerRegex.MatchString(c.Request().RequestURI) {
			c.Set("serviceName", "Project-Manager")
		} else if kafkaManagerRegex.MatchString(c.Request().RequestURI) {
			c.Set("serviceName", "Kafka-Manager")
		} else if logManagerRegex.MatchString(c.Request().RequestURI) {
			c.Set("serviceName", "Logs-Manager")
		} else if gatewayManagerRegex.MatchString(c.Request().RequestURI) {
			c.Set("serviceName", "Gateway-Manager")
		} else if redisManagerRegex.MatchString(c.Request().RequestURI) {
			c.Set("serviceName", "Redis-Manager")
		} else if containerManagerRegex.MatchString(c.Request().RequestURI) {
			c.Set("serviceName", "Container-Manager")
		} else if bucketManagerRegex.MatchString(c.Request().RequestURI) {
			c.Set("serviceName", "Bucket-Manager")
		} else {
			c.Set("serviceName", "N/A")
		}
		return next(c)
	}
}
