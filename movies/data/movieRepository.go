package data

import (
	"time"

	"github.com/mismaeel/moviesapp/movies/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MovieRepository struct {
	C *mgo.Collection
}

func (r *MovieRepository) GetAll() []models.Movie {
	var movies []models.Movie

	iter := r.C.Find(nil).Sort("-rating").Iter()
	result := models.Movie{}
	for iter.Next(&result) {
		movies = append(movies, result)
	}
	return movies
}

func (r *MovieRepository) Create(movie *models.Movie) error {
	obj_id := bson.NewObjectId()
	movie.Id = obj_id
	movie.CreatedOn = time.Now()
	err := r.C.Insert(&movie)
	return err
}

func (r *MovieRepository) Update(id string,movie *models.Movie) error {
	
	err := r.C.Update(bson.M{"_id": bson.ObjectIdHex(id)},&movie)
	return err
}

func (r *MovieRepository) GetById(id string) (movie models.Movie, err error) {
	err = r.C.FindId(bson.ObjectIdHex(id)).One(&movie)
	return
}

