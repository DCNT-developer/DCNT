# dcnt Docker Helper

The dcnt Docker Helper is a simple tool to help build and run dcnt as a container

## Prerequisites

You must have at least Docker v17 installed on your system.

Having this repo cloned helps too 😇

## Build
From wherever you have cloned this repo, run

`docker build -t dcnt_container .`

(yes, you can replace **dcnt_container** with whatever you want to call the container.  e.g. **dcnt**, **foo**, etc.)

#### Cross-Compile
To cross-compile for a different target, you can pass in a `build-arg` as so

`docker build -t dcnt_container --build-arg GOOS=darwin .`

## Run
#### No Persistence
`docker run --rm -p 8090:8090 dcnt_container`
  
* This will start up **dcnt** with no flags.
* The Control Panel is accessible at port 8090  
* When the container terminates, all data will be lost
* **Note** - In the above, replace **dcnt_container** with whatever you called it when you built it - e.g. **dcnt**, **foo**, etc.

#### With Persistence
1. `docker volume create dcnt_volume`
2. `docker run --rm -v $(PWD)/dcnt.conf:/source -v dcnt_volume:/destination busybox /bin/cp /source /destination/dcnt.conf`
3. `docker run --rm -p 8090:8090 -v dcnt_volume:/root/.factom/m2 dcnt_container`

* This will start up **dcnt** with no flags.
* The Control Panel is accessible at port 8090  
* When the container terminates, the data will remain persisted in the volume **dcnt_volume**
* The above copies **dcnt.conf** from the local directory into the container. Put _your_ version in there, or change the path appropriately.
* **Note**.  In the above
   * replace **dcnt_container** with whatever you called it when you built it - e.g. **dcnt**, **foo**, etc.
   * replace **dcnt_volume** with whatever you might want to call it - e.g. **myvolume**, **barbaz**, etc.

#### Additional Flags
In all cases, you can startup with additional flags by passing them at the end of the docker command, e.g.

`docker run --rm -p 8090:8090 dcnt_container -port 9999`


## Copy
So yeah, you want to get your binary _out_ of the container. To do so, you basically mount your target into the container, and copy the binary over, like so


```
docker run --rm --entrypoint='' \
	-v <FULLY_QUALIFIED_PATH_TO_TARGET_DIRECTORY>:/destination dcnt_container \
	/bin/cp /go/bin/dcnt /destination
```

e.g.

```
docker run --rm --entrypoint='' \
-v /tmp:/destination dcnt_container \
/bin/cp /go/bin/dcnt /destination
```

which will copy the binary to `/tmp/dcnt`

**Note** : You should replace **dcnt_container** with whatever you called it in the  [build](#build) section above  e.g. **dcnt**, **foo**, etc.

#### Cross-Compile
If you cross-compiled to a different target, your binary will be in `/go/bin/<target>/dcnt`.  e.g. If you built with `--build-arg GOOS=darwin`, then you can copy out the binary with

```
docker run --rm --entrypoint='' \
-v <FULLY_QUALIFIED_PATH_TO_TARGET_DIRECTORY>:/destination \
dcnt_container \
/bin/cp /go/bin/darwin_amd64/dcnt /destination
```

e.g.

```
docker run --rm --entrypoint='' 
-v /tmp:/destination dcnt_container \
/bin/cp /go/bin/darwin_amd64/dcnt /destination
``` 

which will copy the darwin_amd64 version of the binary to `/tmp/dcnt`

**Note** : You should replace **dcnt_container** with whatever you called it in the **build** section above  e.g. **dcnt**, **foo**, etc.
