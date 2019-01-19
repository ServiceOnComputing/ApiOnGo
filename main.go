package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	//"reflect"
	"ApiOnGo/swapi/swapi"
	"github.com/boltdb/bolt"
	"strconv"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	// 解析url传递的参数
	r.ParseForm()
	for k, v := range r.Form {
		fmt.Println("key:", k)
		// join() 方法用于把数组中的所有元素放入一个字符串。
		// 元素是通过指定的分隔符进行分隔的
		fmt.Println("val:", strings.Join(v, ""))
	}

	// 输出到客户端
	name := r.Form["username"]
	pass := r.Form["password"]
	//fmt.Println("test ",reflect.TypeOf(name[0]))
	//fmt.Println("test")
	b, _ := strconv.Atoi(name[0])

	client := swapi.NewClient()
	result, _, _ := client.GetPersonById(b)

	for i, v := range name {
		fmt.Println(i)
		fmt.Fprintf(w, "用户名:%v\n", v)
	}
	for k, n := range pass {
		fmt.Println(k)
		fmt.Fprintf(w, "密码:%v\n", n)
	}

	fmt.Fprintf(w, "content:%s", result)
	//fmt.Fprintf(w,"content:",result)
}
func person(w http.ResponseWriter, r *http.Request) {
	// 解析url传递的参数
	r.ParseForm()

	name := r.Form["people"]

	b, _ := strconv.Atoi(name[0])

	client := swapi.NewClient()
	result, _, _ := client.GetPersonById(b)
	t, err := template.ParseFiles("./template/person.html")
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}

	if err := t.Execute(w, struct {
		Name      string
		BirthYear string
		EyeColor  string
		Gender    string
		HairColor string
		Height    string
		Homeworld string
		Mass      string
		SkinColor string
		Created   string
		Edited    string
		Url       string
	}{Name: result.Name,
		BirthYear: result.BirthYear,
		EyeColor:  result.EyeColor,
		Gender:    result.Gender,
		HairColor: result.HairColor,
		Height:    result.Height,
		Homeworld: result.Homeworld,
		Mass:      result.Mass,
		SkinColor: result.SkinColor,
		Created:   result.Created,
		Edited:    result.Edited,
		Url:       result.Url}); err != nil {
		fmt.Println("There was an error:", err.Error())
	}

	//fmt.Fprintf(w,"content:%s",result)
	//fmt.Fprintf(w,"<hr/>")
	//fmt.Fprintf(w,"content:",result)
}

func vehicle(w http.ResponseWriter, r *http.Request) {
	// 解析url传递的参数
	r.ParseForm()

	name := r.Form["vehicle"]

	b, _ := strconv.Atoi(name[0])

	client := swapi.NewClient()
	result, _, _ := client.GetVehicleById(b)
	t, err := template.ParseFiles("./template/vehicle.html")
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}

	if err := t.Execute(w, struct {
		Name                 string
		CargoCapacity        string
		Consumables          string
		CostInCredits        string
		Created              string
		Crew                 string
		Edited               string
		Length               string
		Manufacturer         string
		MaxAtmospheringSpeed string
		Model                string
		Passengers           string
		VehicleClass         string
		Url                  string
	}{Name: result.Name,
		CargoCapacity:        result.CargoCapacity,
		Consumables:          result.Consumables,
		CostInCredits:        result.CostInCredits,
		Created:              result.Created,
		Crew:                 result.Crew,
		Edited:               result.Edited,
		Length:               result.Length,
		Manufacturer:         result.Manufacturer,
		MaxAtmospheringSpeed: result.MaxAtmospheringSpeed,
		Model:                result.Model,
		Passengers:           result.Passengers,
		VehicleClass:         result.VehicleClass,
		Url:                  result.Url}); err != nil {
		fmt.Println("There was an error:", err.Error())
	}

	//fmt.Fprintf(w,"content:%s",result)
	//fmt.Fprintf(w,"<hr/>")
	//fmt.Fprintf(w,"content:",result)
}

