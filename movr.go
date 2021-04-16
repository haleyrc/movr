package movr

var DefaultPartitionMap = map[string][]string{
	"us_east":    {"new york", "boston", "washington dc"},
	"us_west":    {"san francisco", "seattle", "log angeles"},
	"us_central": {"chicago", "detroit", "minneapolis"},
	"eu_west":    {"amsterdam", "paris", "rome"},
}
