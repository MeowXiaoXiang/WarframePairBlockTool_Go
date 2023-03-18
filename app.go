package main

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"syscall"
)

var (
	permissions = 0
	portCombo   = []string{"4950 & 4955",
		"4960 & 4965",
		"4970 & 4975",
		"4980 & 4985",
		"4990 & 4995",
		"3074 & 3080"}
)

// App struct
type App struct {
	ctx   context.Context
	color int
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) CreateRule(age int) bool {
	var port1, port2 string
	portSplit := strings.Split(portCombo[age], " & ")
	port1, port2 = portSplit[0], portSplit[1]
	fmt.Printf("netsh advfirewall firewall add rule name=WarframePairBlockPort protocol=UDP dir=out localport=%s-%s action=block \n", port1, port2)
	args := strings.Split(fmt.Sprintf("advfirewall firewall add rule name=WarframePairBlockPort protocol=UDP dir=out localport=%s-%s action=block", port1, port2), " ")
	cmd := exec.Command("netsh", args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	if err := cmd.Run(); err != nil {
		log.Println(err)
	}
	return a.GetRuleStatus()
}

func (a *App) DelRule() bool {
	fmt.Println("netsh advfirewall firewall delete rule name=WarframePairBlockPort")
	args := strings.Split(fmt.Sprint("advfirewall firewall delete rule name=WarframePairBlockPort"), " ")
	cmd := exec.Command("netsh", args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	if err := cmd.Run(); err != nil {
		log.Println(err)
	}
	return a.GetRuleStatus()
}

func (a *App) CheckRule() {
	cmd := exec.Command("mmc", "wf.msc")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	if err := cmd.Run(); err != nil {
		log.Println(err)
		return
	}
}

func (a *App) GetRuleStatus() bool {
	fmt.Println("netsh advfirewall firewall show rule name=WarframePairBlockPort dir=out")
	args := strings.Split(fmt.Sprint("advfirewall firewall show rule name=WarframePairBlockPort dir=out"), " ")
	cmd := exec.Command("netsh", args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	if err := cmd.Run(); err != nil {
		log.Println(err)
		return true
	} else {
		return false
	}
}
