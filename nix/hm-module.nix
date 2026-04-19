flake:
{
  pkgs,
  config,
  lib,
  ...
}:

let
  cfg = config.programs.g-rofi;
in
{
  options.programs.g-rofi = {
    enable = lib.mkEnableOption "g-rofi";

    package = lib.mkOption {
      type = lib.types.package;
      description = "The g-rofi package to install.";
    };

    configPath = lib.mkOption {
      type = lib.types.nullOr lib.types.path;
      default = flake.packages.${pkgs.system}.default;
      description = "Custom path to the rofi config file.";
    };
  };

  config = lib.mkIf cfg.enable {
    home.packages = [ cfg.package ];

    home.sessionVariables = {
      G_ROFI_CONFIG = if cfg.configPath != null then "${cfg.configPath}" else "";
    };
  };
}
