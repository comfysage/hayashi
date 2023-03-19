package util

import "os"

var HOME, _ = os.UserHomeDir()

var HAYASHI_ROOT string = HOME + "/.hayashi"
var CONFIG_NAME string = ".hayashi.yaml"
var STORE_NAME string = ".store.yaml"

var REPO_ROOT string = HAYASHI_ROOT + "/repo"
var PKG_ROOT string = HAYASHI_ROOT + "/pkg"
