# swearjar-go

Install:

    go get github.com/snicol/swearjar-go

# Notes 

This is a very rough project for hobby/testing purposes but I will fix it up when I can. Minimal tests exist.

**FYI:** There's a race condition in Scorecard() and the JSON is embedded in the code (needs go-bindata asset packing). 

# Acknowledgements

This is based on [Swearjar](github.com/joshbuddy/swearjar) for Ruby and uses the JSON config from
 [swearjar-node](https://github.com/raymondjavaxx/swearjar-node).