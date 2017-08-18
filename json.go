package main

import (
  "net/http"
  "encoding/json"
)


type Curso struct{
  Title string `json: "titulo"`
  NumeroVideos int `json: "numero_videos" `
  CantidadEstudiantes int `json: "cantidad_estudiantes" `
}

type Cursos []Curso

func main(){
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
    cursos := Cursos{
      Curso{"Curso de Laravel", 25, 24},
      Curso{"Curso de Go", 30, 28},
      Curso{"Curso de Nodejs", 35, 18},
    }
    json.NewEncoder(w).Encode(cursos)
  })

  http.ListenAndServe(":8000", nil)
}
