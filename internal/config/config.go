package config

type Config struct {
	Workspace WorkspaceConfig `yaml:"workspace"`

	Save SaveConfig `yaml:"save"`

	Backup BackupConfig `yaml:"backup"`

	Compression CompressionConfig `yaml:"compression"`

	Storage StorageConfig `yaml:"storage"`

	S3 S3Config `yaml:"s3"`
}

type WorkspaceConfig struct {
	Path string `yaml:"path"`
}

type SaveConfig struct {
	Path string `yaml:"path"`
}

type BackupConfig struct {
	KeepLocal bool   `yaml:"keep_local"`
	Prefix    string `yaml:"prefix"`
	Interval  string `yaml:"interval"`
}

type CompressionConfig struct {
	Format string `yaml:"format"`
}

type StorageConfig struct {
	Provider string `yaml:"provider"`
}

type S3Config struct {
	Endpoint  string `yaml:"endpoint"`
	Bucket    string `yaml:"bucket"`
	AccessKey string `yaml:"access_key"`
	SecretKey string `yaml:"secret_key"`
}
