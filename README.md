# Hermes Prototype

Hermes Prototype is an experimental CLI automation tool written in Go that prepares my development environment automatically.

The goal of this project is to reduce repetitive tasks performed at the beginning of a work session by launching essential tools, websites, and executing small system automations.

This prototype represents the first step of the **Hermes project**, which aims to evolve into a more advanced automation and productivity assistant.

---

## Features

- CLI launcher for Windows
- Development environment startup automation
- Automatic opening of development tools and websites
- Mouse automation using RobotGo
- Simple terminal interface with multiple modes

---

## Modes

### Dev Mode

Automatically prepares the development environment by:

- Opening Visual Studio Code
- Launching commonly used websites
- Starting automation routines

Example websites opened:

- GitHub
- ChatGPT
- LinkedIn
- YouTube Music

After launching the tools, the program executes small mouse automation tasks using RobotGo.

---

### Game Mode

This function has not yet been implemented; it will be assigned in the future with the Dialog system and saving user settings with UI/UX.

## Technologies

- Go
- RobotGo
- Windows command automation (`cmd /c start`)

---

## Preview

Example CLI interface:





