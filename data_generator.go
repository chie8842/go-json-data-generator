package json_data_generator

import (
  "fmt"
  "reflect"
  "regexp"
  "strconv"
  "strings"
)

func CreateDummy(schema interface{}) interface{}{
  if reflect.TypeOf(schema).String() == "[]interface {}" {
      records := handleRepeatForList(schema.([]interface {}))
      return records
  } else {
  //else if reflect.TypeOf(schema).String() == "map[string]interface {}" {
      records := handleRepeat(schema.(map[string]interface {}), 1)
    return records
  }
}

type funcParams struct {
  strKey string
  fakeFunc string
  funcParam []interface{}
  resultType string
  value interface{}
}

func handleRepeat(schema map[string]interface{}, numRecord int) interface{} {
  fmt.Printf("handleRepeat %v\n", schema)
  keys := reflect.ValueOf(schema).MapKeys()
  strkeys := make([]string, len(keys))
  for i := 0; i < len(keys); i++ {
    strkeys[i] = keys[i].String()
  }

  funcParamsList := make([]funcParams, len(keys))
  fakeFuncRegex := "^\\{\\{(.*)\\(.*\\)\\}\\}$"
  fakeFuncRegexRep := regexp.MustCompile(fakeFuncRegex)
  fakeFuncParamRegex := "^\\{\\{.*\\((.*)\\)\\}\\}$"
  fakeFuncParamRegexRep := regexp.MustCompile(fakeFuncParamRegex)
  for i := 0; i < len(strkeys); i++{
    funcString, ok := schema[strkeys[i]].(string)
    var tempFuncParams funcParams
    tempFuncParams.strKey = strkeys[i]
    if ok {
      fakeFunc := fakeFuncRegexRep.FindStringSubmatch(funcString)
      if len(fakeFunc) == 2 {
        //funcString is func
        fakeParams := fakeFuncParamRegexRep.FindStringSubmatch(funcString)
        tempFuncParams.fakeFunc = fakeFunc[1]
        tempFuncParam := strings.Split(fakeParams[1], ",")
        tempFuncParams.funcParam = make([]interface{}, len(tempFuncParam))
        for i, param := range tempFuncParam {
          tempFuncParams.funcParam[i] = param
        }
      } else {
        //funcString is just a string
        tempFuncParams.value = schema[strkeys[i]]
      }
      funcParamsList[i] = tempFuncParams
      //var x []interface{}
    } else {
      tempFuncParams.value = schema[strkeys[i]]
      funcParamsList[i] = tempFuncParams
    }
  }
  records := createRecords(funcParamsList, numRecord)
  return records
}

func handleRepeatForList(schema []interface{}) interface{} {
  repeated_string := schema[0].(string)
  repeat_regex := "^repeat\\((\\d+)\\)$"
  repeat_rep := regexp.MustCompile(repeat_regex)
  repeatTimes, err := strconv.Atoi(repeat_rep.FindStringSubmatch(repeated_string)[1])
  if err != nil {
    fmt.Printf("%v is not valid repeat statement.\n", repeated_string)
  }

  records := handleRepeat(schema[1].(map[string]interface{}), repeatTimes)
  return records
}
