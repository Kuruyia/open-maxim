package config

import (
	"errors"
	"os"
	"strconv"
)

// TODO use viper maybe someday

type AdvertConfig struct {
	TextPath  string
	ImagePath string
}

type Config struct {
	ChatPort   int
	HttpPort   int
	ChatAddr   string
	AdvertAddr string
	Banner     AdvertConfig
	Big        AdvertConfig
}

func validatePort(srcPort string) (int, error) {
	port, err := strconv.Atoi(srcPort)
	if err != nil {
		return 0, errors.New("invalid port number")
	}
	if port < 1 || port > 65535 {
		return 0, errors.New("port must be between 1 and 65535")
	}
	return port, nil
}

func GetConfig() (Config, error) {
	chatPort := os.Getenv("MAXIM_CHAT_PORT")
	if chatPort == "" {
		chatPort = "2002"
	}

	httpPort := os.Getenv("MAXIM_HTTP_PORT")
	if httpPort == "" {
		httpPort = "80"
	}

	chatAddr := os.Getenv("MAXIM_CHAT_IP")
	if chatAddr == "" {
		chatAddr = "localhost"
	}

	advertAddr := os.Getenv("MAXIM_ADVERT_ADDR")
	if advertAddr == "" {
		advertAddr = "http://localhost"
	}

	bannerTextPath := os.Getenv("MAXIM_BANNER_TEXT_PATH")
	if bannerTextPath == "" {
		bannerTextPath = "./assets/banner.txt"
	}

	bannerImagePath := os.Getenv("MAXIM_BANNER_IMAGE_PATH")
	if bannerImagePath == "" {
		bannerImagePath = "./assets/banner.png"
	}

	bigTextPath := os.Getenv("MAXIM_BIG_TEXT_PATH")
	if bigTextPath == "" {
		bigTextPath = "./assets/big.txt"
	}

	bigImagePath := os.Getenv("MAXIM_BIG_IMAGE_PATH")
	if bigImagePath == "" {
		bigImagePath = "./assets/big.png"
	}

	chatPortInt, err := validatePort(chatPort)
	if err != nil {
		return Config{}, err
	}

	httpPortInt, err := validatePort(httpPort)
	if err != nil {
		return Config{}, err
	}

	return Config{
		ChatPort:   chatPortInt,
		HttpPort:   httpPortInt,
		ChatAddr:   chatAddr,
		AdvertAddr: advertAddr,
		Banner: AdvertConfig{
			TextPath:  bannerTextPath,
			ImagePath: bannerImagePath,
		},
		Big: AdvertConfig{
			TextPath:  bigTextPath,
			ImagePath: bigImagePath,
		},
	}, nil
}

var CONFIG = Config{}
