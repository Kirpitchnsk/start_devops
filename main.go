inport(
  "fmt"
  "net/http"
)

func main()
{
  http.HandleFunc("/",func(w http.ResponseWriter, r*http.Request){
    fmt.Fprintf(w,"Hello World!")
   })
  http.HandleFunc("/hello",func(w http.ResponseWriter, r*http.Request){
    fmt.Fprintf(w, "Hello,user")
  })
  http.LitsenAndServe(":8098",nil)
}
