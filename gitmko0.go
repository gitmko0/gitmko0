package gitmko0

import (
     "time"
     "bytes"
     "encoding/binary"
)

func GetTimestampSeconds() uint64 {
     return uint64(time.Now().UnixNano() / int64(time.Second))
}

func GetTimestampMilliseconds() uint64 {
     return uint64(time.Now().UnixNano() / int64(time.Millisecond))
}

func GetTimestampBytes() []byte{
        buf := new(bytes.Buffer)
        binary.Write(buf, binary.LittleEndian, time.Now().UTC().Unix())
        buf.Truncate(5)
        return buf.Bytes() 
}
