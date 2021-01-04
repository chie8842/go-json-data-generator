package json_data_generator

import (
  "errors"
  "fmt"
  "reflect"
  "strconv"
  "github.com/brianvoe/gofakeit"
)

type  stubMapping map[string]interface{}
var StubStorage = stubMapping{}

func callFakeFunc(funcName string, params []interface{}) (result interface{}, resultType string, err error) {
  StubStorage = map[string]interface{} {
    "AchAccount": gofakeit.AchAccount,
    "AchRouting": gofakeit.AchRouting,
    "Adjective": gofakeit.Adjective,
    "Adverb": gofakeit.Adverb,
    "Animal": gofakeit.Animal,
    "AnimalType": gofakeit.AnimalType,
    "AppAuthor": gofakeit.AppAuthor,
    "AppName": gofakeit.AppName,
    "AppVersion": gofakeit.AppVersion,
    "BS": gofakeit.BS,
    "BeerAlcohol": gofakeit.BeerAlcohol,
    "BeerBlg": gofakeit.BeerBlg,
    "BeerHop": gofakeit.BeerHop,
    "BeerIbu": gofakeit.BeerIbu,
    "BeerMalt": gofakeit.BeerMalt,
    "BeerName": gofakeit.BeerName,
    "BeerStyle": gofakeit.BeerStyle,
    "BeerYeast": gofakeit.BeerYeast,
    "BitcoinAddress": gofakeit.BitcoinAddress,
    "BitcoinPrivateKey": gofakeit.BitcoinPrivateKey,
    "Bool": gofakeit.Bool,
    "Breakfast": gofakeit.Breakfast,
    "BuzzWord": gofakeit.BuzzWord,
    "CarFuelType": gofakeit.CarFuelType,
    "CarMaker": gofakeit.CarMaker,
    "CarModel": gofakeit.CarModel,
    "CarTransmissionType": gofakeit.CarTransmissionType,
    "CarType": gofakeit.CarType,
    "Cat": gofakeit.Cat,
    "ChromeUserAgent": gofakeit.ChromeUserAgent,
    "City": gofakeit.City,
    "Color": gofakeit.Color,
    "Company": gofakeit.Company,
    "CompanySuffix": gofakeit.CompanySuffix,
    "Country": gofakeit.Country,
    "CountryAbr": gofakeit.CountryAbr,
    "CreditCardCvv": gofakeit.CreditCardCvv,
    "CreditCardExp": gofakeit.CreditCardExp,
    "CreditCardNumber": gofakeit.CreditCardNumber,
    "CreditCardType": gofakeit.CreditCardType,
    "CurrencyLong": gofakeit.CurrencyLong,
    "CurrencyShort": gofakeit.CurrencyShort,
    "Date": gofakeit.Date,
    "DateRange": gofakeit.DateRange,
    "Day": gofakeit.Day,
    "Dessert": gofakeit.Dessert,
    "Digit": gofakeit.Digit,
    "DigitN": gofakeit.DigitN,
    "Dinner": gofakeit.Dinner,
    "Dog": gofakeit.Dog,
    "DomainName": gofakeit.DomainName,
    "DomainSuffix": gofakeit.DomainSuffix,
    "Email": gofakeit.Email,
    "Emoji": gofakeit.Emoji,
    "EmojiAlias": gofakeit.EmojiAlias,
    "EmojiCategory": gofakeit.EmojiCategory,
    "EmojiDescription": gofakeit.EmojiDescription,
    "EmojiTag": gofakeit.EmojiTag,
    "FarmAnimal": gofakeit.FarmAnimal,
    "FileExtension": gofakeit.FileExtension,
    "FileMimeType": gofakeit.FileMimeType,
    "FirefoxUserAgent": gofakeit.FirefoxUserAgent,
    "FirstName": gofakeit.FirstName,
    "Float32": gofakeit.Float32,
    "Float32Range": gofakeit.Float32Range,
    "Float64": gofakeit.Float64,
    "Float64Range": gofakeit.Float64Range,
    "Fruit": gofakeit.Fruit,
    "Gamertag": gofakeit.Gamertag,
    "Gender": gofakeit.Gender,
    "Generate": gofakeit.Generate,
    "HTTPMethod": gofakeit.HTTPMethod,
    "HTTPStatusCode": gofakeit.HTTPStatusCode,
    "HTTPStatusCodeSimple": gofakeit.HTTPStatusCodeSimple,
    "HackerAbbreviation": gofakeit.HackerAbbreviation,
    "HackerAdjective": gofakeit.HackerAdjective,
    "HackerNoun": gofakeit.HackerNoun,
    "HackerPhrase": gofakeit.HackerPhrase,
    "HackerVerb": gofakeit.HackerVerb,
    "HackeringVerb": gofakeit.HackeringVerb,
    "HexColor": gofakeit.HexColor,
    "HipsterParagraph": gofakeit.HipsterParagraph,
    "HipsterSentence": gofakeit.HipsterSentence,
    "HipsterWord": gofakeit.HipsterWord,
    "Hour": gofakeit.Hour,
    "IPv4Address": gofakeit.IPv4Address,
    "IPv6Address": gofakeit.IPv6Address,
    "ImageJpeg": gofakeit.ImageJpeg,
    "ImagePng": gofakeit.ImagePng,
    "ImageURL": gofakeit.ImageURL,
    "Int16": gofakeit.Int16,
    "Int32": gofakeit.Int32,
    "Int64": gofakeit.Int64,
    "Int8": gofakeit.Int8,
    "JobDescriptor": gofakeit.JobDescriptor,
    "JobLevel": gofakeit.JobLevel,
    "JobTitle": gofakeit.JobTitle,
    "Language": gofakeit.Language,
    "LanguageAbbreviation": gofakeit.LanguageAbbreviation,
    "LastName": gofakeit.LastName,
    "Latitude": gofakeit.Latitude,
    "Letter": gofakeit.Letter,
    "LetterN": gofakeit.LetterN,
    "Lexify": gofakeit.Lexify,
    "LogLevel": gofakeit.LogLevel,
    "Longitude": gofakeit.Longitude,
    "LoremIpsumParagraph": gofakeit.LoremIpsumParagraph,
    "LoremIpsumSentence": gofakeit.LoremIpsumSentence,
    "LoremIpsumWord": gofakeit.LoremIpsumWord,
    "Lunch": gofakeit.Lunch,
    "MacAddress": gofakeit.MacAddress,
    "Minute": gofakeit.Minute,
    "Month": gofakeit.Month,
    "Name": gofakeit.Name,
    "NamePrefix": gofakeit.NamePrefix,
    "NameSuffix": gofakeit.NameSuffix,
    "NanoSecond": gofakeit.NanoSecond,
    "Noun": gofakeit.Noun,
    "Number": gofakeit.Number,
    "Numerify": gofakeit.Numerify,
    "OperaUserAgent": gofakeit.OperaUserAgent,
    "Paragraph": gofakeit.Paragraph,
    "Password": gofakeit.Password,
    "PetName": gofakeit.PetName,
    "Phone": gofakeit.Phone,
    "PhoneFormatted": gofakeit.PhoneFormatted,
    "Phrase": gofakeit.Phrase,
    "Preposition": gofakeit.Preposition,
    "Price": gofakeit.Price,
    "ProgrammingLanguage": gofakeit.ProgrammingLanguage,
    "ProgrammingLanguageBest": gofakeit.ProgrammingLanguageBest,
    "Question": gofakeit.Question,
    "Quote": gofakeit.Quote,
    "RGBColor": gofakeit.RGBColor,
    "RandomInt": gofakeit.RandomInt,
    "RandomString": gofakeit.RandomString,
    "RandomUint": gofakeit.RandomUint,
    "Regex": gofakeit.Regex,
    "RemoveFuncLookup": gofakeit.RemoveFuncLookup,
    "SSN": gofakeit.SSN,
    "SafariUserAgent": gofakeit.SafariUserAgent,
    "SafeColor": gofakeit.SafeColor,
    "Second": gofakeit.Second,
    "Sentence": gofakeit.Sentence,
    "Snack": gofakeit.Snack,
    "State": gofakeit.State,
    "StateAbr": gofakeit.StateAbr,
    "Street": gofakeit.Street,
    "StreetName": gofakeit.StreetName,
    "StreetNumber": gofakeit.StreetNumber,
    "StreetPrefix": gofakeit.StreetPrefix,
    "StreetSuffix": gofakeit.StreetSuffix,
    "TimeZone": gofakeit.TimeZone,
    "TimeZoneAbv": gofakeit.TimeZoneAbv,
    "TimeZoneFull": gofakeit.TimeZoneFull,
    "TimeZoneOffset": gofakeit.TimeZoneOffset,
    "TimeZoneRegion": gofakeit.TimeZoneRegion,
    "URL": gofakeit.URL,
    "UUID": gofakeit.UUID,
    "Uint16": gofakeit.Uint16,
    "Uint32": gofakeit.Uint32,
    "Uint64": gofakeit.Uint64,
    "Uint8": gofakeit.Uint8,
    "UserAgent": gofakeit.UserAgent,
    "Username": gofakeit.Username,
    "Vegetable": gofakeit.Vegetable,
    "Verb": gofakeit.Verb,
    "WeekDay": gofakeit.WeekDay,
    "Word": gofakeit.Word,
    "Year": gofakeit.Year,
    "Zip": gofakeit.Zip,
  }

  f := reflect.ValueOf(StubStorage[funcName])

  var in []reflect.Value
  if params != nil {
    //引数がない場合、ここのparamsが []interface{}になってparams[0]=""になるのでnilに置き換える
    if params[0] == "" {
      params = nil
    }
    if len(params) > 0 {
      if len(params) != f.Type().NumIn(){
        err = errors.New(fmt.Sprintf("The number of Params(%v) is out of index(%v).", len(params), f.Type().NumIn()))
        return
      }
      in = make([]reflect.Value, len(params))
      for k, param := range params {
        if f.Type().In(k).String() == "string" {
          in[k] = reflect.ValueOf(param)
        } else if f.Type().In(k).String() == "int" {
          s, _ := strconv.Atoi(param.(string))
          in[k] = reflect.ValueOf(s)
        } else if f.Type().In(k).String() == "float64" {
          s, _ := strconv.ParseFloat(param.(string), 64)
          in[k] = reflect.ValueOf(s)
        } else if f.Type().In(k).String() == "uint" {
          s, _ := strconv.ParseUint(param.(string), 10, 64)
          in[k] = reflect.ValueOf(s)
        } else if f.Type().In(k).String() == "bool" {
          s, _ := strconv.ParseBool(param.(string))
          in[k] = reflect.ValueOf(s)
        }
      }
    //} else {
    //  fmt.Println("args 0")
    //  //fmt.Printf("args 0: %v\n", len(params))
    }
  }
  var res []reflect.Value
  res = f.Call(in)
  result = res[0].Interface()
  resultType = f.Type().Out(0).Kind().String()
  return
}
