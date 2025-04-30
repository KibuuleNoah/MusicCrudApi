package database

import "MusicCrudApi/models"

// define in memory database with samples
var Songs = []models.Song{
	{ID: 1, Title: "Omwoyo", Artist: "Liam", Genre: "RnB", Released: 2020},
	{ID: 2, Title: "Bus", Artist: "Liam", Genre: "RnB", Released: 2021},
	{ID: 3, Title: "Adam", Artist: "Unkown", Genre: "RnB", Released: 2025},
}

// last DB index
var DBIdx = len(Songs)
