go install -ldflags "-X github.com/DCNT-developer/dcnt/engine.Build=`git rev-parse HEAD` -X github.com/DCNT-developer/dcnt/engine.dcntVersion=`cat VERSION`" -v
