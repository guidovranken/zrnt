package shuffling

import (
    "encoding/binary"
    "sort"
    "reflect"
    "github.com/protolambda/zrnt/eth2/core"
    "github.com/protolambda/zrnt/eth2/util/shuffling"
)

func Fuzz(data []byte) int {
    if len(data) < 32 {
        return 0
    }
    var seed core.Root
    copy(seed[:], data[:32])

    validatorIndices := []core.ValidatorIndex{}
    for i := 32; i < len(data); i += 8 {
        if i + 8 > len(data) {
            break
        }
        validatorIndices = append(validatorIndices, core.ValidatorIndex(binary.LittleEndian.Uint64(data[i:i+8])))
    }

    validatorIndicesSorted := make([]core.ValidatorIndex, len(validatorIndices))
    copy(validatorIndicesSorted, validatorIndices)
    sort.Slice(validatorIndicesSorted, func(i, j int) bool { return validatorIndicesSorted[i] < validatorIndicesSorted[j] })

    shuffled := make([]core.ValidatorIndex, len(validatorIndices))
    copy(shuffled, validatorIndices)
    shuffling.ShuffleList(shuffled, seed)

    shuffledSorted := make([]core.ValidatorIndex, len(shuffled))
    copy(shuffledSorted, shuffled)
    sort.Slice(shuffledSorted, func(i, j int) bool { return shuffledSorted[i] < shuffledSorted[j] })

    unshuffled := make([]core.ValidatorIndex, len(shuffled))
    copy(unshuffled, shuffled)
    shuffling.UnshuffleList(unshuffled, seed)

    if reflect.DeepEqual(validatorIndices, unshuffled) != true {
        panic("input != unshuffle(shuffle(input))")
    }

    if reflect.DeepEqual(validatorIndicesSorted, shuffledSorted) != true {
        panic("sorted input slice != sorted shuffled slice")
    }

    return 0
}

