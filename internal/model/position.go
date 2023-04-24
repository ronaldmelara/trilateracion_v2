package model

type Position struct{
	x float32
	y float32
}

func(p *Position)Position(x float32, y float32){
	p.x = x;
	p.y = y;
}

func(p *Position)GetX()(float32){
	return p.x
}

func(p *Position)GetY()(float32){
	return p.y
}

func(p *Position)SetX(x float32){
	p.x = x
}
func(p *Position)SetY(y float32){
	p.y = y
}