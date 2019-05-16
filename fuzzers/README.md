# Build

```sh
pushd -n $(pwd)

CC=clang
CXX=clang++
wget https://dl.google.com/go/go1.12.linux-amd64.tar.gz
tar zxvf go1.12.linux-amd64.tar.gz
export GOROOT=`realpath go`
export GOPATH=$GOROOT/packages
mkdir $GOPATH
export PATH=$GOROOT/bin:$PATH
export PATH=$GOROOT/packages/bin:$PATH

go get github.com/dvyukov/go-fuzz
go get golang.org/x/tools/go/packages
go build github.com/dvyukov/go-fuzz/go-fuzz-build

go get gopkg.in/yaml.v2

mkdir -p $GOPATH/src/github.com/protolambda
cd $GOPATH/src/github.com/protolambda

git clone https://github.com/ethereum/eth2.0-specs
cd eth2.0-specs
git checkout v0.6.0
cd ..

git clone https://github.com/guidovranken/zrnt
cd zrnt
git checkout fuzzing
cd eth2/core/
go generate

# Back to root directory
popd

./go-fuzz-build -tags preset_minimal -o SSZ.a -libfuzzer github.com/protolambda/zrnt/fuzzers/SSZ && clang++ -fsanitize=fuzzer SSZ.a
./go-fuzz-build -tags preset_minimal -o shuffling.a -libfuzzer github.com/protolambda/zrnt/fuzzers/shuffling && clang++ -fsanitize=fuzzer shuffling.a
./go-fuzz-build -tags preset_minimal -o validatorset.a -libfuzzer github.com/protolambda/zrnt/fuzzers/validatorset && clang++ -fsanitize=fuzzer validatorset.a

```
