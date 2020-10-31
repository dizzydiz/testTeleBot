package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/gofrs/uuid"
)

const startupMessage = `[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;54;48;5;39m [38;5;54;48;5;39m [38;5;54;48;5;39m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [0m
[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;39;48;5;39m [38;5;21;48;5;45m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;92;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;93;48;5;45m�[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [0m
[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;39;48;5;39m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;92;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;92;48;5;45m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [0m
[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;99;48;5;39m [38;5;39;48;5;39m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;31;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;204;48;5;17m [38;5;204;48;5;17m [38;5;92;48;5;45m [38;5;54;48;5;39m [38;5;92;48;5;45m [38;5;92;48;5;45m [38;5;92;48;5;45m [38;5;92;48;5;45m [38;5;92;48;5;45m [38;5;92;48;5;45m [38;5;212;48;5;24m [38;5;204;48;5;17m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [0m
[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;125;48;5;24m [38;5;39;48;5;39m [38;5;117;48;5;39m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;92;48;5;45m [38;5;55;48;5;39m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;92;48;5;45m [38;5;92;48;5;39m [38;5;62;48;5;17m�[38;5;62;48;5;17m�[38;5;62;48;5;17m�[38;5;62;48;5;17m�[38;5;62;48;5;17m�[38;5;62;48;5;17m�[38;5;62;48;5;17m�[38;5;69;48;5;18m�[38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;99;48;5;45m [0m
[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;39;48;5;39m [38;5;62;48;5;17m�[38;5;62;48;5;17m�[38;5;62;48;5;17m�[38;5;62;48;5;17m�[38;5;62;48;5;17m�[38;5;62;48;5;17m�[38;5;62;48;5;17m�[38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;25;48;5;33m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;38;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;231;48;5;231m�[38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;227;48;5;227m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [0m
[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;105;48;5;122m�[38;5;32;48;5;159m�[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;211;48;5;234m�[38;5;39;48;5;159m�[38;5;1;48;5;16m [38;5;57;48;5;51m�[38;5;57;48;5;51m�[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;27;48;5;159m [38;5;117;48;5;159m�[38;5;39;48;5;159m�[38;5;75;48;5;159m�[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;75;48;5;159m�[38;5;25;48;5;159m�[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;31;48;5;39m [38;5;31;48;5;39m [38;5;56;48;5;239m�[38;5;56;48;5;239m�[38;5;56;48;5;239m�[38;5;56;48;5;239m�[38;5;56;48;5;239m�[38;5;56;48;5;239m�[38;5;56;48;5;239m�[38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;32;48;5;39m [38;5;32;48;5;33m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;45;48;5;81m [38;5;231;48;5;231m�[38;5;226;48;5;226m [38;5;62;48;5;17m�[38;5;4;48;5;17m�[38;5;62;48;5;17m�[38;5;62;48;5;17m�[38;5;226;48;5;226m [38;5;117;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;171;48;5;39m [0m
[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;69;48;5;122m�[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;32;48;5;159m�[38;5;195;48;5;195m�[38;5;33;48;5;159m�[38;5;75;48;5;159m�[38;5;25;48;5;195m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;211;48;5;234m [38;5;1;48;5;16m [38;5;99;48;5;73m�[38;5;99;48;5;159m�[38;5;161;48;5;26m�[38;5;21;48;5;87m�[38;5;159;48;5;159m [38;5;51;48;5;51m [38;5;51;48;5;51m [38;5;51;48;5;51m [38;5;51;48;5;51m [38;5;51;48;5;51m [38;5;51;48;5;51m [38;5;51;48;5;51m [38;5;51;48;5;51m [38;5;51;48;5;51m [38;5;136;48;5;253m�[38;5;136;48;5;253m�[38;5;136;48;5;253m�[38;5;178;48;5;253m�[38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;231;48;5;231m�[38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;25;48;5;33m [38;5;117;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;231;48;5;231m�[38;5;227;48;5;227m [38;5;226;48;5;226m [38;5;62;48;5;17m�[38;5;62;48;5;17m�[38;5;93;48;5;239m�[38;5;230;48;5;230m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;1;48;5;16m [38;5;1;48;5;16m [0m
[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;75;48;5;159m�[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;117;48;5;159m�[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;75;48;5;159m�[38;5;21;48;5;159m�[38;5;39;48;5;159m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;135;48;5;45m�[38;5;57;48;5;87m�[38;5;21;48;5;87m�[38;5;51;48;5;51m [38;5;51;48;5;51m [38;5;51;48;5;51m [38;5;51;48;5;51m [38;5;51;48;5;51m [38;5;51;48;5;51m [38;5;51;48;5;51m [38;5;123;48;5;123m [38;5;195;48;5;195m [38;5;195;48;5;195m [38;5;195;48;5;195m [38;5;195;48;5;195m [38;5;195;48;5;195m [38;5;195;48;5;195m [38;5;195;48;5;195m [38;5;136;48;5;253m�[38;5;136;48;5;253m�[38;5;221;48;5;224m�[38;5;94;48;5;253m�[38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;117;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;117;48;5;39m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [0m
[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;57;48;5;87m�[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;55;48;5;45m�[38;5;57;48;5;87m�[38;5;57;48;5;51m�[38;5;57;48;5;87m�[38;5;99;48;5;87m�[38;5;141;48;5;87m�[38;5;57;48;5;51m�[38;5;166;48;5;249m�[38;5;166;48;5;249m�[38;5;166;48;5;249m�[38;5;43;48;5;145m�[38;5;119;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;226;48;5;226m [38;5;62;48;5;17m�[38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;200;48;5;213m [38;5;200;48;5;213m [38;5;26;48;5;75m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;38;48;5;45m [38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;87;48;5;255m�[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [0m
[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;167;48;5;17m�[38;5;62;48;5;17m�[38;5;62;48;5;17m�[38;5;62;48;5;17m�[38;5;62;48;5;17m�[38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;62;48;5;17m�[38;5;62;48;5;17m�[38;5;62;48;5;17m�[38;5;69;48;5;18m�[38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;200;48;5;213m [38;5;200;48;5;213m [38;5;200;48;5;213m [38;5;200;48;5;213m [38;5;219;48;5;219m [38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;109;48;5;231m�[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [0m
[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;61;48;5;24m�[38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;111;48;5;24m [38;5;62;48;5;17m�[38;5;62;48;5;17m�[38;5;61;48;5;60m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;162;48;5;205m [38;5;162;48;5;205m [38;5;162;48;5;205m [38;5;89;48;5;212m [38;5;200;48;5;213m [38;5;200;48;5;213m [38;5;200;48;5;213m [38;5;200;48;5;213m [38;5;200;48;5;213m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [0m
[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;27;48;5;123m�[38;5;1;48;5;16m [38;5;135;48;5;51m�[38;5;204;48;5;233m [38;5;1;48;5;16m [38;5;135;48;5;51m�[38;5;135;48;5;51m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;165;48;5;32m [38;5;135;48;5;39m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;45;48;5;45m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;62;48;5;17m�[38;5;62;48;5;17m�[38;5;62;48;5;17m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;162;48;5;205m [38;5;212;48;5;205m [38;5;162;48;5;205m [38;5;162;48;5;205m [38;5;200;48;5;213m [38;5;200;48;5;213m [38;5;200;48;5;213m [38;5;84;48;5;212m�[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [0m
[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;135;48;5;51m [38;5;45;48;5;231m�[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;99;48;5;51m�[38;5;1;48;5;16m [38;5;55;48;5;45m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;56;48;5;51m�[38;5;45;48;5;231m�[38;5;56;48;5;51m [38;5;1;48;5;16m [38;5;111;48;5;87m�[38;5;69;48;5;87m�[38;5;1;48;5;16m [38;5;87;48;5;87m [38;5;141;48;5;51m�[38;5;39;48;5;39m [38;5;117;48;5;39m [38;5;39;48;5;39m [38;5;24;48;5;39m [38;5;39;48;5;39m [38;5;117;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;39;48;5;39m [38;5;92;48;5;39m�[38;5;165;48;5;33m [38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;199;48;5;212m [38;5;212;48;5;206m [38;5;225;48;5;225m [38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [0m
[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;56;48;5;51m�[38;5;141;48;5;87m�[38;5;123;48;5;231m�[38;5;31;48;5;195m [38;5;57;48;5;73m�[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;51;48;5;51m [38;5;56;48;5;51m�[38;5;81;48;5;195m�[38;5;55;48;5;45m [38;5;99;48;5;51m�[38;5;177;48;5;38m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;123;48;5;255m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;231;48;5;231m�[38;5;51;48;5;255m�[38;5;168;48;5;241m�[38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [38;5;1;48;5;16m [0m
[0m`

