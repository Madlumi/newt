package voice
import(
   "fmt"
	"os/exec"
)


//Player---------------------------------------------
type Mpv struct{Speed float32 ; Volume int;}
func NewMpv(speed float32, volume int)Mpv{
   var p Mpv;
   p.Set_Volume(volume);
   p.Set_Speed(speed);
   return p;
}
func (p *Mpv)Set_Volume(v int) { p.Volume=v; }
func (p *Mpv)Set_Speed(s float32) { p.Speed=s; }
func (p Mpv)Play(fn string) {
	cmd := exec.Command("mpv","--volume="+fmt.Sprintf("%d",p.Volume), "--speed="+fmt.Sprintf("%.2f",p.Speed), fn);
   e := cmd.Run(); Er("audio playback",e);
}


