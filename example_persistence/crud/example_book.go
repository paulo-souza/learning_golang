package crud

import "database/sql"

// Abaixo segue exemplos de books para inserção no banco de dados

func CodigoLimpoBook(db *sql.DB) Book {
	return Book{
		Name:      "Código Limpo",
		Price:     64.71,
		Author:    "Robert C. Martins",
		Format:    "Físico",
		Genre:     "Informática",
		Publisher: "Alta Books",
		Db:        db,
	}
}

func JavaUseACabeca(db *sql.DB) Book {
	return Book{
		Name:      "Use a Cabeça Java",
		Price:     200.00,
		Author:    "Kathy Sieerra & Bert Bates",
		Format:    "Físico",
		Genre:     "Informática",
		Publisher: "Alta Books",
		Db:        db,
	}
}

func LacosDeAmor(db *sql.DB) Book {
	return Book{
		Name:      "Laços de Amor um Romance no principado de Kiev",
		Genre:     "Romance",
		Format:    "Digital",
		Publisher: "CEAC",
		Author:    "Mônica Dabus",
		Price:     41.10,
		Db:        db,
	}

}

func ALendaDeAng(db *sql.DB) Book {
	return Book{
		Name:      "Avatar: A lenda de Aang - A promessa: 1",
		Format:    "Físico",
		Author:    "Bryan Konietzko & Michael Dante DiMartino",
		Publisher: "Intrínseca",
		Genre:     "Mangá",
		Price:     45.62,
		Db:        db,
	}
}

func AMenteDoEmpreendedor(db *sql.DB) Book {
	return Book{
		Name:      "A mente do empreendedor",
		Publisher: "Astral Cultural",
		Author:    "Kevin D. Johnson ",
		Genre:     "Administração",
		Price:     39.87,
		Format:    "Digital",
		Db:        db,
	}
}

func ExperienciaDoPensamento(db *sql.DB) Book {
	return Book{
		Name:      "Experiência do pensamento ",
		Genre:     "Didático",
		Format:    "Físico",
		Author:    "Sivio Gallo",
		Price:     300.00,
		Publisher: "Scipione Didáticos",
		Db:        db,
	}
}

func TodosLivros() []Book {
	book1 := CodigoLimpoBook(nil)
	book2 := ALendaDeAng(nil)
	book3 := AMenteDoEmpreendedor(nil)
	book4 := ExperienciaDoPensamento(nil)
	book5 := LacosDeAmor(nil)
	book6 := JavaUseACabeca(nil)

	return []Book{book1, book2, book3, book4, book5, book6}
}