func main() {

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		codeParams, ok := r.URL.Query()["code"]
		if ok && len(codeParams) > 0 {
			statusCode, _ := strconv.Atoi(codeParams[0])
			if statusCode >= 200 && statusCode < 600 {
				w.WriteHeader(statusCode)
			}
		}
		requestID := uuid.Must(uuid.NewV4())
		fmt.Fprintf(w, requestID.String())
	})

	http.HandleFunc("/", HandleTelegramWebHook)
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	// for _, encodedRoute := range strings.Split(os.Getenv("ROUTES"), ",") {
	// 	if encodedRoute == "" {
	// 		continue
	// 	}
	// 	pathAndBody := strings.SplitN(encodedRoute, "=", 2)
	// 	path, body := pathAndBody[0], pathAndBody[1]
	// 	http.HandleFunc("/"+path, func(w http.ResponseWriter, r *http.Request) {
	// 		fmt.Fprint(w, body)
	// 	})
	// }

	bindAddr := fmt.Sprintf(":%s", port)
	lines := strings.Split(startupMessage, "\n")
	fmt.Println()
	for _, line := range lines {
		fmt.Println(line)
	}
	fmt.Println()
	fmt.Printf("==> Server listening at %s 🚀\n", bindAddr)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		panic(err)
	}
}

