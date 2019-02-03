package store

import (
	"time"

	"github.com/cargoboat/storage"
	"github.com/nilorg/pkg/logger"
	"github.com/nilorg/sdk/convert"
	"github.com/spf13/viper"
)

var store storage.Storageer

func initDB() {
	var err error
	if viper.GetString("system.db_type") == "leveldb" {
		store, err = storage.NewLevelDBStorage(viper.GetString("leveldb.path"))
		if err != nil {
			logger.Fatalln(err)
		}
	}
}

// Start 启动存储
func Start() {
	initDB()
	err := SetVersion(time.Now().Unix())
	if err != nil {
		logger.Errorln(err)
	}
}

// Close 关闭
func Close() {
	err := store.Close()
	if err != nil {
		logger.Errorln(err)
	}
}

// Get ...
func Get(key string) (value string) {
	value, _ = store.Get(key)
	return
}

// Set ...
func Set(key string, value string) error {
	err := store.Set(key, value)
	if err == nil {
		err = SetVersion(time.Now().Unix())
	}
	return err
}

const cargoboatConfigVersionKey = "cargoboat.config.version"

// GetVersion ...
func GetVersion() (value int64) {
	var result string
	result, _ = store.Get(cargoboatConfigVersionKey)
	value = convert.ToInt64(result)
	return
}

// SetVersion ...
func SetVersion(value int64) error {
	return store.Set(cargoboatConfigVersionKey, convert.ToString(value))
}

// GetAllKeys ...
func GetAllKeys() (keys []string) {
	var err error
	keys, err = store.GetAllKeys()
	if err != nil {
		keys = []string{}
	}
	return
}

// GetAll ...
func GetAll() (values map[string]string) {
	var err error
	values, err = store.GetAll()
	if err != nil {
		values = make(map[string]string)
	}
	return
}

// Delete ...
func Delete(key string) error {
	err := store.Delete(key)
	if err == nil {
		err = SetVersion(time.Now().Unix())
	}
	return err
}
