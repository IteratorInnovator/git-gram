package telegram

func HandleCommands(command string, chatId int64) {
	switch (command) {
	case "/start":
		handleStart()
	case "/status":
		handleStatus()
	case "/mute":
		handleMute()
	case "/unmute":
		handleUnmute()
	case "/unlink":
		handleUnlink()
	case "/help":
		handleHelp()
	default: 
		handleInvalidCommand()
	}
}

func handleStart() {

}

func handleStatus() {

}

func handleMute() {

}

func handleUnmute() {

}

func handleUnlink() {

}

func handleHelp() {

}

func handleInvalidCommand() {

}
