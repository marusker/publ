package publ

import (
  "io/ioutil"
  "path/filepath"
  "net/http"
  "strings"
  "github.com/gorilla/mux"
)

//FileURLs has file pathes of files. So that it can be used for the other muxes.
var FileURLs []string
//publicFolder has public folder path
var publicFolder string


func SetFolder(publicFolder string){
  if publicFolder[(len(publicFolder)-1):] == "/" {
    panic("publ: The directory given to publ.SetFoler ends with '/'. Please remove '/'.")
  }
  publicFolder = publicFolder
  //Assign file directories to FileURLs
	FileURLs = searchDir(publicFolder)

  //publicFolder starts with "./" needs special treatment, because "./" disapplear from directory
  if publicFolder[:2] == "./" {
    for i,file := range FileURLs {
      //replace file directories to root path url
      FileURLs[i] = strings.Replace(file,publicFolder[2:],"",1)
    }
  }else{
    for i,file := range FileURLs {
      //replace file directories to root path url
      FileURLs[i] = strings.Replace(file,publicFolder,"",1)
    }
  }
}

func searchDir(dir string) []string {

  files, err := ioutil.ReadDir(dir)
  if err != nil {
      panic(err)
  }

  var paths []string
  for _, file := range files {
    if file.IsDir() {
      paths = append(paths, searchDir(filepath.Join(dir, file.Name()))...)
      continue
    }
    //exclude mac's .DS_Store file
    if string(file.Name()) != ".DS_Store"{
      paths = append(paths, filepath.Join(dir, file.Name()))
    }
  }
  return paths
}

//for net/http
func Activate(){
    for _,fileName := range FileURLs {
      serveSingle(fileName, publicFolder + fileName)
   }
}


func serveSingle(fileName string, fileDir string) {
  http.HandleFunc(fileName, func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, fileDir)
  })
}

//for gorilla mux
func ActivateGorilla(router *mux.Router){
    for _,fileName := range FileURLs {
      serveSingleGorilla(router, fileName, publicFolder + fileName)
   }
}


func serveSingleGorilla(router *mux.Router, fileName string, fileDir string) {
  router.HandleFunc(fileName, func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, fileDir)
  })
}


