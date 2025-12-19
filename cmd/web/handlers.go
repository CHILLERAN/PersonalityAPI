package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/CHILLERAN/PersonalityAPI/internal/config"
)

func home(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Welcome to the app, these are the categories: animal, food and color")
	}
}

func getAllTraitsByCategory(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		category := r.PathValue("category")

		allTraits, err := app.TraitModel.GetTraitsByCategory(category)

		if errors.Is(err, sql.ErrNoRows){
			notFoundError(app, w, r, err)
			return 
		} 	

		if err != nil {
			serverError(app, w, r, err)
			return 
		}

		for _, trait := range allTraits {
			fmt.Fprintf(w, "category: %v\ntrait: %v\n\n", trait.Category, trait.Trait)
		}
	}
}

func getUniqueTraitByCategory(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		category := r.PathValue("category")
		trait := r.PathValue("trait")

		foundTrait, err :=  app.TraitModel.GetTraitsByCategoryAndTrait(category, trait)

		if errors.Is(err, sql.ErrNoRows){
			notFoundError(app, w, r, err)
			return 
		} 	

		if err != nil {
			serverError(app, w, r, err)
			return 
		}

		fmt.Fprintf(w, "Trait: %v\nCategory: %v\nDescription: %v\n\n", foundTrait.Trait, foundTrait.Category, foundTrait.Description)
	}
}