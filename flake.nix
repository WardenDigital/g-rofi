{
  description = "Go-client wrapper over rofi";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs =
    { self, nixpkgs }:
    let
      supportedSystems = [
        "x86_64-linux"
        "aarch64-linux"
        "x86_64-darwin"
        "aarch64-darwin"
      ];

      forAllSystems =
        f:
        builtins.listToAttrs (
          map (system: {
            name = system;
            value = f system;
          }) supportedSystems
        );

    in
    {
      packages = forAllSystems (
        system:
        let
          pkgs = import nixpkgs { inherit system; };
          inherit (pkgs) lib;
        in
        {
          g-rofi = pkgs.buildGoModule (finalAttrs: {
            pname = "g-rofi";
            version = "0.1.0";
            src = ./.;
            vendorHash = "sha256-7K17JaXFsjf163g5PXCb5ng2gYdotnZ2IDKk8KFjNj0=";

            nativeBuildInputs = with pkgs; [
              makeWrapper
            ];
            disallowedRequisites = [ ];

            postInstall = ''
              mkdir -p $out/share/config
              cp -r ./config/* $out/share/config/

              wrapProgram $out/bin/g-rofi \
                --prefix PATH : ${pkgs.lib.makeBinPath [ pkgs.rofi ]} \
                --set THEME_DIR "$out/share/config" \
                --set XDG_DATA_DIRS "$out/share"
            '';

            meta = {
              description = "Go-wrapper for rofi";
              homepage = "https://github.com/WardenDigital/g-rofi";
              mainProgram = "g-rofi";
              binaryNativeCode = true;
              license = lib.licenses.mit;
              platforms = lib.platforms.unix;
              maintainers = with lib.maintainers; [ WardenDigital ];
            };
          });

          default = self.packages.${system}.g-rofi;
        }
      );
    };
}
