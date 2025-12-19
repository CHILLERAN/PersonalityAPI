package models

import (
	"database/sql"
)

type Trait struct {
	ID          int
	Category    string
	Trait      string
	Description string
}

type TraitModel struct {
	DB *sql.DB
}

func (tm TraitModel) GetTraitsByCategory(category string) ([]Trait, error) {
	query := `SELECT *
	FROM TRAITS
	WHERE LOWER(CATEGORY) = LOWER(?)`

	rows, err := tm.DB.Query(query, category)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var traits []Trait

	for rows.Next(){
		var trait Trait

		err = rows.Scan(&trait.ID, &trait.Category, &trait.Trait, &trait.Description)

		if err != nil {
			return nil, err
		}

		traits = append(traits, trait)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return traits, nil
}

func (tm TraitModel) GetTraitsByCategoryAndTrait(category, trait string) (Trait, error) {
	query := `SELECT *
	FROM TRAITS
	WHERE LOWER(CATEGORY) = LOWER(?)
	AND LOWER(TRAIT) = LOWER(?)`

	var foundTrait Trait

	err := tm.DB.QueryRow(query, category, trait).Scan(&foundTrait.ID, &foundTrait.Category, &foundTrait.Trait ,&foundTrait.Description)

	if err != nil {
		return Trait{}, err
	}

	return foundTrait, nil
}