package models

type Orange struct {
    Size float64 `json:"size"`
}

type Basket struct {
    Small  int
    Medium int
    Large  int
}