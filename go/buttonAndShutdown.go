package main

import (
    "fmt"
    "log"
    "log/syslog"
    "time"
    "os"
    "os/exec"
    "encoding/json"
    "io/ioutil"

    "periph.io/x/conn/v3/gpio"
    "periph.io/x/conn/v3/gpio/gpioreg"
    "periph.io/x/host/v3"
)

const duration int64 = 1
const buttonDuration time.Duration = 2
const ledDuration time.Duration = 5
const defaultConfigFileName string = "/usr/local/etc/buttonAndShutdown.cfg"

type Config struct {
    UseSyslog   bool   `json:"usesyslog"`
    UseStdout   bool   `json:"usestdout"`
    LogFileName string `json:"logfilename"`
    ButtonPin   string `json:"buttonpin"`
    LedPin      string `json:"ledpin"`
}

type Logger struct {
    syslogFlag bool
    sysLog *syslog.Writer
    stdoutFlag bool
    logFile *os.File
}

func (out *Logger) init(sFlag bool, s *syslog.Writer, stdFlag bool, fp *os.File) {
    out.syslogFlag = sFlag
    out.sysLog = s
    out.stdoutFlag = stdFlag
    out.logFile = fp
}

func (logger *Logger) log(msg string) {
    if logger.syslogFlag {
        logger.sysLog.Err(msg)
    }
    if logger.stdoutFlag {
        log.Println(msg)
    }
    if logger.logFile != nil {
        logger.logFile.WriteString(time.Now().String()+" : "+ msg+ "\n")
    }
}

func shutdown(logger *Logger) {
    time.Sleep(ledDuration * time.Second)
    logger.log("shutdown start")
    err := exec.Command("shutdown", "-h", "now").Run()
    //err := exec.Command("shutdown", "-h now").Run()
    if err != nil {
        logger.log("Error: can not exec shutdown")
	os.Exit(1)
    }
}

func loop(ButtonPin string, LedPin string, logger *Logger){
    // Load all the drivers:
    if _, err := host.Init(); err != nil {
        logger.log("GPIO initialization error")
        os.Exit(1)
    }

    // Lookup a pin by its number:
    buttonPin := gpioreg.ByName(ButtonPin)
    if buttonPin == nil {
        logger.log("Failed to find Button pin "+ButtonPin)
        os.Exit(1)
    }
    ledPin := gpioreg.ByName(LedPin)
    if ledPin == nil {
        logger.log("Failed to find LED pin "+LedPin)
        os.Exit(1)
    }

    // Set it as input, with an internal pull down resistor:
    if err := buttonPin.In(gpio.PullDown, gpio.BothEdges); err != nil {
        logger.log("Failed to initialize Button pin "+ButtonPin)
        os.Exit(1)
    }
    // pull down pin voltage:
    if err := ledPin.Out(gpio.Low); err != nil {
        logger.log("Failed to initialize LED pin "+LedPin)
        os.Exit(1)
    }

    for {
        buttonPin.WaitForEdge(-1)
	if gpio.High == buttonPin.Read() {
            if false == buttonPin.WaitForEdge(buttonDuration * time.Second) {
                if err := ledPin.Out(gpio.High); err != nil {
                    logger.log("Failed to pull up LED pin "+LedPin)
                    os.Exit(1)
                }
                shutdown(logger)
	    } else {
                if err := ledPin.Out(gpio.Low); err != nil {
                    logger.log("Failed to pull down LED pin "+LedPin)
                    os.Exit(1)
                }
            }
	} else {
            // pull down pin voltage:
            if err := ledPin.Out(gpio.Low); err != nil {
                logger.log("Failed to pull down LED pin "+LedPin)
                os.Exit(1)
            }
	}
    }
}

func Usage() {
    fmt.Println("Usage: button [ConfigFile]")
    os.Exit(1)
}

func main() {
    var configFileName string
    argv := os.Args
    if 2 < len(os.Args) { Usage() }
    if 2 == len(os.Args) {
        configFileName = argv[1]
	if _, err := os.Stat(configFileName); err != nil {
            fmt.Printf("Error: configfile \"%s\" does not exist.\n",configFileName)
            Usage()
	}
    } else {
        configFileName = defaultConfigFileName
    }
    // JSON形式configファイル読み込み
    texts, err := ioutil.ReadFile(configFileName)
    if err != nil {
        log.Fatal(err)
	os.Exit(1)
    }
    // configデータ(JSON)デコード
    var config Config
    if err := json.Unmarshal(texts, &config); err != nil {
        log.Fatal(err)
	os.Exit(1)
    }
    var sysLog *syslog.Writer
    if (true == config.UseSyslog) {
        sysLog, err = syslog.Dial("tcp", "localhost:514",
		syslog.LOG_WARNING|syslog.LOG_DAEMON, "ButtonAndShutdown")
        if err != nil {
            log.Fatal(err)
	    config.UseSyslog = false
        }
    }
    var filePointer *os.File
    if "" != config.LogFileName {
        filePointer, err = os.OpenFile(config.LogFileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0664)
        if err != nil {
            log.Fatal(err)
	    filePointer = nil
        }
    }
    var logger Logger
    logger.init(config.UseSyslog, sysLog, config.UseStdout, filePointer)
    loop(config.ButtonPin, config.LedPin, &logger)
}
