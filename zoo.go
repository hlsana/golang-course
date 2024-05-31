package main

type Zookeeper struct {
	FirstName  string
	LastName   string
	Occupation string
}

type Animal struct {
	Species  string
	Name     string
	IsInCage bool
}

func main() {

	Lion := Animal{
		Species:  "Lion",
		Name:     "Alex",
		IsInCage: false,
	}

	Zebra := Animal{
		Species:  "Zebra",
		Name:     "Marty",
		IsInCage: false,
	}
	Hippo := Animal{
		Species:  "Hippopotamus",
		Name:     "Gloria",
		IsInCage: false,
	}
	Penguin := Animal{
		Species:  "Penguin",
		Name:     "Skipper",
		IsInCage: false,
	}
	Giraffe := Animal{
		Species:  "Giraffe",
		Name:     "Melman",
		IsInCage: false,
	}

	Zookeeper := Zookeeper{
		FirstName:  "Henry",
		LastName:   "Zoolover",
		Occupation: "Zookeeper",
	}

	if !Lion.IsInCage {
		Lion.IsInCage = true
	}
	if !Zebra.IsInCage {
		Zebra.IsInCage = true
	}
	if !Hippo.IsInCage {
		Hippo.IsInCage = true
	}
	if !Penguin.IsInCage {
		Penguin.IsInCage = true
	}
	if !Giraffe.IsInCage {
		Giraffe.IsInCage = true
	}

	println(Zookeeper.Occupation, Zookeeper.FirstName, Zookeeper.LastName, "starts catching runaway animals.")
	if Lion.IsInCage {
		println(Lion.Species, Lion.Name, "has been caught.")
	} else {
		println(Lion.Species, Lion.Name, "hasn't been caught.")
	}

	if Zebra.IsInCage {
		println(Zebra.Species, Zebra.Name, "has been caught.")
	} else {
		println(Zebra.Species, Zebra.Name, "hasn't been caught.")
	}

	if Hippo.IsInCage {
		println(Hippo.Species, Hippo.Name, "has been caught.")
	} else {
		println(Hippo.Species, Hippo.Name, "hasn't been caught,")
	}

	if Penguin.IsInCage {
		println(Penguin.Species, Penguin.Name, "has been caught.")
	} else {
		println(Penguin.Species, Penguin.Name, "hasn't been caught.")
	}

	if Giraffe.IsInCage {
		println(Giraffe.Species, Giraffe.Name, "has been caught.")
	} else {
		println(Giraffe.Species, Giraffe.Name, "hasn't been caught.")
	}

}
