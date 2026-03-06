 #  sudo npm install -g @marp-team/marp-cli
 #  go install github.com/radovskyb/watcher/cmd/watcher@latest
 #
 # In markdownmerge folder:
 #
 #     go install
 #
 # In slides folder:
 #
 #     watcher -cmd="markdownmerge -i  slides.txt -o ../main.md"
 PORT=8484 marp --config marp.yml -s .
