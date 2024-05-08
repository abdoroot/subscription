# subscription
subscription system using golang and htmx
# export go/bin folder: for local air and templ
export PATH=$PATH:/home/abdoroot/go/bin
abdoroot == user
# Hotreload using templ 
templ generate --watch --proxy="http://localhost:4000" --cmd="air air -c .air.tom"
# create migration using go-migrate
migrate create -ext sql -dir migrations -seq create_users_table