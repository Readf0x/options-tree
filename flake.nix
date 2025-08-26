rec {
  description = "Description for the project";

  inputs = {
    flake-parts.url = "github:hercules-ci/flake-parts";
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    templating-engine = {
      url = "github:readf0x/templating-engine";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs = inputs @ {flake-parts, ...}:
    flake-parts.lib.mkFlake {inherit inputs;} {
      systems = ["x86_64-linux"];
      perSystem = {
        system,
        pkgs,
        ...
      }: let
        info = {
          projectName = "optionstree"; # Don't forget to change this!
          # You can set the module name as well
          # moduleName = "github.com/example/${projectName}";
        };
      in
        (
          {
            projectName,
            moduleName ? projectName,
          }: rec {
            devShells.default = pkgs.mkShell {
              packages = with pkgs; [
                go
                delve
                htmlq
                inputs.templating-engine.packages.${system}.default
              ];
            };
            packages = {
              ${projectName} = pkgs.buildGoModule rec {
                name = projectName;
                pname = name;
                version = "0.1";

                src = ./.;

                vendorHash = "sha256-AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=";

                meta = {
                  inherit description;
                  # homepage = "";
                  # license = lib.licenses.;
                  # maintainers = with lib.maintainers; [  ];
                };
              };
              default = packages.${projectName};
            };
          }
        )
        info;
    };
}
