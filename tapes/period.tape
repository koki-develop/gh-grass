Output assets/period.gif

Require gh

Set Shell "bash"
Set FontSize 32
Set Width 2200
Set Height 600
Set PlaybackSpeed 0.6

Hide
Type "go build ." Enter
Type "gh extension remove grass" Enter
Type "gh extension install ." Enter
Type "gh grass --help" Enter
Ctrl+l
Sleep 3s
Show

Type 'gh grass --from "2023-03-01" --to "2023-03-31"' Sleep 500ms Enter

Sleep 5s
