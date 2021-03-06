Versions       = ["~> 1.10.0", "~> 1.9.0"]
DefaultVersion = "1.10.3"
FileExts       = [".go"]
IgnoreFiles    = [".a", "vendor/"]
DevImages      = ["golang"]
PubImages      = ["alpine"]
ScaffoldFiles  = [
    "cmd/main.go",
    "Makefile",
    "dockerfile",
    "README.md",
    "Gopkg.toml",
    "Gopkg.lock"
]
