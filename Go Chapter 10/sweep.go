package main

func Sweep(){

   var objects *[]object

   objects = GetObjects()

   var object *object

   for _, object = range objects {

   var markedAlready bool

   markedAlready = IfMarked(object)
   if markedAlready {
        map[object] = true
   }
       Release(object)
   }
}