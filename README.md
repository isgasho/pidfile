# pidfile
pidfile package for Go

```
import github.com/tinystack/pidfile

p := pidfile.New("/var/pidfile.pid")
if err := p.Create(); err != nil {
    log.Fatal(err.Error())
}
defer p.Remove()

currentPid := p.ProcessID()
```
