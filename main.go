package main

import (
	"bufio"
	_ "encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	_ "path/filepath"
	"strconv"
	_ "time"

	"github.com/go-vgo/robotgo"
	"github.com/tidwall/gjson"
)

var PathJSON = filepath.Join(os.Getenv("APPDATA"), "..", "Local", "Hermes", "Data.json")
var ContentJSON string

type Config struct {
	Taxa_de_aprovacao string   `json:"name"`
	Localizacao       []string `json:"return x and y in the EyesHermes OBJECT"`
	Encontrado        bool     `json:"Return True or False"`
}

func LoadJSON() {
	if _, err := os.Stat(PathJSON); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Erro: o arquivo JSON não existe.")
			return
		}
		fmt.Println("Erro:", err)
		return
	}

	contentFile, err := os.ReadFile(PathJSON)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	if !gjson.ValidBytes(contentFile) {
		fmt.Println("Erro: JSON inválido.")
		return
	}

	ContentJSON = string(contentFile)
}

func Core(command string) {
	if command == "1" {
		exec.Command("cmd", "/c", "start code -r").Run()
		exec.Command("cmd", "/c", "start", "https://music.youtube.com/playlist?list=PLuX-LhW3ddzHCYBo92tdtkcNINHyAIHJJ").Run()
		exec.Command("cmd", "/c", "start", "https://github.com/dashboard").Run()
		exec.Command("cmd", "/c", "start", "https://chatgpt.com/").Run()
		exec.Command("cmd", "/c", "start", "https://www.linkedin.com/feed/").Run()
		exec.Command("cmd", "/c", "start", "https://patorjk.com/software/taag/#p=display&f=Big+Money-nw&t=Hermes&x=none&v=4&h=4&w=80&we=false").Run()
		AutomationMouse()
	}
	if command == "2" {
		exec.Command("cmd", "/c", "start https://www.youtube.com/").Run()
	}
	if command == "3" {
		CleanTerminal()
		PainelCreateTask()
	}
	if command == "4" {
		os.Exit(0)
	}
}
func CleanTerminal() {
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
}
func AutomationMouse() {
	EyeHermesPATH := filepath.Join("")
	EyehermesProcess := exec.Command("cmd", "/c", "start", EyeHermesPATH)
	fmt.Print(EyehermesProcess)
	if _, err := os.Stat(PathJSON); err == nil {
		ContentFile, err := os.ReadFile(PathJSON)
		if err != nil {
			fmt.Println("Erro ao abrir o arquivo ", err)
			return
		}
		x := gjson.Get(string(ContentFile), "Localizacao.x")
		y := gjson.Get(string(ContentFile), "Localizacao.y")

		strX := x.String()
		strY := y.String()

		numX, err := strconv.Atoi(strX)
		if err != nil {
			fmt.Println("Erro na conversão para o tipo inteiro")
		}

		numY, err := strconv.Atoi(strY)
		if err != nil {
			fmt.Println("Erro na conversão para o tipo inteiro")
		}
		robotgo.Move(numX, numY)
		robotgo.Click()
	} else if os.IsNotExist(err) {

	} else {
		fmt.Println("Erro ao verificar o arquivo", err)
	}
}

func PainelCreateTask() {
	CleanTerminal()
	fmt.Print(`
	==================================
	--Commands for creating a process
	==================================

	1-args- Set Name Process
	2-args- Set Image recognition(S/N)
	3-args- Set Path to the image; note: If you are not going to use image recognition, just leave it blank.
	4-args- Set Mouse Position x (You don't need to specify if you selected Image Recognition).
	5-args- Set Mouse Position y (You don't need to specify if you selected Image Recognition.)
	6-args- Set Mouse Button(LEFT/RIGHT)
	0-args- Finish
`)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\nCommand (1-5, 0 to finish): ")
		if !scanner.Scan() {
			return
		}
		command := scanner.Text()

		switch command {
		case "0":
			return
		case "1":
			fmt.Print("Set Name Process: ")
		case "2":
			fmt.Print("Image recognition (S/N): ")
		case "3":
			fmt.Print("Path: ")
		case "4":
			fmt.Print("Mouse Position X: ")
		case "5":
			fmt.Print("Mouse Position Y: ")
		case "6":
			fmt.Print("Mouse Button (LEFT/RIGHT): ")
		default:
			fmt.Println("Comando inválido.")
			continue
		}

		if !scanner.Scan() {
			return
		}
		value := scanner.Text()

	}
}

func ProgramPainel() {
	fmt.Println(`
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
	------------------------------------------------------------------------------------------ `)
}

func CommandsPainelInit() {
	fmt.Print(`
	===========
	--Comands
	===========
	1 => Modo Dev
	2 => Modo Jogo
	3 => Definir Novo Processo
	4 => Sair do Programa
	
	Command:`)
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		ProgramPainel()
		CommandsPainelInit()
		if !scanner.Scan() {
			break
		}

		command := scanner.Text()
		Core(command)
	}

}
