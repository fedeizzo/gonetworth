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

      # frontend
    ] ++ config.pre-commit.settings.enabledPackages;

    devshell.startup.pre-commit-hooks.text = config.pre-commit.installationScript;
  };
}
