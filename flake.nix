{
  description = "Diceware passphrase generator in French";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-parts.url = "github:hercules-ci/flake-parts";
  };

  outputs =
    inputs@{
      self,
      nixpkgs,
      flake-parts,
      ...
    }:
    flake-parts.lib.mkFlake { inherit inputs; } {
      systems = [
        "x86_64-linux"
        "aarch64-linux"
        "aarch64-darwin"
        "x86_64-darwin"
      ];

      perSystem = { pkgs, self', ... }: {
        packages.default = pkgs.buildGoModule {
          pname = "diceware-fr";
          version = "1.0.0";
          src = ./.;
          vendorHash = null;
        };

        apps.default = {
          type = "app";
          program = "${self'.packages.default}/bin/diceware-fr";
        };
      };
    };
}