func starship(w http.ResponseWriter, r *http.Request) {
	// 解析url传递的参数
	r.ParseForm()

	name := r.Form["starship"]

	b, _ := strconv.Atoi(name[0])

	client := swapi.NewClient()
	result, _, _ := client.GetStarshipById(b)
	t, err := template.ParseFiles("./template/starship.html")
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}

	if err := t.Execute(w, struct {
		Name                 string
		MGLT                 string
		CargoCapacity        string
		Consumables          string
		CostInCredits        string
		Created              string
		Crew                 string
		Edited               string
		HyperdriveRating     string
		Length               string
		Manufacturer         string
		MaxAtmospheringSpeed string
		Model                string
		Passengers           string
		StarshipClass        string
		Url                  string
	}{Name: result.Name,
		MGLT:                 result.MGLT,
		CargoCapacity:        result.CargoCapacity,
		Consumables:          result.Consumables,
		CostInCredits:        result.CostInCredits,
		Created:              result.Created,
		Crew:                 result.Crew,
		Edited:               result.Edited,
		HyperdriveRating:     result.HyperdriveRating,
		Length:               result.Length,
		Manufacturer:         result.Manufacturer,
		MaxAtmospheringSpeed: result.MaxAtmospheringSpeed,
		Model:                result.Model,
		Passengers:           result.Passengers,
		StarshipClass:        result.StarshipClass,
		Url:                  result.Url}); err != nil {
		fmt.Println("There was an error:", err.Error())
	}

	//fmt.Fprintf(w,"content:%s",result)
	//fmt.Fprintf(w,"<hr/>")
	//fmt.Fprintf(w,"content:",result)
}
func film(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	name := r.Form["film"]

	b, _ := strconv.Atoi(name[0])

	client := swapi.NewClient()
	result, _, _ := client.GetFilmById(b)
	t, err := template.ParseFiles("./template/film.html")
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}

	if err := t.Execute(w, struct {
		Created   string
		Director  string
		Edited    string
		EpisodeId int
		Producer  string
		Title     string
		Url       string
	}{Created: result.Created,
		Director:  result.Director,
		Edited:    result.Edited,
		EpisodeId: result.EpisodeId,
		Producer:  result.Producer,
		Title:     result.Title,
		Url:       result.Url}); err != nil {
		fmt.Println("There was an error:", err.Error())
	}
}
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

func species(w http.ResponseWriter, r *http.Request) {
	// 解析url传递的参数
	r.ParseForm()

	name := r.Form["species"]

	b, _ := strconv.Atoi(name[0])

	client := swapi.NewClient()
	result, _, _ := client.GetSpeciesById(b)
	t, err := template.ParseFiles("./template/species.html")
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}

	if err := t.Execute(w, struct {
		Name            string
		Homeworld       string
		Created         string
		Edited          string
		Url             string
		Classification  string
		Designation     string
		AverageHeight   string
		AverageLifespan string
	}{Name: result.Name,
		Homeworld:       result.Homeworld,
		Created:         result.Created,
		Edited:          result.Edited,
		Classification:  result.Classification,
		Designation:     result.Designation,
		AverageHeight:   result.AverageHeight,
		AverageLifespan: result.AverageLifespan,
		Url:             result.Url}); err != nil {
		fmt.Println("There was an error:", err.Error())
	}

	//fmt.Fprintf(w,"content:%s",result)
	//fmt.Fprintf(w,"<hr/>")
	//fmt.Fprintf(w,"content:",result)
}
func login(w http.ResponseWriter, r *http.Request) {

    fmt.Println("method:", r.Method)
    r.ParseForm()
    username :=r.Form["username"]
    password :=r.Form["password"]
    db, _ := bolt.Open("swapi.db", 0600, nil)
    defer db.Close()
    var v []byte
    db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte("users"))
        v = b.Get([]byte(string(username[0])))
        // fmt.Printf("The answer is: %s\n", v)
        return nil
    })
    if string(v)==password[0] {
        t, err := template.ParseFiles("./template/page.html")
	    if err != nil {
			fmt.Println("parse file err:", err)
			return
		}
	    if err := t.Execute(w, struct {}{}); err != nil {
			fmt.Println("There was an error:", err.Error())
		}
    } else {
        fmt.Fprintf(w,"登陆失败")
    }
}

