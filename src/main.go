package main

import (
	"fmt"
	"newt/src/voice"
	"os"
)
   
func main(){
   v := voice.NewElevelabsVoice("JZA3aoJ27TgFUC1cDh5o",os.Getenv("ELEVENLABS_API_KEY") )
   p := voice.NewMpv(2.0,80)
   if len(os.Args) < 2 {
		fmt.Println("Please provide the path to the file as an argument")
		return
	}

	filePath := os.Args[1]

	cont, e := os.ReadFile(filePath)
	if e != nil { fmt.Println("Error reading file:", e); return; }

	t := string(cont)

   fmt.Println(t)
   fn := voice.GenerateFilename("/home/madi/proj/newt/v","mp3")
   v.Generate(t, fn)
   p.Play(fn);

}
