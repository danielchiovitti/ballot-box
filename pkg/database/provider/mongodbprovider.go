package provider

import (
	"context"
	"fmt"
	"github.com/danielchiovitti/ballot-box/pkg/domain/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/url"
	"strconv"
	"sync"
)

var lockProvider sync.Mutex
var lockMongoDbClient sync.Mutex
var mongoDbProviderInstance *MongoDbProvider
var mongoDbClientInstance *mongo.Client

type MongoDbProvider struct {
	*model.DatabaseOptions
}

func NewMongoDbProvider(opts ...model.DatabaseOptionsFunc) *MongoDbProvider {
	if mongoDbProviderInstance == nil {
		lockProvider.Lock()
		defer lockProvider.Unlock()
		if mongoDbProviderInstance == nil {
			o := DatabaseDefaultOpts()
			for _, fn := range opts {
				fn(o)
			}
			mongoDbProviderInstance = &MongoDbProvider{
				DatabaseOptions: o,
			}
		}
	}
	return mongoDbProviderInstance
}

func WithHost(host string) model.DatabaseOptionsFunc {
	return func(opt *model.DatabaseOptions) {
		opt.Host = host
	}
}

func WithPort(port int) model.DatabaseOptionsFunc {
	return func(opt *model.DatabaseOptions) {
		opt.Port = port
	}
}

func WithDatabaseName(databaseName string) model.DatabaseOptionsFunc {
	return func(opt *model.DatabaseOptions) {
		opt.DatabaseName = databaseName
	}
}

func WithUser(user string) model.DatabaseOptionsFunc {
	return func(opt *model.DatabaseOptions) {
		opt.User = user
	}
}

func WithPassword(password string) model.DatabaseOptionsFunc {
	return func(opt *model.DatabaseOptions) {
		opt.Password = password
	}
}

func WithMinPoolSize(minPoolSize int) model.DatabaseOptionsFunc {
	return func(opt *model.DatabaseOptions) {
		opt.MinPoolSize = minPoolSize
	}
}

func WithMaxPoolSize(maxPoolSize int) model.DatabaseOptionsFunc {
	return func(opt *model.DatabaseOptions) {
		opt.MaxPoolSize = maxPoolSize
	}
}

func WithMaxIdleTimeMS(maxIdleTimeMS int) model.DatabaseOptionsFunc {
	return func(opt *model.DatabaseOptions) {
		opt.MaxIdleTimeMS = maxIdleTimeMS
	}
}

func WithConnectTimeoutMS(connectTimeoutMS int) model.DatabaseOptionsFunc {
	return func(opt *model.DatabaseOptions) {
		opt.ConnectTimeoutMS = connectTimeoutMS
	}
}

func WithWaitQueueTimeoutMS(waitQueueTimeoutMS int) model.DatabaseOptionsFunc {
	return func(opt *model.DatabaseOptions) {
		opt.WaitQueueTimeoutMS = waitQueueTimeoutMS
	}
}

func WithAuthSource(authSource string) model.DatabaseOptionsFunc {
	return func(opt *model.DatabaseOptions) {
		opt.AuthSource = authSource
	}
}

func DatabaseDefaultOpts() *model.DatabaseOptions {
	return &model.DatabaseOptions{
		Host: "127.0.0.1",
		Port: 3306,
	}
}

func (d *MongoDbProvider) GetMongoDbClient() (*mongo.Client, error) {
	if mongoDbClientInstance != nil {
		return mongoDbClientInstance, nil
	}

	query := url.Values{}

	if d.MinPoolSize != 0 {
		query.Add("minPoolSize", strconv.Itoa(d.MinPoolSize))
	}

	if d.MaxPoolSize > 0 {
		query.Add("maxPoolSize", strconv.Itoa(d.MaxPoolSize))
	}

	if d.MaxIdleTimeMS > 0 {
		query.Add("maxIdleTimeMS", strconv.Itoa(d.MaxIdleTimeMS))
	}

	if d.ConnectTimeoutMS > 0 {
		query.Add("connectTimeoutMS", strconv.Itoa(d.ConnectTimeoutMS))
	}

	if d.WaitQueueTimeoutMS > 0 {
		query.Add("waitQueueTimeoutMS", strconv.Itoa(d.WaitQueueTimeoutMS))
	}

	if d.AuthSource != "" {
		query.Add("authSource", d.AuthSource)
	}

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s?%s", d.User, d.Password, d.Host, d.Port, d.DatabaseName, query.Encode())

	serverApi := options.ServerAPI(options.ServerAPIVersion1)
	mongoOptions := options.Client().ApplyURI(uri).SetServerAPIOptions(serverApi)

	client, err := mongo.Connect(context.Background(), mongoOptions)
	if err != nil {
		return nil, err
	}

	if mongoDbClientInstance == nil {
		lockMongoDbClient.Lock()
		defer lockMongoDbClient.Unlock()
		if mongoDbClientInstance == nil {
			mongoDbClientInstance = client
		}
	}

	return mongoDbClientInstance, nil
}
