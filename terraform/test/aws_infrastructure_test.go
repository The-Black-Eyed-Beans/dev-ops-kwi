package aws_infrastructure

type MainJson struct {
	AppLB  AppJson  `json:"alb"`
	GateLB GateJson `json:"glb"`
	R53    R53Json  `json:"route53"`
}

type AppJson struct {
	Name string `json:"name"`
}

type GateJson struct {
	Name string `json:"name"`
}

type R53Json struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
