package md5sum

import (
        "crypto/md5"
        "encoding/hex"
        "fmt"
        "io"
        "os"
)

func md5sum(filePath string) (string, error) {
        var returnMD5String string
        file, err := os.Open(filePath)
        if err != nil {
                return returnMD5String, err
        }
        defer file.Close()
        hash := md5.New()
        if _, err := io.Copy(hash, file); err != nil {
                return returnMD5String, err
        }
        hashInBytes := hash.Sum(nil)[:16]
        returnMD5String = hex.EncodeToString(hashInBytes)
        return returnMD5String, nil

}

func main() {
        hash, err := md5sum(os.Args[1])
        if err == nil {
                fmt.Println(hash)
        }
}

