package store

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"shortener/internal/app/config"
	"shortener/internal/app/encoder"
	"shortener/internal/app/randgen"
	"shortener/pkg/utils"
	"strings"
	"time"
)

type Postgres struct {
	pool *pgxpool.Pool
}

func (st *Postgres) Print() {
	return
}

func (st *Postgres) FindInStore(ctx context.Context, shortURL string, config *config.Config) (string, error) {
	var longURL string
	q := `SELECT 
				longurl
		  FROM 
				urls
          WHERE 
				shortURL = $1;`
	row := st.pool.QueryRow(context.Background(), q, config.Options.Schema+"://"+config.Options.Prefix+"/"+shortURL)
	err := row.Scan(&longURL)
	if err != nil {
		return "", nil
	}
	return longURL, nil
}

type Row struct {
	id       uint64
	longURL  string
	shortURL string
}

func (st *Postgres) PostStore(ctx context.Context, longURL string, config *config.Config) (string, error) {
	if err := validateURL(longURL); err != nil {
		return "", err
	}
	q := `INSERT INTO urls (id, longURL, shortURL) VALUES ($1, $2, $3);`
	for {
		id := randgen.Generate()
		shortURL := config.Options.Schema + "://" + config.Options.Prefix + "/" + encoder.Encode(id)

		_, err := st.pool.Exec(context.Background(), q, id, longURL, shortURL)
		if err != nil {

			if strings.Contains(err.Error(), "urls_longurl_key") == true {
				return "", errors.New("URL is already in base")
			} else if strings.Contains(err.Error(), "urls_pkey") == true {
				continue
			} else {
				return "", err
			}

		} else {
			return shortURL, nil
		}
	}
}

func NewPostgres(config *config.Config) (*Postgres, error) {
	postgres := new(Postgres)
	err, pool := NewClient(context.Background(), config)
	if err != nil {
		return nil, err
	}
	postgres.pool = pool
	return postgres, nil
}

func NewClient(ctx context.Context, config *config.Config) (error, *pgxpool.Pool) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		config.Storage.Username,
		config.Storage.Password,
		config.Storage.Host,
		config.Storage.Port,
		config.Storage.Database)
	var pool *pgxpool.Pool

	err := utils.DoWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		var err error
		pool, err = pgxpool.Connect(ctx, dsn)
		if err != nil {
			return err
		}
		return nil
	}, config.Storage.Attempts, 5*time.Second)
	if err != nil {
		logrus.Fatal(err.Error())
	}
	return nil, pool
}
