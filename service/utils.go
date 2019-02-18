package service

func InterfaceSlice2BytesSlice(ifcs []interface{}) [][]byte {
    bytes := make([][]byte, len(ifcs))
    for i, ifc := range ifcs {
        bytes[i] = ifc.([]byte)
    }
    return bytes
}
