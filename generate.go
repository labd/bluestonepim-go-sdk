package bluestonepim_go_sdk

// Generates the clients from the specs in specs/. Run `task generate`, which
// downloads the specs first (see tools/fixspec for how they are prepared).
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config pim/config.yaml specs/pim.json
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config notification_external/config.yaml specs/notification-external.json
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config global_settings/config.yaml specs/global-settings.json
