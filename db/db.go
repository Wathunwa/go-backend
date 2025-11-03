package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// ประกาศ Pool เป็น connection pool ของ pgxpool ทั้ง app
type Pool = pgxpool.Pool

func Connect(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	//Custom pool
	cfg.MaxConns = 10
	cfg.MinConns = 0
	cfg.MaxConnIdleTime = 30 * time.Minute
	cfg.MaxConnLifetime = 5 * time.Minute
	cfg.HealthCheckPeriod = 60 * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}

	//ทดสอบการเชื่อมต่อ
	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := pool.Ping(pingCtx); err != nil {
		pool.Close()
		return nil, err
	}

	return pool, nil
}

func Close(p *pgxpool.Pool) {
	if p != nil {
		p.Close()
	}
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}
