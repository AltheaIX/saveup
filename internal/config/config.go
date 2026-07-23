package config

type Config struct {
	Workspace WorkspaceConfig `yaml:"workspace"`

	Save SaveConfig `yaml:"save"`

	Backup BackupConfig `yaml:"backup"`

	Compression CompressionConfig `yaml:"compression"`

	R2 R2Config `yaml:"r2"`
}

type WorkspaceConfig struct {
	Path string `yaml:"path"`
}

type SaveConfig struct {
	Path string `yaml:"path"`
}

type BackupConfig struct {
	Interval string `yaml:"interval"`
}

type CompressionConfig struct {
	Format string `yaml:"format"`
}

type R2Config struct {
	Endpoint  string `yaml:"endpoint"`
	Bucket    string `yaml:"bucket"`
	AccessKey string `yaml:"access_key"`
	SecretKey string `yaml:"secret_key"`
}