// Define a few constants and variable to handle different commands
const punchCommand string = "/punch"

var lenPunchCommand int = len(punchCommand)

const startCommand string = "/start"

var lenStartCommand int = len(startCommand)

const botTag string = "@HemdaBot"

var lenBotTag int = len(botTag)

// Pass token and sensible APIs through environment variables
const telegramApiBaseUrl string = "https://api.telegram.org/bot"
const telegramApiSendMessage string = "/sendMessage"
const telegramTokenEnv string = "TELEGRAM_BOT_TOKEN"

var telegramApi string = telegramApiBaseUrl + os.Getenv(telegramTokenEnv) + telegramApiSendMessage

// Update is a Telegram object that we receive every time an user interacts with the bot.
type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

// Implements the fmt.String interface to get the representation of an Update as a string.
func (u Update) String() string {
	return fmt.Sprintf("(update id: %d, message: %s)", u.UpdateId, u.Message)
}

// Message is a Telegram object that can be found in an update.
// Note that not all Update contains a Message. Update for an Inline Query doesn't.
type Message struct {
	Text     string   `json:"text"`
	Chat     Chat     `json:"chat"`
	Audio    Audio    `json:"audio"`
	Voice    Voice    `json:"voice"`
	Document Document `json:"document"`
}

// Implements the fmt.String interface to get the representation of a Message as a string.
func (m Message) String() string {
	return fmt.Sprintf("(text: %s, chat: %s, audio %s)", m.Text, m.Chat, m.Audio)
}

// Audio message has extra attributes
type Audio struct {
	FileId   string `json:"file_id"`
	Duration int    `json:"duration"`
}

