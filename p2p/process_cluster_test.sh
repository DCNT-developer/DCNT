#!/bin/bash
echo "run this from DCNT-developer/dcnt eg:"
echo "$ ./p2p/process_cluster_test.sh"
echo
echo "changing directory to dcnt"
cd "$GOPATH/src/github.com/DCNT-developer/dcnt"
rm "$GOPATH/bin/dcnt"
echo "Compiling..."
go install -ldflags "-X github.com/DCNT-developer/dcnt/engine.Build=`git rev-parse HEAD` -X github.com/DCNT-developer/dcnt/engine.dcntVersion=`cat VERSION`"
if [ $? -eq 0 ]; then
     echo "was binary updated? Current:`date`"
    ls -G -lh "$GOPATH/bin/dcnt"

    echo "changing directory to dcnt/p2p"
    cd "$GOPATH/src/github.com/DCNT-developer/dcnt/p2p"
    pkill dcnt
 
    echo "Running..."
    dcnt -exclusive=true -blktime=10 -network="TEST" -networkPort=8118 -peers="127.0.0.1:8119" > testing/node1.out & node0=$!
    sleep 6 
    dcnt -exclusive=true -blktime=10 -network="TEST" -prefix="test2-" -port=9121 -networkPort=8119 -peers="127.0.0.1:8118" > testing/node2.out  & node1=$!
    # sleep 6
    # dcnt -network="TEST" -prefix="test3-" -port=9122 -networkPort=8120 -peers="127.0.0.1:8119" -db=MAP  & node2=$!
    # sleep 6
    # dcnt -network="TEST" -prefix="test4-" -port=9123 -networkPort=8121  -peers="127.0.0.1:8120" -db=MAP  & node3=$!

    tail -f testing/node1.out testing/node2.out  | grep -B 3 -A 15 -e "Network Controller Status Report"  -e "InMsgQueue" & ncsp=$!
    echo
    echo
    sleep 120
    echo
    echo
    echo "Killing processes now..."
    echo
    # kill -2 $node0 $node1 $node2 $node3
    kill -2 $node1 # Kill this first to see how node0 handles it.
    kill -2 $node0 $node2 $node3
    kill  $ncsp
fi

    ./testing/analyze.sh
