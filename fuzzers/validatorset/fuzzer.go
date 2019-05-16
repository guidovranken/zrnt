package validatorset

import (
    "encoding/binary"
    "reflect"
    "github.com/protolambda/zrnt/eth2/core"
)

func Fuzz(data []byte) int {
    var vs1 core.ValidatorSet
    var vs2 core.ValidatorSet

    i := 0

    /* n = 0: update vs1, n = 1: update v2 */
    for n := 0; n < 2; n++ {
        for ; i < len(data); i += 9 {
            if i > len(data) {
                break
            }

            /* Decide whether to proceed or move to vs2 */
            if n == 0 && data[i] & 1 == 1 {
                break
            }

            if i + 1 + 8 > len(data) {
                break
            }

            if n == 0 {
                vs1 = append(vs1, core.ValidatorIndex(binary.LittleEndian.Uint64(data[i+1:i+1+8])))
            } else {
                vs2 = append(vs2, core.ValidatorIndex(binary.LittleEndian.Uint64(data[i+1:i+1+8])))
            }
        }
    }

    if len(vs1) != 0 {
        vs1Copy:= make(core.ValidatorSet, len(vs1))
        copy(vs1Copy, vs1)
        vs1Copy.Dedup()
    }

    if len(vs1) != 0 && len(vs2) != 0 {
        /* XXX Crashes (index out of range) */
        vs1Copy:= make(core.ValidatorSet, len(vs1))
        copy(vs1Copy, vs1)

        vs2Copy:= make(core.ValidatorSet, len(vs2))
        copy(vs2Copy, vs2)

        vs1Copy.Intersects(vs2Copy)

        if reflect.DeepEqual(vs1, vs1Copy) != true || reflect.DeepEqual(vs1, vs1Copy) != true {
            panic("ValidatorSet.Intersects altered input")
        }
    }

    if len(vs1) != 0 && len(vs2) != 0 {
        /* XXX Hang */
        vs1.ZigZagJoin(vs2, nil, nil)
    }

    return 0
}
