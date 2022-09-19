package test

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	redisv7 "github.com/go-redis/redis/v7"
	redisv8 "github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"gorm.io/datatypes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println(aws.String("hello world"))
	fmt.Println(cors.DefaultSchemas)
	fmt.Println(gin.ContextKey)
	fmt.Println(redisv8.XAutoClaimCmd{})
	fmt.Println(redisv7.ClusterSlotsCmd{})
	fmt.Println(uuid.New().String())
	fmt.Println(newrelic.Application{})
	fmt.Println(logrus.DebugLevel)
	fmt.Println(viper.AllKeys())
	fmt.Println(assert.CallerInfo())
	fmt.Println(datatypes.URL{})
	fmt.Println(postgres.Index{TableName: "hello"})
	fmt.Println(gorm.DeletedAt{})

}