func regist(w http.ResponseWriter, r *http.Request) {

    fmt.Println("method:", r.Method)
    r.ParseForm()
    username :=r.Form["username"]
    password :=r.Form["password"]
    db, _ := bolt.Open("swapi.db", 0600, nil)
    defer db.Close()
    db.Update(func(tx *bolt.Tx) error {
        bu, err := tx.CreateBucketIfNotExists([]byte("users"))
        if err != nil {
            fmt.Println("open bucket error")
            return err
        }
        bu.Put([]byte(string(username[0])), []byte(string(password[0])))
        return nil
    })

    t, err := template.ParseFiles("./template/log.html")
    if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
    if err := t.Execute(w, struct {Success string}{Success:"注册成功请登录："}); err != nil {
		fmt.Println("There was an error:", err.Error())
	}
    
    

}

func logg(w http.ResponseWriter, r *http.Request) {

    fmt.Println("method:", r.Method)
    t, err := template.ParseFiles("./template/log.html")
    if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
    if err := t.Execute(w, struct {Success string}{Success:""}); err != nil {
		fmt.Println("There was an error:", err.Error())
	}

}

func reg(w http.ResponseWriter, r *http.Request) {

    fmt.Println("method:", r.Method)
    t, err := template.ParseFiles("./template/reg.html")
    if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
    if err := t.Execute(w, struct {}{}); err != nil {
		fmt.Println("There was an error:", err.Error())
	}


}
func test1(w http.ResponseWriter, r *http.Request) {

	client := swapi.NewClient()
	result := client.GetpersonById(2)

    fmt.Fprintf(w, result)


}

func test2(w http.ResponseWriter, r *http.Request) {

	client := swapi.NewClient()
	result := client.GetfilmById(2)

    fmt.Fprintf(w, result)


}

func apiroot(w http.ResponseWriter, r *http.Request) {

	

    fmt.Fprintf(w, "localhost:9090/people/id\n"+
    				"localhost:9090/film/id\n"+
    				"localhost:9090/vehicles/id\n"+
    				"localhost:9090/starships/id\n"+
    				"localhost:9090/species/id\n"+
    				"localhost:9090/planets/id\n")


}

func main() {
	// http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
	//     io.WriteString(w, "这里写上你的json数据就行了")
	// })
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
    http.HandleFunc("/regist",regist)
	http.HandleFunc("/people", person)
	http.HandleFunc("/film", film)
	http.HandleFunc("/planet", planet)
	http.HandleFunc("/log",logg)
	http.HandleFunc("/reg",reg)
	http.HandleFunc("/species", species)
	http.HandleFunc("/vehicle", vehicle)
	http.HandleFunc("/starship", starship)
	http.HandleFunc("/people/2",test1)
	http.HandleFunc("/film/2",test2)
	http.HandleFunc("/apiroot",apiroot)
	fmt.Println("listening port: 9090")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndserve:", err)
	}
}

// package main

// import (
//     "log"

//     "github.com/adampresley/swapi-go/swapi"
// )

// func main() {
//     client := swapi.NewClient()
//     result, status, err := client.GetPersonById(1)

//     if err != nil {
//         log.Fatalf("Cannot get person: %s", err.Error())
//     }

//     if status != 200 {
//         log.Println("HTTP error with message", result.ErrorMessage)
//     }

//     log.Println(result)
// }
