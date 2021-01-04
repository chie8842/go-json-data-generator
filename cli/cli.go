package main
import (
  "encoding/json"
  "fmt"
  "flag"
  "os"
  "reflect"
  "io/ioutil"
  //"reflect"
  "github.com/chie8842/go-json-data-generator"
)

func main() {
    var (
      sourceJsonFile = flag.String("s", "", "source json file")
      destinationJsonFile = flag.String("d", "", "destination json file")
    )
    flag.Parse()
    parsed := parseSourceJson(*sourceJsonFile)
    records := json_data_generator.CreateDummy(parsed)
    dumpJson(records, *destinationJsonFile)
}

func parseSourceJson(sourceJsonFile string) interface{}{
    sourceJson, err := os.Open(sourceJsonFile)
    if err != nil {
      fmt.Println(err)
    }
    fmt.Printf("Successfully Opend %v, %v\n", sourceJsonFile)
    defer sourceJson.Close()

    byteValue, _ := ioutil.ReadAll(sourceJson)

    var result interface{}
    err = json.Unmarshal([]byte(byteValue), &result)
    if err != nil {
        fmt.Println(err)
    }
    return result
}

func dumpJson(records interface{}, destinationJsonFile string) {
  f, _ := os.OpenFile(destinationJsonFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
  defer f.Close()
  if reflect.TypeOf(records).String() == "[]map[string]interface {}" {
    for i:=0; i<len(records.([]map[string]interface {}));i++ {
      bytes, _ := json.Marshal(records.([]map[string]interface {})[i])
      f.Write(bytes)
      f.WriteString("\n")
    }
  } else {
    bytes, _ := json.Marshal(records)
    ioutil.WriteFile(destinationJsonFile, bytes, os.ModePerm)
  }
}
