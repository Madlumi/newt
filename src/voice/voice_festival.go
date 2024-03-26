package voice;

import (
)


type FestivalVoice struct{ VoiceId string; }

func (v FestivalVoice)Tell(text string) {
  /* text = removeSpecialChars(text);
   fileName := Vc.OutputFolder + generateFilename("wav");

   cmd := exec.Command("text2wave", "-f", "32000", "-o", fileName, "-eval", "("+v.VoiceId+")")
   cmd.Stdin = strings.NewReader(text+":.........:")

   e := cmd.Run()
   if(m.Er("run tts cmd ",e)){return;}

*/
}


