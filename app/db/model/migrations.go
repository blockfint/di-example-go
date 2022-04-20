package model

func NewMigrations() []interface{} {
	var migrations []interface{}

	migrations = append(migrations, &Todo{})
	migrations = append(migrations, &IDAddress{})
	migrations = append(migrations, &ContactAddress{})
	migrations = append(migrations, &IDCardAddress{})
	migrations = append(migrations, &IDCard{})
	migrations = append(migrations, &Customer{})

	return migrations
}
