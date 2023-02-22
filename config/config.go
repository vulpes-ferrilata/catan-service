package config

type Config struct {
	Server ServerConfig `mapstructure:"server"`
	Mongo  MongoConfig  `mapstructure:"mongo"`
	Kafka  KafkaConfig  `mapstructure:"kafka"`
}
