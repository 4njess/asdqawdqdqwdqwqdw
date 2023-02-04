package main

type Cat struct {
	tableName struct{} `pg:"cats"`
	ID        string   `json:"id" pg:"id"`
	Name      string   `json:"name" pg:"name"`
	IsStrip   bool     `json:"is_strip" pg:"is_strip"`
	Color     string   `json:"color" pg:"color"`
}

// получить список котов
// Переменная cats содержит массив структуры Cat
func findAllCats() []Cat {
	var cats []Cat

	pgConnect := PostgresConnect()

	err := pgConnect.Model(&cats).Select()
	if err != nil {
		panic(err)
	}
	pgConnect.Close()

	return cats
}

// CreateCat - создание кота
func CreateCat(cat Cat) Cat {
	pgConnect := PostgresConnect()
	_, err := pgConnect.Model(&cat).Insert()
	if err != nil {
		panic(err)
	}
	pgConnect.Close()
	return cat
}

// FindCatById - ПОЛУЧИТЬ КОТА ПО ID !!!!!!!
func FindCatById(id string) Cat {
	var cat Cat
	pgConnect := PostgresConnect()

	err := pgConnect.Model(&cat).
		Where("id = ?", id).
		First()
	if err != nil {
		panic(err)
	}
	pgConnect.Close()
	return cat
}

// DeleteCatById - УНИЧТОЖИТЬ КОТА ПО ID !!!!!!!
func DeleteCatById(id string) Cat {
	var cat Cat
	pgConnect := PostgresConnect()

	_, err := pgConnect.Model(&cat).
		Where("id = ?", id).
		Delete()
	if err != nil {
		panic(err)
	}
	pgConnect.Close()
	return cat
}

// UpdateCat - ОБНОВИТЬ КОТА !!!!!
func UpdateCat(cat Cat) Cat {
	pgConnect := PostgresConnect()

	oldCat := FindCatById(cat.ID)
	oldCat.Name = cat.Name
	oldCat.IsStrip = cat.IsStrip
	oldCat.Color = cat.Color

	_, err := pgConnect.Model(&oldCat).Where("id = ?", oldCat.ID).Update()
	if err != nil {
		panic(err)
	}

	pgConnect.Close()
	return oldCat
}
