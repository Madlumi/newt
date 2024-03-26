package voice;

import (
   "fmt"
	"bytes"
	"io"
	"net/http"
	"os"
)
//ElevenlabsVoice------------------------------
type ElevenlabsVoice struct{ VoiceID string; Key string; }
const ( ChunkSize = 1024;)
func NewElevelabsVoice(voice string, apiKey string)ElevenlabsVoice{
   var v ElevenlabsVoice;
   v.VoiceID=voice;
   v.Key    =apiKey;
   return v;
}
func (v ElevenlabsVoice)Generate(text string, fn string) {
   //build request
   var Url = "https://api.elevenlabs.io/v1/text-to-speech/"+v.VoiceID
   text = removeSpecialChars(text);
	payload := fmt.Sprintf(`{"text": "%s", "model_id": "eleven_monolingual_v1", "voice_settings": {"stability": 0.5, "similarity_boost": 0.5}}`, text)
	req, e := http.NewRequest("POST", Url, bytes.NewBufferString(payload)); 
   if(Er("creating request",e)){return;}
	req.Header.Set("Content-Type", "application/json"); req.Header.Set("Accept", "audio/mpeg"); req.Header.Set("xi-api-key", v.Key);
	
   //http
   client := &http.Client{}
	resp, e := client.Do(req)
   if(Er("sending request",e)){return;}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK { fmt.Println("Error:", resp.Status); return; }

   //file
   //TODO check is mp3
	out, e := os.Create(fn)
   if(Er("making file",e)){return;}
	defer out.Close()
   _, e = io.Copy(out, resp.Body)
   if(Er("cp resp to file",e)){return;}
}

