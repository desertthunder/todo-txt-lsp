echo "building server"
go build
echo "installing server"
go install && asdf reshim
echo "compiling client"
pnpm compile

echo "opening log file..."
tail -f "$(echo "console.log(require('os').tmpdir())" | node)/lsp.log"
