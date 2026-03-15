package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/go-vgo/robotgo"
)

func Core(command string) {
	if command == "1" {
		exec.Command("cmd", "/c", "start code -r").Run()
		exec.Command("cmd", "/c", "start", "https://music.youtube.com/playlist?list=PLuX-LhW3ddzHCYBo92tdtkcNINHyAIHJJ").Run()
		exec.Command("cmd", "/c", "start", "https://github.com/dashboard").Run()
		exec.Command("cmd", "/c", "start", "https://chatgpt.com/").Run()
		exec.Command("cmd", "/c", "start", "https://www.linkedin.com/feed/").Run()
		exec.Command("cmd", "/c", "start", "https://patorjk.com/software/taag/#p=display&f=Big+Money-nw&t=Hermes&x=none&v=4&h=4&w=80&we=false").Run()
		Teste()
	}
	if command == "2" {
		exec.Command("cmd", "/c", "start https://www.youtube.com/").Run()
	}
}

func Teste() {
	time.Sleep(6 * time.Second)
	robotgo.MoveMouse(150, 20)
	robotgo.MouseClick("left", false)
	time.Sleep(1 * time.Second)
	robotgo.MoveMouse(512, 660)
	robotgo.MouseClick("left", false)
	time.Sleep(1 * time.Second)
	robotgo.MoveMouse(610, 260)
	robotgo.MouseClick("left", false)
}

func ProgramPainel() {
	fmt.Print(`
	------------------------------------------------------------------------------------------
	------------------------------------------------------------------------------------------
                   $$\   $$\                                                       
                   $$ |  $$ |                                                      
                   $$ |  $$ | $$$$$$\   $$$$$$\  $$$$$$\$$$$\   $$$$$$\   $$$$$$$\ 
                   $$$$$$$$ |$$  __$$\ $$  __$$\ $$  _$$  _$$\ $$  __$$\ $$  _____|
                   $$  __$$ |$$$$$$$$ |$$ |  \__|$$ / $$ / $$ |$$$$$$$$ |\$$$$$$\  
                   $$ |  $$ |$$   ____|$$ |      $$ | $$ | $$ |$$   ____| \____$$\ 
                   $$ |  $$ |\$$$$$$$\ $$ |      $$ | $$ | $$ |\$$$$$$$\ $$$$$$$  |
                   \__|  \__| \_______|\__|      \__| \__| \__| \_______|\_______/ 
                                                                
	------------------------------------------------------------------------------------------
	------------------------------------------------------------------------------------------
     
	===========
	--Comands
	===========
	1 => Modo Dev
	2 => Modo Jogo
	3 => Sair do Programa
	
	Command:`)
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ProgramPainel()
	if scanner.Scan() {
		command := scanner.Text()
		Core(command)
	}
}
