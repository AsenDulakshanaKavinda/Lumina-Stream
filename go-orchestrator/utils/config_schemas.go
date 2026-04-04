package utils

// todo -- // - change this according to the YAML structure - // todo -- //

type Config struct {
    Logging struct {
        LogFile    string `mapstructure:"log_file"`
        LogDir     string `mapstructure:"log_dir"`
        MaxSize    int    `mapstructure:"max_size"`
        MaxAge     int    `mapstructure:"max_age"`
        MaxBackups int    `mapstructure:"max_backups"`
        Compress   bool   `mapstructure:"compress"`
    } `mapstructure:"logging"`

    GRPCConfig struct {
        ServerAddr string `mapstructure:"server_address"`
    } `mapstructure:"grpc_config"`

}
