package config

import (
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/mohitsinghs/wormholes/constants"
	"github.com/oschwald/geoip2-golang"
)

type PipeConfig struct {
	Streams   int
	BatchSize int
}

func DefaultPipe() PipeConfig {
	return PipeConfig{
		Streams:   constants.STREAMS,
		BatchSize: constants.BATCH_SIZE,
	}
}

func (p *PipeConfig) OpenDB() *geoip2.Reader {
	dbPath := cityPath()
	db, err := geoip2.Open(dbPath)
	if err != nil {
		log.Panicln("Missing GeoLite2-City.mmdb")
	}
	return db
}

// Get path for GeoLite2-City.mmdb
func cityPath() string {
	home, err := os.UserConfigDir()
	if err != nil {
		log.Printf("Error getting home dir : %v", err)
	}
	cfgDir := filepath.Join(home, constants.DEFAULT_DIR)
	_, err = os.Stat(cfgDir)
	log.Println(err)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(cfgDir, constants.DIR_PERM); err != nil {
			return ""
		}
	}
	return path.Join(cfgDir, constants.CITY_DB)
}
