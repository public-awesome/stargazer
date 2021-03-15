module github.com/public-awesome/stakewatcher

go 1.15

require (
	github.com/cosmos/cosmos-sdk v0.42.1
	github.com/friendsofgo/errors v0.9.2
	github.com/gofrs/uuid v3.4.0+incompatible
	github.com/kat-co/vala v0.0.0-20170210184112-42e1d8b61f12
	github.com/lib/pq v1.2.1-0.20191011153232-f91d3411e481
	github.com/markbates/pkger v0.17.1
	github.com/mitchellh/mapstructure v1.3.1 // indirect
	github.com/public-awesome/stargaze v0.6.0
	github.com/rs/zerolog v1.20.0
	github.com/rubenv/sql-migrate v0.0.0-20200616145509-8d140a17f351
	github.com/spf13/viper v1.7.1
	github.com/tendermint/tendermint v0.34.8
	github.com/volatiletech/null/v8 v8.1.1
	github.com/volatiletech/randomize v0.0.1
	github.com/volatiletech/sqlboiler/v4 v4.4.0
	github.com/volatiletech/strmangle v0.0.1
	gopkg.in/ini.v1 v1.57.0 // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
