## Goenv
[ -d "$HOME/.goenv" ] && export GOENV_ROOT="$HOME/.goenv"

[ -d "$GOENV_ROOT/bin" ] && export PATH="$GOENV_ROOT/bin:$PATH"

if which goenv > /dev/null; then
    eval "$(goenv init -)"
fi
[ -d "$GOROOT/bin" ] && export PATH="$GOROOT/bin:$PATH"
[ -d "$GOPATH/bin" ] && export PATH="$PATH:$GOPATH/bin"
