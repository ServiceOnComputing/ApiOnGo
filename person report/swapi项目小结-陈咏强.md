
初始设计

由于有多个数据种类，我把每一个数据种类放在不同的文件夹中，每一个文件夹中存放一个包含数据种类的结构体，便于获取的json数据以结构体的方式存放。

接口函数

首先定义一个get函数，从"http://swapi.co/api" 获取资源，该函数通过调用http.Get函数体请求远程网页，获取资源，然后通过

body, err := ioutil.ReadAll(response.Body)

ioutil.ReadAll函数解析出相应的body中的内容，获取到的body是byte数组类型。

在该接口函数中，首先定义一个Film结构体类型的变量，将输入的id进行格式化后作为参数传入get函数中，通过http.Get函数获取http://swapi.co/api/films/id的数据，以byte类型返回到body中，然后使用json的反序列化函数json.Unmarshal(body, result)将body中的数据传入result变量中，保存为结构体的形式，并返回。

main函数调用

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

首先是获取表单提交的请求数据，将是数据转为整数类型，然后调用对应的接口函数返回对应的结构体类型的变量。这个变量中保存着结构体为Film的数据，通过template.ParseFiles函数获取结构体Film对应的HTML展示模板，用t.Execute函数将结构体的数据写入到模板对应的字段中。
```
http.HandleFunc("/film", film)
```

调用http.HandleFunc("/film", film) 来为http的server端做相应的处理，这里的话就是调用之前写的film函数，跳转到对应的film数据显示的HTML界面。

至此，我的工作基本完成。