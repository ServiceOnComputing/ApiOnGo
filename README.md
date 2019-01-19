# 简单 web 服务与客户端开发

实验人：周伟标 韩承希 陈咏强

## 实验要求
**任务目标***

1. 选择合适的 API 风格，实现从接口或资源（领域）建模，到 API 设计的过程
2. 使用 API 工具，编制 API 描述文件，编译生成服务器、客户端原型
3. 使用 Github 建立一个组织，通过 API 文档，实现 客户端项目 与 RESTful 服务项目同步开发
4. 使用 API 设计工具提供 Mock 服务，两个团队独立测试 API
5. 使用 travis 测试相关模块

## 实验过程

### API的设计
首先是各种数据类型的设计，每个数据类型实现在一个文件中，存储为json格式，有利于数据的调用。

前端访问localhost:port/apiroot，服务器的实现如下：
```
func apiroot(w http.ResponseWriter, r *http.Request) {
 
    fmt.Fprintf(w, "localhost:9090/people/id\n"+
    				"localhost:9090/film/id\n"+
    				"localhost:9090/vehicles/id\n"+
    				"localhost:9090/starships/id\n"+
    				"localhost:9090/species/id\n"+
    				"localhost:9090/planets/id\n")
    				}
```
上面是响应函数，w是http.ResponseWriter，我们在w中写入字符串，就完成了响应过程。绑定响应函数：
```
    http.HandleFunc("/apiroot",apiroot)
```
接口函数：
```
result := &film.Film{}

status, body, err := this.get(fmt.Sprintf("/films/%d", id))
if err != nil {
    return result, 0, err
}

err = json.Unmarshal(body, result)
return result, status, err
```
运行的结果如下：

```
    $cur1-i1 localhost:9090/ aplroot
    HTTP/1.1 200 OK
    Date: Fri. t14 Dec 2018 11: 33: 52 GMT
    Content-length: 155
    Content-type: text/plain; charset=utf-8
    localhost: 9090/people/id
    localhost: 9090/film/id
    localhost: 9090/vehicles/id 
    localhost: 9090/starships/id
    localhost: 9090/species/id
    localhost: 9090/planets/id
```

### main函数
在主函数中，主要是通过调用接口函数，编写能将数据传输到HTML中的函数，我主要负责编写结构体planet和结构体film对应的函数。
```
func planet(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	name := r.Form["planet"]

	b, _ := strconv.Atoi(name[0])

	client := swapi.NewClient()
	result, _, _ := client.GetPlanetById(b)
	t, err := template.ParseFiles("./template/planet.html")
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}

	if err := t.Execute(w, struct {
		Climate        string
		Created        string
		Diameter       string
		Edited         string
		Gravity        string
		Name           string
		OrbitalPeriod  string
		Population     string
		RotationPeriod string
		SurfaceWater   string
		Terrain        string
		Url            string
	}{Climate: result.Climate,
		Created:        result.Created,
		Diameter:       result.Diameter,
		Edited:         result.Edited,
		Gravity:        result.Gravity,
		Name:           result.Name,
		OrbitalPeriod:  result.OrbitalPeriod,
		Population:     result.Population,
		RotationPeriod: result.RotationPeriod,
		SurfaceWater:   result.SurfaceWater,
		Terrain:        result.Terrain,
		Url:            result.Url}); err != nil {
		fmt.Println("There was an error:", err.Error())
	}
}
```
### json数据读取
我们完成的项目是swapi，提供了查询星球大战各种信息的api，源数据在swapi官方github找到[http://github.com/phalt/swapi/tree/master/resources/fixtures](http://github.com/phalt/swapi/tree/master/resources/fixtures "源数据")

我们做的第一个工作是先进行分割，讲里面的json数据根据不同的种类分割开来保存。然后根据前端请求的种类和id来返回对用的json。我们为了方便起见，讲读取后的json封装成了一个struct,struct的内容对应json的内容。然后Go语言提供了函数可以直接进行转换。

文件读取json的过程如下：
```
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
            return
        }
     
        //读取的数据为json格式，需要进行解码
        err = json.Unmarshal(data, v)
        if err != nil {
            return
        }
```
到这里工作基本完成。

### boltDB数据库操作
用blotdb实现保存注册的用户的数据：
```
//插入数据 
db, _ := bolt.Open("swapi.db", 0600, nil)//开启数据库，如果没有的话则新建
defer db.Close()
db.Update(func(tx *bolt.Tx) error {
    bu, err := tx.CreateBucketIfNotExists([]byte("users"))//新建数据表，users是表名    
	if err != nil {
        fmt.Println("open bucket error")
        return err    
	}    
	bu.Put([]byte(string(username[0])), []byte(string(password[0])))//在数据表中插入键值对    
	return nil}) 
//查询数据
db, _ := bolt.Open("swapi.db", 0600, nil)
defer db.Close()
var v []byte
db.View(func(tx *bolt.Tx) error {
    b := tx.Bucket([]byte("users"))//打开数据表
    v = b.Get([]byte(string(username[0])))//通过键获取对应的值
    return nil
})

```
