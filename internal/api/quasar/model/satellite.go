package model

type Satellite struct{
	id int
	name string
	position Position
}

func(s *Satellite)Satelite(id int, name string){
	s.id = id
	s.name = name
}

func(s *Satellite)SetPosition(x float32, y float32){
	s.position.Position(x,y)
}

func(s *Satellite)GetPosiion()(float32, float32){
	x:= s.position.GetX()
	y:= s.position.GetY()
	return x, y
}