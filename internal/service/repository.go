package service

import (
	"fmt"

	"github.com/PrakharSrivastav/artist-service-grpc/internal/model"
	"github.com/brianvoe/gofakeit"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type repository struct {
	db *sqlx.DB
}

func (r *repository) get(ID string) (*model.Artist, error) {
	var artist model.Artist
	if err := r.db.Get(&artist, "Select * from artists where id = $1", ID); err != nil {
		return nil, err
	}
	return &artist, nil
}

func (r *repository) getAll() ([]*model.Artist, error) {
	artists := []*model.Artist{}
	if err := r.db.Select(&artists, "SELECT * FROM artists"); err != nil {
		return nil, err
	}
	return artists, nil
}

// where in ('','','') clause
func (r *repository) getIn(ids []string) ([]*model.Artist, error) {
	// {"SELECT * FROM foo WHERE x in (?)",[]interface{}{[]int{1, 2, 3, 4, 5, 6, 7, 8},8}
	//err := r.db.Select(&albums, "Select * from artists  where id in (?) ", []int{0, 5, 7, 2, 9});

	var artists []*model.Artist
	query := "Select * from artists  where id in (?) "
	q, vs, err := sqlx.In(query, ids)
	err = r.db.Select(&artists, q, vs...)
	if err != nil {
		fmt.Println("getIn", err.Error())
		return nil, err
	}
	return artists, nil
}

func (r *repository) setupDatabase() error {
	r.db.MustExec(schema)

	var artists []model.Artist
	artists = append(artists, model.Artist{Description: gofakeit.Sentence(5), Id: "artist_id_1", Name: gofakeit.Name()})
	artists = append(artists, model.Artist{Description: gofakeit.Sentence(6), Id: "artist_id_2", Name: gofakeit.Name()})
	artists = append(artists, model.Artist{Description: gofakeit.Sentence(4), Id: "artist_id_3", Name: gofakeit.Name()})
	artists = append(artists, model.Artist{Description: gofakeit.Sentence(12), Id: "artist_id_4", Name: gofakeit.Name()})
	artists = append(artists, model.Artist{Description: gofakeit.Sentence(12), Id: "artist_id_5", Name: gofakeit.Name()})
	artists = append(artists, model.Artist{Description: gofakeit.Sentence(12), Id: "artist_id_6", Name: gofakeit.Name()})
	artists = append(artists, model.Artist{Description: gofakeit.Sentence(12), Id: "artist_id_7", Name: gofakeit.Name()})
	artists = append(artists, model.Artist{Description: gofakeit.Sentence(12), Id: "artist_id_8", Name: gofakeit.Name()})
	artists = append(artists, model.Artist{Description: gofakeit.Sentence(12), Id: "artist_id_9", Name: gofakeit.Name()})
	artists = append(artists, model.Artist{Description: gofakeit.Sentence(12), Id: "artist_id_10", Name: gofakeit.Name()})
	tx := r.db.MustBegin()
	for _, u := range artists {
		fmt.Printf("%#v\n", &u)
		_, err := tx.NamedExec("INSERT into artists (id,name,description) values(:id,:name,:description)", u)
		if err != nil {
			return err
		}
	}
	tx.Commit()
	return nil
}

var schema = `
Drop Table if exists artists;
CREATE TABLE artists (
    id text,
    name text,
    description text
);`
