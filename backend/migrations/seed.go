package migrations

import (
	"github.com/Amierza/warasin-mobile/backend/entity"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	err := SeedFromJSON[entity.Role](db, "./migrations/json/roles.json", entity.Role{}, "Name")
	if err != nil {
		return err
	}

	err = SeedFromJSON[entity.Province](db, "./migrations/json/provinces.json", entity.Province{}, "Name")
	if err != nil {
		return err
	}

	err = SeedFromJSON[entity.City](db, "./migrations/json/cities.json", entity.City{}, "Name")
	if err != nil {
		return err
	}

	err = SeedFromJSON[entity.User](db, "./migrations/json/users.json", entity.User{}, "Email")
	if err != nil {
		return err
	}

	err = SeedFromJSON[entity.Psycholog](db, "./migrations/json/psychologs.json", entity.Psycholog{}, "Email")
	if err != nil {
		return err
	}

	err = SeedFromJSON[entity.Consulation](db, "./migrations/json/consulations.json", entity.Consulation{}, "UserID", "PsychologID")
	if err != nil {
		return err
	}

	err = SeedFromJSON[entity.Education](db, "./migrations/json/educations.json", entity.Education{}, "PsychologID", "Major", "Degree", "Institution")
	if err != nil {
		return err
	}

	err = SeedFromJSON[entity.MotivationCategory](db, "./migrations/json/motivation_categories.json", entity.MotivationCategory{}, "Name")
	if err != nil {
		return err
	}

	err = SeedFromJSON[entity.Motivation](db, "./migrations/json/motivations.json", entity.Motivation{}, "MotivationCategoryID")
	if err != nil {
		return err
	}

	err = SeedFromJSON[entity.UserMotivation](db, "./migrations/json/user_motivations.json", entity.UserMotivation{}, "UserID", "MotivationID")
	if err != nil {
		return err
	}

	err = SeedFromJSON[entity.News](db, "./migrations/json/news.json", entity.News{}, "Title", "Date")
	if err != nil {
		return err
	}

	err = SeedFromJSON[entity.NewsDetail](db, "./migrations/json/news_details.json", entity.NewsDetail{}, "NewsID", "UserID")
	if err != nil {
		return err
	}

	err = SeedFromJSON[entity.LanguageMaster](db, "./migrations/json/language_masters.json", entity.LanguageMaster{}, "Name")
	if err != nil {
		return err
	}

	err = SeedFromJSON[entity.PsychologLanguage](db, "./migrations/json/psycholog_languages.json", entity.PsychologLanguage{}, "LanguageMasterID", "PsychologID")
	if err != nil {
		return err
	}

	err = SeedFromJSON[entity.Specialization](db, "./migrations/json/specializations.json", entity.Specialization{}, "Name")
	if err != nil {
		return err
	}

	err = SeedFromJSON[entity.SpecializationDetail](db, "./migrations/json/specialization_details.json", entity.SpecializationDetail{}, "SpecializationID", "Name")
	if err != nil {
		return err
	}

	err = SeedFromJSON[entity.PsychologSpecialization](db, "./migrations/json/psycholog_specializations.json", entity.PsychologSpecialization{}, "SpecializationID", "PsychologID")
	if err != nil {
		return err
	}

	err = SeedFromJSON[entity.Practice](db, "./migrations/json/practices.json", entity.Practice{}, "PsychologID", "Name")
	if err != nil {
		return err
	}

	err = SeedFromJSON[entity.PracticeSchedule](db, "./migrations/json/practice_schedules.json", entity.PracticeSchedule{}, "PracticeID", "Day")
	if err != nil {
		return err
	}

	err = SeedFromJSON[entity.AvailableSlot](db, "./migrations/json/available_slots.json", entity.AvailableSlot{}, "PracticeID", "Date")
	if err != nil {
		return err
	}

	err = SeedFromJSON[entity.Permission](db, "./migrations/json/permissions.json", entity.Permission{}, "RoleID", "Endpoint")
	if err != nil {
		return err
	}

	return nil
}
