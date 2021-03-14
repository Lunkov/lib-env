package env

import (
  "os"
  "time"
  "strconv"
  "io/ioutil"
  "path/filepath"
  "github.com/golang/glog"
  "github.com/jinzhu/gorm"
)

func SetDef(value string, fallback string) string {
  if len(value) > 0 {
    return value
  }
  return fallback
}

func Get(key, fallback string) string {
  if value, ok := os.LookupEnv(key); ok {
    return value
  }
  return fallback
}

func GetInt(key string, fallback int) int {
  if value, ok := os.LookupEnv(key); ok {
    i, err := strconv.Atoi(value)
    if err != nil {
      return fallback
    }
    return i
  }
  return fallback
}

func fileExists(filename string) bool {
  info, err := os.Stat(filename)
  if os.IsNotExist(err) {
    return false
  }
  return !info.IsDir()
}

func WaitFile(filename string, timeout int) bool {
  step := 5
  d := 0
  for d < timeout {
    if fileExists(filename) {
      glog.Infof("LOG: Load File: %s", filename)
      return true
    }
    glog.Infof("LOG: Wait File: %s", filename)
    time.Sleep(time.Duration(step) * time.Second)
    d += step
  }
  return false
}

type FuncParseFileDB func (*gorm.DB, string, []byte) int
type FuncParseFile func (string, []byte) int

func LoadFromFilesDB(dbh *gorm.DB, scanPath string, extension string, fParse FuncParseFileDB) int {
  count := 0
  errScan := filepath.Walk(scanPath, func(filename string, f os.FileInfo, err error) error {
    if f != nil && f.IsDir() == false && (extension == "" || extension == filepath.Ext(filename)) {
      if glog.V(2) {
        glog.Infof("LOG: ReadFile: %s", filename)
      }
      var err error
      yamlFile, err := ioutil.ReadFile(filename)
      if err != nil {
        glog.Errorf("ERR: ReadFile.yamlFile(%s)  #%v ", filename, err)
      } else {
        count += fParse(dbh, filename, yamlFile)
      }
    }
    return nil
  })
  if glog.V(2) {
    glog.Infof("LOG: Scan Path: %s, Items: %d\n", scanPath, count)
  }
  if errScan != nil {
    glog.Errorf("ERR: ScanPath(%s): %s", scanPath, errScan)
  }

  return count
}

func LoadFromFiles(scanPath string, extension string, fParse FuncParseFile) int {
  count := 0
  errScan := filepath.Walk(scanPath, func(filename string, f os.FileInfo, err error) error {
    if f != nil && f.IsDir() == false && (extension == "" || extension == filepath.Ext(filename)) {
      if glog.V(2) {
        glog.Infof("LOG: ReadFile: %s\n", filename)
      }
      var err error
      yamlFile, err := ioutil.ReadFile(filename)
      if err != nil {
        glog.Errorf("ERR: ReadFile.yamlFile(%s)  #%v ", filename, err)
      } else {
        count += fParse(filename, yamlFile)
      }
    }
    return nil
  })
  if glog.V(2) {
    glog.Infof("LOG: Scan Path: %s, Items: %d\n", scanPath, count)
  }
  if errScan != nil {
    glog.Errorf("ERR: ScanPath(%s): %s", scanPath, errScan)
  }

  return count
}
