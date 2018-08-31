package main

import (
	"bytes"
	"log"
	"os/exec"

	"github.com/atotto/clipboard"
	"github.com/gen2brain/beeep"
)

func main() {
	cmd := exec.Command("uuidgen")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	uuid := out.String()
	err = clipboard.WriteAll(uuid)
	if err != nil {
		log.Fatalln(err)
	}

	err = beeep.Notify("UUID copied to the clipboard", uuid, "assets/information.png")
	if err != nil {
		panic(err)
	}

	err = beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	if err != nil {
		panic(err)
	}
}
