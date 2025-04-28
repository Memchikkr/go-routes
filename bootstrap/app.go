package bootstrap

type Application struct {
	Env *Env
	DB  *PostgresClient
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.DB = NewPostgresClient(app.Env)
	return *app
}

func (app *Application) Close() {
	app.DB.Close()
}
