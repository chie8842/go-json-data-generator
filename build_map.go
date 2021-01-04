package json_data_generator

func createRecords(funcParamsList []funcParams, numRecord int) interface{} {
  //var err error
  if numRecord == 1 {
    //var records map[string]interface{}
    records := make(map[string]interface{})
    for i:=0; i<len(funcParamsList); i++ {
      if funcParamsList[i].fakeFunc != "" {
        funcParamsList[i].value, funcParamsList[i].resultType, _ = callFakeFunc(
          funcParamsList[i].fakeFunc, funcParamsList[i].funcParam)
      }
      records[funcParamsList[i].strKey] = funcParamsList[i].value
    }
    return records
  } else {
    var records []map[string]interface{}
    //records := make(map[string]interface{}, numRecord)
    for i:= 0; i<numRecord; i++{
      record := make(map[string]interface{})
      //var record map[string]interface{}
      for j:= 0; j<len(funcParamsList); j++{
        if funcParamsList[j].fakeFunc != "" {
          funcParamsList[j].value, funcParamsList[j].resultType, _ = callFakeFunc(
            funcParamsList[j].fakeFunc, funcParamsList[j].funcParam)
        }
        record[funcParamsList[j].strKey] = funcParamsList[j].value
      }
      records = append(records, record)
    }
    return records
  }
}
