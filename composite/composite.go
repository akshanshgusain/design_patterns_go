package composite

//Composite is a structural design pattern that allows composing objects into a tree-like structure and work with the
//it as if it was a singular object.

//Composite became a pretty popular solution for the most problems that require building a tree structure.
//Composite’s great feature is the ability to run methods recursively over the whole tree structure and sum up
//the results.

/*
Let’s try to understand the Composite pattern with an example of an operating system’s file system. In the file system,
there are two types of objects: files and folders. There are cases when files and folders should be treated to be the
same way. This is where the Composite pattern comes in handy.

Imagine that you need to run a search for a particular keyword in your file system. This search operation applies to
both files and folders. For a file, it will just look into the contents of the file; for a folder, it will go through
all files of that folder to find that keyword.*/

func Composite() {
	file01 := &file{name: "file01"}
	file02 := &file{name: "file02"}
	file03 := &file{name: "file03"}

	folder01 := &folder{name: "folder01"}
	folder02 := &folder{name: "folder02"}

	folder01.add(file01)
	folder01.add(file02)

	folder02.add(file03)
	folder02.add(folder01)

	folder02.search("rose")
}
