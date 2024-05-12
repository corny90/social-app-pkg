# STEPTS IN UPDATING THE PKG

### AFTER CHANGES IN THE COMMON PKG PUSH TO GITHUB
```
git status
git add .
git commit -m "Update changes"
git push origin main
```

### TAG NEW VERSION / RELEASE
```
git tag -a v1.0.0 -m "Release version 1.0.0"
git push origin v1.0.0
```

### IN THE MICROSERVICE PROJECT CHANGE THE VERSION IN GO.MOD
```
go get github.com/corny90/social-app-pkg@v1.0.0
go mod tidy
```

#### SOMETIMES IS NECESSARY TO REMOVE THE GO.SUM FILE, CLEAN CACHE AND RE-RUN GO MOD TIDY
```
rm go.sum
go clean -modcache
go mod tidy
```

### GENERATE THE PROTO FILES
```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative grpc-proto/*.proto
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative *.proto
```