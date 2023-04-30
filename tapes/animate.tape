Output assets/animate.gif

Require gh

Set Shell "bash"
Set FontSize 32
Set Width 2200
Set Height 600
Set PlaybackSpeed 0.5

Hide
Type "go build ." Enter
Type "gh extension remove grass" Enter
Type "gh extension install ." Enter
Type "gh grass --help" Enter
Ctrl+l
Sleep 3s
Show

Type "gh grass --animate" Sleep 500ms Enter
Sleep 3650ms

Sleep 5s
