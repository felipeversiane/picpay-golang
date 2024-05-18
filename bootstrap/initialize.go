package bootstrap

func Initialize() error {
	err := LoadEnv()
	if err != nil {
		return err
	}
	return nil
}
