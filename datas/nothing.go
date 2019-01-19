package datas

import (
	//"encoding/json"
	"fmt"
    // "os"
    // "io"
    // "bufio"
    // "strings"
    "strconv"
    "path"
    "runtime"
    "encoding/json"
    "io/ioutil"
	//"bytes"
	//"log"
	// "github.com/adampresley/swapi-go/swapi/film"
	 "github.com/adampresley/swapi-go/swapi/people"
	// "github.com/adampresley/swapi-go/swapi/planet"
	// "github.com/adampresley/swapi-go/swapi/root"
	// "github.com/adampresley/swapi-go/swapi/species"
	// "github.com/adampresley/swapi-go/swapi/starship"
	// "github.com/adampresley/swapi-go/swapi/vehicle"
	
)

func Getjson(kind string,id int) (*people.People) {

	result := &people.People{}
	fileName := "./datas/"+kind+"/"+kind+strconv.Itoa(id)+".json"
	_, filen, _, _ := runtime.Caller(1)
	datapath := path.Join(path.Dir(filen), fileName)
	JsonParse := NewJsonStruct()
	JsonParse.Load(datapath, &result)
	fmt.Println(result)
	return result



       
}

type JsonStruct struct {
}

func NewJsonStruct() *JsonStruct {
    return &JsonStruct{}
}

func (jst *JsonStruct) Load(filename string, v interface{}) {
    //ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
    data, err := ioutil.ReadFile(filename)
    if err != nil {
    	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
        return
    }

    //读取的数据为json格式，需要进行解码
    err = json.Unmarshal(data, v)
    if err != nil {
    	fmt.Println("??????????????????????????????????????????????????????")
        return
    }
}
