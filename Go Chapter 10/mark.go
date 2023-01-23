package main

func Mark( root *object){

   var markedAlready bool

   markedAlready = IfMarked(root)
   if !markedAlready {
        map[root] = true
   }

   var references *[]object

   references = GetReferences(root)

   var reference *object

   for _, reference = range references {
       Mark(reference)
   }
}
