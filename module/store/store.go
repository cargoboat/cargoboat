package store

import (
	"fmt"
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
func Get(group, key string) (value string) {
	value, _ = store.Get(fmt.Sprintf("%s.%s", group, key))
	return
}

// Set ...
func Set(group, key, value string) error {
	err := store.Set(fmt.Sprintf("%s.%s", group, key), value)
	if err == nil {
		err = SetVersion(time.Now().Unix())
	}
	return err
}

// CargoboatConfigVersionKey ...
const CargoboatConfigVersionKey = "cargoboat.config.version"

// GetVersion ...
func GetVersion() (value int64) {
	var result string
	result, _ = store.Get(CargoboatConfigVersionKey)
	value = convert.ToInt64(result)
	return
}

// SetVersion ...
func SetVersion(value int64) error {
	return store.Set(CargoboatConfigVersionKey, convert.ToString(value))
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

// GetAllKeysByPrefix ...
func GetAllKeysByPrefix(prefix string) (keys []string) {
	var err error
	keys, err = store.GetAllKeysByPrefix(prefix)
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

// GetAllByPrefix ...
func GetAllByPrefix(prefix string) (values map[string]string) {
	var err error
	values, err = store.GetAllByPrefix(prefix)
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
