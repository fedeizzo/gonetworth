{ pkgs, config, ... }:

{
  devshells.default = {
    devshell = {
      motd = ''
        {202}gonetworth dev shell{reset}
        $(type -p menu &>/dev/null && menu)
      '';
    };
    commands = [
      {
        help = "Build the tool";
        name = "build";
        command = "go build -o gonetworth cmd/main.go";
      }
      {
        help = "Run all tests";
        name = "test";
        command = ''
          gotestsum --format testdox
        '';
      }
      {
        help = "Run the tool";
        name = "run";
        command = "go run cmd/main.go";
      }
      {
        help = "Generate code from postgres schema/queries using sqlc";
        name = "codegen-db";
        command = "${pkgs.sqlc}/bin/sqlc generate";
      }
      {
        help = "Generate mocks from using mockery";
        name = "codegen-mock";
        command = "${pkgs.go-mockery}/bin/mockery";
      }
      {
        help = "Generate server from openapi specs";
        name = "codegen-openapi-server";
        command = "${pkgs.oapi-codegen}/bin/oapi-codegen --config=./internal/web/api/cfg.yaml openapi/openapi.yaml";
      }
    ];

    packages = with pkgs; [
      # backend
      go
      gopls
      gomodifytags
      gotestsum
      gotestdox
      sqlc
      go-mockery
      oapi-codegen

      # frontend
      nodejs_22
    ] ++ config.pre-commit.settings.enabledPackages;

    devshell.startup.pre-commit-hooks.text = config.pre-commit.installationScript;
  };
}
