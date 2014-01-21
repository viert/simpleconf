package simpleconf

import (
  "testing"
  "os"
)

func TestParseConfig(t *testing.T) {
  testConfig := "key1 = value1\nkey2 = 3000 # some comment\n # comment line\n key3 = value3"
  fileName := "simpleconf.conf"
  tempFileName := os.TempDir() + "/" + fileName
  tempFile, err := os.Create(tempFileName)
  if (err != nil) { t.Errorf("Can't create temp file for test: %s", err) }
  defer os.Remove(tempFile.Name())
  tempFile.WriteString(testConfig)
  tempFile.Close()
  
  parser, err := ParseConfig(tempFile.Name())
  if (err != nil) { t.Errorf("Error parsing config: %s", err) }

  var v1, v3 string
  var v2 int

  v1, err = parser.GetString("key1")
  if (err != nil) { t.Errorf("Error retrieving key1: %s", err) }

  v2, err = parser.GetInt("key2")
  if (err != nil) { t.Errorf("Error retrieving key2: %s", err) }

  v3, err = parser.GetString("key3")
  if (err != nil) { t.Errorf("Error retrieving key3: %s", err) }

  if (v1 != "value1") { t.Errorf("%q expected, but %q got", "value1", v1) }
  if (v2 != 3000) { t.Errorf("%d expected, but %d got", 3000, v2) }
  if (v3 != "value3") { t.Errorf("%q expected, but %q got", "value3", v3) }
}
