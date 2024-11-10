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
        command = "${pkgs.gotestdox}/bin/gotestdox --format testdox";
      }
      {
        help = "Run the tool";
        name = "run";
        command = "go run cmd/main.go";
      }
    ];

    packages = with pkgs; [
      go
      gopls
      gotestsum
      gotestdox
    ] ++ config.pre-commit.settings.enabledPackages;

    devshell.startup.pre-commit-hooks.text = config.pre-commit.installationScript;
  };
}