// Implements the fmt.String interface to get the representation of an Audio as a string.
func (a Audio) String() string {
	return fmt.Sprintf("(file id: %s, duration: %d)", a.FileId, a.Duration)
}

// Voice Message can be summarized with similar attribute as an Audio message for our use case.
type Voice Audio

// Document Message refer to a file sent.
type Document struct {
	FileId   string `json:"file_id"`
	FileName string `json:"file_name"`
}

// Implements the fmt.String interface to get the representation of an Document as a string.
func (d Document) String() string {
	return fmt.Sprintf("(file id: %s, file name: %s)", d.FileId, d.FileName)
}

// A Chat indicates the conversation to which the Message belongs.
type Chat struct {
	Id int `json:"id"`
}

// Implements the fmt.String interface to get the representation of a Chat as a string.
func (c Chat) String() string {
	return fmt.Sprintf("(id: %d)", c.Id)
}

// A Lyric is generated by the RapLyrics service.
type Lyric struct {
	Punch string `json:"output"`
}

// HandleTelegramWebHook sends a message back to the chat with a punchline starting by the message provided by the user.
func HandleTelegramWebHook(w http.ResponseWriter, r *http.Request) {

	// Parse incoming request
	var update, err = parseTelegramRequest(r)
	if err != nil {
		log.Printf("error parsing update, %s", err.Error())
		return
	}

	// Sanitize input
	// var sanitizedSeed = sanitize(update.Message.Text)

	// Call RapLyrics to get a punchline
	// var lyric, errRapLyrics = getPunchline(sanitizedSeed)
	// if errRapLyrics != nil {
	// 	log.Printf("got error when calling RapLyrics API %s", errRapLyrics.Error())
	// 	return
	// }

	lyric := update.Message.Text

	// Send the punchline back to Telegram
	var telegramResponseBody, errTelegram = sendTextToTelegramChat(update.Message.Chat.Id, lyric)
	if errTelegram != nil {
		log.Printf("got error %s from telegram, response body is %s", errTelegram.Error(), telegramResponseBody)
	} else {
		log.Printf("punchline %s successfully distributed to chat id %d", lyric, update.Message.Chat.Id)
	}
}

// parseTelegramRequest handles incoming update from the Telegram web hook
func parseTelegramRequest(r *http.Request) (*Update, error) {
	var update Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		log.Printf("could not decode incoming update %s", err.Error())
		return nil, err
	}
	if update.UpdateId == 0 {
		log.Printf("invalid update id, got update id = 0")
		return nil, errors.New("invalid update id of 0 indicates failure to parse incoming update")
	}
	return &update, nil
}

// sanitize remove clutter like /start /punch or the bot name from the string s passed as input
func sanitize(s string) string {
	if len(s) >= lenStartCommand {
		if s[:lenStartCommand] == startCommand {
			s = s[lenStartCommand:]
		}
	}

	if len(s) >= lenPunchCommand {
		if s[:lenPunchCommand] == punchCommand {
			s = s[lenPunchCommand:]
		}
	}
	if len(s) >= lenBotTag {
		if s[:lenBotTag] == botTag {
			s = s[lenBotTag:]
		}
	}
	return s
}

// sendTextToTelegramChat sends a text message to the Telegram chat identified by its chat Id
func sendTextToTelegramChat(chatId int, text string) (string, error) {

	log.Printf("Sending %s to chat_id: %d", text, chatId)
	response, err := http.PostForm(
		telegramApi,
		url.Values{
			"chat_id": {strconv.Itoa(chatId)},
			"text":    {text},
		})

	if err != nil {
		log.Printf("error when posting text to the chat: %s", err.Error())
		return "", err
	}
	defer response.Body.Close()

	var bodyBytes, errRead = ioutil.ReadAll(response.Body)
	if errRead != nil {
		log.Printf("error in parsing telegram answer %s", errRead.Error())
		return "", err
	}
	bodyString := string(bodyBytes)
	log.Printf("Body of Telegram Response: %s", bodyString)

	return bodyString, nil
}
