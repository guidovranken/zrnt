package SSZ

import "github.com/protolambda/zrnt/eth2/util/ssz"

func Fuzz(data []byte) int {
    ssz.SSZEncode(data)
    return 0
}
