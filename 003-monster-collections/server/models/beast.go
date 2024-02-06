package models

type Beast struct {
    Class       string `bson:"class"`
    Name        string `bson:"name"`
    AttackPower int    `bson:"attack_power"`
    HealthPoints float64   `bson:"health_points"`
    Speed       int    `bson:"speed"`
}
