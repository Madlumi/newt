package voice;
import (
   "strings"
   "unicode"
   "time"
   "fmt"
)

type Voice  interface { Generate(text string, fn string); }
type Player interface { Play(fn string); }
//helpers-------------
func removeSpecialChars(s string) string {
   var clean string
   allowed := " .?,!*"
   for _, char := range s {
      // Check if the character is alphanumeric or is an allowed special character
      if unicode.IsLetter(char) || unicode.IsNumber(char) || strings.ContainsRune(allowed, char) {
         clean += string(char)
      }
   }
   return clean
}

func GenerateFilename(dir string,filetype string) string {
   t := time.Now()
   return fmt.Sprintf("%s/%d-%02d-%02d_%02d-%02d-%02d.%s",dir, t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), filetype)
}

func Er(msg string, e error)bool{ if(e!=nil){ fmt.Println("error: ", msg, ";\n ", e); return true;}; return false; }
